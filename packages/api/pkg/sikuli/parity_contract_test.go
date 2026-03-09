package sikuli

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	core "github.com/smysnk/sikuligo/internal/core"
	pb "github.com/smysnk/sikuligo/internal/grpcv1/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestAPIParityContracts(t *testing.T) {
	t.Run("SearchSemantics", func(t *testing.T) {
		source, err := NewImageFromMatrix("contract-source", [][]uint8{
			{10, 10, 10, 10, 10},
			{10, 0, 255, 10, 10},
			{10, 255, 0, 10, 10},
			{10, 10, 10, 10, 10},
		})
		if err != nil {
			t.Fatalf("new image: %v", err)
		}
		finder, err := NewFinder(source)
		if err != nil {
			t.Fatalf("new finder: %v", err)
		}
		pattern := parityContractPattern(t, [][]uint8{{0, 255}, {255, 0}})
		missing := parityContractPattern(t, [][]uint8{{7, 7}, {7, 7}})

		match, ok, err := finder.Exists(pattern)
		if err != nil || !ok {
			t.Fatalf("image exists mismatch ok=%v err=%v", ok, err)
		}
		if match.X != 1 || match.Y != 1 {
			t.Fatalf("image exists coordinates mismatch: %+v", match)
		}
		if _, ok, err := finder.Exists(missing); err != nil || ok {
			t.Fatalf("image missing exists mismatch ok=%v err=%v", ok, err)
		}
		if _, err := finder.Wait(missing, 10*time.Millisecond); !errors.Is(err, ErrTimeout) {
			t.Fatalf("image wait timeout mismatch: %v", err)
		}
		vanished, err := finder.WaitVanish(missing, 10*time.Millisecond)
		if err != nil || !vanished {
			t.Fatalf("image wait vanish mismatch vanished=%v err=%v", vanished, err)
		}

		runtime := newParityContractRuntime(t)
		screen, err := runtime.Screen(2)
		if err != nil {
			t.Fatalf("runtime screen lookup failed: %v", err)
		}
		liveMatch, ok, err := screen.Exists(pattern, 20*time.Millisecond)
		if err != nil || !ok {
			t.Fatalf("live exists mismatch ok=%v err=%v", ok, err)
		}
		if liveMatch.X != 7 || liveMatch.Y != 1 {
			t.Fatalf("live exists coordinates mismatch: %+v", liveMatch)
		}
		if _, err := screen.Wait(missing, 10*time.Millisecond); !errors.Is(err, ErrTimeout) {
			t.Fatalf("live wait timeout mismatch: %v", err)
		}
		vanished, err = screen.WaitVanish(missing, 10*time.Millisecond)
		if err != nil || !vanished {
			t.Fatalf("live wait vanish mismatch vanished=%v err=%v", vanished, err)
		}
	})

	t.Run("LiveScreenAndRegionSurface", func(t *testing.T) {
		runtime := newParityContractRuntime(t)
		screens, err := runtime.Screens()
		if err != nil {
			t.Fatalf("screens failed: %v", err)
		}
		if len(screens) != 2 {
			t.Fatalf("screen count mismatch: %d", len(screens))
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
		if screen.TargetPoint() != NewPoint(9, 3) {
			t.Fatalf("screen target point mismatch: %+v", screen.TargetPoint())
		}
		if bounds := screen.FullRegion().Bounds(); bounds != NewRegion(6, 0, 6, 6) {
			t.Fatalf("full region bounds mismatch: %+v", bounds)
		}
		full, err := runtime.Capture()
		if err != nil {
			t.Fatalf("runtime capture failed: %v", err)
		}
		if full.Width() != 12 || full.Height() != 6 {
			t.Fatalf("runtime capture size mismatch: %dx%d", full.Width(), full.Height())
		}
		screenCapture, err := screen.Capture()
		if err != nil {
			t.Fatalf("screen capture failed: %v", err)
		}
		if screenCapture.Width() != 6 || screenCapture.Height() != 6 {
			t.Fatalf("screen capture size mismatch: %dx%d", screenCapture.Width(), screenCapture.Height())
		}
		regionCapture, err := screen.Region(1, 1, 2, 2).Capture()
		if err != nil {
			t.Fatalf("region capture failed: %v", err)
		}
		if regionCapture.Width() != 2 || regionCapture.Height() != 2 {
			t.Fatalf("region capture size mismatch: %dx%d", regionCapture.Width(), regionCapture.Height())
		}
		globalCapture, err := runtime.CaptureRegion(NewRegion(7, 1, 2, 2))
		if err != nil {
			t.Fatalf("global region capture failed: %v", err)
		}
		if globalCapture.Width() != 2 || globalCapture.Height() != 2 {
			t.Fatalf("global region capture size mismatch: %dx%d", globalCapture.Width(), globalCapture.Height())
		}
	})

	t.Run("MatchAsActionTarget", func(t *testing.T) {
		runtime := newParityContractRuntime(t)
		screen, err := runtime.Screen(2)
		if err != nil {
			t.Fatalf("screen lookup failed: %v", err)
		}
		pattern := parityContractPattern(t, [][]uint8{{0, 255}, {255, 0}})
		missing := parityContractPattern(t, [][]uint8{{7, 7}, {7, 7}})

		match, err := screen.Find(pattern)
		if err != nil {
			t.Fatalf("screen find failed: %v", err)
		}
		if !match.Live() {
			t.Fatalf("expected live match")
		}
		if match.Region() != NewRegion(7, 1, 2, 2) {
			t.Fatalf("match region mismatch: %+v", match.Region())
		}
		if match.TargetPoint() != NewPoint(8, 2) {
			t.Fatalf("match target mismatch: %+v", match.TargetPoint())
		}
		capture, err := match.Capture()
		if err != nil {
			t.Fatalf("match capture failed: %v", err)
		}
		if capture.Width() != 2 || capture.Height() != 2 {
			t.Fatalf("match capture size mismatch: %dx%d", capture.Width(), capture.Height())
		}
		matchWithin, ok, err := match.Exists(pattern, 20*time.Millisecond)
		if err != nil || !ok {
			t.Fatalf("match exists mismatch ok=%v err=%v", ok, err)
		}
		if matchWithin.X != 7 || matchWithin.Y != 1 {
			t.Fatalf("match exists coordinates mismatch: %+v", matchWithin)
		}
		if _, err := match.Wait(missing, 10*time.Millisecond); !errors.Is(err, ErrTimeout) {
			t.Fatalf("match wait timeout mismatch: %v", err)
		}
	})

	t.Run("DirectActionSurface", func(t *testing.T) {
		client := &recordingRuntimeClient{}
		runtime := &Runtime{client: client, rpcTimeout: time.Second}
		screen := Screen{ID: 1, Name: "primary", Bounds: NewRect(0, 0, 100, 100), Primary: true, runtime: runtime}
		region := screen.Region(10, 20, 30, 40)
		match := Match{Rect: NewRect(60, 70, 20, 20), Target: NewPoint(65, 75), runtime: runtime, screenID: 1, hasScreenID: true}

		if err := screen.Hover(InputOptions{}); err != nil {
			t.Fatalf("screen hover failed: %v", err)
		}
		if err := region.Click(InputOptions{}); err != nil {
			t.Fatalf("region click failed: %v", err)
		}
		if err := region.RightClick(InputOptions{}); err != nil {
			t.Fatalf("region right click failed: %v", err)
		}
		if err := match.DoubleClick(InputOptions{}); err != nil {
			t.Fatalf("match double click failed: %v", err)
		}
		if err := region.MouseDown(InputOptions{}); err != nil {
			t.Fatalf("region mouse down failed: %v", err)
		}
		if err := region.MouseUp(InputOptions{}); err != nil {
			t.Fatalf("region mouse up failed: %v", err)
		}
		if err := screen.TypeText("typed", InputOptions{}); err != nil {
			t.Fatalf("screen type failed: %v", err)
		}
		if err := screen.Paste("pasted", InputOptions{}); err != nil {
			t.Fatalf("screen paste failed: %v", err)
		}
		if err := region.Wheel(WheelDirectionDown, 2, InputOptions{}); err != nil {
			t.Fatalf("region wheel failed: %v", err)
		}
		if err := screen.KeyDown("cmd", "shift"); err != nil {
			t.Fatalf("screen key down failed: %v", err)
		}
		if err := screen.KeyUp("cmd", "shift"); err != nil {
			t.Fatalf("screen key up failed: %v", err)
		}
		if err := region.DragDrop(match, InputOptions{}); err != nil {
			t.Fatalf("region drag drop failed: %v", err)
		}

		if len(client.moveRequests) != 3 {
			t.Fatalf("move request count mismatch: %d", len(client.moveRequests))
		}
		assertMoveRequest(t, client.moveRequests[0], 50, 50)
		assertMoveRequest(t, client.moveRequests[2], 65, 75)
		if len(client.clickRequests) != 6 {
			t.Fatalf("click request count mismatch: %d", len(client.clickRequests))
		}
		assertClickRequest(t, client.clickRequests[0], 25, 40, string(MouseButtonLeft))
		assertClickRequest(t, client.clickRequests[1], 25, 40, string(MouseButtonRight))
		assertClickRequest(t, client.clickRequests[2], 65, 75, string(MouseButtonLeft))
		assertClickRequest(t, client.clickRequests[3], 65, 75, string(MouseButtonLeft))
		assertClickRequest(t, client.clickRequests[4], 50, 50, string(MouseButtonLeft))
		assertClickRequest(t, client.clickRequests[5], 50, 50, string(MouseButtonLeft))
		if len(client.mouseDownRequests) != 2 || len(client.mouseUpRequests) != 2 {
			t.Fatalf("mouse lifecycle counts mismatch down=%d up=%d", len(client.mouseDownRequests), len(client.mouseUpRequests))
		}
		assertClickRequest(t, client.mouseDownRequests[1], 25, 40, string(MouseButtonLeft))
		assertClickRequest(t, client.mouseUpRequests[1], 65, 75, string(MouseButtonLeft))
		if len(client.typeRequests) != 1 || client.typeRequests[0].GetText() != "typed" {
			t.Fatalf("type requests mismatch: %+v", client.typeRequests)
		}
		if len(client.pasteRequests) != 1 || client.pasteRequests[0].GetText() != "pasted" {
			t.Fatalf("paste requests mismatch: %+v", client.pasteRequests)
		}
		if len(client.scrollWheelRequest) != 1 || client.scrollWheelRequest[0].GetSteps() != 2 {
			t.Fatalf("wheel requests mismatch: %+v", client.scrollWheelRequest)
		}
		if len(client.keyDownRequests) != 1 || len(client.keyUpRequests) != 1 {
			t.Fatalf("key lifecycle counts mismatch down=%d up=%d", len(client.keyDownRequests), len(client.keyUpRequests))
		}
	})

	t.Run("FinderTraversalAndLifecycle", func(t *testing.T) {
		finder, pattern := testFinderWithTwoMatches(t)
		if err := finder.IterateAll(pattern); err != nil {
			t.Fatalf("iterate all failed: %v", err)
		}
		first, ok := finder.Next()
		if !ok || first.Index != 0 {
			t.Fatalf("first iterate result mismatch ok=%v match=%+v", ok, first)
		}
		second, ok := finder.Next()
		if !ok || second.Index != 1 {
			t.Fatalf("second iterate result mismatch ok=%v match=%+v", ok, second)
		}
		finder.Reset()
		if !finder.HasNext() {
			t.Fatalf("expected iterator to rewind after reset")
		}
		finder.Destroy()
		if finder.HasNext() || finder.LastMatches() != nil {
			t.Fatalf("destroy should clear iterator state")
		}
		missing := parityContractPattern(t, [][]uint8{{255, 255}, {255, 255}})
		if err := finder.Iterate(missing); err != nil {
			t.Fatalf("iterate miss should not fail: %v", err)
		}
		if finder.HasNext() {
			t.Fatalf("iterate miss should produce no matches")
		}
	})

	t.Run("MultiTargetSearchHelpers", func(t *testing.T) {
		source, err := NewImageFromMatrix("multi-source", [][]uint8{
			{10, 10, 10, 10, 10, 10, 10},
			{10, 0, 255, 10, 255, 0, 10},
			{10, 255, 0, 10, 0, 255, 10},
			{10, 10, 10, 10, 10, 10, 10},
		})
		if err != nil {
			t.Fatalf("new multi-source image: %v", err)
		}
		finder, err := NewFinder(source)
		if err != nil {
			t.Fatalf("new finder failed: %v", err)
		}
		patternA := parityContractPattern(t, [][]uint8{{0, 255}, {255, 0}})
		patternB := parityContractPattern(t, [][]uint8{{255, 0}, {0, 255}})

		matches, err := finder.FindAnyList([]*Pattern{patternB, patternA})
		if err != nil {
			t.Fatalf("finder find any failed: %v", err)
		}
		if len(matches) != 2 || matches[0].Index != 0 || matches[1].Index != 1 {
			t.Fatalf("finder find any mismatch: %+v", matches)
		}
		best, err := finder.FindBestList([]*Pattern{patternB, patternA})
		if err != nil {
			t.Fatalf("finder find best failed: %v", err)
		}
		if best.Index != 0 {
			t.Fatalf("finder best mismatch: %+v", best)
		}

		runtime := newParityContractRuntime(t)
		screen, err := runtime.Screen(2)
		if err != nil {
			t.Fatalf("screen lookup failed: %v", err)
		}
		liveMatches, err := screen.FindAnyList([]*Pattern{patternB, patternA})
		if err != nil {
			t.Fatalf("screen find any failed: %v", err)
		}
		if len(liveMatches) != 2 || liveMatches[0].Index != 0 || liveMatches[1].Index != 1 {
			t.Fatalf("screen find any mismatch: %+v", liveMatches)
		}
		waitBest, err := screen.WaitBestList([]*Pattern{patternB, patternA}, 20*time.Millisecond)
		if err != nil {
			t.Fatalf("screen wait best failed: %v", err)
		}
		if waitBest.Index != 0 || waitBest.X != 9 || waitBest.Y != 1 {
			t.Fatalf("screen wait best mismatch: %+v", waitBest)
		}
	})

	t.Run("OCRCollectionSurface", func(t *testing.T) {
		stub := &stubOCRBackend{result: coreContractOCRResult()}
		img, err := NewImageFromMatrix("ocr-src", [][]uint8{{1, 1, 1, 1}, {1, 1, 1, 1}, {1, 1, 1, 1}, {1, 1, 1, 1}})
		if err != nil {
			t.Fatalf("new image: %v", err)
		}
		finder, err := NewFinder(img)
		if err != nil {
			t.Fatalf("new finder: %v", err)
		}
		finder.SetOCRBackend(stub)
		words, err := finder.CollectWords(OCRParams{})
		if err != nil {
			t.Fatalf("finder collect words failed: %v", err)
		}
		if len(words) != 4 || words[0].Text != "Search" {
			t.Fatalf("finder collect words mismatch: %+v", words)
		}
		lines, err := finder.CollectLines(OCRParams{})
		if err != nil {
			t.Fatalf("finder collect lines failed: %v", err)
		}
		if len(lines) != 2 || lines[0].Text != "Search Pane" {
			t.Fatalf("finder collect lines mismatch: %+v", lines)
		}

		prevFactory := newOCRBackend
		newOCRBackend = func() core.OCR { return &stubOCRBackend{result: coreContractOCRResult()} }
		defer func() { newOCRBackend = prevFactory }()

		client := &fakeRuntimeOCRClient{image: mustOCRTestImage(t)}
		runtime := &Runtime{address: "fake", rpcTimeout: time.Second, matcherEngine: MatcherEngineHybrid, client: client}
		region := Screen{ID: 1, Name: "primary", Bounds: NewRect(10, 20, 20, 20), Primary: true, runtime: runtime}.Region(2, 3, 12, 10)
		liveWords, err := region.CollectWords(OCRParams{})
		if err != nil {
			t.Fatalf("live collect words failed: %v", err)
		}
		if len(liveWords) != 4 || liveWords[0].X != 13 || liveWords[0].Y != 24 {
			t.Fatalf("live collect words mismatch: %+v", liveWords)
		}
		liveLines, err := region.CollectLines(OCRParams{})
		if err != nil {
			t.Fatalf("live collect lines failed: %v", err)
		}
		if len(liveLines) != 2 || liveLines[1].Text != "Send Mail" {
			t.Fatalf("live collect lines mismatch: %+v", liveLines)
		}
	})

	t.Run("AppWindowSurface", func(t *testing.T) {
		backend := &stubAppBackend{result: core.AppResult{Windows: []core.WindowInfo{
			{ID: "mail-1", App: "Mail", PID: 1001, Title: "Inbox", X: 10, Y: 20, W: 300, H: 200, Focused: false},
			{ID: "mail-2", App: "Mail", PID: 1001, Title: "Compose", X: 20, Y: 30, W: 320, H: 220, Focused: true},
		}}}
		controller := NewAppController()
		controller.SetBackend(backend)

		windows, err := controller.FindWindows("Mail", WindowQuery{TitleContains: "mp"}, AppOptions{})
		if err != nil {
			t.Fatalf("find windows failed: %v", err)
		}
		if len(windows) != 1 || windows[0].ID != "mail-2" {
			t.Fatalf("find windows mismatch: %+v", windows)
		}
		window, ok, err := controller.GetWindow("Mail", WindowQuery{TitleContains: "mp"}, AppOptions{})
		if err != nil || !ok {
			t.Fatalf("get window mismatch ok=%v err=%v", ok, err)
		}
		if window.Bounds != NewRect(20, 30, 320, 220) || window.PID != 1001 || window.App != "Mail" {
			t.Fatalf("get window metadata mismatch: %+v", window)
		}
		focused, ok, err := controller.FocusedWindow("Mail", AppOptions{})
		if err != nil || !ok || !focused.Focused || focused.ID != "mail-2" {
			t.Fatalf("focused window mismatch ok=%v err=%v window=%+v", ok, err, focused)
		}
	})
}

