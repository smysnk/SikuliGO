package sikuli

import (
	"errors"
	"testing"
	"time"

	"github.com/sikulix/portgo/internal/core"
)

func TestRegionDefaultsAndSetters(t *testing.T) {
	r := NewRegion(10, 20, 30, 40)
	if !r.ThrowException {
		t.Fatalf("ThrowException default should be true")
	}
	if r.AutoWaitTimeout != DefaultAutoWaitTimeout {
		t.Fatalf("AutoWaitTimeout default mismatch: got=%v", r.AutoWaitTimeout)
	}
	if r.WaitScanRate != DefaultWaitScanRate {
		t.Fatalf("WaitScanRate default mismatch: got=%v", r.WaitScanRate)
	}
	if r.ObserveScanRate != DefaultObserveScanRate {
		t.Fatalf("ObserveScanRate default mismatch: got=%v", r.ObserveScanRate)
	}

	r.SetThrowException(false)
	if r.ThrowException {
		t.Fatalf("SetThrowException(false) failed")
	}
	r.ResetThrowException()
	if !r.ThrowException {
		t.Fatalf("ResetThrowException() failed")
	}

	r.SetAutoWaitTimeout(-1)
	if r.AutoWaitTimeout != 0 {
		t.Fatalf("SetAutoWaitTimeout clamp failed: got=%v", r.AutoWaitTimeout)
	}
	r.SetWaitScanRate(0)
	if r.WaitScanRate != DefaultWaitScanRate {
		t.Fatalf("SetWaitScanRate default fallback failed: got=%v", r.WaitScanRate)
	}
	r.SetObserveScanRate(-2)
	if r.ObserveScanRate != DefaultObserveScanRate {
		t.Fatalf("SetObserveScanRate default fallback failed: got=%v", r.ObserveScanRate)
	}
}

func TestRegionGeometryHelpers(t *testing.T) {
	r := NewRegion(10, 20, 30, 40)
	c := r.Center()
	if c.X != 25 || c.Y != 40 {
		t.Fatalf("center mismatch: got=(%d,%d)", c.X, c.Y)
	}

	grown := r.Grow(5, 10)
	if grown.X != 5 || grown.Y != 10 || grown.W != 40 || grown.H != 60 {
		t.Fatalf("grow mismatch: got=%+v", grown)
	}

	offset := r.Offset(-3, 4)
	if offset.X != 7 || offset.Y != 24 || offset.W != 30 || offset.H != 40 {
		t.Fatalf("offset mismatch: got=%+v", offset)
	}

	moved := r.MoveTo(1, 2)
	if moved.X != 1 || moved.Y != 2 || moved.W != 30 || moved.H != 40 {
		t.Fatalf("move mismatch: got=%+v", moved)
	}

	size := r.SetSize(-1, 7)
	if size.W != 0 || size.H != 7 {
		t.Fatalf("set size clamp mismatch: got=%+v", size)
	}

	if !r.Contains(NewPoint(10, 20)) {
		t.Fatalf("expected top-left point to be contained")
	}
	if r.Contains(NewPoint(40, 60)) {
		t.Fatalf("expected right/bottom edge point to be excluded")
	}

	inside := NewRegion(12, 22, 5, 6)
	if !r.ContainsRegion(inside) {
		t.Fatalf("expected region to contain nested region")
	}
	outside := NewRegion(0, 0, 5, 5)
	if r.ContainsRegion(outside) {
		t.Fatalf("expected region not to contain outside region")
	}

	other := NewRegion(25, 35, 20, 20)
	union := r.Union(other)
	if union.X != 10 || union.Y != 20 || union.W != 35 || union.H != 40 {
		t.Fatalf("union mismatch: got=%+v", union)
	}

	inter := r.Intersection(other)
	if inter.X != 25 || inter.Y != 35 || inter.W != 15 || inter.H != 20 {
		t.Fatalf("intersection mismatch: got=%+v", inter)
	}

	noOverlap := r.Intersection(NewRegion(200, 200, 10, 10))
	if noOverlap.W != 0 || noOverlap.H != 0 {
		t.Fatalf("expected empty intersection: got=%+v", noOverlap)
	}
}

