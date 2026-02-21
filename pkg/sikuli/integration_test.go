package sikuli

import (
	"image"
	"image/color"
	"testing"
	"time"

	"github.com/sikulix/portgo/internal/core"
)

type integrationOCRBackend struct {
	result core.OCRResult
	err    error
}

func (b *integrationOCRBackend) Read(req core.OCRRequest) (core.OCRResult, error) {
	if err := req.Validate(); err != nil {
		return core.OCRResult{}, err
	}
	if b.err != nil {
		return core.OCRResult{}, b.err
	}
	return b.result, nil
}

type integrationInputBackend struct {
	source   *image.Gray
	pattern  *image.Gray
	placeX   int
	placeY   int
	requests []core.InputRequest
}

func (b *integrationInputBackend) Execute(req core.InputRequest) error {
	if err := req.Validate(); err != nil {
		return err
	}
	b.requests = append(b.requests, req)
	if req.Action != core.InputActionClick || b.source == nil || b.pattern == nil {
		return nil
	}
	pb := b.pattern.Bounds()
	for y := 0; y < pb.Dy(); y++ {
		for x := 0; x < pb.Dx(); x++ {
			b.source.SetGray(b.placeX+x, b.placeY+y, b.pattern.GrayAt(pb.Min.X+x, pb.Min.Y+y))
		}
	}
	return nil
}

type integrationAppBackend struct {
	running bool
	windows []core.WindowInfo
	log     []core.AppRequest
}

func (b *integrationAppBackend) Execute(req core.AppRequest) (core.AppResult, error) {
	if err := req.Validate(); err != nil {
		return core.AppResult{}, err
	}
	b.log = append(b.log, req)

	switch req.Action {
	case core.AppActionOpen:
		b.running = true
		b.windows = []core.WindowInfo{
			{Title: req.Name, X: 100, Y: 200, W: 640, H: 480, Focused: false},
		}
		return core.AppResult{Running: true, Windows: cloneWindows(b.windows)}, nil
	case core.AppActionFocus:
		for i := range b.windows {
			b.windows[i].Focused = true
		}
		return core.AppResult{Running: b.running, Windows: cloneWindows(b.windows)}, nil
	case core.AppActionClose:
		b.running = false
		b.windows = nil
		return core.AppResult{Running: false, Windows: nil}, nil
	case core.AppActionIsRunning:
		return core.AppResult{Running: b.running, Windows: cloneWindows(b.windows)}, nil
	case core.AppActionListWindow:
		return core.AppResult{Running: b.running, Windows: cloneWindows(b.windows)}, nil
	default:
		return core.AppResult{}, nil
	}
}

func cloneWindows(in []core.WindowInfo) []core.WindowInfo {
	if len(in) == 0 {
		return nil
	}
	out := make([]core.WindowInfo, len(in))
	copy(out, in)
	return out
}