func coreContractOCRResult() core.OCRResult {
	return core.OCRResult{
		Text: "Search Pane\nSend Mail",
		Words: []core.OCRWord{
			{Text: "Search", X: 1, Y: 1, W: 6, H: 2, Confidence: 0.92},
			{Text: "Pane", X: 8, Y: 1, W: 4, H: 2, Confidence: 0.87},
			{Text: "Send", X: 2, Y: 6, W: 4, H: 2, Confidence: 0.91},
			{Text: "Mail", X: 7, Y: 6, W: 4, H: 2, Confidence: 0.89},
		},
	}
}

func newParityContractRuntime(t *testing.T) *Runtime {
	t.Helper()

	screenImage, err := NewImageFromMatrix("contract-screen", [][]uint8{
		{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10},
		{10, 10, 10, 10, 10, 10, 10, 0, 255, 255, 0, 10},
		{10, 10, 10, 10, 10, 10, 10, 255, 0, 0, 255, 10},
		{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10},
		{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10},
		{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10},
	})
	if err != nil {
		t.Fatalf("new screen image failed: %v", err)
	}
	screens := []Screen{
		{ID: 1, Name: "left", Bounds: NewRect(0, 0, 6, 6), Primary: true},
		{ID: 2, Name: "right", Bounds: NewRect(6, 0, 6, 6), Primary: false},
	}
	client := &fakeParityRuntimeClient{image: screenImage, screens: screens}
	return &Runtime{
		address:       "parity-fake",
		rpcTimeout:    250 * time.Millisecond,
		matcherEngine: MatcherEngineTemplate,
		client:        client,
	}
}

