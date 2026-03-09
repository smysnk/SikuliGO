package grpcv1

import (
	"context"
	"image"
	"testing"
	"time"

	"github.com/smysnk/sikuligo/internal/cv"
	pb "github.com/smysnk/sikuligo/internal/grpcv1/pb"
	"github.com/smysnk/sikuligo/pkg/sikuli"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func TestFindAll(t *testing.T) {
	srv := NewServer()

	req := &pb.FindRequest{
		Source: grayImage("source", [][]uint8{
			{10, 10, 10, 10, 10, 10, 10, 10},
			{10, 0, 255, 10, 10, 10, 10, 10},
			{10, 255, 0, 10, 0, 255, 10, 10},
			{10, 10, 10, 10, 255, 0, 10, 10},
			{10, 10, 10, 10, 10, 10, 10, 10},
		}),
		Pattern: &pb.Pattern{
			Image: grayImage("needle", [][]uint8{
				{0, 255},
				{255, 0},
			}),
			Exact: boolPtr(true),
		},
	}

	res, err := srv.FindAll(context.Background(), req)
	if err != nil {
		t.Fatalf("find all failed: %v", err)
	}
	if got := len(res.GetMatches()); got != 2 {
		t.Fatalf("expected 2 matches, got %d", got)
	}
	if m := res.GetMatches()[0].GetRect(); m.GetX() != 1 || m.GetY() != 1 {
		t.Fatalf("first match mismatch: %+v", res.GetMatches()[0])
	}
	if m := res.GetMatches()[1].GetRect(); m.GetX() != 4 || m.GetY() != 2 {
		t.Fatalf("second match mismatch: %+v", res.GetMatches()[1])
	}
}

func TestFindNotFoundMapsToNotFound(t *testing.T) {
	srv := NewServer()

	req := &pb.FindRequest{
		Source: grayImage("source", [][]uint8{
			{1, 1, 1, 1},
			{1, 1, 1, 1},
			{1, 1, 1, 1},
			{1, 1, 1, 1},
		}),
		Pattern: &pb.Pattern{
			Image: grayImage("needle", [][]uint8{
				{0, 255},
				{255, 0},
			}),
			Exact: boolPtr(true),
		},
	}

	_, err := srv.Find(context.Background(), req)
	if err == nil {
		t.Fatalf("expected not found error")
	}
	if code := status.Code(err); code != codes.NotFound {
		t.Fatalf("expected not found code, got %s", code)
	}
}

func TestFindInvalidImageMapsToInvalidArgument(t *testing.T) {
	srv := NewServer()

	req := &pb.FindRequest{
		Source: &pb.GrayImage{
			Name:   "bad",
			Width:  2,
			Height: 2,
			Pix:    []byte{0, 1, 2},
		},
		Pattern: &pb.Pattern{
			Image: grayImage("needle", [][]uint8{
				{1, 1},
				{1, 1},
			}),
		},
	}

	_, err := srv.Find(context.Background(), req)
	if err == nil {
		t.Fatalf("expected invalid argument error")
	}
	if code := status.Code(err); code != codes.InvalidArgument {
		t.Fatalf("expected invalid argument code, got %s", code)
	}
}

func TestFindTextEmptyQueryMapsToInvalidArgument(t *testing.T) {
	srv := NewServer()

	req := &pb.FindTextRequest{
		Source: grayImage("source", [][]uint8{
			{1, 1},
			{1, 1},
		}),
		Query: "   ",
	}

	_, err := srv.FindText(context.Background(), req)
	if err == nil {
		t.Fatalf("expected invalid argument error")
	}
	if code := status.Code(err); code != codes.InvalidArgument {
		t.Fatalf("expected invalid argument code, got %s", code)
	}
}

func TestScreenRpcMissingPatternMapsToInvalidArgument(t *testing.T) {
	srv := NewServer()
	ctx := context.Background()

	cases := []struct {
		name string
		call func() error
	}{
		{
			name: "find_on_screen",
			call: func() error {
				_, err := srv.FindOnScreen(ctx, &pb.FindOnScreenRequest{})
				return err
			},
		},
		{
			name: "exists_on_screen",
			call: func() error {
				_, err := srv.ExistsOnScreen(ctx, &pb.ExistsOnScreenRequest{})
				return err
			},
		},
		{
			name: "wait_on_screen",
			call: func() error {
				_, err := srv.WaitOnScreen(ctx, &pb.WaitOnScreenRequest{})
				return err
			},
		},
		{
			name: "click_on_screen",
			call: func() error {
				_, err := srv.ClickOnScreen(ctx, &pb.ClickOnScreenRequest{})
				return err
			},
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			err := tc.call()
			if err == nil {
				t.Fatalf("expected invalid argument error")
			}
			if code := status.Code(err); code != codes.InvalidArgument {
				t.Fatalf("expected invalid argument code, got %s", code)
			}
		})
	}
}

