package sikuli

import (
	"image"
	"time"
)

// This file intentionally defines stable workstream-1 public signatures.
// If these interfaces are changed, update the generated API reference docs.

// ImageAPI describes immutable image primitives used by matching and OCR.
// This aligns with the SikuliX notion of image snapshots used by Region/Finder.
type ImageAPI interface {
	Name() string
	Width() int
	Height() int
	Gray() *image.Gray
	Clone() *Image
	Crop(rect Rect) (*Image, error)
}

// TargetPointProvider describes values that can resolve to a concrete action target.
type TargetPointProvider interface {
	TargetPoint() Point
}

// PatternAPI configures how a target image should be matched on screen.
// It mirrors SikuliX Pattern behavior such as similar(), exact(), and targetOffset().
type PatternAPI interface {
	Image() *Image
	Similar(sim float64) *Pattern
	Similarity() float64
	Exact() *Pattern
	TargetOffset(dx, dy int) *Pattern
	Offset() Point
	Resize(factor float64) *Pattern
	ResizeFactor() float64
	Mask() *image.Gray
}

// FinderAPI performs match/OCR operations against a source image.
// Miss and timeout handling follows the canonical sikuli-go parity contract:
// Find returns ErrFindFailed on a miss, Exists returns (Match{}, false, nil),
// Wait returns ErrTimeout after the wait budget is exhausted, and WaitVanish
// returns false,nil on timeout.
// The additive compatibility iterator surface provides SikuliX-style traversal
// without removing the existing value/slice-oriented methods.
type FinderAPI interface {
	Find(pattern *Pattern) (Match, error)
	FindAll(pattern *Pattern) ([]Match, error)
	FindAllByRow(pattern *Pattern) ([]Match, error)
	FindAllByColumn(pattern *Pattern) ([]Match, error)
	FindAnyList(patterns []*Pattern) ([]Match, error)
	FindBestList(patterns []*Pattern) (Match, error)
	Iterate(pattern *Pattern) error
	IterateAll(pattern *Pattern) error
	Exists(pattern *Pattern) (Match, bool, error)
	Has(pattern *Pattern) (bool, error)
	Wait(pattern *Pattern, timeout time.Duration) (Match, error)
	WaitAnyList(patterns []*Pattern, timeout time.Duration) ([]Match, error)
	WaitBestList(patterns []*Pattern, timeout time.Duration) (Match, error)
	WaitVanish(pattern *Pattern, timeout time.Duration) (bool, error)
	CollectWords(params OCRParams) ([]OCRWord, error)
	CollectLines(params OCRParams) ([]OCRLine, error)
	HasNext() bool
	Next() (Match, bool)
	Reset()
	Destroy()
	ReadText(params OCRParams) (string, error)
	FindText(query string, params OCRParams) ([]TextMatch, error)
	LastMatches() []Match
}

// RegionAPI defines region geometry and region-scoped automation operations.
// It maps to familiar SikuliX Region methods while keeping the same explicit
// Go miss/timeout contract used by FinderAPI.
type RegionAPI interface {
	Center() Point
	Grow(dx, dy int) Region
	Offset(dx, dy int) Region
	MoveTo(x, y int) Region
	SetSize(w, h int) Region
	Contains(p Point) bool
	ContainsRegion(other Region) bool
	Union(other Region) Region
	Intersection(other Region) Region
	Find(source *Image, pattern *Pattern) (Match, error)
	Exists(source *Image, pattern *Pattern, timeout time.Duration) (Match, bool, error)
	Has(source *Image, pattern *Pattern, timeout time.Duration) (bool, error)
	Wait(source *Image, pattern *Pattern, timeout time.Duration) (Match, error)
	WaitVanish(source *Image, pattern *Pattern, timeout time.Duration) (bool, error)
	FindAll(source *Image, pattern *Pattern) ([]Match, error)
	FindAllByRow(source *Image, pattern *Pattern) ([]Match, error)
	FindAllByColumn(source *Image, pattern *Pattern) ([]Match, error)
	FindAnyList(source *Image, patterns []*Pattern) ([]Match, error)
	FindBestList(source *Image, patterns []*Pattern) (Match, error)
	WaitAnyList(source *Image, patterns []*Pattern, timeout time.Duration) ([]Match, error)
	WaitBestList(source *Image, patterns []*Pattern, timeout time.Duration) (Match, error)
	ReadText(source *Image, params OCRParams) (string, error)
	FindText(source *Image, query string, params OCRParams) ([]TextMatch, error)
	CollectWords(source *Image, params OCRParams) ([]OCRWord, error)
	CollectLines(source *Image, params OCRParams) ([]OCRLine, error)
}