func parityContractPattern(t *testing.T, rows [][]uint8) *Pattern {
	t.Helper()
	img, err := NewImageFromMatrix("contract-pattern", rows)
	if err != nil {
		t.Fatalf("new pattern image failed: %v", err)
	}
	pattern, err := NewPattern(img)
	if err != nil {
		t.Fatalf("new pattern failed: %v", err)
	}
	return pattern.Exact()
}

type fakeParityRuntimeClient struct {
	pb.SikuliServiceClient
	recorder recordingRuntimeClient
	image    *Image
	screens  []Screen
}

func (c *fakeParityRuntimeClient) ListScreens(_ context.Context, _ *pb.ListScreensRequest, _ ...grpc.CallOption) (*pb.ListScreensResponse, error) {
	out := &pb.ListScreensResponse{Screens: make([]*pb.ScreenDescriptor, 0, len(c.screens))}
	for _, screen := range c.screens {
		out.Screens = append(out.Screens, parityScreenToProto(screen))
	}
	return out, nil
}

func (c *fakeParityRuntimeClient) GetPrimaryScreen(_ context.Context, _ *pb.GetPrimaryScreenRequest, _ ...grpc.CallOption) (*pb.GetPrimaryScreenResponse, error) {
	for _, screen := range c.screens {
		if screen.Primary {
			return &pb.GetPrimaryScreenResponse{Screen: parityScreenToProto(screen)}, nil
		}
	}
	return nil, status.Error(codes.NotFound, "primary screen missing")
}

