package sikuli_test

import (
	"context"
	"errors"
	"net"
	"strings"
	"testing"
	"time"

	grpcv1 "github.com/smysnk/sikuligo/internal/grpcv1"
	pb "github.com/smysnk/sikuligo/internal/grpcv1/pb"
	"github.com/smysnk/sikuligo/pkg/sikuli"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

func TestRuntimeScreensPrimaryAndLookup(t *testing.T) {
	runtime := newTestRuntime(t)

	screens, err := runtime.Screens()
	if err != nil {
		t.Fatalf("screens failed: %v", err)
	}
	if got := len(screens); got != 2 {
		t.Fatalf("screen count mismatch got=%d want=2", got)
	}
	if !screens[0].Live() || !screens[1].Live() {
		t.Fatalf("expected live screens")
	}

	primary, err := runtime.PrimaryScreen()
	if err != nil {
		t.Fatalf("primary screen failed: %v", err)
	}
	if primary.ID != 1 || !primary.Primary {
		t.Fatalf("primary screen mismatch: %+v", primary)
	}

	screen, err := runtime.Screen(2)
	if err != nil {
		t.Fatalf("screen lookup failed: %v", err)
	}
	if screen.ID != 2 || screen.Bounds.X != 4 {
		t.Fatalf("screen lookup mismatch: %+v", screen)
	}
	if _, err := runtime.Screen(99); err == nil {
		t.Fatalf("expected missing screen lookup to fail")
	}
}

func TestRuntimeCaptureAndRegionCapture(t *testing.T) {
	runtime := newTestRuntime(t)

	full, err := runtime.Capture()
	if err != nil {
		t.Fatalf("capture failed: %v", err)
	}
	if full.Width() != 8 || full.Height() != 4 {
		t.Fatalf("full capture dimensions mismatch got=%dx%d", full.Width(), full.Height())
	}

	screen, err := runtime.Screen(2)
	if err != nil {
		t.Fatalf("screen lookup failed: %v", err)
	}

	screenCapture, err := screen.Capture()
	if err != nil {
		t.Fatalf("screen capture failed: %v", err)
	}
	if screenCapture.Width() != 4 || screenCapture.Height() != 4 {
		t.Fatalf("screen capture dimensions mismatch got=%dx%d", screenCapture.Width(), screenCapture.Height())
	}

	regionCapture, err := screen.Region(1, 1, 2, 2).Capture()
	if err != nil {
		t.Fatalf("screen region capture failed: %v", err)
	}
	assertGrayPixels(t, regionCapture, 2, 2, []byte{0, 255, 255, 0})

	globalCapture, err := runtime.CaptureRegion(sikuli.NewRegion(5, 1, 2, 2))
	if err != nil {
		t.Fatalf("global region capture failed: %v", err)
	}
	assertGrayPixels(t, globalCapture, 2, 2, []byte{0, 255, 255, 0})
}

func TestRuntimeScreenRegionFindExistsAndWait(t *testing.T) {
	runtime := newTestRuntime(t)
	screen, err := runtime.Screen(2)
	if err != nil {
		t.Fatalf("screen lookup failed: %v", err)
	}
	pattern := mustPattern(t, [][]uint8{
		{0, 255},
		{255, 0},
	})

	match, err := screen.Find(pattern)
	if err != nil {
		t.Fatalf("screen find failed: %v", err)
	}
	if match.X != 5 || match.Y != 1 {
		t.Fatalf("screen find coordinates mismatch: %+v", match)
	}

	match, ok, err := screen.Region(1, 1, 2, 2).Exists(pattern, 20*time.Millisecond)
	if err != nil {
		t.Fatalf("screen region exists failed: %v", err)
	}
	if !ok || match.X != 5 || match.Y != 1 {
		t.Fatalf("screen region exists mismatch ok=%v match=%+v", ok, match)
	}

	match, err = screen.Region(1, 1, 2, 2).Wait(pattern, 20*time.Millisecond)
	if err != nil {
		t.Fatalf("screen region wait failed: %v", err)
	}
	if match.X != 5 || match.Y != 1 {
		t.Fatalf("screen region wait mismatch: %+v", match)
	}

	missing := mustPattern(t, [][]uint8{
		{7, 7},
		{7, 7},
	})
	_, err = screen.Wait(missing, 10*time.Millisecond)
	if err == nil {
		t.Fatalf("expected timeout for missing pattern")
	}
	if !isTimeout(err) {
		t.Fatalf("expected timeout error, got %v", err)
	}
}

func newTestRuntime(t *testing.T) *sikuli.Runtime {
	t.Helper()

	screenImage := mustImage(t, "screen", [][]uint8{
		{10, 10, 10, 10, 10, 10, 10, 10},
		{10, 0, 255, 10, 10, 0, 255, 10},
		{10, 255, 0, 10, 10, 255, 0, 10},
		{10, 10, 10, 10, 10, 10, 10, 10},
	})
	screens := []sikuli.Screen{
		{ID: 1, Name: "left", Bounds: sikuli.NewRect(0, 0, 4, 4), Primary: true},
		{ID: 2, Name: "right", Bounds: sikuli.NewRect(4, 0, 4, 4), Primary: false},
	}

	lis := bufconn.Listen(1024 * 1024)
	srv := grpc.NewServer(
		grpc.ChainUnaryInterceptor(grpcv1.UnaryInterceptors("", nil, grpcv1.NewMetricsRegistry(), nil)...),
	)
	pb.RegisterSikuliServiceServer(srv, grpcv1.NewServer(
		grpcv1.WithCaptureScreen(func(_ context.Context, _ string) (*sikuli.Image, error) { return screenImage, nil }),
		grpcv1.WithScreenLister(func(context.Context) ([]sikuli.Screen, error) { return screens, nil }),
	))
	go func() {
		_ = srv.Serve(lis)
	}()
	t.Cleanup(func() {
		srv.Stop()
		_ = lis.Close()
	})

	dialer := func(ctx context.Context, _ string) (net.Conn, error) {
		return lis.DialContext(ctx)
	}
	runtime, err := sikuli.NewRuntime("bufnet",
		sikuli.WithRuntimeContextDialer(dialer),
		sikuli.WithRuntimeDialTimeout(time.Second),
		sikuli.WithRuntimeRPCTimeout(250*time.Millisecond),
		sikuli.WithRuntimeMatcherEngine(sikuli.MatcherEngineTemplate),
	)
	if err != nil {
		t.Fatalf("new runtime failed: %v", err)
	}
	t.Cleanup(func() {
		_ = runtime.Close()
	})
	return runtime
}

func mustImage(t *testing.T, name string, rows [][]uint8) *sikuli.Image {
	t.Helper()
	img, err := sikuli.NewImageFromMatrix(name, rows)
	if err != nil {
		t.Fatalf("new image: %v", err)
	}
	return img
}

func mustPattern(t *testing.T, rows [][]uint8) *sikuli.Pattern {
	t.Helper()
	img := mustImage(t, "pattern", rows)
	pattern, err := sikuli.NewPattern(img)
	if err != nil {
		t.Fatalf("new pattern: %v", err)
	}
	return pattern.Exact()
}

func assertGrayPixels(t *testing.T, img *sikuli.Image, wantW, wantH int, want []byte) {
	t.Helper()
	if img == nil || img.Gray() == nil {
		t.Fatalf("expected image pixels")
	}
	if img.Width() != wantW || img.Height() != wantH {
		t.Fatalf("image dimensions mismatch got=%dx%d want=%dx%d", img.Width(), img.Height(), wantW, wantH)
	}
	got := img.Gray().Pix
	if len(got) != len(want) {
		t.Fatalf("image pix length mismatch got=%d want=%d", len(got), len(want))
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("image pix mismatch at %d got=%d want=%d full=%v", i, got[i], want[i], got)
		}
	}
}

func isTimeout(err error) bool {
	return errors.Is(err, sikuli.ErrTimeout) || strings.Contains(strings.ToLower(err.Error()), "timeout")
}