func TestFindOnScreenUsesCapturedScreenImage(t *testing.T) {
	srv := NewServer()
	withMockScreenCapture(
		t,
		sikuliImageFromRows(t, "screen", [][]uint8{
			{10, 10, 10, 10, 10, 10, 10, 10},
			{10, 0, 255, 10, 10, 10, 10, 10},
			{10, 255, 0, 10, 0, 255, 10, 10},
			{10, 10, 10, 10, 255, 0, 10, 10},
			{10, 10, 10, 10, 10, 10, 10, 10},
		}),
	)

	res, err := srv.FindOnScreen(context.Background(), &pb.FindOnScreenRequest{
		Pattern: &pb.Pattern{
			Image: grayImage("needle", [][]uint8{
				{0, 255},
				{255, 0},
			}),
			Exact: boolPtr(true),
		},
	})
	if err != nil {
		t.Fatalf("find_on_screen failed: %v", err)
	}
	if res.GetMatch().GetRect().GetX() != 1 || res.GetMatch().GetRect().GetY() != 1 {
		t.Fatalf("unexpected match: %+v", res.GetMatch())
	}
}

func TestExistsOnScreenFalseWhenNotFound(t *testing.T) {
	srv := NewServer()
	withMockScreenCapture(
		t,
		sikuliImageFromRows(t, "screen", [][]uint8{
			{1, 1, 1, 1},
			{1, 1, 1, 1},
			{1, 1, 1, 1},
			{1, 1, 1, 1},
		}),
	)

	res, err := srv.ExistsOnScreen(context.Background(), &pb.ExistsOnScreenRequest{
		Pattern: &pb.Pattern{
			Image: grayImage("needle", [][]uint8{
				{0, 255},
				{255, 0},
			}),
			Exact: boolPtr(true),
		},
	})
	if err != nil {
		t.Fatalf("exists_on_screen failed: %v", err)
	}
	if res.GetExists() {
		t.Fatalf("expected exists=false, got true")
	}
}

func TestWaitOnScreenTimeoutMapsToDeadlineExceeded(t *testing.T) {
	srv := NewServer()
	withMockScreenCapture(
		t,
		sikuliImageFromRows(t, "screen", [][]uint8{
			{1, 1, 1, 1},
			{1, 1, 1, 1},
			{1, 1, 1, 1},
			{1, 1, 1, 1},
		}),
	)

	_, err := srv.WaitOnScreen(context.Background(), &pb.WaitOnScreenRequest{
		Pattern: &pb.Pattern{
			Image: grayImage("needle", [][]uint8{
				{0, 255},
				{255, 0},
			}),
			Exact: boolPtr(true),
		},
		Opts: &pb.ScreenQueryOptions{
			TimeoutMillis:  int64Ptr(1),
			IntervalMillis: int64Ptr(1),
		},
	})
	if err == nil {
		t.Fatalf("expected deadline exceeded error")
	}
	if code := status.Code(err); code != codes.DeadlineExceeded {
		t.Fatalf("expected deadline exceeded, got %s", code)
	}
}

func TestScreenSearchSemanticsIgnoreLegacyThrowSetting(t *testing.T) {
	t.Cleanup(func() {
		sikuli.ResetSettings()
	})
	sikuli.UpdateSettings(func(in *sikuli.RuntimeSettings) {
		in.FindFailedThrows = false
	})

	srv := NewServer()
	withMockScreenCapture(
		t,
		sikuliImageFromRows(t, "screen", [][]uint8{
			{1, 1, 1, 1},
			{1, 1, 1, 1},
			{1, 1, 1, 1},
			{1, 1, 1, 1},
		}),
	)

	pattern := &pb.Pattern{
		Image: grayImage("needle", [][]uint8{
			{0, 255},
			{255, 0},
		}),
		Exact: boolPtr(true),
	}

	_, err := srv.FindOnScreen(context.Background(), &pb.FindOnScreenRequest{Pattern: pattern})
	if code := status.Code(err); code != codes.NotFound {
		t.Fatalf("find_on_screen miss mismatch: got %s err=%v", code, err)
	}

	existsRes, err := srv.ExistsOnScreen(context.Background(), &pb.ExistsOnScreenRequest{
		Pattern: pattern,
		Opts: &pb.ScreenQueryOptions{
			TimeoutMillis:  int64Ptr(int64((2 * time.Millisecond) / time.Millisecond)),
			IntervalMillis: int64Ptr(1),
		},
	})
	if err != nil {
		t.Fatalf("exists_on_screen failed: %v", err)
	}
	if existsRes.GetExists() {
		t.Fatalf("expected exists=false on miss")
	}

	_, err = srv.WaitOnScreen(context.Background(), &pb.WaitOnScreenRequest{
		Pattern: pattern,
		Opts: &pb.ScreenQueryOptions{
			TimeoutMillis:  int64Ptr(int64((2 * time.Millisecond) / time.Millisecond)),
			IntervalMillis: int64Ptr(1),
		},
	})
	if code := status.Code(err); code != codes.DeadlineExceeded {
		t.Fatalf("wait_on_screen miss mismatch: got %s err=%v", code, err)
	}
}