func runCrossProtocolIntegrationFlow(t *testing.T, appName string) {
	t.Helper()

	source, err := NewImageFromMatrix("integration-src", [][]uint8{
		{0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0},
	})
	if err != nil {
		t.Fatalf("new source image failed: %v", err)
	}
	patternImage, err := NewImageFromMatrix("integration-pattern", [][]uint8{
		{10, 200},
		{220, 15},
	})
	if err != nil {
		t.Fatalf("new pattern image failed: %v", err)
	}
	pattern, err := NewPattern(patternImage)
	if err != nil {
		t.Fatalf("new pattern failed: %v", err)
	}
	pattern.Exact()

	finder, err := NewFinder(source)
	if err != nil {
		t.Fatalf("new finder failed: %v", err)
	}
	finder.SetOCRBackend(&integrationOCRBackend{
		result: core.OCRResult{
			Text: "Open Demo",
			Words: []core.OCRWord{
				{Text: "Open", X: 1, Y: 1, W: 2, H: 1, Confidence: 0.99},
				{Text: "Demo", X: 4, Y: 1, W: 2, H: 1, Confidence: 0.95},
			},
		},
	})

	textMatches, err := finder.FindText("open", OCRParams{})
	if err != nil {
		t.Fatalf("find text failed: %v", err)
	}
	if len(textMatches) != 1 {
		t.Fatalf("expected one OCR match, got=%d", len(textMatches))
	}

	appBackend := &integrationAppBackend{}
	appController := NewAppController()
	appController.SetBackend(appBackend)
	if err := appController.Open(appName, []string{"--smoke"}, AppOptions{Timeout: 100 * time.Millisecond}); err != nil {
		t.Fatalf("app open failed: %v", err)
	}
	running, err := appController.IsRunning(appName, AppOptions{})
	if err != nil {
		t.Fatalf("app is-running failed: %v", err)
	}
	if !running {
		t.Fatalf("expected running=true after open")
	}

	observer := NewObserverController()
	region := NewRegion(0, 0, source.Width(), source.Height())
	before, err := observer.ObserveAppear(source, region, pattern, ObserveOptions{Interval: 1 * time.Millisecond, Timeout: 0})
	if err != nil {
		t.Fatalf("observe before click failed: %v", err)
	}
	if len(before) != 0 {
		t.Fatalf("expected no observe events before click, got=%+v", before)
	}

	inputBackend := &integrationInputBackend{
		source:  source.Gray(),
		pattern: patternImage.Gray(),
		placeX:  2,
		placeY:  2,
	}
	inputController := NewInputController()
	inputController.SetBackend(inputBackend)

	clickX := textMatches[0].X + textMatches[0].W/2
	clickY := textMatches[0].Y + textMatches[0].H/2
	if err := inputController.Click(clickX, clickY, InputOptions{Button: MouseButtonLeft}); err != nil {
		t.Fatalf("click failed: %v", err)
	}
	if len(inputBackend.requests) != 1 || inputBackend.requests[0].Action != core.InputActionClick {
		t.Fatalf("input request log mismatch: %+v", inputBackend.requests)
	}

	after, err := observer.ObserveAppear(source, region, pattern, ObserveOptions{Interval: 1 * time.Millisecond, Timeout: 0})
	if err != nil {
		t.Fatalf("observe after click failed: %v", err)
	}
	if len(after) != 1 {
		t.Fatalf("expected one observe appear event after click, got=%+v", after)
	}
	if after[0].Type != ObserveEventAppear || after[0].Match.X != 2 || after[0].Match.Y != 2 {
		t.Fatalf("observe event mismatch: %+v", after[0])
	}

	clearPattern(source.Gray(), patternImage.Gray(), 2, 2)
	vanish, err := observer.ObserveVanish(source, region, pattern, ObserveOptions{Interval: 1 * time.Millisecond, Timeout: 0})
	if err != nil {
		t.Fatalf("observe vanish after clear failed: %v", err)
	}
	if len(vanish) != 1 || vanish[0].Type != ObserveEventVanish {
		t.Fatalf("expected one vanish event after clearing pattern, got=%+v", vanish)
	}

	if err := appController.Focus(appName, AppOptions{}); err != nil {
		t.Fatalf("app focus failed: %v", err)
	}
	windows, err := appController.ListWindows(appName, AppOptions{})
	if err != nil {
		t.Fatalf("list windows failed: %v", err)
	}
	if len(windows) != 1 || !windows[0].Focused {
		t.Fatalf("expected focused window after focus call, got=%+v", windows)
	}

	if err := appController.Close(appName, AppOptions{}); err != nil {
		t.Fatalf("app close failed: %v", err)
	}
	running, err = appController.IsRunning(appName, AppOptions{})
	if err != nil {
		t.Fatalf("is-running after close failed: %v", err)
	}
	if running {
		t.Fatalf("expected running=false after close")
	}
}

func TestCrossProtocolIntegrationFlow(t *testing.T) {
	runCrossProtocolIntegrationFlow(t, "DemoApp")
}

func clearPattern(dst, pattern *image.Gray, atX, atY int) {
	if dst == nil || pattern == nil {
		return
	}
	pb := pattern.Bounds()
	for y := 0; y < pb.Dy(); y++ {
		for x := 0; x < pb.Dx(); x++ {
			dst.SetGray(atX+x, atY+y, color.Gray{Y: 0})
		}
	}
}