func TestPatternNormalization(t *testing.T) {
	img, err := NewImageFromMatrix("needle", [][]uint8{
		{1, 2},
		{3, 4},
	})
	if err != nil {
		t.Fatalf("new image: %v", err)
	}
	p, err := NewPattern(img)
	if err != nil {
		t.Fatalf("new pattern: %v", err)
	}

	p.Similar(2)
	if p.Similarity() != 1 {
		t.Fatalf("similarity upper clamp failed: %v", p.Similarity())
	}
	p.Similar(-0.5)
	if p.Similarity() != 0 {
		t.Fatalf("similarity lower clamp failed: %v", p.Similarity())
	}

	p.Resize(0)
	if p.ResizeFactor() != 1 {
		t.Fatalf("resize fallback failed: %v", p.ResizeFactor())
	}
}

func TestFinderExistsAndHas(t *testing.T) {
	hay, err := NewImageFromMatrix("hay", [][]uint8{
		{10, 10, 10, 10, 10},
		{10, 0, 255, 10, 10},
		{10, 255, 0, 10, 10},
		{10, 10, 10, 10, 10},
	})
	if err != nil {
		t.Fatalf("new hay image: %v", err)
	}
	needle, err := NewImageFromMatrix("needle", [][]uint8{
		{0, 255},
		{255, 0},
	})
	if err != nil {
		t.Fatalf("new needle image: %v", err)
	}
	p, err := NewPattern(needle)
	if err != nil {
		t.Fatalf("new pattern: %v", err)
	}
	p.Exact()

	f, err := NewFinder(hay)
	if err != nil {
		t.Fatalf("new finder: %v", err)
	}
	m, ok, err := f.Exists(p)
	if err != nil {
		t.Fatalf("exists should not fail: %v", err)
	}
	if !ok {
		t.Fatalf("expected Exists to report match")
	}
	if m.X != 1 || m.Y != 1 {
		t.Fatalf("exists match coordinates mismatch: got=(%d,%d)", m.X, m.Y)
	}
	has, err := f.Has(p)
	if err != nil {
		t.Fatalf("has should not fail: %v", err)
	}
	if !has {
		t.Fatalf("expected Has to be true")
	}

	missingNeedle, err := NewImageFromMatrix("missing", [][]uint8{
		{4, 4},
		{4, 4},
	})
	if err != nil {
		t.Fatalf("new missing needle: %v", err)
	}
	missingPattern, err := NewPattern(missingNeedle)
	if err != nil {
		t.Fatalf("new missing pattern: %v", err)
	}
	missingPattern.Exact()
	_, ok, err = f.Exists(missingPattern)
	if err != nil {
		t.Fatalf("exists missing should not hard fail: %v", err)
	}
	if ok {
		t.Fatalf("expected missing pattern not to exist")
	}

	_, err = f.Wait(missingPattern, 10*time.Millisecond)
	if !errors.Is(err, ErrTimeout) {
		t.Fatalf("wait timeout mismatch: got=%v", err)
	}
	vanished, err := f.WaitVanish(missingPattern, 10*time.Millisecond)
	if err != nil {
		t.Fatalf("wait vanish should not fail: %v", err)
	}
	if !vanished {
		t.Fatalf("expected missing pattern to be vanished")
	}
	vanished, err = f.WaitVanish(p, 10*time.Millisecond)
	if err != nil {
		t.Fatalf("wait vanish present should not fail: %v", err)
	}
	if vanished {
		t.Fatalf("expected present pattern not to vanish")
	}
}

func TestRuntimeSettingsLifecycle(t *testing.T) {
	ResetSettings()
	s := GetSettings()
	if s.MinSimilarity != DefaultSimilarity {
		t.Fatalf("default MinSimilarity mismatch: %v", s.MinSimilarity)
	}

	s = UpdateSettings(func(in *RuntimeSettings) {
		in.MinSimilarity = 0.9
		in.ShowActions = true
	})
	if s.MinSimilarity != 0.9 || !s.ShowActions {
		t.Fatalf("update did not apply: %+v", s)
	}

	s = ResetSettings()
	if s.MinSimilarity != DefaultSimilarity || s.ShowActions {
		t.Fatalf("reset failed: %+v", s)
	}
}