func TestClickOnScreenInvokesClickBackend(t *testing.T) {
	srv := NewServer()
	withMockScreenCapture(
		t,
		sikuliImageFromRows(t, "screen", [][]uint8{
			{10, 10, 10, 10, 10, 10, 10, 10},
			{10, 0, 255, 10, 10, 10, 10, 10},
			{10, 255, 0, 10, 0, 255, 10, 10},
			{10, 10, 10, 10, 255, 0, 10, 10},
			{10, 10, 10, 10, 10, 10, 10, 10},
		}),
	)

	var clicked struct {
		x int
		y int
	}
	original := clickOnScreenFn
	clickOnScreenFn = func(x, y int, _ sikuli.InputOptions) error {
		clicked.x = x
		clicked.y = y
		return nil
	}
	t.Cleanup(func() {
		clickOnScreenFn = original
	})

	res, err := srv.ClickOnScreen(context.Background(), &pb.ClickOnScreenRequest{
		Pattern: &pb.Pattern{
			Image: grayImage("needle", [][]uint8{
				{0, 255},
				{255, 0},
			}),
			Exact: boolPtr(true),
		},
	})
	if err != nil {
		t.Fatalf("click_on_screen failed: %v", err)
	}
	if res.GetMatch() == nil || res.GetMatch().GetTarget() == nil {
		t.Fatalf("expected match target in response")
	}
	if clicked.x != int(res.GetMatch().GetTarget().GetX()) || clicked.y != int(res.GetMatch().GetTarget().GetY()) {
		t.Fatalf("clicked coordinates mismatch clicked=(%d,%d) target=(%d,%d)", clicked.x, clicked.y, res.GetMatch().GetTarget().GetX(), res.GetMatch().GetTarget().GetY())
	}
}

func TestFindInvalidMatcherEngineMapsToInvalidArgument(t *testing.T) {
	srv := NewServer()
	req := &pb.FindRequest{
		Source: grayImage("source", [][]uint8{
			{10, 10, 10},
			{10, 10, 10},
			{10, 10, 10},
		}),
		Pattern: &pb.Pattern{
			Image: grayImage("needle", [][]uint8{
				{10},
			}),
			Exact: boolPtr(true),
		},
		MatcherEngine: pb.MatcherEngine(99),
	}
	_, err := srv.Find(context.Background(), req)
	if err == nil {
		t.Fatalf("expected invalid argument error")
	}
	if code := status.Code(err); code != codes.InvalidArgument {
		t.Fatalf("expected invalid argument code, got %s", code)
	}
}

func TestFindHybridMatcherViaProtoField(t *testing.T) {
	srv := NewServer()
	req := &pb.FindRequest{
		Source: grayImage("source", [][]uint8{
			{10, 10, 10, 10},
			{10, 0, 255, 10},
			{10, 255, 0, 10},
			{10, 10, 10, 10},
		}),
		Pattern: &pb.Pattern{
			Image: grayImage("needle", [][]uint8{
				{0, 255},
				{255, 0},
			}),
			Exact: boolPtr(true),
		},
		MatcherEngine: pb.MatcherEngine_MATCHER_ENGINE_HYBRID,
	}
	res, err := srv.Find(context.Background(), req)
	if err != nil {
		t.Fatalf("find failed: %v", err)
	}
	if res.GetMatch() == nil {
		t.Fatalf("expected match in response")
	}
}

func TestFindIgnoresLegacyMatcherHeader(t *testing.T) {
	srv := NewServer()
	req := &pb.FindRequest{
		Source: grayImage("source", [][]uint8{
			{10, 10, 10, 10},
			{10, 0, 255, 10},
			{10, 255, 0, 10},
			{10, 10, 10, 10},
		}),
		Pattern: &pb.Pattern{
			Image: grayImage("needle", [][]uint8{
				{0, 255},
				{255, 0},
			}),
			Exact: boolPtr(true),
		},
	}
	ctx := metadata.NewIncomingContext(context.Background(), metadata.Pairs("x-sikuli-go-engine", "not-a-valid-engine"))
	res, err := srv.Find(ctx, req)
	if err != nil {
		t.Fatalf("find failed: %v", err)
	}
	if res.GetMatch() == nil {
		t.Fatalf("expected match in response")
	}
}