func (c *fakeParityRuntimeClient) CaptureScreen(_ context.Context, in *pb.CaptureScreenRequest, _ ...grpc.CallOption) (*pb.CaptureScreenResponse, error) {
	capture, _, err := c.capture(in.ScreenId, in.GetRegion())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	return &pb.CaptureScreenResponse{Image: grayImageFromImage(capture)}, nil
}

func (c *fakeParityRuntimeClient) MoveMouse(ctx context.Context, in *pb.MoveMouseRequest, opts ...grpc.CallOption) (*pb.ActionResponse, error) {
	return c.recorder.MoveMouse(ctx, in, opts...)
}

func (c *fakeParityRuntimeClient) Click(ctx context.Context, in *pb.ClickRequest, opts ...grpc.CallOption) (*pb.ActionResponse, error) {
	return c.recorder.Click(ctx, in, opts...)
}

func (c *fakeParityRuntimeClient) TypeText(ctx context.Context, in *pb.TypeTextRequest, opts ...grpc.CallOption) (*pb.ActionResponse, error) {
	return c.recorder.TypeText(ctx, in, opts...)
}

func (c *fakeParityRuntimeClient) PasteText(ctx context.Context, in *pb.TypeTextRequest, opts ...grpc.CallOption) (*pb.ActionResponse, error) {
	return c.recorder.PasteText(ctx, in, opts...)
}