func TestOptionsLifecycle(t *testing.T) {
	o := NewOptions()
	if o.Has("alpha") {
		t.Fatalf("expected empty options")
	}
	if got := o.GetString("missing", "default"); got != "default" {
		t.Fatalf("default string mismatch: %q", got)
	}
	o.SetString("alpha", "one")
	o.SetInt("n", 7)
	o.SetFloat64("ratio", 1.25)
	o.SetBool("enabled", true)

	if !o.Has("alpha") || o.GetString("alpha", "") != "one" {
		t.Fatalf("string set/get mismatch")
	}
	if o.GetInt("n", 0) != 7 {
		t.Fatalf("int set/get mismatch")
	}
	if o.GetFloat64("ratio", 0) != 1.25 {
		t.Fatalf("float set/get mismatch")
	}
	if !o.GetBool("enabled", false) {
		t.Fatalf("bool set/get mismatch")
	}
	if o.GetInt("bad", 9) != 9 {
		t.Fatalf("int default fallback mismatch")
	}

	clone := o.Clone()
	clone.SetString("alpha", "two")
	if o.GetString("alpha", "") != "one" {
		t.Fatalf("clone should not mutate original")
	}

	merged := NewOptionsFromMap(map[string]string{"x": "1"})
	merged.Merge(o)
	if merged.GetString("alpha", "") != "one" || merged.GetString("x", "") != "1" {
		t.Fatalf("merge mismatch")
	}
	merged.Delete("x")
	if merged.Has("x") {
		t.Fatalf("delete mismatch")
	}
}

func TestImageCrop(t *testing.T) {
	img, err := NewImageFromMatrix("src", [][]uint8{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	})
	if err != nil {
		t.Fatalf("new image: %v", err)
	}
	crop, err := img.Crop(NewRect(1, 1, 2, 2))
	if err != nil {
		t.Fatalf("crop failed: %v", err)
	}
	if crop.Width() != 2 || crop.Height() != 2 {
		t.Fatalf("crop dimensions mismatch: got=%dx%d", crop.Width(), crop.Height())
	}
	if crop.Gray().GrayAt(1, 1).Y != 6 || crop.Gray().GrayAt(2, 2).Y != 11 {
		t.Fatalf("crop pixel mismatch")
	}
	_, err = img.Crop(NewRect(100, 100, 3, 3))
	if err == nil {
		t.Fatalf("expected crop outside bounds to fail")
	}
}

func TestRegionFindExistsWaitParityScaffold(t *testing.T) {
	hay, err := NewImageFromMatrix("hay", [][]uint8{
		{10, 10, 10, 10, 10, 10},
		{10, 10, 0, 255, 10, 10},
		{10, 10, 255, 0, 10, 10},
		{10, 10, 10, 10, 10, 10},
	})
	if err != nil {
		t.Fatalf("new hay image: %v", err)
	}
	needle, err := NewImageFromMatrix("needle", [][]uint8{
		{0, 255},
		{255, 0},
	})
	if err != nil {
		t.Fatalf("new needle image: %v", err)
	}
	p, err := NewPattern(needle)
	if err != nil {
		t.Fatalf("new pattern: %v", err)
	}
	p.Exact()

	region := NewRegion(1, 0, 4, 4)
	m, err := region.Find(hay, p)
	if err != nil {
		t.Fatalf("region find failed: %v", err)
	}
	if m.X != 2 || m.Y != 1 {
		t.Fatalf("region find coordinates mismatch: got=(%d,%d)", m.X, m.Y)
	}

	existsMatch, ok, err := region.Exists(hay, p, 0)
	if err != nil {
		t.Fatalf("region exists failed: %v", err)
	}
	if !ok || existsMatch.X != 2 || existsMatch.Y != 1 {
		t.Fatalf("region exists mismatch: ok=%v match=%+v", ok, existsMatch)
	}

	has, err := region.Has(hay, p, 0)
	if err != nil {
		t.Fatalf("region has failed: %v", err)
	}
	if !has {
		t.Fatalf("region has should be true")
	}

	missingRegion := NewRegion(0, 0, 2, 2)
	_, ok, err = missingRegion.Exists(hay, p, 20*time.Millisecond)
	if err != nil {
		t.Fatalf("missing exists should not error: %v", err)
	}
	if ok {
		t.Fatalf("missing region should not contain target")
	}

	_, err = missingRegion.Wait(hay, p, 20*time.Millisecond)
	if !errors.Is(err, ErrTimeout) {
		t.Fatalf("wait timeout mismatch: got=%v", err)
	}

	vanished, err := missingRegion.WaitVanish(hay, p, 20*time.Millisecond)
	if err != nil {
		t.Fatalf("wait vanish missing should not fail: %v", err)
	}
	if !vanished {
		t.Fatalf("expected vanish=true for missing target")
	}

	vanished, err = region.WaitVanish(hay, p, 20*time.Millisecond)
	if err != nil {
		t.Fatalf("wait vanish present should not fail: %v", err)
	}
	if vanished {
		t.Fatalf("expected vanish=false for present target")
	}

	all, err := region.FindAll(hay, p)
	if err != nil {
		t.Fatalf("region findall failed: %v", err)
	}
	if len(all) != 1 || all[0].X != 2 || all[0].Y != 1 {
		t.Fatalf("region findall mismatch: %+v", all)
	}
	byRow, err := region.FindAllByRow(hay, p)
	if err != nil {
		t.Fatalf("region findall by row failed: %v", err)
	}
	byCol, err := region.FindAllByColumn(hay, p)
	if err != nil {
		t.Fatalf("region findall by column failed: %v", err)
	}
	if len(byRow) != 1 || len(byCol) != 1 {
		t.Fatalf("region sorted findall mismatch: row=%d col=%d", len(byRow), len(byCol))
	}
}

