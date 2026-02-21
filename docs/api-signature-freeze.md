# API Signature Freeze

This freeze covers the complete currently exported GoLang API in `pkg/sikuli`.

## Exported constants

```go
const DefaultSimilarity = 0.70
const ExactSimilarity = 0.99
const DefaultAutoWaitTimeout = 3.0
const DefaultWaitScanRate = 3.0
const DefaultObserveScanRate = 3.0
const DefaultOCRLanguage = "eng"
```

## Exported sentinel errors

```go
var ErrFindFailed error
var ErrTimeout error
var ErrInvalidTarget error
var ErrBackendUnsupported error
```

## Exported interfaces

```go
type ImageAPI interface {
  Name() string
  Width() int
  Height() int
  Gray() *image.Gray
  Clone() *Image
  Crop(rect Rect) (*Image, error)
}

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
```

## Exported data types and methods

### Point

```go
type Point struct {
  X int
  Y int
}
func NewPoint(x, y int) Point
```

### Location

```go
type Location struct {
  X int
  Y int
}
func NewLocation(x, y int) Location
func (l Location) ToPoint() Point
func (l Location) Move(dx, dy int) Location
func (l Location) String() string
```

### Offset

```go
type Offset struct {
  X int
  Y int
}
func NewOffset(x, y int) Offset
func (o Offset) ToPoint() Point
func (o Offset) String() string
```

### Rect

```go
type Rect struct {
  X int
  Y int
  W int
  H int
}
func NewRect(x, y, w, h int) Rect
func (r Rect) Empty() bool
func (r Rect) Contains(p Point) bool
func (r Rect) String() string
```

### Region

```go
type Region struct {
  Rect
  ThrowException  bool
  AutoWaitTimeout float64
  WaitScanRate    float64
  ObserveScanRate float64
}
func NewRegion(x, y, w, h int) Region
func (r Region) Center() Point
func (r Region) Grow(dx, dy int) Region
func (r Region) Offset(dx, dy int) Region
func (r Region) MoveTo(x, y int) Region
func (r Region) SetSize(w, h int) Region
func (r Region) Contains(p Point) bool
func (r Region) ContainsRegion(other Region) bool
func (r Region) Union(other Region) Region
func (r Region) Intersection(other Region) Region
func (r *Region) SetThrowException(flag bool)
func (r *Region) ResetThrowException()
func (r *Region) SetAutoWaitTimeout(sec float64)
func (r *Region) SetWaitScanRate(rate float64)
func (r *Region) SetObserveScanRate(rate float64)
func (r Region) Find(source *Image, pattern *Pattern) (Match, error)
func (r Region) Exists(source *Image, pattern *Pattern, timeout time.Duration) (Match, bool, error)
func (r Region) Has(source *Image, pattern *Pattern, timeout time.Duration) (bool, error)
func (r Region) Wait(source *Image, pattern *Pattern, timeout time.Duration) (Match, error)
func (r Region) WaitVanish(source *Image, pattern *Pattern, timeout time.Duration) (bool, error)
func (r Region) FindAll(source *Image, pattern *Pattern) ([]Match, error)
func (r Region) FindAllByRow(source *Image, pattern *Pattern) ([]Match, error)
func (r Region) FindAllByColumn(source *Image, pattern *Pattern) ([]Match, error)
func (r Region) ReadText(source *Image, params OCRParams) (string, error)
func (r Region) FindText(source *Image, query string, params OCRParams) ([]TextMatch, error)
```

### Screen

```go
type Screen struct {
  ID     int
  Bounds Rect
}
func NewScreen(id int, bounds Rect) Screen
```

### Image

```go
type Image struct
func NewImageFromGray(name string, src *image.Gray) (*Image, error)
func NewImageFromAny(name string, src image.Image) (*Image, error)
func NewImageFromMatrix(name string, rows [][]uint8) (*Image, error)
func (i *Image) Name() string
func (i *Image) Width() int
func (i *Image) Height() int
func (i *Image) Gray() *image.Gray
func (i *Image) Clone() *Image
func (i *Image) Crop(rect Rect) (*Image, error)
```

### Pattern

```go
type Pattern struct
func NewPattern(img *Image) (*Pattern, error)
func (p *Pattern) Image() *Image
func (p *Pattern) Similar(sim float64) *Pattern
func (p *Pattern) Similarity() float64
func (p *Pattern) Exact() *Pattern
func (p *Pattern) TargetOffset(dx, dy int) *Pattern
func (p *Pattern) Offset() Point
func (p *Pattern) Resize(factor float64) *Pattern
func (p *Pattern) ResizeFactor() float64
func (p *Pattern) WithMask(mask *image.Gray) (*Pattern, error)
func (p *Pattern) WithMaskMatrix(rows [][]uint8) (*Pattern, error)
func (p *Pattern) Mask() *image.Gray
```

### Match

```go
type Match struct {
  Rect
  Score  float64
  Target Point
  Index  int
}
func NewMatch(x, y, w, h int, score float64, off Point) Match
func (m Match) String() string
```

### TextMatch

```go
type TextMatch struct {
  Rect
  Text       string
  Confidence float64
  Index      int
}
```

### OCRParams

```go
type OCRParams struct {
  Language         string
  TrainingDataPath string
  MinConfidence    float64
  Timeout          time.Duration
  CaseSensitive    bool
}
```

### Finder

```go
type Finder struct
func NewFinder(source *Image) (*Finder, error)
func (f *Finder) SetMatcher(m core.Matcher)
func (f *Finder) SetOCRBackend(ocr core.OCR)
func (f *Finder) Find(pattern *Pattern) (Match, error)
func (f *Finder) FindAll(pattern *Pattern) ([]Match, error)
func (f *Finder) FindAllByRow(pattern *Pattern) ([]Match, error)
func (f *Finder) FindAllByColumn(pattern *Pattern) ([]Match, error)
func (f *Finder) Exists(pattern *Pattern) (Match, bool, error)
func (f *Finder) Has(pattern *Pattern) (bool, error)
func (f *Finder) Wait(pattern *Pattern, timeout time.Duration) (Match, error)
func (f *Finder) WaitVanish(pattern *Pattern, timeout time.Duration) (bool, error)
func (f *Finder) ReadText(params OCRParams) (string, error)
func (f *Finder) FindText(query string, params OCRParams) ([]TextMatch, error)
func (f *Finder) LastMatches() []Match
func SortMatchesByRowColumn(matches []Match)
func SortMatchesByColumnRow(matches []Match)
```

### Options

```go
type Options struct
func NewOptions() *Options
func NewOptionsFromMap(entries map[string]string) *Options
func (o *Options) Has(key string) bool
func (o *Options) GetString(key, def string) string
func (o *Options) SetString(key, value string)
func (o *Options) GetInt(key string, def int) int
func (o *Options) SetInt(key string, value int)
func (o *Options) GetFloat64(key string, def float64) float64
func (o *Options) SetFloat64(key string, value float64)
func (o *Options) GetBool(key string, def bool) bool
func (o *Options) SetBool(key string, value bool)
func (o *Options) Delete(key string)
func (o *Options) Entries() map[string]string
func (o *Options) Merge(other *Options)
func (o *Options) Clone() *Options
```

### Runtime settings

```go
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
func UpdateSettings(apply func(*RuntimeSettings)) RuntimeSettings
func ResetSettings() RuntimeSettings
```

## Compatibility rule

Any signature change to exported members listed above requires an explicit update to this document. Additions are allowed when they are non-breaking.
