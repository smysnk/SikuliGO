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
// Semantics follow SikuliX Finder style calls for find/findAll/exists/wait flows.
type FinderAPI interface {
	Find(pattern *Pattern) (Match, error)
	FindAll(pattern *Pattern) ([]Match, error)
	FindAllByRow(pattern *Pattern) ([]Match, error)
	FindAllByColumn(pattern *Pattern) ([]Match, error)
	Exists(pattern *Pattern) (Match, bool, error)
	Has(pattern *Pattern) (bool, error)
	Wait(pattern *Pattern, timeout time.Duration) (Match, error)
	WaitVanish(pattern *Pattern, timeout time.Duration) (bool, error)
	ReadText(params OCRParams) (string, error)
	FindText(query string, params OCRParams) ([]TextMatch, error)
	LastMatches() []Match
}

// RegionAPI defines region geometry and region-scoped automation operations.
// It maps to familiar SikuliX Region methods (find, exists, wait, findAll, readText).
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
	ReadText(source *Image, params OCRParams) (string, error)
	FindText(source *Image, query string, params OCRParams) ([]TextMatch, error)
}

// InputAPI exposes desktop input actions.
// This is the compatibility layer for click/type/hotkey style operations.
type InputAPI interface {
	MoveMouse(x, y int, opts InputOptions) error
	Click(x, y int, opts InputOptions) error
	TypeText(text string, opts InputOptions) error
	Hotkey(keys ...string) error
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
}

var (
	_ ImageAPI   = (*Image)(nil)
	_ PatternAPI = (*Pattern)(nil)
	_ FinderAPI  = (*Finder)(nil)
	_ RegionAPI  = (*Region)(nil)
	_ InputAPI   = (*InputController)(nil)
	_ ObserveAPI = (*ObserverController)(nil)
	_ AppAPI     = (*AppController)(nil)
)