func TestFinderFindAllSortedHelpers(t *testing.T) {
	hay, err := NewImageFromMatrix("hay", [][]uint8{
		{10, 10, 10, 10, 10, 10, 10, 10},
		{10, 0, 255, 10, 10, 10, 10, 10},
		{10, 255, 0, 10, 0, 255, 10, 10},
		{10, 10, 10, 10, 255, 0, 10, 10},
		{10, 10, 10, 10, 10, 10, 10, 10},
	})
	if err != nil {
		t.Fatalf("new hay image: %v", err)
	}
	needle, err := NewImageFromMatrix("needle", [][]uint8{
		{0, 255},
		{255, 0},
	})
	if err != nil {
		t.Fatalf("new needle image: %v", err)
	}
	p, err := NewPattern(needle)
	if err != nil {
		t.Fatalf("new pattern: %v", err)
	}
	p.Exact()

	f, err := NewFinder(hay)
	if err != nil {
		t.Fatalf("new finder: %v", err)
	}
	row, err := f.FindAllByRow(p)
	if err != nil {
		t.Fatalf("findall by row failed: %v", err)
	}
	col, err := f.FindAllByColumn(p)
	if err != nil {
		t.Fatalf("findall by column failed: %v", err)
	}
	if len(row) != 2 || len(col) != 2 {
		t.Fatalf("expected 2 matches each, row=%d col=%d", len(row), len(col))
	}
	if row[0].Index != 0 || row[1].Index != 1 {
		t.Fatalf("row reindex mismatch: %+v", row)
	}
	if col[0].Index != 0 || col[1].Index != 1 {
		t.Fatalf("col reindex mismatch: %+v", col)
	}
}

func TestLocationAndOffsetBasics(t *testing.T) {
	l := NewLocation(10, 20)
	if l.X != 10 || l.Y != 20 {
		t.Fatalf("location mismatch: %+v", l)
	}
	lp := l.ToPoint()
	if lp.X != 10 || lp.Y != 20 {
		t.Fatalf("location to point mismatch: %+v", lp)
	}
	moved := l.Move(-3, 4)
	if moved.X != 7 || moved.Y != 24 {
		t.Fatalf("location move mismatch: %+v", moved)
	}

	o := NewOffset(5, -2)
	op := o.ToPoint()
	if op.X != 5 || op.Y != -2 {
		t.Fatalf("offset to point mismatch: %+v", op)
	}
}

type stubOCRBackend struct {
	result  core.OCRResult
	err     error
	lastReq core.OCRRequest
}

func (s *stubOCRBackend) Read(req core.OCRRequest) (core.OCRResult, error) {
	s.lastReq = req
	if s.err != nil {
		return core.OCRResult{}, s.err
	}
	return s.result, nil
}

