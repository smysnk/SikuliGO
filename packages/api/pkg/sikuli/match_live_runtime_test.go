package sikuli_test

import (
	"context"
	"errors"
	"net"
	"testing"
	"time"

	grpcv1 "github.com/smysnk/sikuligo/internal/grpcv1"
	pb "github.com/smysnk/sikuligo/internal/grpcv1/pb"
	"github.com/smysnk/sikuligo/pkg/sikuli"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

func TestLiveMatchDelegatesSearchAndCapture(t *testing.T) {
	runtime := newMatchTestRuntime(t)
	screen, err := runtime.Screen(2)
	if err != nil {
		t.Fatalf("screen lookup failed: %v", err)
	}
	pattern := mustMatchPattern(t, [][]uint8{
		{0, 255},
		{255, 0},
	})

	match, err := screen.Find(pattern)
	if err != nil {
		t.Fatalf("screen find failed: %v", err)
	}
	if !match.Live() {
		t.Fatalf("expected live match")
	}
	if got := match.Bounds(); got.X != 5 || got.Y != 1 || got.W != 2 || got.H != 2 {
		t.Fatalf("match bounds mismatch: %+v", got)
	}
	if got := match.Center(); got != sikuli.NewPoint(6, 2) {
		t.Fatalf("match center mismatch: %+v", got)
	}
	if got := match.TargetPoint(); got != sikuli.NewPoint(6, 2) {
		t.Fatalf("match target mismatch: %+v", got)
	}

	capture, err := match.Capture()
	if err != nil {
		t.Fatalf("match capture failed: %v", err)
	}
	assertMatchGrayPixels(t, capture, 2, 2, []byte{0, 255, 255, 0})

	foundAgain, err := match.Find(pattern)
	if err != nil {
		t.Fatalf("match find failed: %v", err)
	}
	if foundAgain.X != 5 || foundAgain.Y != 1 {
		t.Fatalf("nested match mismatch: %+v", foundAgain)
	}

	existsMatch, ok, err := match.Exists(pattern, 20*time.Millisecond)
	if err != nil {
		t.Fatalf("match exists failed: %v", err)
	}
	if !ok || existsMatch.X != 5 || existsMatch.Y != 1 {
		t.Fatalf("match exists mismatch ok=%v match=%+v", ok, existsMatch)
	}

	waitMatch, err := match.Wait(pattern, 20*time.Millisecond)
	if err != nil {
		t.Fatalf("match wait failed: %v", err)
	}
	if waitMatch.X != 5 || waitMatch.Y != 1 {
		t.Fatalf("match wait mismatch: %+v", waitMatch)
	}
}

func TestPlainMatchRejectsLiveOnlyHelpers(t *testing.T) {
	match := sikuli.NewMatch(10, 20, 30, 40, 0.95, sikuli.NewPoint(0, 0))
	if match.Live() {
		t.Fatalf("plain match should not be live")
	}
	if _, err := match.Capture(); !errors.Is(err, sikuli.ErrRuntimeUnavailable) {
		t.Fatalf("expected runtime unavailable from capture, got %v", err)
	}
	if err := match.Click(sikuli.InputOptions{}); !errors.Is(err, sikuli.ErrRuntimeUnavailable) {
		t.Fatalf("expected runtime unavailable from click, got %v", err)
	}
}

func newMatchTestRuntime(t *testing.T) *sikuli.Runtime {
	t.Helper()

	screenImage := mustMatchImage(t, "screen", [][]uint8{
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

func mustMatchImage(t *testing.T, name string, rows [][]uint8) *sikuli.Image {
	t.Helper()
	img, err := sikuli.NewImageFromMatrix(name, rows)
	if err != nil {
		t.Fatalf("new image: %v", err)
	}
	return img
}

func mustMatchPattern(t *testing.T, rows [][]uint8) *sikuli.Pattern {
	t.Helper()
	img := mustMatchImage(t, "pattern", rows)
	pattern, err := sikuli.NewPattern(img)
	if err != nil {
		t.Fatalf("new pattern: %v", err)
	}
	return pattern.Exact()
}

func assertMatchGrayPixels(t *testing.T, img *sikuli.Image, wantW, wantH int, want []byte) {
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