func (c *fakeParityRuntimeClient) Hotkey(_ context.Context, _ *pb.HotkeyRequest, _ ...grpc.CallOption) (*pb.ActionResponse, error) {
	return nil, unusedRuntimeClientError()
}

func (c *fakeParityRuntimeClient) MouseDown(ctx context.Context, in *pb.ClickRequest, opts ...grpc.CallOption) (*pb.ActionResponse, error) {
	return c.recorder.MouseDown(ctx, in, opts...)
}

func (c *fakeParityRuntimeClient) MouseUp(ctx context.Context, in *pb.ClickRequest, opts ...grpc.CallOption) (*pb.ActionResponse, error) {
	return c.recorder.MouseUp(ctx, in, opts...)
}

func (c *fakeParityRuntimeClient) KeyDown(ctx context.Context, in *pb.HotkeyRequest, opts ...grpc.CallOption) (*pb.ActionResponse, error) {
	return c.recorder.KeyDown(ctx, in, opts...)
}

func (c *fakeParityRuntimeClient) KeyUp(ctx context.Context, in *pb.HotkeyRequest, opts ...grpc.CallOption) (*pb.ActionResponse, error) {
	return c.recorder.KeyUp(ctx, in, opts...)
}

func (c *fakeParityRuntimeClient) ScrollWheel(ctx context.Context, in *pb.ScrollWheelRequest, opts ...grpc.CallOption) (*pb.ActionResponse, error) {
	return c.recorder.ScrollWheel(ctx, in, opts...)
}