type stubInputBackend struct {
	err      error
	requests []core.InputRequest
}

func (s *stubInputBackend) Execute(req core.InputRequest) error {
	s.requests = append(s.requests, req)
	if s.err != nil {
		return s.err
	}
	return nil
}

type stubObserveBackend struct {
	err      error
	events   []core.ObserveEvent
	requests []core.ObserveRequest
}

func (s *stubObserveBackend) Observe(req core.ObserveRequest) ([]core.ObserveEvent, error) {
	s.requests = append(s.requests, req)
	if s.err != nil {
		return nil, s.err
	}
	return s.events, nil
}

type stubAppBackend struct {
	err      error
	result   core.AppResult
	requests []core.AppRequest
}

func (s *stubAppBackend) Execute(req core.AppRequest) (core.AppResult, error) {
	s.requests = append(s.requests, req)
	if s.err != nil {
		return core.AppResult{}, s.err
	}
	return s.result, nil
}

func TestFinderOCRUnsupportedByDefault(t *testing.T) {
	img, err := NewImageFromMatrix("ocr-src", [][]uint8{
		{255, 255},
		{255, 255},
	})
	if err != nil {
		t.Fatalf("new image: %v", err)
	}
	f, err := NewFinder(img)
	if err != nil {
		t.Fatalf("new finder: %v", err)
	}
	_, err = f.ReadText(OCRParams{})
	if !errors.Is(err, ErrBackendUnsupported) {
		t.Fatalf("expected ErrBackendUnsupported, got=%v", err)
	}
}

func TestFinderOCRWithStubBackend(t *testing.T) {
	img, err := NewImageFromMatrix("ocr-src", [][]uint8{
		{1, 1, 1, 1},
		{1, 1, 1, 1},
		{1, 1, 1, 1},
		{1, 1, 1, 1},
	})
	if err != nil {
		t.Fatalf("new image: %v", err)
	}
	f, err := NewFinder(img)
	if err != nil {
		t.Fatalf("new finder: %v", err)
	}

	stub := &stubOCRBackend{
		result: core.OCRResult{
			Text: "Sikuli Go OCR",
			Words: []core.OCRWord{
				{Text: "OCR", X: 2, Y: 0, W: 1, H: 1, Confidence: 0.80},
				{Text: "Sikuli", X: 0, Y: 0, W: 2, H: 1, Confidence: 0.95},
				{Text: "Go", X: 0, Y: 1, W: 2, H: 1, Confidence: 0.90},
			},
		},
	}
	f.SetOCRBackend(stub)

	text, err := f.ReadText(OCRParams{
		MinConfidence: -2,
		Timeout:       -1,
	})
	if err != nil {
		t.Fatalf("read text failed: %v", err)
	}
	if text != "Sikuli Go OCR" {
		t.Fatalf("read text mismatch: %q", text)
	}
	if stub.lastReq.Language != DefaultOCRLanguage {
		t.Fatalf("default language mismatch: %q", stub.lastReq.Language)
	}
	if stub.lastReq.MinConfidence != 0 {
		t.Fatalf("min confidence clamp mismatch: %v", stub.lastReq.MinConfidence)
	}
	if stub.lastReq.Timeout != 0 {
		t.Fatalf("timeout clamp mismatch: %v", stub.lastReq.Timeout)
	}

	matches, err := f.FindText("go", OCRParams{})
	if err != nil {
		t.Fatalf("find text failed: %v", err)
	}
	if len(matches) != 1 {
		t.Fatalf("expected one match, got=%d", len(matches))
	}
	if matches[0].Text != "Go" || matches[0].X != 0 || matches[0].Y != 1 {
		t.Fatalf("match mismatch: %+v", matches[0])
	}

	matches, err = f.FindText("Sikuli", OCRParams{})
	if err != nil {
		t.Fatalf("find text case-insensitive failed: %v", err)
	}
	if len(matches) != 1 || matches[0].Index != 0 {
		t.Fatalf("expected one indexed match, got=%+v", matches)
	}

	_, err = f.FindText("sikuli", OCRParams{CaseSensitive: true})
	if !errors.Is(err, ErrFindFailed) {
		t.Fatalf("expected ErrFindFailed for case-sensitive miss, got=%v", err)
	}
	_, err = f.FindText("", OCRParams{})
	if !errors.Is(err, ErrInvalidTarget) {
		t.Fatalf("expected ErrInvalidTarget for empty query, got=%v", err)
	}
}