// LiveRegionAPI defines screen-backed region operations that route through the API runtime.
type LiveRegionAPI interface {
	Bounds() Region
	Center() Point
	TargetPoint() Point
	Grow(dx, dy int) LiveRegion
	Offset(dx, dy int) LiveRegion
	MoveTo(x, y int) LiveRegion
	SetSize(w, h int) LiveRegion
	WithMatcherEngine(engine MatcherEngine) LiveRegion
	Capture() (*Image, error)
	Find(pattern *Pattern) (Match, error)
	FindAnyList(patterns []*Pattern) ([]Match, error)
	FindBestList(patterns []*Pattern) (Match, error)
	Exists(pattern *Pattern, timeout time.Duration) (Match, bool, error)
	Has(pattern *Pattern, timeout time.Duration) (bool, error)
	Wait(pattern *Pattern, timeout time.Duration) (Match, error)
	WaitAnyList(patterns []*Pattern, timeout time.Duration) ([]Match, error)
	WaitBestList(patterns []*Pattern, timeout time.Duration) (Match, error)
	WaitVanish(pattern *Pattern, timeout time.Duration) (bool, error)
	ReadText(params OCRParams) (string, error)
	FindText(query string, params OCRParams) ([]TextMatch, error)
	CollectWords(params OCRParams) ([]OCRWord, error)
	CollectLines(params OCRParams) ([]OCRLine, error)
	Hover(opts InputOptions) error
	Click(opts InputOptions) error
	RightClick(opts InputOptions) error
	DoubleClick(opts InputOptions) error
	MouseDown(opts InputOptions) error
	MouseUp(opts InputOptions) error
	TypeText(text string, opts InputOptions) error
	Paste(text string, opts InputOptions) error
	DragDrop(target TargetPointProvider, opts InputOptions) error
	Wheel(direction WheelDirection, steps int, opts InputOptions) error
	KeyDown(keys ...string) error
	KeyUp(keys ...string) error
}

// MatchAPI defines a live match result that can be used directly for follow-up
// search, capture, and current direct input verbs without manual point routing.
type MatchAPI interface {
	Bounds() Region
	Region() Region
	Center() Point
	TargetPoint() Point
	Live() bool
	Capture() (*Image, error)
	Find(pattern *Pattern) (Match, error)
	FindAnyList(patterns []*Pattern) ([]Match, error)
	FindBestList(patterns []*Pattern) (Match, error)
	Exists(pattern *Pattern, timeout time.Duration) (Match, bool, error)
	Has(pattern *Pattern, timeout time.Duration) (bool, error)
	Wait(pattern *Pattern, timeout time.Duration) (Match, error)
	WaitAnyList(patterns []*Pattern, timeout time.Duration) ([]Match, error)
	WaitBestList(patterns []*Pattern, timeout time.Duration) (Match, error)
	WaitVanish(pattern *Pattern, timeout time.Duration) (bool, error)
	ReadText(params OCRParams) (string, error)
	FindText(query string, params OCRParams) ([]TextMatch, error)
	CollectWords(params OCRParams) ([]OCRWord, error)
	CollectLines(params OCRParams) ([]OCRLine, error)
	MoveMouse(opts InputOptions) error
	Hover(opts InputOptions) error
	Click(opts InputOptions) error
	RightClick(opts InputOptions) error
	DoubleClick(opts InputOptions) error
	MouseDown(opts InputOptions) error
	MouseUp(opts InputOptions) error
	TypeText(text string, opts InputOptions) error
	Paste(text string, opts InputOptions) error
	DragDrop(target TargetPointProvider, opts InputOptions) error
	Wheel(direction WheelDirection, steps int, opts InputOptions) error
	KeyDown(keys ...string) error
	KeyUp(keys ...string) error
}