func TestFindOnScreenInvalidMatcherEngineMapsToInvalidArgument(t *testing.T) {
	srv := NewServer()
	req := &pb.FindOnScreenRequest{
		Pattern: &pb.Pattern{
			Image: grayImage("needle", [][]uint8{
				{10},
			}),
		},
		Opts: &pb.ScreenQueryOptions{
			MatcherEngine: pb.MatcherEngine(99),
		},
	}
	_, err := srv.FindOnScreen(context.Background(), req)
	if err == nil {
		t.Fatalf("expected invalid argument error")
	}
	if code := status.Code(err); code != codes.InvalidArgument {
		t.Fatalf("expected invalid argument code, got %s", code)
	}
}

func TestMatcherEngineDefaultsToHybrid(t *testing.T) {
	engine, err := matcherEngineFromFindRequest(nil)
	if err != nil {
		t.Fatalf("unexpected error for nil find request: %v", err)
	}
	if engine != cv.MatcherEngineHybrid {
		t.Fatalf("nil find request default mismatch got=%q want=%q", engine, cv.MatcherEngineHybrid)
	}

	engine, err = matcherEngineFromScreenOptions(nil)
	if err != nil {
		t.Fatalf("unexpected error for nil screen options: %v", err)
	}
	if engine != cv.MatcherEngineHybrid {
		t.Fatalf("nil screen options default mismatch got=%q want=%q", engine, cv.MatcherEngineHybrid)
	}

	engine, err = matcherEngineFromProto(pb.MatcherEngine_MATCHER_ENGINE_UNSPECIFIED)
	if err != nil {
		t.Fatalf("unexpected error for unspecified enum: %v", err)
	}
	if engine != cv.MatcherEngineHybrid {
		t.Fatalf("unspecified enum default mismatch got=%q want=%q", engine, cv.MatcherEngineHybrid)
	}
}

func TestMatcherEngineEnumMappings(t *testing.T) {
	tests := []struct {
		name string
		in   pb.MatcherEngine
		want cv.MatcherEngine
	}{
		{name: "template", in: pb.MatcherEngine_MATCHER_ENGINE_TEMPLATE, want: cv.MatcherEngineTemplate},
		{name: "orb", in: pb.MatcherEngine_MATCHER_ENGINE_ORB, want: cv.MatcherEngineORB},
		{name: "akaze", in: pb.MatcherEngine_MATCHER_ENGINE_AKAZE, want: cv.MatcherEngineAKAZE},
		{name: "brisk", in: pb.MatcherEngine_MATCHER_ENGINE_BRISK, want: cv.MatcherEngineBRISK},
		{name: "kaze", in: pb.MatcherEngine_MATCHER_ENGINE_KAZE, want: cv.MatcherEngineKAZE},
		{name: "sift", in: pb.MatcherEngine_MATCHER_ENGINE_SIFT, want: cv.MatcherEngineSIFT},
		{name: "hybrid", in: pb.MatcherEngine_MATCHER_ENGINE_HYBRID, want: cv.MatcherEngineHybrid},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			engine, err := matcherEngineFromProto(tc.in)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if engine != tc.want {
				t.Fatalf("engine enum mismatch got=%q want=%q", engine, tc.want)
			}
		})
	}
}

func withMockScreenCapture(t *testing.T, img *sikuli.Image) {
	t.Helper()
	original := captureScreenFn
	captureScreenFn = func(_ context.Context, _ string) (*sikuli.Image, error) {
		return img, nil
	}
	t.Cleanup(func() {
		captureScreenFn = original
	})
}

func grayImage(name string, rows [][]uint8) *pb.GrayImage {
	h := len(rows)
	w := len(rows[0])
	pix := make([]byte, 0, w*h)
	for _, row := range rows {
		pix = append(pix, row...)
	}
	return &pb.GrayImage{
		Name:   name,
		Width:  int32(w),
		Height: int32(h),
		Pix:    pix,
	}
}

func boolPtr(v bool) *bool {
	return &v
}

func int64Ptr(v int64) *int64 {
	return &v
}

func sikuliImageFromRows(t *testing.T, name string, rows [][]uint8) *sikuli.Image {
	t.Helper()
	h := len(rows)
	w := len(rows[0])
	gray := image.NewGray(image.Rect(0, 0, w, h))
	for y := 0; y < h; y++ {
		copy(gray.Pix[y*w:(y+1)*w], rows[y])
	}
	out, err := sikuli.NewImageFromGray(name, gray)
	if err != nil {
		t.Fatalf("build gray image: %v", err)
	}
	return out
}