func TestRegionOCRMethods(t *testing.T) {
	src, err := NewImageFromMatrix("region-ocr", [][]uint8{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	})
	if err != nil {
		t.Fatalf("new source: %v", err)
	}
	region := NewRegion(1, 1, 2, 2)

	stub := &stubOCRBackend{
		result: core.OCRResult{
			Text: "Zone",
			Words: []core.OCRWord{
				{Text: "Zone", X: 1, Y: 1, W: 2, H: 1, Confidence: 0.9},
			},
		},
	}

	prevFactory := newOCRBackend
	newOCRBackend = func() core.OCR { return stub }
	defer func() {
		newOCRBackend = prevFactory
	}()

	text, err := region.ReadText(src, OCRParams{})
	if err != nil {
		t.Fatalf("region read text failed: %v", err)
	}
	if text != "Zone" {
		t.Fatalf("region read text mismatch: %q", text)
	}

	matches, err := region.FindText(src, "zone", OCRParams{})
	if err != nil {
		t.Fatalf("region find text failed: %v", err)
	}
	if len(matches) != 1 || matches[0].Text != "Zone" {
		t.Fatalf("region find text mismatch: %+v", matches)
	}
}

func TestInputControllerUnsupportedByDefault(t *testing.T) {
	c := NewInputController()
	err := c.Click(10, 20, InputOptions{})
	if !errors.Is(err, ErrBackendUnsupported) {
		t.Fatalf("expected ErrBackendUnsupported, got=%v", err)
	}
}

func TestInputControllerDispatchWithStub(t *testing.T) {
	c := NewInputController()
	stub := &stubInputBackend{}
	c.SetBackend(stub)

	if err := c.MoveMouse(10, 20, InputOptions{Delay: -10 * time.Millisecond}); err != nil {
		t.Fatalf("move mouse failed: %v", err)
	}
	if err := c.Click(30, 40, InputOptions{Button: MouseButtonRight}); err != nil {
		t.Fatalf("click failed: %v", err)
	}
	if err := c.TypeText("  hello  ", InputOptions{}); err != nil {
		t.Fatalf("type text failed: %v", err)
	}
	if err := c.Hotkey("CMD", "SHIFT", "P"); err != nil {
		t.Fatalf("hotkey failed: %v", err)
	}

	if len(stub.requests) != 4 {
		t.Fatalf("expected 4 requests, got=%d", len(stub.requests))
	}
	if stub.requests[0].Action != core.InputActionMouseMove || stub.requests[0].Delay != 0 {
		t.Fatalf("move request mismatch: %+v", stub.requests[0])
	}
	if stub.requests[1].Action != core.InputActionClick || stub.requests[1].Button != string(MouseButtonRight) {
		t.Fatalf("click request mismatch: %+v", stub.requests[1])
	}
	if stub.requests[2].Action != core.InputActionTypeText || stub.requests[2].Text != "hello" {
		t.Fatalf("type request mismatch: %+v", stub.requests[2])
	}
	if stub.requests[3].Action != core.InputActionHotkey || len(stub.requests[3].Keys) != 3 {
		t.Fatalf("hotkey request mismatch: %+v", stub.requests[3])
	}
}

func TestInputControllerValidation(t *testing.T) {
	c := NewInputController()
	stub := &stubInputBackend{}
	c.SetBackend(stub)

	if err := c.TypeText("   ", InputOptions{}); !errors.Is(err, ErrInvalidTarget) {
		t.Fatalf("expected ErrInvalidTarget for empty type text, got=%v", err)
	}
	if err := c.Hotkey(); !errors.Is(err, ErrInvalidTarget) {
		t.Fatalf("expected ErrInvalidTarget for empty hotkey, got=%v", err)
	}

	stub.err = errors.New("custom backend error")
	if err := c.Click(1, 2, InputOptions{}); err == nil || errors.Is(err, ErrInvalidTarget) {
		t.Fatalf("expected raw backend error, got=%v", err)
	}
}