// ScreenAPI defines live screen selection, capture, and search operations.
type ScreenAPI interface {
	Live() bool
	TargetPoint() Point
	FullRegion() LiveRegion
	Region(x, y, w, h int) LiveRegion
	RegionRect(rect Rect) LiveRegion
	Capture() (*Image, error)
	Find(pattern *Pattern) (Match, error)
	FindAnyList(patterns []*Pattern) ([]Match, error)
	FindBestList(patterns []*Pattern) (Match, error)
	Exists(pattern *Pattern, timeout time.Duration) (Match, bool, error)
	Has(pattern *Pattern, timeout time.Duration) (bool, error)
	Wait(pattern *Pattern, timeout time.Duration) (Match, error)
	WaitAnyList(patterns []*Pattern, timeout time.Duration) ([]Match, error)
	WaitBestList(patterns []*Pattern, timeout time.Duration) (Match, error)
	WaitVanish(pattern *Pattern, timeout time.Duration) (bool, error)
	ReadText(params OCRParams) (string, error)
	FindText(query string, params OCRParams) ([]TextMatch, error)
	CollectWords(params OCRParams) ([]OCRWord, error)
	CollectLines(params OCRParams) ([]OCRLine, error)
	Hover(opts InputOptions) error
	Click(opts InputOptions) error
	RightClick(opts InputOptions) error
	DoubleClick(opts InputOptions) error
	MouseDown(opts InputOptions) error
	MouseUp(opts InputOptions) error
	TypeText(text string, opts InputOptions) error
	Paste(text string, opts InputOptions) error
	DragDrop(target TargetPointProvider, opts InputOptions) error
	Wheel(direction WheelDirection, steps int, opts InputOptions) error
	KeyDown(keys ...string) error
	KeyUp(keys ...string) error
}

// RuntimeAPI defines the public live runtime client used to discover screens and capture/search them.
type RuntimeAPI interface {
	Address() string
	Close() error
	Screens() ([]Screen, error)
	PrimaryScreen() (Screen, error)
	Screen(id int) (Screen, error)
	Capture() (*Image, error)
	CaptureRegion(region Region) (*Image, error)
	Region(region Region) LiveRegion
}

// InputAPI exposes desktop input actions.
// This is the compatibility layer for click/type/hotkey style operations.
type InputAPI interface {
	MoveMouse(x, y int, opts InputOptions) error
	Hover(x, y int, opts InputOptions) error
	Click(x, y int, opts InputOptions) error
	RightClick(x, y int, opts InputOptions) error
	DoubleClick(x, y int, opts InputOptions) error
	MouseDown(x, y int, opts InputOptions) error
	MouseUp(x, y int, opts InputOptions) error
	TypeText(text string, opts InputOptions) error
	Paste(text string, opts InputOptions) error
	Hotkey(keys ...string) error
	KeyDown(keys ...string) error
	KeyUp(keys ...string) error
	Wheel(x, y int, direction WheelDirection, steps int, opts InputOptions) error
	DragDrop(fromX, fromY, toX, toY int, opts InputOptions) error
}

// ObserveAPI exposes appear/vanish/change polling contracts for a region.
type ObserveAPI interface {
	ObserveAppear(source *Image, region Region, pattern *Pattern, opts ObserveOptions) ([]ObserveEvent, error)
	ObserveVanish(source *Image, region Region, pattern *Pattern, opts ObserveOptions) ([]ObserveEvent, error)
	ObserveChange(source *Image, region Region, opts ObserveOptions) ([]ObserveEvent, error)
}

// AppAPI exposes lightweight app lifecycle helpers used by script flows.
type AppAPI interface {
	Open(name string, args []string, opts AppOptions) error
	Focus(name string, opts AppOptions) error
	Close(name string, opts AppOptions) error
	IsRunning(name string, opts AppOptions) (bool, error)
	ListWindows(name string, opts AppOptions) ([]Window, error)
	FindWindows(name string, query WindowQuery, opts AppOptions) ([]Window, error)
	GetWindow(name string, query WindowQuery, opts AppOptions) (Window, bool, error)
	FocusedWindow(name string, opts AppOptions) (Window, bool, error)
}

var (
	_ ImageAPI      = (*Image)(nil)
	_ PatternAPI    = (*Pattern)(nil)
	_ FinderAPI     = (*Finder)(nil)
	_ RegionAPI     = (*Region)(nil)
	_ LiveRegionAPI = (*LiveRegion)(nil)
	_ MatchAPI      = Match{}
	_ ScreenAPI     = (*Screen)(nil)
	_ RuntimeAPI    = (*Runtime)(nil)
	_ InputAPI      = (*InputController)(nil)
	_ ObserveAPI    = (*ObserverController)(nil)
	_ AppAPI        = (*AppController)(nil)
)
