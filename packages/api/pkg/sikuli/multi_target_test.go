package sikuli_test

import (
	"context"
	"errors"
	"image"
	"image/color"
	"net"
	"testing"
	"time"

	grpcv1 "github.com/smysnk/sikuligo/internal/grpcv1"
	pb "github.com/smysnk/sikuligo/internal/grpcv1/pb"
	"github.com/smysnk/sikuligo/pkg/sikuli"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

func TestFinderMultiTargetHelpers(t *testing.T) {
	source := mustImage(t, "multi-source", [][]uint8{
		{10, 10, 10, 10, 10, 10, 10},
		{10, 0, 255, 10, 255, 0, 10},
		{10, 255, 0, 10, 0, 255, 10},
		{10, 10, 10, 10, 10, 10, 10},
	})
	finder, err := sikuli.NewFinder(source)
	if err != nil {
		t.Fatalf("new finder failed: %v", err)
	}
	patternA := mustPattern(t, [][]uint8{{0, 255}, {255, 0}})
	patternB := mustPattern(t, [][]uint8{{255, 0}, {0, 255}})
	missing := mustPattern(t, [][]uint8{{7, 7}, {7, 7}})

	matches, err := finder.FindAnyList([]*sikuli.Pattern{patternB, missing, patternA})
	if err != nil {
		t.Fatalf("find any list failed: %v", err)
	}
	if len(matches) != 2 {
		t.Fatalf("find any list count mismatch got=%d want=2", len(matches))
	}
	if matches[0].Index != 0 || matches[0].X != 4 || matches[0].Y != 1 {
		t.Fatalf("first multi-target match mismatch: %+v", matches[0])
	}
	if matches[1].Index != 2 || matches[1].X != 1 || matches[1].Y != 1 {
		t.Fatalf("second multi-target match mismatch: %+v", matches[1])
	}
	last := finder.LastMatches()
	if len(last) != 2 || last[0].Index != 0 || last[1].Index != 2 {
		t.Fatalf("last matches mismatch: %+v", last)
	}

	best, err := finder.FindBestList([]*sikuli.Pattern{patternB, patternA})
	if err != nil {
		t.Fatalf("find best list failed: %v", err)
	}
	if best.Index != 0 || best.X != 4 || best.Y != 1 {
		t.Fatalf("best multi-target match mismatch: %+v", best)
	}

	waitMatches, err := finder.WaitAnyList([]*sikuli.Pattern{missing, patternA}, 20*time.Millisecond)
	if err != nil {
		t.Fatalf("wait any list failed: %v", err)
	}
	if len(waitMatches) != 1 || waitMatches[0].Index != 1 || waitMatches[0].X != 1 || waitMatches[0].Y != 1 {
		t.Fatalf("wait any list mismatch: %+v", waitMatches)
	}

	waitBest, err := finder.WaitBestList([]*sikuli.Pattern{missing, patternB}, 20*time.Millisecond)
	if err != nil {
		t.Fatalf("wait best list failed: %v", err)
	}
	if waitBest.Index != 1 || waitBest.X != 4 || waitBest.Y != 1 {
		t.Fatalf("wait best list mismatch: %+v", waitBest)
	}

	if _, err := finder.FindAnyList([]*sikuli.Pattern{missing}); !errors.Is(err, sikuli.ErrFindFailed) {
		t.Fatalf("expected find any miss to return ErrFindFailed, got %v", err)
	}
	if _, err := finder.WaitAnyList([]*sikuli.Pattern{missing}, 10*time.Millisecond); !errors.Is(err, sikuli.ErrTimeout) {
		t.Fatalf("expected wait any miss to return ErrTimeout, got %v", err)
	}
}

func TestLiveMultiTargetHelpersUseSingleCaptureAndAbsoluteCoords(t *testing.T) {
	captureCalls := 0
	runtime := newPhase6Runtime(t, &captureCalls)
	screen, err := runtime.Screen(1)
	if err != nil {
		t.Fatalf("screen lookup failed: %v", err)
	}
	patternA := mustPattern(t, [][]uint8{{0, 255}, {255, 0}})
	patternB := mustPattern(t, [][]uint8{{255, 0}, {0, 255}})

	matches, err := screen.FindAnyList([]*sikuli.Pattern{patternB, patternA})
	if err != nil {
		t.Fatalf("screen find any list failed: %v", err)
	}
	if captureCalls != 1 {
		t.Fatalf("expected one capture for find any list, got %d", captureCalls)
	}
	if len(matches) != 2 {
		t.Fatalf("screen find any count mismatch got=%d want=2", len(matches))
	}
	if !matches[0].Live() || !matches[1].Live() {
		t.Fatalf("expected live matches from screen multi-target helper")
	}
	if matches[0].Index != 0 || matches[0].X != 24 || matches[0].Y != 31 {
		t.Fatalf("first live multi-target match mismatch: %+v", matches[0])
	}
	if matches[1].Index != 1 || matches[1].X != 21 || matches[1].Y != 31 {
		t.Fatalf("second live multi-target match mismatch: %+v", matches[1])
	}

	captureCalls = 0
	best, err := screen.FindBestList([]*sikuli.Pattern{patternB, patternA})
	if err != nil {
		t.Fatalf("screen find best list failed: %v", err)
	}
	if captureCalls != 1 {
		t.Fatalf("expected one capture for find best list, got %d", captureCalls)
	}
	if best.Index != 0 || best.X != 24 || best.Y != 31 {
		t.Fatalf("screen best multi-target match mismatch: %+v", best)
	}

	captureCalls = 0
	waitMatches, err := screen.WaitAnyList([]*sikuli.Pattern{patternA, patternB}, 20*time.Millisecond)
	if err != nil {
		t.Fatalf("screen wait any list failed: %v", err)
	}
	if captureCalls != 1 {
		t.Fatalf("expected one capture for wait any list success, got %d", captureCalls)
	}
	if len(waitMatches) != 2 {
		t.Fatalf("screen wait any count mismatch got=%d want=2", len(waitMatches))
	}

	captureCalls = 0
	matchAny, err := best.FindAnyList([]*sikuli.Pattern{patternB, patternA})
	if err != nil {
		t.Fatalf("live match find any list failed: %v", err)
	}
	if captureCalls != 1 {
		t.Fatalf("expected one capture for live match helper, got %d", captureCalls)
	}
	if len(matchAny) != 1 || matchAny[0].Index != 0 || matchAny[0].X != 24 || matchAny[0].Y != 31 {
		t.Fatalf("live match multi-target mismatch: %+v", matchAny)
	}
}

func newPhase6Runtime(t *testing.T, captureCalls *int) *sikuli.Runtime {
	t.Helper()
	desktop := image.NewGray(image.Rect(0, 0, 30, 40))
	for y := 0; y < 40; y++ {
		for x := 0; x < 30; x++ {
			desktop.SetGray(x, y, color.Gray{Y: 10})
		}
	}
	setPatch := func(x, y int, rows [][]uint8) {
		for dy := range rows {
			for dx := range rows[dy] {
				desktop.SetGray(x+dx, y+dy, color.Gray{Y: rows[dy][dx]})
			}
		}
	}
	setPatch(21, 31, [][]uint8{{0, 255}, {255, 0}})
	setPatch(24, 31, [][]uint8{{255, 0}, {0, 255}})
	full, err := sikuli.NewImageFromGray("desktop", desktop)
	if err != nil {
		t.Fatalf("new image from gray failed: %v", err)
	}
	screens := []sikuli.Screen{{ID: 1, Name: "primary", Bounds: sikuli.NewRect(20, 30, 7, 4), Primary: true}}

	lis := bufconn.Listen(1024 * 1024)
	srv := grpc.NewServer(
		grpc.ChainUnaryInterceptor(grpcv1.UnaryInterceptors("", nil, grpcv1.NewMetricsRegistry(), nil)...),
	)
	pb.RegisterSikuliServiceServer(srv, grpcv1.NewServer(
		grpcv1.WithCaptureScreen(func(_ context.Context, _ string) (*sikuli.Image, error) {
			*captureCalls++
			return full, nil
		}),
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