func TestObserverControllerUnsupportedByDefault(t *testing.T) {
	source, err := NewImageFromMatrix("obs-src", [][]uint8{
		{1, 1, 1},
		{1, 1, 1},
		{1, 1, 1},
	})
	if err != nil {
		t.Fatalf("new source: %v", err)
	}
	patternImage, err := NewImageFromMatrix("obs-needle", [][]uint8{{1}})
	if err != nil {
		t.Fatalf("new pattern image: %v", err)
	}
	pattern, err := NewPattern(patternImage)
	if err != nil {
		t.Fatalf("new pattern: %v", err)
	}
	observer := NewObserverController()
	_, err = observer.ObserveAppear(source, NewRegion(0, 0, 3, 3), pattern, ObserveOptions{})
	if !errors.Is(err, ErrBackendUnsupported) {
		t.Fatalf("expected ErrBackendUnsupported, got=%v", err)
	}
}

func TestObserverControllerDispatchWithStub(t *testing.T) {
	source, err := NewImageFromMatrix("obs-src", [][]uint8{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	})
	if err != nil {
		t.Fatalf("new source: %v", err)
	}
	patternImage, err := NewImageFromMatrix("obs-needle", [][]uint8{{6}})
	if err != nil {
		t.Fatalf("new pattern image: %v", err)
	}
	pattern, err := NewPattern(patternImage)
	if err != nil {
		t.Fatalf("new pattern: %v", err)
	}

	stub := &stubObserveBackend{
		events: []core.ObserveEvent{
			{
				Event:     core.ObserveEventAppear,
				X:         1,
				Y:         1,
				W:         1,
				H:         1,
				Score:     0.9,
				Timestamp: time.Now(),
			},
		},
	}
	observer := NewObserverController()
	observer.SetBackend(stub)

	events, err := observer.ObserveAppear(source, NewRegion(0, 0, 4, 4), pattern, ObserveOptions{
		Interval: -1,
		Timeout:  -1,
	})
	if err != nil {
		t.Fatalf("observe appear failed: %v", err)
	}
	if len(events) != 1 || events[0].Type != ObserveEventAppear {
		t.Fatalf("observe events mismatch: %+v", events)
	}
	if len(stub.requests) != 1 {
		t.Fatalf("expected one backend request, got=%d", len(stub.requests))
	}
	req := stub.requests[0]
	if req.Event != core.ObserveEventAppear {
		t.Fatalf("expected appear event, got=%q", req.Event)
	}
	if req.Interval <= 0 {
		t.Fatalf("expected normalized interval > 0, got=%v", req.Interval)
	}
	if req.Timeout != 0 {
		t.Fatalf("expected normalized timeout 0, got=%v", req.Timeout)
	}
	if req.Region.Dx() != 4 || req.Region.Dy() != 4 {
		t.Fatalf("region bounds mismatch: %+v", req.Region)
	}

	_, err = observer.ObserveChange(source, NewRegion(0, 0, 4, 4), ObserveOptions{})
	if err != nil {
		t.Fatalf("observe change failed: %v", err)
	}
}

func TestObserverControllerValidation(t *testing.T) {
	source, err := NewImageFromMatrix("obs-src", [][]uint8{
		{1, 1},
		{1, 1},
	})
	if err != nil {
		t.Fatalf("new source: %v", err)
	}
	patternImage, err := NewImageFromMatrix("obs-needle", [][]uint8{{1}})
	if err != nil {
		t.Fatalf("new pattern image: %v", err)
	}
	pattern, err := NewPattern(patternImage)
	if err != nil {
		t.Fatalf("new pattern: %v", err)
	}

	observer := NewObserverController()
	stub := &stubObserveBackend{}
	observer.SetBackend(stub)

	_, err = observer.ObserveAppear(nil, NewRegion(0, 0, 2, 2), pattern, ObserveOptions{})
	if !errors.Is(err, ErrInvalidTarget) {
		t.Fatalf("expected ErrInvalidTarget for nil source, got=%v", err)
	}
	_, err = observer.ObserveAppear(source, NewRegion(0, 0, 2, 2), nil, ObserveOptions{})
	if !errors.Is(err, ErrInvalidTarget) {
		t.Fatalf("expected ErrInvalidTarget for nil pattern, got=%v", err)
	}
	_, err = observer.ObserveChange(source, NewRegion(0, 0, 0, 0), ObserveOptions{})
	if !errors.Is(err, ErrInvalidTarget) {
		t.Fatalf("expected ErrInvalidTarget for empty region, got=%v", err)
	}

	stub.err = core.ErrObserveUnsupported
	_, err = observer.ObserveChange(source, NewRegion(0, 0, 2, 2), ObserveOptions{})
	if !errors.Is(err, ErrBackendUnsupported) {
		t.Fatalf("expected ErrBackendUnsupported for unsupported observe backend, got=%v", err)
	}
}

