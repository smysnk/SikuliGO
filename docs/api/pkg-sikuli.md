# API: `pkg/sikuli`

[Back to API Index](./)

## Full Package Doc

```text
package sikuli // import "github.com/sikulix/portgo/pkg/sikuli"


CONSTANTS

const (
	// DefaultSimilarity matches classic Sikuli behavior for image search.
	DefaultSimilarity = 0.70

	// ExactSimilarity is used by Pattern.Exact().
	ExactSimilarity = 0.99

	// DefaultAutoWaitTimeout is the baseline timeout for wait/find loops.
	DefaultAutoWaitTimeout = 3.0

	// DefaultWaitScanRate controls wait polling frequency.
	DefaultWaitScanRate = 3.0

	// DefaultObserveScanRate controls observe polling frequency.
	DefaultObserveScanRate = 3.0
)
const DefaultOCRLanguage = "eng"

VARIABLES

var (
	ErrFindFailed         = errors.New("sikuli: find failed")
	ErrTimeout            = errors.New("sikuli: timeout")
	ErrInvalidTarget      = errors.New("sikuli: invalid target")
	ErrBackendUnsupported = errors.New("sikuli: backend unsupported")
)

FUNCTIONS

func SortMatchesByColumnRow(matches []Match)
    SortMatchesByColumnRow keeps parity with Java helper behavior for "by
    column".

func SortMatchesByRowColumn(matches []Match)
    SortMatchesByRowColumn keeps parity with Java helper behavior for "by row".


TYPES

type Finder struct {
	// Has unexported fields.
}

func NewFinder(source *Image) (*Finder, error)

func (f *Finder) Exists(pattern *Pattern) (Match, bool, error)
    Exists returns the first match when present. Missing targets return
    (Match{}, false, nil).

func (f *Finder) Find(pattern *Pattern) (Match, error)

func (f *Finder) FindAll(pattern *Pattern) ([]Match, error)

func (f *Finder) FindAllByColumn(pattern *Pattern) ([]Match, error)

func (f *Finder) FindAllByRow(pattern *Pattern) ([]Match, error)

func (f *Finder) FindText(query string, params OCRParams) ([]TextMatch, error)

func (f *Finder) Has(pattern *Pattern) (bool, error)
    Has reports whether the target exists and bubbles non-find errors.

func (f *Finder) LastMatches() []Match

func (f *Finder) ReadText(params OCRParams) (string, error)

func (f *Finder) SetMatcher(m core.Matcher)

func (f *Finder) SetOCRBackend(ocr core.OCR)

func (f *Finder) Wait(pattern *Pattern, timeout time.Duration) (Match, error)

func (f *Finder) WaitVanish(pattern *Pattern, timeout time.Duration) (bool, error)

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

type Image struct {
	// Has unexported fields.
}

func NewImageFromAny(name string, src image.Image) (*Image, error)

func NewImageFromGray(name string, src *image.Gray) (*Image, error)

func NewImageFromMatrix(name string, rows [][]uint8) (*Image, error)

func (i *Image) Clone() *Image

func (i *Image) Crop(rect Rect) (*Image, error)

func (i *Image) Gray() *image.Gray

func (i *Image) Height() int

func (i *Image) Name() string

func (i *Image) Width() int

type ImageAPI interface {
	Name() string
	Width() int
	Height() int
	Gray() *image.Gray
	Clone() *Image
	Crop(rect Rect) (*Image, error)
}

type InputAPI interface {
	MoveMouse(x, y int, opts InputOptions) error
	Click(x, y int, opts InputOptions) error
	TypeText(text string, opts InputOptions) error
	Hotkey(keys ...string) error
}

type InputController struct {
	// Has unexported fields.
}

func NewInputController() *InputController

func (c *InputController) Click(x, y int, opts InputOptions) error

func (c *InputController) Hotkey(keys ...string) error

func (c *InputController) MoveMouse(x, y int, opts InputOptions) error

func (c *InputController) SetBackend(backend core.Input)

func (c *InputController) TypeText(text string, opts InputOptions) error

type InputOptions struct {
	Delay  time.Duration
	Button MouseButton
}

type Location struct {
	X int
	Y int
}

func NewLocation(x, y int) Location

func (l Location) Move(dx, dy int) Location

func (l Location) String() string

func (l Location) ToPoint() Point

type Match struct {
	Rect
	Score  float64
	Target Point
	Index  int
}

func NewMatch(x, y, w, h int, score float64, off Point) Match

func (m Match) String() string

type MouseButton string

const (
	MouseButtonLeft   MouseButton = "left"
	MouseButtonRight  MouseButton = "right"
	MouseButtonMiddle MouseButton = "middle"
)
type OCRParams struct {
	Language         string
	TrainingDataPath string
	MinConfidence    float64
	Timeout          time.Duration
	CaseSensitive    bool
}

type Offset struct {
	X int
	Y int
}

func NewOffset(x, y int) Offset

func (o Offset) String() string

func (o Offset) ToPoint() Point

type Options struct {
	// Has unexported fields.
}

func NewOptions() *Options

func NewOptionsFromMap(entries map[string]string) *Options

func (o *Options) Clone() *Options

func (o *Options) Delete(key string)

func (o *Options) Entries() map[string]string

func (o *Options) GetBool(key string, def bool) bool

func (o *Options) GetFloat64(key string, def float64) float64

func (o *Options) GetInt(key string, def int) int

func (o *Options) GetString(key, def string) string

func (o *Options) Has(key string) bool

func (o *Options) Merge(other *Options)

func (o *Options) SetBool(key string, value bool)

func (o *Options) SetFloat64(key string, value float64)

func (o *Options) SetInt(key string, value int)

func (o *Options) SetString(key, value string)

type Pattern struct {
	// Has unexported fields.
}

func NewPattern(img *Image) (*Pattern, error)

func (p *Pattern) Exact() *Pattern

func (p *Pattern) Image() *Image

func (p *Pattern) Mask() *image.Gray

func (p *Pattern) Offset() Point

func (p *Pattern) Resize(factor float64) *Pattern

func (p *Pattern) ResizeFactor() float64

func (p *Pattern) Similar(sim float64) *Pattern

func (p *Pattern) Similarity() float64

func (p *Pattern) TargetOffset(dx, dy int) *Pattern

func (p *Pattern) WithMask(mask *image.Gray) (*Pattern, error)

func (p *Pattern) WithMaskMatrix(rows [][]uint8) (*Pattern, error)

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

type Point struct {
	X int
	Y int
}

func NewPoint(x, y int) Point

type Rect struct {
	X int
	Y int
	W int
	H int
}

func NewRect(x, y, w, h int) Rect

func (r Rect) Contains(p Point) bool

func (r Rect) Empty() bool

func (r Rect) String() string

type Region struct {
	Rect
	ThrowException  bool
	AutoWaitTimeout float64
	WaitScanRate    float64
	ObserveScanRate float64
}

func NewRegion(x, y, w, h int) Region

func (r Region) Center() Point

func (r Region) Contains(p Point) bool

func (r Region) ContainsRegion(other Region) bool

func (r Region) Exists(source *Image, pattern *Pattern, timeout time.Duration) (Match, bool, error)

func (r Region) Find(source *Image, pattern *Pattern) (Match, error)

func (r Region) FindAll(source *Image, pattern *Pattern) ([]Match, error)

func (r Region) FindAllByColumn(source *Image, pattern *Pattern) ([]Match, error)

func (r Region) FindAllByRow(source *Image, pattern *Pattern) ([]Match, error)

func (r Region) FindText(source *Image, query string, params OCRParams) ([]TextMatch, error)

func (r Region) Grow(dx, dy int) Region

func (r Region) Has(source *Image, pattern *Pattern, timeout time.Duration) (bool, error)

func (r Region) Intersection(other Region) Region

func (r Region) MoveTo(x, y int) Region

func (r Region) Offset(dx, dy int) Region

func (r Region) ReadText(source *Image, params OCRParams) (string, error)

func (r *Region) ResetThrowException()

func (r *Region) SetAutoWaitTimeout(sec float64)

func (r *Region) SetObserveScanRate(rate float64)

func (r Region) SetSize(w, h int) Region

func (r *Region) SetThrowException(flag bool)

func (r *Region) SetWaitScanRate(rate float64)

func (r Region) Union(other Region) Region

func (r Region) Wait(source *Image, pattern *Pattern, timeout time.Duration) (Match, error)

func (r Region) WaitVanish(source *Image, pattern *Pattern, timeout time.Duration) (bool, error)

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

type RuntimeSettings struct {
	ImageCache       int
	ShowActions      bool
	WaitScanRate     float64
	ObserveScanRate  float64
	AutoWaitTimeout  float64
	MinSimilarity    float64
	FindFailedThrows bool
}

func GetSettings() RuntimeSettings

func ResetSettings() RuntimeSettings

func UpdateSettings(apply func(*RuntimeSettings)) RuntimeSettings

type Screen struct {
	ID     int
	Bounds Rect
}

func NewScreen(id int, bounds Rect) Screen

type TextMatch struct {
	Rect
	Text       string
	Confidence float64
	Index      int
}

```
