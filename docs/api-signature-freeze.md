# API Signature Freeze (Workstream 1)

This freeze covers the first GoLang-compatible core for:

- `Image`
- `Pattern`
- `Match`
- `Finder`
- `Region`
- `Screen`
- settings defaults

## Public constructors

```go
func NewImageFromGray(name string, src *image.Gray) (*Image, error)
func NewImageFromAny(name string, src image.Image) (*Image, error)
func NewImageFromMatrix(name string, rows [][]uint8) (*Image, error)
func NewPattern(img *Image) (*Pattern, error)
func NewFinder(source *Image) (*Finder, error)
func NewRegion(x, y, w, h int) Region
func NewScreen(id int, bounds Rect) Screen
```

## Frozen `Image` surface

```go
func (i *Image) Name() string
func (i *Image) Width() int
func (i *Image) Height() int
func (i *Image) Gray() *image.Gray
func (i *Image) Clone() *Image
```

## Frozen `Pattern` surface

```go
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

## Frozen `Finder` surface

```go
func (f *Finder) SetMatcher(m core.Matcher)
func (f *Finder) Find(pattern *Pattern) (Match, error)
func (f *Finder) FindAll(pattern *Pattern) ([]Match, error)
func (f *Finder) LastMatches() []Match
func SortMatchesByRowColumn(matches []Match)
func SortMatchesByColumnRow(matches []Match)
```

## Frozen `Region` and `Screen` baseline

```go
func (r Region) Center() Point
func (r Region) Grow(dx, dy int) Region
func (r Region) Offset(dx, dy int) Region
```

Any additions are non-breaking. Any signature change to these functions requires
an explicit API freeze update.