func TestAppControllerUnsupportedByDefault(t *testing.T) {
	controller := NewAppController()
	err := controller.Open("Demo", nil, AppOptions{})
	if !errors.Is(err, ErrBackendUnsupported) {
		t.Fatalf("expected ErrBackendUnsupported, got=%v", err)
	}
}

func TestAppControllerDispatchWithStub(t *testing.T) {
	stub := &stubAppBackend{
		result: core.AppResult{
			Running: true,
			PID:     42,
			Windows: []core.WindowInfo{
				{Title: "Demo", X: 1, Y: 2, W: 300, H: 200, Focused: true},
			},
		},
	}
	controller := NewAppController()
	controller.SetBackend(stub)

	if err := controller.Open("Demo", []string{"--flag"}, AppOptions{Timeout: -time.Second}); err != nil {
		t.Fatalf("open failed: %v", err)
	}
	if err := controller.Focus("Demo", AppOptions{}); err != nil {
		t.Fatalf("focus failed: %v", err)
	}
	if err := controller.Close("Demo", AppOptions{}); err != nil {
		t.Fatalf("close failed: %v", err)
	}
	running, err := controller.IsRunning("Demo", AppOptions{})
	if err != nil {
		t.Fatalf("is running failed: %v", err)
	}
	if !running {
		t.Fatalf("expected running=true")
	}
	windows, err := controller.ListWindows("Demo", AppOptions{})
	if err != nil {
		t.Fatalf("list windows failed: %v", err)
	}
	if len(windows) != 1 || windows[0].Title != "Demo" || !windows[0].Focused {
		t.Fatalf("windows mismatch: %+v", windows)
	}
	if len(stub.requests) != 5 {
		t.Fatalf("expected 5 backend requests, got=%d", len(stub.requests))
	}
	if stub.requests[0].Action != core.AppActionOpen || stub.requests[0].Timeout != 0 || len(stub.requests[0].Args) != 1 {
		t.Fatalf("open request mismatch: %+v", stub.requests[0])
	}
	if stub.requests[1].Action != core.AppActionFocus {
		t.Fatalf("focus request mismatch: %+v", stub.requests[1])
	}
	if stub.requests[2].Action != core.AppActionClose {
		t.Fatalf("close request mismatch: %+v", stub.requests[2])
	}
	if stub.requests[3].Action != core.AppActionIsRunning {
		t.Fatalf("is running request mismatch: %+v", stub.requests[3])
	}
	if stub.requests[4].Action != core.AppActionListWindow {
		t.Fatalf("list windows request mismatch: %+v", stub.requests[4])
	}
}

func TestAppControllerValidation(t *testing.T) {
	controller := NewAppController()
	stub := &stubAppBackend{}
	controller.SetBackend(stub)

	if err := controller.Open("   ", nil, AppOptions{}); !errors.Is(err, ErrInvalidTarget) {
		t.Fatalf("expected ErrInvalidTarget for empty app name, got=%v", err)
	}

	stub.err = errors.New("unsupported app action \"bad\"")
	if err := controller.Focus("Demo", AppOptions{}); !errors.Is(err, ErrInvalidTarget) {
		t.Fatalf("expected ErrInvalidTarget for validation-like backend error, got=%v", err)
	}

	stub.err = errors.New("custom backend error")
	if err := controller.Close("Demo", AppOptions{}); err == nil || errors.Is(err, ErrInvalidTarget) {
		t.Fatalf("expected raw backend error, got=%v", err)
	}
}
