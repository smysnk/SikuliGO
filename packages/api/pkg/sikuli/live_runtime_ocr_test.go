package sikuli

import (
	"context"
	"errors"
	"testing"
	"time"

	core "github.com/smysnk/sikuligo/internal/core"
	pb "github.com/smysnk/sikuligo/internal/grpcv1/pb"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

func TestLiveRegionOCRHelpers(t *testing.T) {
	stub := &stubOCRBackend{result: core.OCRResult{
		Text: "Search Pane\nSend Mail",
		Words: []core.OCRWord{
			{Text: "Search", X: 1, Y: 1, W: 6, H: 2, Confidence: 0.92},
			{Text: "Pane", X: 8, Y: 1, W: 4, H: 2, Confidence: 0.87},
			{Text: "Send", X: 2, Y: 6, W: 4, H: 2, Confidence: 0.91},
			{Text: "Mail", X: 7, Y: 6, W: 4, H: 2, Confidence: 0.89},
		},
	}}
	prevFactory := newOCRBackend
	newOCRBackend = func() core.OCR { return stub }
	defer func() { newOCRBackend = prevFactory }()

	client := &fakeRuntimeOCRClient{
		image: mustOCRTestImage(t),
	}
	runtime := &Runtime{
		address:       "fake",
		rpcTimeout:    time.Second,
		matcherEngine: MatcherEngineHybrid,
		client:        client,
	}
	screen := Screen{
		ID:      1,
		Name:    "primary",
		Bounds:  NewRect(10, 20, 20, 20),
		Primary: true,
		runtime: runtime,
	}
	region := screen.Region(2, 3, 12, 10)

	text, err := region.ReadText(OCRParams{Language: "eng", TrainingDataPath: "/tmp/tessdata", MinConfidence: 0.4, Timeout: 50 * time.Millisecond})
	if err != nil {
		t.Fatalf("live read text failed: %v", err)
	}
	if text != "Search Pane\nSend Mail" {
		t.Fatalf("live read text mismatch: %q", text)
	}
	if client.captureCalls != 1 {
		t.Fatalf("expected one capture for read text, got=%d", client.captureCalls)
	}
	if client.lastCapture.ScreenId == nil || *client.lastCapture.ScreenId != 1 {
		t.Fatalf("expected screen id=1 capture, got=%v", client.lastCapture.ScreenId)
	}
	if rect := client.lastCapture.GetRegion(); rect == nil || rect.GetX() != 2 || rect.GetY() != 3 || rect.GetW() != 12 || rect.GetH() != 10 {
		t.Fatalf("capture rect mismatch: %+v", rect)
	}
	if stub.lastReq.Language != "eng" || stub.lastReq.TrainingDataPath != "/tmp/tessdata" || stub.lastReq.MinConfidence != 0.4 || stub.lastReq.Timeout != 50*time.Millisecond {
		t.Fatalf("ocr params did not propagate: %+v", stub.lastReq)
	}

	client.captureCalls = 0
	words, err := region.CollectWords(OCRParams{})
	if err != nil {
		t.Fatalf("live collect words failed: %v", err)
	}
	if client.captureCalls != 1 {
		t.Fatalf("expected one capture for collect words, got=%d", client.captureCalls)
	}
	if len(words) != 4 || words[0].X != 13 || words[0].Y != 24 || words[3].X != 19 || words[3].Y != 29 {
		t.Fatalf("live collect words mismatch: %+v", words)
	}

	client.captureCalls = 0
	lines, err := region.CollectLines(OCRParams{})
	if err != nil {
		t.Fatalf("live collect lines failed: %v", err)
	}
	if client.captureCalls != 1 {
		t.Fatalf("expected one capture for collect lines, got=%d", client.captureCalls)
	}
	if len(lines) != 2 || lines[0].X != 13 || lines[0].Y != 24 || lines[1].X != 14 || lines[1].Y != 29 {
		t.Fatalf("live collect lines mismatch: %+v", lines)
	}

	client.captureCalls = 0
	matches, err := region.FindText("send", OCRParams{})
	if err != nil {
		t.Fatalf("live find text failed: %v", err)
	}
	if client.captureCalls != 1 {
		t.Fatalf("expected one capture for find text, got=%d", client.captureCalls)
	}
	if len(matches) != 1 || matches[0].Text != "Send" || matches[0].X != 14 || matches[0].Y != 29 {
		t.Fatalf("live find text mismatch: %+v", matches)
	}

	match := Match{
		Rect:          NewRect(12, 23, 12, 10),
		Target:        NewPoint(18, 28),
		runtime:       runtime,
		screenID:      1,
		hasScreenID:   true,
		screenBounds:  screen.Bounds,
		matcherEngine: MatcherEngineHybrid,
		waitScanRate:  DefaultWaitScanRate,
	}
	client.captureCalls = 0
	matchLines, err := match.CollectLines(OCRParams{})
	if err != nil {
		t.Fatalf("live match collect lines failed: %v", err)
	}
	if client.captureCalls != 1 || len(matchLines) != 2 {
		t.Fatalf("live match line capture mismatch calls=%d lines=%+v", client.captureCalls, matchLines)
	}
	if matchLines[0].X != 13 || matchLines[0].Y != 24 || matchLines[1].X != 14 || matchLines[1].Y != 29 {
		t.Fatalf("live match collect lines mismatch: %+v", matchLines)
	}
}

func mustOCRTestImage(t *testing.T) *Image {
	t.Helper()
	img, err := NewImageFromMatrix("ocr-screen", [][]uint8{
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	})
	if err != nil {
		t.Fatalf("new image failed: %v", err)
	}
	return img
}

type fakeRuntimeOCRClient struct {
	image        *Image
	lastCapture  *pb.CaptureScreenRequest
	captureCalls int
}

func (c *fakeRuntimeOCRClient) CaptureScreen(_ context.Context, in *pb.CaptureScreenRequest, _ ...grpc.CallOption) (*pb.CaptureScreenResponse, error) {
	c.captureCalls++
	c.lastCapture = in
	return &pb.CaptureScreenResponse{Image: grayImageFromImage(c.image)}, nil
}

func (*fakeRuntimeOCRClient) ListScreens(context.Context, *pb.ListScreensRequest, ...grpc.CallOption) (*pb.ListScreensResponse, error) {
	return nil, unusedRuntimeClientError()
}

func (*fakeRuntimeOCRClient) GetPrimaryScreen(context.Context, *pb.GetPrimaryScreenRequest, ...grpc.CallOption) (*pb.GetPrimaryScreenResponse, error) {
	return nil, unusedRuntimeClientError()
}

func (*fakeRuntimeOCRClient) Find(context.Context, *pb.FindRequest, ...grpc.CallOption) (*pb.FindResponse, error) {
	return nil, unusedRuntimeClientError()
}

func (*fakeRuntimeOCRClient) FindAll(context.Context, *pb.FindRequest, ...grpc.CallOption) (*pb.FindAllResponse, error) {
	return nil, unusedRuntimeClientError()
}

func (*fakeRuntimeOCRClient) FindOnScreen(context.Context, *pb.FindOnScreenRequest, ...grpc.CallOption) (*pb.FindResponse, error) {
	return nil, unusedRuntimeClientError()
}

func (*fakeRuntimeOCRClient) ExistsOnScreen(context.Context, *pb.ExistsOnScreenRequest, ...grpc.CallOption) (*pb.ExistsOnScreenResponse, error) {
	return nil, unusedRuntimeClientError()
}

func (*fakeRuntimeOCRClient) WaitOnScreen(context.Context, *pb.WaitOnScreenRequest, ...grpc.CallOption) (*pb.FindResponse, error) {
	return nil, unusedRuntimeClientError()
}

func (*fakeRuntimeOCRClient) ClickOnScreen(context.Context, *pb.ClickOnScreenRequest, ...grpc.CallOption) (*pb.FindResponse, error) {
	return nil, unusedRuntimeClientError()
}

func (*fakeRuntimeOCRClient) ReadText(context.Context, *pb.ReadTextRequest, ...grpc.CallOption) (*pb.ReadTextResponse, error) {
	return nil, unusedRuntimeClientError()
}

func (*fakeRuntimeOCRClient) FindText(context.Context, *pb.FindTextRequest, ...grpc.CallOption) (*pb.FindTextResponse, error) {
	return nil, unusedRuntimeClientError()
}

func (*fakeRuntimeOCRClient) MoveMouse(context.Context, *pb.MoveMouseRequest, ...grpc.CallOption) (*pb.ActionResponse, error) {
	return nil, unusedRuntimeClientError()
}

func (*fakeRuntimeOCRClient) Click(context.Context, *pb.ClickRequest, ...grpc.CallOption) (*pb.ActionResponse, error) {
	return nil, unusedRuntimeClientError()
}

func (*fakeRuntimeOCRClient) TypeText(context.Context, *pb.TypeTextRequest, ...grpc.CallOption) (*pb.ActionResponse, error) {
	return nil, unusedRuntimeClientError()
}

func (*fakeRuntimeOCRClient) PasteText(context.Context, *pb.TypeTextRequest, ...grpc.CallOption) (*pb.ActionResponse, error) {
	return nil, unusedRuntimeClientError()
}

func (*fakeRuntimeOCRClient) Hotkey(context.Context, *pb.HotkeyRequest, ...grpc.CallOption) (*pb.ActionResponse, error) {
	return nil, unusedRuntimeClientError()
}

func (*fakeRuntimeOCRClient) MouseDown(context.Context, *pb.ClickRequest, ...grpc.CallOption) (*pb.ActionResponse, error) {
	return nil, unusedRuntimeClientError()
}

func (*fakeRuntimeOCRClient) MouseUp(context.Context, *pb.ClickRequest, ...grpc.CallOption) (*pb.ActionResponse, error) {
	return nil, unusedRuntimeClientError()
}

func (*fakeRuntimeOCRClient) KeyDown(context.Context, *pb.HotkeyRequest, ...grpc.CallOption) (*pb.ActionResponse, error) {
	return nil, unusedRuntimeClientError()
}

func (*fakeRuntimeOCRClient) KeyUp(context.Context, *pb.HotkeyRequest, ...grpc.CallOption) (*pb.ActionResponse, error) {
	return nil, unusedRuntimeClientError()
}

func (*fakeRuntimeOCRClient) ScrollWheel(context.Context, *pb.ScrollWheelRequest, ...grpc.CallOption) (*pb.ActionResponse, error) {
	return nil, unusedRuntimeClientError()
}

func (*fakeRuntimeOCRClient) ObserveAppear(context.Context, *pb.ObserveRequest, ...grpc.CallOption) (*pb.ObserveResponse, error) {
	return nil, unusedRuntimeClientError()
}

func (*fakeRuntimeOCRClient) ObserveVanish(context.Context, *pb.ObserveRequest, ...grpc.CallOption) (*pb.ObserveResponse, error) {
	return nil, unusedRuntimeClientError()
}

func (*fakeRuntimeOCRClient) ObserveChange(context.Context, *pb.ObserveChangeRequest, ...grpc.CallOption) (*pb.ObserveResponse, error) {
	return nil, unusedRuntimeClientError()
}

func (*fakeRuntimeOCRClient) OpenApp(context.Context, *pb.AppActionRequest, ...grpc.CallOption) (*pb.ActionResponse, error) {
	return nil, unusedRuntimeClientError()
}

func (*fakeRuntimeOCRClient) FocusApp(context.Context, *pb.AppActionRequest, ...grpc.CallOption) (*pb.ActionResponse, error) {
	return nil, unusedRuntimeClientError()
}

func (*fakeRuntimeOCRClient) CloseApp(context.Context, *pb.AppActionRequest, ...grpc.CallOption) (*pb.ActionResponse, error) {
	return nil, unusedRuntimeClientError()
}

func (*fakeRuntimeOCRClient) IsAppRunning(context.Context, *pb.AppActionRequest, ...grpc.CallOption) (*pb.IsAppRunningResponse, error) {
	return nil, unusedRuntimeClientError()
}

func (*fakeRuntimeOCRClient) ListWindows(context.Context, *pb.AppActionRequest, ...grpc.CallOption) (*pb.ListWindowsResponse, error) {
	return nil, unusedRuntimeClientError()
}

func (*fakeRuntimeOCRClient) FindWindows(context.Context, *pb.WindowQueryRequest, ...grpc.CallOption) (*pb.ListWindowsResponse, error) {
	return nil, unusedRuntimeClientError()
}

func (*fakeRuntimeOCRClient) GetWindow(context.Context, *pb.WindowQueryRequest, ...grpc.CallOption) (*pb.GetWindowResponse, error) {
	return nil, unusedRuntimeClientError()
}

func (*fakeRuntimeOCRClient) GetFocusedWindow(context.Context, *pb.AppActionRequest, ...grpc.CallOption) (*pb.GetWindowResponse, error) {
	return nil, unusedRuntimeClientError()
}

func unusedRuntimeClientError() error {
	return status.Error(codes.Unimplemented, "unused in test")
}

func TestLiveRegionOCRErrorsWhenCaptureFails(t *testing.T) {
	prevFactory := newOCRBackend
	newOCRBackend = func() core.OCR { return &stubOCRBackend{} }
	defer func() { newOCRBackend = prevFactory }()

	runtime := &Runtime{
		address: "fake",
		client: &fakeRuntimeErrorClient{
			err: status.Error(codes.DeadlineExceeded, "capture timed out"),
		},
		rpcTimeout: time.Second,
	}
	_, err := runtime.Region(NewRegion(1, 2, 3, 4)).CollectWords(OCRParams{})
	if !errors.Is(err, ErrTimeout) {
		t.Fatalf("expected mapped timeout error, got=%v", err)
	}
}

type fakeRuntimeErrorClient struct {
	err error
}

func (c *fakeRuntimeErrorClient) CaptureScreen(context.Context, *pb.CaptureScreenRequest, ...grpc.CallOption) (*pb.CaptureScreenResponse, error) {
	return nil, c.err
}

func (*fakeRuntimeErrorClient) ListScreens(context.Context, *pb.ListScreensRequest, ...grpc.CallOption) (*pb.ListScreensResponse, error) {
	return nil, unusedRuntimeClientError()
}

func (*fakeRuntimeErrorClient) GetPrimaryScreen(context.Context, *pb.GetPrimaryScreenRequest, ...grpc.CallOption) (*pb.GetPrimaryScreenResponse, error) {
	return nil, unusedRuntimeClientError()
}

func (*fakeRuntimeErrorClient) Find(context.Context, *pb.FindRequest, ...grpc.CallOption) (*pb.FindResponse, error) {
	return nil, unusedRuntimeClientError()
}

func (*fakeRuntimeErrorClient) FindAll(context.Context, *pb.FindRequest, ...grpc.CallOption) (*pb.FindAllResponse, error) {
	return nil, unusedRuntimeClientError()
}

func (*fakeRuntimeErrorClient) FindOnScreen(context.Context, *pb.FindOnScreenRequest, ...grpc.CallOption) (*pb.FindResponse, error) {
	return nil, unusedRuntimeClientError()
}

func (*fakeRuntimeErrorClient) ExistsOnScreen(context.Context, *pb.ExistsOnScreenRequest, ...grpc.CallOption) (*pb.ExistsOnScreenResponse, error) {
	return nil, unusedRuntimeClientError()
}

func (*fakeRuntimeErrorClient) WaitOnScreen(context.Context, *pb.WaitOnScreenRequest, ...grpc.CallOption) (*pb.FindResponse, error) {
	return nil, unusedRuntimeClientError()
}

func (*fakeRuntimeErrorClient) ClickOnScreen(context.Context, *pb.ClickOnScreenRequest, ...grpc.CallOption) (*pb.FindResponse, error) {
	return nil, unusedRuntimeClientError()
}

func (*fakeRuntimeErrorClient) ReadText(context.Context, *pb.ReadTextRequest, ...grpc.CallOption) (*pb.ReadTextResponse, error) {
	return nil, unusedRuntimeClientError()
}

func (*fakeRuntimeErrorClient) FindText(context.Context, *pb.FindTextRequest, ...grpc.CallOption) (*pb.FindTextResponse, error) {
	return nil, unusedRuntimeClientError()
}

func (*fakeRuntimeErrorClient) MoveMouse(context.Context, *pb.MoveMouseRequest, ...grpc.CallOption) (*pb.ActionResponse, error) {
	return nil, unusedRuntimeClientError()
}

func (*fakeRuntimeErrorClient) Click(context.Context, *pb.ClickRequest, ...grpc.CallOption) (*pb.ActionResponse, error) {
	return nil, unusedRuntimeClientError()
}

func (*fakeRuntimeErrorClient) TypeText(context.Context, *pb.TypeTextRequest, ...grpc.CallOption) (*pb.ActionResponse, error) {
	return nil, unusedRuntimeClientError()
}

func (*fakeRuntimeErrorClient) PasteText(context.Context, *pb.TypeTextRequest, ...grpc.CallOption) (*pb.ActionResponse, error) {
	return nil, unusedRuntimeClientError()
}

func (*fakeRuntimeErrorClient) Hotkey(context.Context, *pb.HotkeyRequest, ...grpc.CallOption) (*pb.ActionResponse, error) {
	return nil, unusedRuntimeClientError()
}

func (*fakeRuntimeErrorClient) MouseDown(context.Context, *pb.ClickRequest, ...grpc.CallOption) (*pb.ActionResponse, error) {
	return nil, unusedRuntimeClientError()
}

func (*fakeRuntimeErrorClient) MouseUp(context.Context, *pb.ClickRequest, ...grpc.CallOption) (*pb.ActionResponse, error) {
	return nil, unusedRuntimeClientError()
}

func (*fakeRuntimeErrorClient) KeyDown(context.Context, *pb.HotkeyRequest, ...grpc.CallOption) (*pb.ActionResponse, error) {
	return nil, unusedRuntimeClientError()
}

func (*fakeRuntimeErrorClient) KeyUp(context.Context, *pb.HotkeyRequest, ...grpc.CallOption) (*pb.ActionResponse, error) {
	return nil, unusedRuntimeClientError()
}

func (*fakeRuntimeErrorClient) ScrollWheel(context.Context, *pb.ScrollWheelRequest, ...grpc.CallOption) (*pb.ActionResponse, error) {
	return nil, unusedRuntimeClientError()
}

func (*fakeRuntimeErrorClient) ObserveAppear(context.Context, *pb.ObserveRequest, ...grpc.CallOption) (*pb.ObserveResponse, error) {
	return nil, unusedRuntimeClientError()
}

func (*fakeRuntimeErrorClient) ObserveVanish(context.Context, *pb.ObserveRequest, ...grpc.CallOption) (*pb.ObserveResponse, error) {
	return nil, unusedRuntimeClientError()
}

func (*fakeRuntimeErrorClient) ObserveChange(context.Context, *pb.ObserveChangeRequest, ...grpc.CallOption) (*pb.ObserveResponse, error) {
	return nil, unusedRuntimeClientError()
}

func (*fakeRuntimeErrorClient) OpenApp(context.Context, *pb.AppActionRequest, ...grpc.CallOption) (*pb.ActionResponse, error) {
	return nil, unusedRuntimeClientError()
}

func (*fakeRuntimeErrorClient) FocusApp(context.Context, *pb.AppActionRequest, ...grpc.CallOption) (*pb.ActionResponse, error) {
	return nil, unusedRuntimeClientError()
}

func (*fakeRuntimeErrorClient) CloseApp(context.Context, *pb.AppActionRequest, ...grpc.CallOption) (*pb.ActionResponse, error) {
	return nil, unusedRuntimeClientError()
}

func (*fakeRuntimeErrorClient) IsAppRunning(context.Context, *pb.AppActionRequest, ...grpc.CallOption) (*pb.IsAppRunningResponse, error) {
	return nil, unusedRuntimeClientError()
}

func (*fakeRuntimeErrorClient) ListWindows(context.Context, *pb.AppActionRequest, ...grpc.CallOption) (*pb.ListWindowsResponse, error) {
	return nil, unusedRuntimeClientError()
}

func (*fakeRuntimeErrorClient) FindWindows(context.Context, *pb.WindowQueryRequest, ...grpc.CallOption) (*pb.ListWindowsResponse, error) {
	return nil, unusedRuntimeClientError()
}

func (*fakeRuntimeErrorClient) GetWindow(context.Context, *pb.WindowQueryRequest, ...grpc.CallOption) (*pb.GetWindowResponse, error) {
	return nil, unusedRuntimeClientError()
}

func (*fakeRuntimeErrorClient) GetFocusedWindow(context.Context, *pb.AppActionRequest, ...grpc.CallOption) (*pb.GetWindowResponse, error) {
	return nil, unusedRuntimeClientError()
}