func (c *fakeParityRuntimeClient) FindOnScreen(_ context.Context, in *pb.FindOnScreenRequest, _ ...grpc.CallOption) (*pb.FindResponse, error) {
	match, err := c.findOnScreen(in.GetPattern(), in.GetOpts())
	if err != nil {
		if errors.Is(err, ErrFindFailed) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	return &pb.FindResponse{Match: parityMatchToProto(match)}, nil
}

func (c *fakeParityRuntimeClient) ExistsOnScreen(_ context.Context, in *pb.ExistsOnScreenRequest, _ ...grpc.CallOption) (*pb.ExistsOnScreenResponse, error) {
	match, err := c.findOnScreen(in.GetPattern(), in.GetOpts())
	if err != nil {
		if errors.Is(err, ErrFindFailed) {
			return &pb.ExistsOnScreenResponse{Exists: false}, nil
		}
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	return &pb.ExistsOnScreenResponse{Exists: true, Match: parityMatchToProto(match)}, nil
}

func (c *fakeParityRuntimeClient) WaitOnScreen(_ context.Context, in *pb.WaitOnScreenRequest, _ ...grpc.CallOption) (*pb.FindResponse, error) {
	match, err := c.findOnScreen(in.GetPattern(), in.GetOpts())
	if err != nil {
		if errors.Is(err, ErrFindFailed) {
			return nil, status.Error(codes.DeadlineExceeded, err.Error())
		}
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	return &pb.FindResponse{Match: parityMatchToProto(match)}, nil
}

func (c *fakeParityRuntimeClient) capture(screenID *int32, region *pb.Rect) (*Image, Rect, error) {
	if c == nil || c.image == nil {
		return nil, Rect{}, fmt.Errorf("runtime image unavailable")
	}
	captureRect := NewRect(0, 0, c.image.Width(), c.image.Height())
	if screenID != nil {
		screen, ok := c.screenByID(int(*screenID))
		if !ok {
			return nil, Rect{}, fmt.Errorf("screen id %d not found", *screenID)
		}
		captureRect = screen.Bounds
	}
	if region != nil {
		if screenID != nil {
			captureRect = NewRect(captureRect.X+int(region.GetX()), captureRect.Y+int(region.GetY()), int(region.GetW()), int(region.GetH()))
		} else {
			captureRect = rectFromProto(region)
		}
	}
	cropped, err := c.image.Crop(captureRect)
	if err != nil {
		return nil, Rect{}, err
	}
	normalized, err := imageFromProtoGray(grayImageFromImage(cropped), "capture")
	if err != nil {
		return nil, Rect{}, err
	}
	return normalized, captureRect, nil
}

func (c *fakeParityRuntimeClient) findOnScreen(patternProto *pb.Pattern, opts *pb.ScreenQueryOptions) (Match, error) {
	capture, offset, err := c.capture(nil, nil)
	if opts != nil {
		capture, offset, err = c.capture(opts.ScreenId, opts.Region)
	}
	if err != nil {
		return Match{}, err
	}
	pattern, err := parityPatternFromProto(patternProto)
	if err != nil {
		return Match{}, err
	}
	finder, err := NewFinder(capture)
	if err != nil {
		return Match{}, err
	}
	match, err := finder.Find(pattern)
	if err != nil {
		return Match{}, err
	}
	match.Rect = NewRect(match.X+offset.X, match.Y+offset.Y, match.W, match.H)
	match.Target = NewPoint(match.Target.X+offset.X, match.Target.Y+offset.Y)
	return match, nil
}

func (c *fakeParityRuntimeClient) screenByID(id int) (Screen, bool) {
	for _, screen := range c.screens {
		if screen.ID == id {
			return screen, true
		}
	}
	return Screen{}, false
}

func parityScreenToProto(screen Screen) *pb.ScreenDescriptor {
	return &pb.ScreenDescriptor{
		Id:      int32(screen.ID),
		Name:    screen.Name,
		Bounds:  rectToProto(screen.Bounds),
		Primary: screen.Primary,
	}
}

func parityPatternFromProto(in *pb.Pattern) (*Pattern, error) {
	if in == nil || in.GetImage() == nil {
		return nil, fmt.Errorf("%w: pattern image is nil", ErrInvalidTarget)
	}
	img, err := imageFromProtoGray(in.GetImage(), "pattern.image")
	if err != nil {
		return nil, err
	}
	pattern, err := NewPattern(img)
	if err != nil {
		return nil, err
	}
	if in.Exact != nil && in.GetExact() {
		pattern.Exact()
	} else if in.Similarity != nil {
		pattern.Similar(in.GetSimilarity())
	}
	if in.ResizeFactor != nil {
		pattern.Resize(in.GetResizeFactor())
	}
	if target := in.GetTargetOffset(); target != nil {
		pattern.TargetOffset(int(target.GetX()), int(target.GetY()))
	}
	if in.GetMask() != nil {
		mask, err := imageFromProtoGray(in.GetMask(), "pattern.mask")
		if err != nil {
			return nil, err
		}
		if _, err := pattern.WithMask(mask.Gray()); err != nil {
			return nil, err
		}
	}
	return pattern, nil
}

func parityMatchToProto(match Match) *pb.Match {
	return &pb.Match{
		Rect:  rectToProto(match.Rect),
		Score: match.Score,
		Target: &pb.Point{
			X: int32(match.Target.X),
			Y: int32(match.Target.Y),
		},
		Index: int32(match.Index),
	}
}

func (c *fakeParityRuntimeClient) Find(_ context.Context, _ *pb.FindRequest, _ ...grpc.CallOption) (*pb.FindResponse, error) {
	return nil, unusedRuntimeClientError()
}

func (c *fakeParityRuntimeClient) FindAll(_ context.Context, _ *pb.FindRequest, _ ...grpc.CallOption) (*pb.FindAllResponse, error) {
	return nil, unusedRuntimeClientError()
}

func (c *fakeParityRuntimeClient) ClickOnScreen(_ context.Context, _ *pb.ClickOnScreenRequest, _ ...grpc.CallOption) (*pb.FindResponse, error) {
	return nil, unusedRuntimeClientError()
}

func (c *fakeParityRuntimeClient) ReadText(_ context.Context, _ *pb.ReadTextRequest, _ ...grpc.CallOption) (*pb.ReadTextResponse, error) {
	return nil, unusedRuntimeClientError()
}

func (c *fakeParityRuntimeClient) FindText(_ context.Context, _ *pb.FindTextRequest, _ ...grpc.CallOption) (*pb.FindTextResponse, error) {
	return nil, unusedRuntimeClientError()
}

func (c *fakeParityRuntimeClient) ObserveAppear(_ context.Context, _ *pb.ObserveRequest, _ ...grpc.CallOption) (*pb.ObserveResponse, error) {
	return nil, unusedRuntimeClientError()
}

func (c *fakeParityRuntimeClient) ObserveVanish(_ context.Context, _ *pb.ObserveRequest, _ ...grpc.CallOption) (*pb.ObserveResponse, error) {
	return nil, unusedRuntimeClientError()
}

func (c *fakeParityRuntimeClient) ObserveChange(_ context.Context, _ *pb.ObserveChangeRequest, _ ...grpc.CallOption) (*pb.ObserveResponse, error) {
	return nil, unusedRuntimeClientError()
}

func (c *fakeParityRuntimeClient) OpenApp(_ context.Context, _ *pb.AppActionRequest, _ ...grpc.CallOption) (*pb.ActionResponse, error) {
	return nil, unusedRuntimeClientError()
}

func (c *fakeParityRuntimeClient) FocusApp(_ context.Context, _ *pb.AppActionRequest, _ ...grpc.CallOption) (*pb.ActionResponse, error) {
	return nil, unusedRuntimeClientError()
}

func (c *fakeParityRuntimeClient) CloseApp(_ context.Context, _ *pb.AppActionRequest, _ ...grpc.CallOption) (*pb.ActionResponse, error) {
	return nil, unusedRuntimeClientError()
}

func (c *fakeParityRuntimeClient) IsAppRunning(_ context.Context, _ *pb.AppActionRequest, _ ...grpc.CallOption) (*pb.IsAppRunningResponse, error) {
	return nil, unusedRuntimeClientError()
}

func (c *fakeParityRuntimeClient) ListWindows(_ context.Context, _ *pb.AppActionRequest, _ ...grpc.CallOption) (*pb.ListWindowsResponse, error) {
	return nil, unusedRuntimeClientError()
}

func (c *fakeParityRuntimeClient) FindWindows(_ context.Context, _ *pb.WindowQueryRequest, _ ...grpc.CallOption) (*pb.ListWindowsResponse, error) {
	return nil, unusedRuntimeClientError()
}

func (c *fakeParityRuntimeClient) GetWindow(_ context.Context, _ *pb.WindowQueryRequest, _ ...grpc.CallOption) (*pb.GetWindowResponse, error) {
	return nil, unusedRuntimeClientError()
}

func (c *fakeParityRuntimeClient) GetFocusedWindow(_ context.Context, _ *pb.AppActionRequest, _ ...grpc.CallOption) (*pb.GetWindowResponse, error) {
	return nil, unusedRuntimeClientError()
}
