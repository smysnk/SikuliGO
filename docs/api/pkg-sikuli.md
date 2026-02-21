# API: `pkg/sikuli`

[Back to API Index](./)

<style>
  .api-type { color: #0f766e; font-weight: 700; }
  .api-func { color: #1d4ed8; font-weight: 700; }
  .api-method { color: #7c3aed; font-weight: 700; }
  .api-signature { font-family: ui-monospace, SFMono-Regular, Menlo, Consolas, "Liberation Mono", monospace; }
</style>

Legend: <span class="api-type">Type</span>, <span class="api-func">Function</span>, <span class="api-method">Method</span>

Package: `package sikuli // import "github.com/sikulix/portgo/pkg/sikuli"`

## Symbol Index

### Types

- <span class="api-type">[`AppAPI`](#type-appapi)</span>
- <span class="api-type">[`AppController`](#type-appcontroller)</span>
- <span class="api-type">[`AppOptions`](#type-appoptions)</span>
- <span class="api-type">[`Finder`](#type-finder)</span>
- <span class="api-type">[`FinderAPI`](#type-finderapi)</span>
- <span class="api-type">[`Image`](#type-image)</span>
- <span class="api-type">[`ImageAPI`](#type-imageapi)</span>
- <span class="api-type">[`InputAPI`](#type-inputapi)</span>
- <span class="api-type">[`InputController`](#type-inputcontroller)</span>
- <span class="api-type">[`InputOptions`](#type-inputoptions)</span>
- <span class="api-type">[`Location`](#type-location)</span>
- <span class="api-type">[`Match`](#type-match)</span>
- <span class="api-type">[`MouseButton`](#type-mousebutton)</span>
- <span class="api-type">[`OCRParams`](#type-ocrparams)</span>
- <span class="api-type">[`ObserveAPI`](#type-observeapi)</span>
- <span class="api-type">[`ObserveEvent`](#type-observeevent)</span>
- <span class="api-type">[`ObserveEventType`](#type-observeeventtype)</span>
- <span class="api-type">[`ObserveOptions`](#type-observeoptions)</span>
- <span class="api-type">[`ObserverController`](#type-observercontroller)</span>
- <span class="api-type">[`Offset`](#type-offset)</span>
- <span class="api-type">[`Options`](#type-options)</span>
- <span class="api-type">[`Pattern`](#type-pattern)</span>
- <span class="api-type">[`PatternAPI`](#type-patternapi)</span>
- <span class="api-type">[`Point`](#type-point)</span>
- <span class="api-type">[`Rect`](#type-rect)</span>
- <span class="api-type">[`Region`](#type-region)</span>
- <span class="api-type">[`RegionAPI`](#type-regionapi)</span>
- <span class="api-type">[`RuntimeSettings`](#type-runtimesettings)</span>
- <span class="api-type">[`Screen`](#type-screen)</span>
- <span class="api-type">[`TextMatch`](#type-textmatch)</span>
- <span class="api-type">[`Window`](#type-window)</span>

### Functions

- <span class="api-func">[`SortMatchesByColumnRow`](#func-sortmatchesbycolumnrow)</span>
- <span class="api-func">[`SortMatchesByRowColumn`](#func-sortmatchesbyrowcolumn)</span>
- <span class="api-func">[`NewAppController`](#func-newappcontroller)</span>
- <span class="api-func">[`NewFinder`](#func-newfinder)</span>
- <span class="api-func">[`NewImageFromAny`](#func-newimagefromany)</span>
- <span class="api-func">[`NewImageFromGray`](#func-newimagefromgray)</span>
- <span class="api-func">[`NewImageFromMatrix`](#func-newimagefrommatrix)</span>
- <span class="api-func">[`NewInputController`](#func-newinputcontroller)</span>
- <span class="api-func">[`NewLocation`](#func-newlocation)</span>
- <span class="api-func">[`NewMatch`](#func-newmatch)</span>
- <span class="api-func">[`NewObserverController`](#func-newobservercontroller)</span>
- <span class="api-func">[`NewOffset`](#func-newoffset)</span>
- <span class="api-func">[`NewOptions`](#func-newoptions)</span>
- <span class="api-func">[`NewOptionsFromMap`](#func-newoptionsfrommap)</span>
- <span class="api-func">[`NewPattern`](#func-newpattern)</span>
- <span class="api-func">[`NewPoint`](#func-newpoint)</span>
- <span class="api-func">[`NewRect`](#func-newrect)</span>
- <span class="api-func">[`NewRegion`](#func-newregion)</span>
- <span class="api-func">[`GetSettings`](#func-getsettings)</span>
- <span class="api-func">[`ResetSettings`](#func-resetsettings)</span>
- <span class="api-func">[`UpdateSettings`](#func-updatesettings)</span>
- <span class="api-func">[`NewScreen`](#func-newscreen)</span>

### Methods

- <span class="api-method">[`AppController.Close`](#method-appcontroller-close)</span>
- <span class="api-method">[`AppController.Focus`](#method-appcontroller-focus)</span>
- <span class="api-method">[`AppController.IsRunning`](#method-appcontroller-isrunning)</span>
- <span class="api-method">[`AppController.ListWindows`](#method-appcontroller-listwindows)</span>
- <span class="api-method">[`AppController.Open`](#method-appcontroller-open)</span>
- <span class="api-method">[`AppController.SetBackend`](#method-appcontroller-setbackend)</span>
- <span class="api-method">[`Finder.Exists`](#method-finder-exists)</span>
- <span class="api-method">[`Finder.Find`](#method-finder-find)</span>
- <span class="api-method">[`Finder.FindAll`](#method-finder-findall)</span>
- <span class="api-method">[`Finder.FindAllByColumn`](#method-finder-findallbycolumn)</span>
- <span class="api-method">[`Finder.FindAllByRow`](#method-finder-findallbyrow)</span>
- <span class="api-method">[`Finder.FindText`](#method-finder-findtext)</span>
- <span class="api-method">[`Finder.Has`](#method-finder-has)</span>
- <span class="api-method">[`Finder.LastMatches`](#method-finder-lastmatches)</span>
- <span class="api-method">[`Finder.ReadText`](#method-finder-readtext)</span>
- <span class="api-method">[`Finder.SetMatcher`](#method-finder-setmatcher)</span>
- <span class="api-method">[`Finder.SetOCRBackend`](#method-finder-setocrbackend)</span>
- <span class="api-method">[`Finder.Wait`](#method-finder-wait)</span>
- <span class="api-method">[`Finder.WaitVanish`](#method-finder-waitvanish)</span>
- <span class="api-method">[`Image.Clone`](#method-image-clone)</span>
- <span class="api-method">[`Image.Crop`](#method-image-crop)</span>
- <span class="api-method">[`Image.Gray`](#method-image-gray)</span>
- <span class="api-method">[`Image.Height`](#method-image-height)</span>
- <span class="api-method">[`Image.Name`](#method-image-name)</span>
- <span class="api-method">[`Image.Width`](#method-image-width)</span>
- <span class="api-method">[`InputController.Click`](#method-inputcontroller-click)</span>
- <span class="api-method">[`InputController.Hotkey`](#method-inputcontroller-hotkey)</span>
- <span class="api-method">[`InputController.MoveMouse`](#method-inputcontroller-movemouse)</span>
- <span class="api-method">[`InputController.SetBackend`](#method-inputcontroller-setbackend)</span>
- <span class="api-method">[`InputController.TypeText`](#method-inputcontroller-typetext)</span>
- <span class="api-method">[`Location.Move`](#method-location-move)</span>
- <span class="api-method">[`Location.String`](#method-location-string)</span>
- <span class="api-method">[`Location.ToPoint`](#method-location-topoint)</span>
- <span class="api-method">[`Match.String`](#method-match-string)</span>
- <span class="api-method">[`ObserverController.ObserveAppear`](#method-observercontroller-observeappear)</span>
- <span class="api-method">[`ObserverController.ObserveChange`](#method-observercontroller-observechange)</span>
- <span class="api-method">[`ObserverController.ObserveVanish`](#method-observercontroller-observevanish)</span>
- <span class="api-method">[`ObserverController.SetBackend`](#method-observercontroller-setbackend)</span>
- <span class="api-method">[`Offset.String`](#method-offset-string)</span>
- <span class="api-method">[`Offset.ToPoint`](#method-offset-topoint)</span>
- <span class="api-method">[`Options.Clone`](#method-options-clone)</span>
- <span class="api-method">[`Options.Delete`](#method-options-delete)</span>
- <span class="api-method">[`Options.Entries`](#method-options-entries)</span>
- <span class="api-method">[`Options.GetBool`](#method-options-getbool)</span>
- <span class="api-method">[`Options.GetFloat64`](#method-options-getfloat64)</span>
- <span class="api-method">[`Options.GetInt`](#method-options-getint)</span>
- <span class="api-method">[`Options.GetString`](#method-options-getstring)</span>
- <span class="api-method">[`Options.Has`](#method-options-has)</span>
- <span class="api-method">[`Options.Merge`](#method-options-merge)</span>
- <span class="api-method">[`Options.SetBool`](#method-options-setbool)</span>
- <span class="api-method">[`Options.SetFloat64`](#method-options-setfloat64)</span>
- <span class="api-method">[`Options.SetInt`](#method-options-setint)</span>
- <span class="api-method">[`Options.SetString`](#method-options-setstring)</span>
- <span class="api-method">[`Pattern.Exact`](#method-pattern-exact)</span>
- <span class="api-method">[`Pattern.Image`](#method-pattern-image)</span>
- <span class="api-method">[`Pattern.Mask`](#method-pattern-mask)</span>
- <span class="api-method">[`Pattern.Offset`](#method-pattern-offset)</span>
- <span class="api-method">[`Pattern.Resize`](#method-pattern-resize)</span>
- <span class="api-method">[`Pattern.ResizeFactor`](#method-pattern-resizefactor)</span>
- <span class="api-method">[`Pattern.Similar`](#method-pattern-similar)</span>
- <span class="api-method">[`Pattern.Similarity`](#method-pattern-similarity)</span>
- <span class="api-method">[`Pattern.TargetOffset`](#method-pattern-targetoffset)</span>
- <span class="api-method">[`Pattern.WithMask`](#method-pattern-withmask)</span>
- <span class="api-method">[`Pattern.WithMaskMatrix`](#method-pattern-withmaskmatrix)</span>
- <span class="api-method">[`Rect.Contains`](#method-rect-contains)</span>
- <span class="api-method">[`Rect.Empty`](#method-rect-empty)</span>
- <span class="api-method">[`Rect.String`](#method-rect-string)</span>
- <span class="api-method">[`Region.Center`](#method-region-center)</span>
- <span class="api-method">[`Region.Contains`](#method-region-contains)</span>
- <span class="api-method">[`Region.ContainsRegion`](#method-region-containsregion)</span>
- <span class="api-method">[`Region.Exists`](#method-region-exists)</span>
- <span class="api-method">[`Region.Find`](#method-region-find)</span>
- <span class="api-method">[`Region.FindAll`](#method-region-findall)</span>
- <span class="api-method">[`Region.FindAllByColumn`](#method-region-findallbycolumn)</span>
- <span class="api-method">[`Region.FindAllByRow`](#method-region-findallbyrow)</span>
- <span class="api-method">[`Region.FindText`](#method-region-findtext)</span>
- <span class="api-method">[`Region.Grow`](#method-region-grow)</span>
- <span class="api-method">[`Region.Has`](#method-region-has)</span>
- <span class="api-method">[`Region.Intersection`](#method-region-intersection)</span>
- <span class="api-method">[`Region.MoveTo`](#method-region-moveto)</span>
- <span class="api-method">[`Region.Offset`](#method-region-offset)</span>
- <span class="api-method">[`Region.ReadText`](#method-region-readtext)</span>
- <span class="api-method">[`Region.ResetThrowException`](#method-region-resetthrowexception)</span>
- <span class="api-method">[`Region.SetAutoWaitTimeout`](#method-region-setautowaittimeout)</span>
- <span class="api-method">[`Region.SetObserveScanRate`](#method-region-setobservescanrate)</span>
- <span class="api-method">[`Region.SetSize`](#method-region-setsize)</span>
- <span class="api-method">[`Region.SetThrowException`](#method-region-setthrowexception)</span>
- <span class="api-method">[`Region.SetWaitScanRate`](#method-region-setwaitscanrate)</span>
- <span class="api-method">[`Region.Union`](#method-region-union)</span>
- <span class="api-method">[`Region.Wait`](#method-region-wait)</span>
- <span class="api-method">[`Region.WaitVanish`](#method-region-waitvanish)</span>

## Declarations

### Types

#### <a id="type-appapi"></a><span class="api-type">Type</span> `AppAPI`

- Signature: <span class="api-signature">`type AppAPI interface {`</span>

#### <a id="type-appcontroller"></a><span class="api-type">Type</span> `AppController`

- Signature: <span class="api-signature">`type AppController struct {`</span>

#### <a id="type-appoptions"></a><span class="api-type">Type</span> `AppOptions`

- Signature: <span class="api-signature">`type AppOptions struct {`</span>

#### <a id="type-finder"></a><span class="api-type">Type</span> `Finder`

- Signature: <span class="api-signature">`type Finder struct {`</span>

#### <a id="type-finderapi"></a><span class="api-type">Type</span> `FinderAPI`

- Signature: <span class="api-signature">`type FinderAPI interface {`</span>

#### <a id="type-image"></a><span class="api-type">Type</span> `Image`

- Signature: <span class="api-signature">`type Image struct {`</span>

#### <a id="type-imageapi"></a><span class="api-type">Type</span> `ImageAPI`

- Signature: <span class="api-signature">`type ImageAPI interface {`</span>

#### <a id="type-inputapi"></a><span class="api-type">Type</span> `InputAPI`

- Signature: <span class="api-signature">`type InputAPI interface {`</span>

#### <a id="type-inputcontroller"></a><span class="api-type">Type</span> `InputController`

- Signature: <span class="api-signature">`type InputController struct {`</span>

#### <a id="type-inputoptions"></a><span class="api-type">Type</span> `InputOptions`

- Signature: <span class="api-signature">`type InputOptions struct {`</span>

#### <a id="type-location"></a><span class="api-type">Type</span> `Location`

- Signature: <span class="api-signature">`type Location struct {`</span>

#### <a id="type-match"></a><span class="api-type">Type</span> `Match`

- Signature: <span class="api-signature">`type Match struct {`</span>

#### <a id="type-mousebutton"></a><span class="api-type">Type</span> `MouseButton`

- Signature: <span class="api-signature">`type MouseButton string`</span>

#### <a id="type-ocrparams"></a><span class="api-type">Type</span> `OCRParams`

- Signature: <span class="api-signature">`type OCRParams struct {`</span>

#### <a id="type-observeapi"></a><span class="api-type">Type</span> `ObserveAPI`

- Signature: <span class="api-signature">`type ObserveAPI interface {`</span>

#### <a id="type-observeevent"></a><span class="api-type">Type</span> `ObserveEvent`

- Signature: <span class="api-signature">`type ObserveEvent struct {`</span>

#### <a id="type-observeeventtype"></a><span class="api-type">Type</span> `ObserveEventType`

- Signature: <span class="api-signature">`type ObserveEventType string`</span>

#### <a id="type-observeoptions"></a><span class="api-type">Type</span> `ObserveOptions`

- Signature: <span class="api-signature">`type ObserveOptions struct {`</span>

#### <a id="type-observercontroller"></a><span class="api-type">Type</span> `ObserverController`

- Signature: <span class="api-signature">`type ObserverController struct {`</span>

#### <a id="type-offset"></a><span class="api-type">Type</span> `Offset`

- Signature: <span class="api-signature">`type Offset struct {`</span>

#### <a id="type-options"></a><span class="api-type">Type</span> `Options`

- Signature: <span class="api-signature">`type Options struct {`</span>

#### <a id="type-pattern"></a><span class="api-type">Type</span> `Pattern`

- Signature: <span class="api-signature">`type Pattern struct {`</span>

#### <a id="type-patternapi"></a><span class="api-type">Type</span> `PatternAPI`

- Signature: <span class="api-signature">`type PatternAPI interface {`</span>

#### <a id="type-point"></a><span class="api-type">Type</span> `Point`

- Signature: <span class="api-signature">`type Point struct {`</span>

#### <a id="type-rect"></a><span class="api-type">Type</span> `Rect`

- Signature: <span class="api-signature">`type Rect struct {`</span>

#### <a id="type-region"></a><span class="api-type">Type</span> `Region`

- Signature: <span class="api-signature">`type Region struct {`</span>

#### <a id="type-regionapi"></a><span class="api-type">Type</span> `RegionAPI`

- Signature: <span class="api-signature">`type RegionAPI interface {`</span>

#### <a id="type-runtimesettings"></a><span class="api-type">Type</span> `RuntimeSettings`

- Signature: <span class="api-signature">`type RuntimeSettings struct {`</span>

#### <a id="type-screen"></a><span class="api-type">Type</span> `Screen`

- Signature: <span class="api-signature">`type Screen struct {`</span>

#### <a id="type-textmatch"></a><span class="api-type">Type</span> `TextMatch`

- Signature: <span class="api-signature">`type TextMatch struct {`</span>

#### <a id="type-window"></a><span class="api-type">Type</span> `Window`

- Signature: <span class="api-signature">`type Window struct {`</span>

### Functions

#### <a id="func-sortmatchesbycolumnrow"></a><span class="api-func">Function</span> `SortMatchesByColumnRow`

- Signature: <span class="api-signature">`func SortMatchesByColumnRow(matches []Match)`</span>
- Uses: [`Match`](#type-match)
- Notes: SortMatchesByColumnRow keeps parity with Java helper behavior for "by column".

#### <a id="func-sortmatchesbyrowcolumn"></a><span class="api-func">Function</span> `SortMatchesByRowColumn`

- Signature: <span class="api-signature">`func SortMatchesByRowColumn(matches []Match)`</span>
- Uses: [`Match`](#type-match)
- Notes: SortMatchesByRowColumn keeps parity with Java helper behavior for "by row".

#### <a id="func-newappcontroller"></a><span class="api-func">Function</span> `NewAppController`

- Signature: <span class="api-signature">`func NewAppController() *AppController`</span>
- Uses: [`AppController`](#type-appcontroller)

#### <a id="func-newfinder"></a><span class="api-func">Function</span> `NewFinder`

- Signature: <span class="api-signature">`func NewFinder(source *Image) (*Finder, error)`</span>
- Uses: [`Image`](#type-image), [`Finder`](#type-finder)

#### <a id="func-newimagefromany"></a><span class="api-func">Function</span> `NewImageFromAny`

- Signature: <span class="api-signature">`func NewImageFromAny(name string, src image.Image) (*Image, error)`</span>
- Uses: [`Image`](#type-image)

#### <a id="func-newimagefromgray"></a><span class="api-func">Function</span> `NewImageFromGray`

- Signature: <span class="api-signature">`func NewImageFromGray(name string, src *image.Gray) (*Image, error)`</span>
- Uses: [`Image`](#type-image)

#### <a id="func-newimagefrommatrix"></a><span class="api-func">Function</span> `NewImageFromMatrix`

- Signature: <span class="api-signature">`func NewImageFromMatrix(name string, rows [][]uint8) (*Image, error)`</span>
- Uses: [`Image`](#type-image)

#### <a id="func-newinputcontroller"></a><span class="api-func">Function</span> `NewInputController`

- Signature: <span class="api-signature">`func NewInputController() *InputController`</span>
- Uses: [`InputController`](#type-inputcontroller)

#### <a id="func-newlocation"></a><span class="api-func">Function</span> `NewLocation`

- Signature: <span class="api-signature">`func NewLocation(x, y int) Location`</span>
- Uses: [`Location`](#type-location)

#### <a id="func-newmatch"></a><span class="api-func">Function</span> `NewMatch`

- Signature: <span class="api-signature">`func NewMatch(x, y, w, h int, score float64, off Point) Match`</span>
- Uses: [`Point`](#type-point), [`Match`](#type-match)

#### <a id="func-newobservercontroller"></a><span class="api-func">Function</span> `NewObserverController`

- Signature: <span class="api-signature">`func NewObserverController() *ObserverController`</span>
- Uses: [`ObserverController`](#type-observercontroller)

#### <a id="func-newoffset"></a><span class="api-func">Function</span> `NewOffset`

- Signature: <span class="api-signature">`func NewOffset(x, y int) Offset`</span>
- Uses: [`Offset`](#type-offset)

#### <a id="func-newoptions"></a><span class="api-func">Function</span> `NewOptions`

- Signature: <span class="api-signature">`func NewOptions() *Options`</span>
- Uses: [`Options`](#type-options)

#### <a id="func-newoptionsfrommap"></a><span class="api-func">Function</span> `NewOptionsFromMap`

- Signature: <span class="api-signature">`func NewOptionsFromMap(entries map[string]string) *Options`</span>
- Uses: [`Options`](#type-options)

#### <a id="func-newpattern"></a><span class="api-func">Function</span> `NewPattern`

- Signature: <span class="api-signature">`func NewPattern(img *Image) (*Pattern, error)`</span>
- Uses: [`Image`](#type-image), [`Pattern`](#type-pattern)

#### <a id="func-newpoint"></a><span class="api-func">Function</span> `NewPoint`

- Signature: <span class="api-signature">`func NewPoint(x, y int) Point`</span>
- Uses: [`Point`](#type-point)

#### <a id="func-newrect"></a><span class="api-func">Function</span> `NewRect`

- Signature: <span class="api-signature">`func NewRect(x, y, w, h int) Rect`</span>
- Uses: [`Rect`](#type-rect)

#### <a id="func-newregion"></a><span class="api-func">Function</span> `NewRegion`

- Signature: <span class="api-signature">`func NewRegion(x, y, w, h int) Region`</span>
- Uses: [`Region`](#type-region)

#### <a id="func-getsettings"></a><span class="api-func">Function</span> `GetSettings`

- Signature: <span class="api-signature">`func GetSettings() RuntimeSettings`</span>
- Uses: [`RuntimeSettings`](#type-runtimesettings)

#### <a id="func-resetsettings"></a><span class="api-func">Function</span> `ResetSettings`

- Signature: <span class="api-signature">`func ResetSettings() RuntimeSettings`</span>
- Uses: [`RuntimeSettings`](#type-runtimesettings)

#### <a id="func-updatesettings"></a><span class="api-func">Function</span> `UpdateSettings`

- Signature: <span class="api-signature">`func UpdateSettings(apply func(*RuntimeSettings)) RuntimeSettings`</span>
- Uses: [`RuntimeSettings`](#type-runtimesettings)

#### <a id="func-newscreen"></a><span class="api-func">Function</span> `NewScreen`

- Signature: <span class="api-signature">`func NewScreen(id int, bounds Rect) Screen`</span>
- Uses: [`Rect`](#type-rect), [`Screen`](#type-screen)

### Methods

#### <a id="method-appcontroller-close"></a><span class="api-method">Method</span> `AppController.Close`

- Signature: <span class="api-signature">`func (c *AppController) Close(name string, opts AppOptions) error`</span>
- Uses: [`AppOptions`](#type-appoptions)

#### <a id="method-appcontroller-focus"></a><span class="api-method">Method</span> `AppController.Focus`

- Signature: <span class="api-signature">`func (c *AppController) Focus(name string, opts AppOptions) error`</span>
- Uses: [`AppOptions`](#type-appoptions)

#### <a id="method-appcontroller-isrunning"></a><span class="api-method">Method</span> `AppController.IsRunning`

- Signature: <span class="api-signature">`func (c *AppController) IsRunning(name string, opts AppOptions) (bool, error)`</span>
- Uses: [`AppOptions`](#type-appoptions)

#### <a id="method-appcontroller-listwindows"></a><span class="api-method">Method</span> `AppController.ListWindows`

- Signature: <span class="api-signature">`func (c *AppController) ListWindows(name string, opts AppOptions) ([]Window, error)`</span>
- Uses: [`Window`](#type-window), [`AppOptions`](#type-appoptions)

#### <a id="method-appcontroller-open"></a><span class="api-method">Method</span> `AppController.Open`

- Signature: <span class="api-signature">`func (c *AppController) Open(name string, args []string, opts AppOptions) error`</span>
- Uses: [`AppOptions`](#type-appoptions)

#### <a id="method-appcontroller-setbackend"></a><span class="api-method">Method</span> `AppController.SetBackend`

- Signature: <span class="api-signature">`func (c *AppController) SetBackend(backend core.App)`</span>

#### <a id="method-finder-exists"></a><span class="api-method">Method</span> `Finder.Exists`

- Signature: <span class="api-signature">`func (f *Finder) Exists(pattern *Pattern) (Match, bool, error)`</span>
- Uses: [`Pattern`](#type-pattern), [`Match`](#type-match)
- Notes: Exists returns the first match when present. Missing targets return (Match{}, false, nil).

#### <a id="method-finder-find"></a><span class="api-method">Method</span> `Finder.Find`

- Signature: <span class="api-signature">`func (f *Finder) Find(pattern *Pattern) (Match, error)`</span>
- Uses: [`Pattern`](#type-pattern), [`Match`](#type-match)

#### <a id="method-finder-findall"></a><span class="api-method">Method</span> `Finder.FindAll`

- Signature: <span class="api-signature">`func (f *Finder) FindAll(pattern *Pattern) ([]Match, error)`</span>
- Uses: [`Pattern`](#type-pattern), [`Match`](#type-match)

#### <a id="method-finder-findallbycolumn"></a><span class="api-method">Method</span> `Finder.FindAllByColumn`

- Signature: <span class="api-signature">`func (f *Finder) FindAllByColumn(pattern *Pattern) ([]Match, error)`</span>
- Uses: [`Pattern`](#type-pattern), [`Match`](#type-match)

#### <a id="method-finder-findallbyrow"></a><span class="api-method">Method</span> `Finder.FindAllByRow`

- Signature: <span class="api-signature">`func (f *Finder) FindAllByRow(pattern *Pattern) ([]Match, error)`</span>
- Uses: [`Pattern`](#type-pattern), [`Match`](#type-match)

#### <a id="method-finder-findtext"></a><span class="api-method">Method</span> `Finder.FindText`

- Signature: <span class="api-signature">`func (f *Finder) FindText(query string, params OCRParams) ([]TextMatch, error)`</span>
- Uses: [`OCRParams`](#type-ocrparams), [`TextMatch`](#type-textmatch)

#### <a id="method-finder-has"></a><span class="api-method">Method</span> `Finder.Has`

- Signature: <span class="api-signature">`func (f *Finder) Has(pattern *Pattern) (bool, error)`</span>
- Uses: [`Pattern`](#type-pattern)
- Notes: Has reports whether the target exists and bubbles non-find errors.

#### <a id="method-finder-lastmatches"></a><span class="api-method">Method</span> `Finder.LastMatches`

- Signature: <span class="api-signature">`func (f *Finder) LastMatches() []Match`</span>
- Uses: [`Match`](#type-match)

#### <a id="method-finder-readtext"></a><span class="api-method">Method</span> `Finder.ReadText`

- Signature: <span class="api-signature">`func (f *Finder) ReadText(params OCRParams) (string, error)`</span>
- Uses: [`OCRParams`](#type-ocrparams)

#### <a id="method-finder-setmatcher"></a><span class="api-method">Method</span> `Finder.SetMatcher`

- Signature: <span class="api-signature">`func (f *Finder) SetMatcher(m core.Matcher)`</span>

#### <a id="method-finder-setocrbackend"></a><span class="api-method">Method</span> `Finder.SetOCRBackend`

- Signature: <span class="api-signature">`func (f *Finder) SetOCRBackend(ocr core.OCR)`</span>

#### <a id="method-finder-wait"></a><span class="api-method">Method</span> `Finder.Wait`

- Signature: <span class="api-signature">`func (f *Finder) Wait(pattern *Pattern, timeout time.Duration) (Match, error)`</span>
- Uses: [`Pattern`](#type-pattern), [`Match`](#type-match)

#### <a id="method-finder-waitvanish"></a><span class="api-method">Method</span> `Finder.WaitVanish`

- Signature: <span class="api-signature">`func (f *Finder) WaitVanish(pattern *Pattern, timeout time.Duration) (bool, error)`</span>
- Uses: [`Pattern`](#type-pattern)

#### <a id="method-image-clone"></a><span class="api-method">Method</span> `Image.Clone`

- Signature: <span class="api-signature">`func (i *Image) Clone() *Image`</span>

#### <a id="method-image-crop"></a><span class="api-method">Method</span> `Image.Crop`

- Signature: <span class="api-signature">`func (i *Image) Crop(rect Rect) (*Image, error)`</span>
- Uses: [`Rect`](#type-rect)

#### <a id="method-image-gray"></a><span class="api-method">Method</span> `Image.Gray`

- Signature: <span class="api-signature">`func (i *Image) Gray() *image.Gray`</span>

#### <a id="method-image-height"></a><span class="api-method">Method</span> `Image.Height`

- Signature: <span class="api-signature">`func (i *Image) Height() int`</span>

#### <a id="method-image-name"></a><span class="api-method">Method</span> `Image.Name`

- Signature: <span class="api-signature">`func (i *Image) Name() string`</span>

#### <a id="method-image-width"></a><span class="api-method">Method</span> `Image.Width`

- Signature: <span class="api-signature">`func (i *Image) Width() int`</span>

#### <a id="method-inputcontroller-click"></a><span class="api-method">Method</span> `InputController.Click`

- Signature: <span class="api-signature">`func (c *InputController) Click(x, y int, opts InputOptions) error`</span>
- Uses: [`InputOptions`](#type-inputoptions)

#### <a id="method-inputcontroller-hotkey"></a><span class="api-method">Method</span> `InputController.Hotkey`

- Signature: <span class="api-signature">`func (c *InputController) Hotkey(keys ...string) error`</span>

#### <a id="method-inputcontroller-movemouse"></a><span class="api-method">Method</span> `InputController.MoveMouse`

- Signature: <span class="api-signature">`func (c *InputController) MoveMouse(x, y int, opts InputOptions) error`</span>
- Uses: [`InputOptions`](#type-inputoptions)

#### <a id="method-inputcontroller-setbackend"></a><span class="api-method">Method</span> `InputController.SetBackend`

- Signature: <span class="api-signature">`func (c *InputController) SetBackend(backend core.Input)`</span>

#### <a id="method-inputcontroller-typetext"></a><span class="api-method">Method</span> `InputController.TypeText`

- Signature: <span class="api-signature">`func (c *InputController) TypeText(text string, opts InputOptions) error`</span>
- Uses: [`InputOptions`](#type-inputoptions)

#### <a id="method-location-move"></a><span class="api-method">Method</span> `Location.Move`

- Signature: <span class="api-signature">`func (l Location) Move(dx, dy int) Location`</span>

#### <a id="method-location-string"></a><span class="api-method">Method</span> `Location.String`

- Signature: <span class="api-signature">`func (l Location) String() string`</span>

#### <a id="method-location-topoint"></a><span class="api-method">Method</span> `Location.ToPoint`

- Signature: <span class="api-signature">`func (l Location) ToPoint() Point`</span>
- Uses: [`Point`](#type-point)

#### <a id="method-match-string"></a><span class="api-method">Method</span> `Match.String`

- Signature: <span class="api-signature">`func (m Match) String() string`</span>

#### <a id="method-observercontroller-observeappear"></a><span class="api-method">Method</span> `ObserverController.ObserveAppear`

- Signature: <span class="api-signature">`func (c *ObserverController) ObserveAppear(source *Image, region Region, pattern *Pattern, opts ObserveOptions) ([]ObserveEvent, error)`</span>
- Uses: [`Image`](#type-image), [`Region`](#type-region), [`ObserveOptions`](#type-observeoptions), [`ObserveEvent`](#type-observeevent), [`Pattern`](#type-pattern)

#### <a id="method-observercontroller-observechange"></a><span class="api-method">Method</span> `ObserverController.ObserveChange`

- Signature: <span class="api-signature">`func (c *ObserverController) ObserveChange(source *Image, region Region, opts ObserveOptions) ([]ObserveEvent, error)`</span>
- Uses: [`Image`](#type-image), [`Region`](#type-region), [`ObserveOptions`](#type-observeoptions), [`ObserveEvent`](#type-observeevent)

#### <a id="method-observercontroller-observevanish"></a><span class="api-method">Method</span> `ObserverController.ObserveVanish`

- Signature: <span class="api-signature">`func (c *ObserverController) ObserveVanish(source *Image, region Region, pattern *Pattern, opts ObserveOptions) ([]ObserveEvent, error)`</span>
- Uses: [`Image`](#type-image), [`Region`](#type-region), [`ObserveOptions`](#type-observeoptions), [`ObserveEvent`](#type-observeevent), [`Pattern`](#type-pattern)

#### <a id="method-observercontroller-setbackend"></a><span class="api-method">Method</span> `ObserverController.SetBackend`

- Signature: <span class="api-signature">`func (c *ObserverController) SetBackend(backend core.Observer)`</span>

#### <a id="method-offset-string"></a><span class="api-method">Method</span> `Offset.String`

- Signature: <span class="api-signature">`func (o Offset) String() string`</span>

#### <a id="method-offset-topoint"></a><span class="api-method">Method</span> `Offset.ToPoint`

- Signature: <span class="api-signature">`func (o Offset) ToPoint() Point`</span>
- Uses: [`Point`](#type-point)

#### <a id="method-options-clone"></a><span class="api-method">Method</span> `Options.Clone`

- Signature: <span class="api-signature">`func (o *Options) Clone() *Options`</span>

#### <a id="method-options-delete"></a><span class="api-method">Method</span> `Options.Delete`

- Signature: <span class="api-signature">`func (o *Options) Delete(key string)`</span>

#### <a id="method-options-entries"></a><span class="api-method">Method</span> `Options.Entries`

- Signature: <span class="api-signature">`func (o *Options) Entries() map[string]string`</span>

#### <a id="method-options-getbool"></a><span class="api-method">Method</span> `Options.GetBool`

- Signature: <span class="api-signature">`func (o *Options) GetBool(key string, def bool) bool`</span>

#### <a id="method-options-getfloat64"></a><span class="api-method">Method</span> `Options.GetFloat64`

- Signature: <span class="api-signature">`func (o *Options) GetFloat64(key string, def float64) float64`</span>

#### <a id="method-options-getint"></a><span class="api-method">Method</span> `Options.GetInt`

- Signature: <span class="api-signature">`func (o *Options) GetInt(key string, def int) int`</span>

#### <a id="method-options-getstring"></a><span class="api-method">Method</span> `Options.GetString`

- Signature: <span class="api-signature">`func (o *Options) GetString(key, def string) string`</span>

#### <a id="method-options-has"></a><span class="api-method">Method</span> `Options.Has`

- Signature: <span class="api-signature">`func (o *Options) Has(key string) bool`</span>

#### <a id="method-options-merge"></a><span class="api-method">Method</span> `Options.Merge`

- Signature: <span class="api-signature">`func (o *Options) Merge(other *Options)`</span>

#### <a id="method-options-setbool"></a><span class="api-method">Method</span> `Options.SetBool`

- Signature: <span class="api-signature">`func (o *Options) SetBool(key string, value bool)`</span>

#### <a id="method-options-setfloat64"></a><span class="api-method">Method</span> `Options.SetFloat64`

- Signature: <span class="api-signature">`func (o *Options) SetFloat64(key string, value float64)`</span>

#### <a id="method-options-setint"></a><span class="api-method">Method</span> `Options.SetInt`

- Signature: <span class="api-signature">`func (o *Options) SetInt(key string, value int)`</span>

#### <a id="method-options-setstring"></a><span class="api-method">Method</span> `Options.SetString`

- Signature: <span class="api-signature">`func (o *Options) SetString(key, value string)`</span>

#### <a id="method-pattern-exact"></a><span class="api-method">Method</span> `Pattern.Exact`

- Signature: <span class="api-signature">`func (p *Pattern) Exact() *Pattern`</span>

#### <a id="method-pattern-image"></a><span class="api-method">Method</span> `Pattern.Image`

- Signature: <span class="api-signature">`func (p *Pattern) Image() *Image`</span>
- Uses: [`Image`](#type-image)

#### <a id="method-pattern-mask"></a><span class="api-method">Method</span> `Pattern.Mask`

- Signature: <span class="api-signature">`func (p *Pattern) Mask() *image.Gray`</span>

#### <a id="method-pattern-offset"></a><span class="api-method">Method</span> `Pattern.Offset`

- Signature: <span class="api-signature">`func (p *Pattern) Offset() Point`</span>
- Uses: [`Point`](#type-point), [`Offset`](#type-offset)

#### <a id="method-pattern-resize"></a><span class="api-method">Method</span> `Pattern.Resize`

- Signature: <span class="api-signature">`func (p *Pattern) Resize(factor float64) *Pattern`</span>

#### <a id="method-pattern-resizefactor"></a><span class="api-method">Method</span> `Pattern.ResizeFactor`

- Signature: <span class="api-signature">`func (p *Pattern) ResizeFactor() float64`</span>

#### <a id="method-pattern-similar"></a><span class="api-method">Method</span> `Pattern.Similar`

- Signature: <span class="api-signature">`func (p *Pattern) Similar(sim float64) *Pattern`</span>

#### <a id="method-pattern-similarity"></a><span class="api-method">Method</span> `Pattern.Similarity`

- Signature: <span class="api-signature">`func (p *Pattern) Similarity() float64`</span>

#### <a id="method-pattern-targetoffset"></a><span class="api-method">Method</span> `Pattern.TargetOffset`

- Signature: <span class="api-signature">`func (p *Pattern) TargetOffset(dx, dy int) *Pattern`</span>

#### <a id="method-pattern-withmask"></a><span class="api-method">Method</span> `Pattern.WithMask`

- Signature: <span class="api-signature">`func (p *Pattern) WithMask(mask *image.Gray) (*Pattern, error)`</span>

#### <a id="method-pattern-withmaskmatrix"></a><span class="api-method">Method</span> `Pattern.WithMaskMatrix`

- Signature: <span class="api-signature">`func (p *Pattern) WithMaskMatrix(rows [][]uint8) (*Pattern, error)`</span>

#### <a id="method-rect-contains"></a><span class="api-method">Method</span> `Rect.Contains`

- Signature: <span class="api-signature">`func (r Rect) Contains(p Point) bool`</span>
- Uses: [`Point`](#type-point)

#### <a id="method-rect-empty"></a><span class="api-method">Method</span> `Rect.Empty`

- Signature: <span class="api-signature">`func (r Rect) Empty() bool`</span>

#### <a id="method-rect-string"></a><span class="api-method">Method</span> `Rect.String`

- Signature: <span class="api-signature">`func (r Rect) String() string`</span>

#### <a id="method-region-center"></a><span class="api-method">Method</span> `Region.Center`

- Signature: <span class="api-signature">`func (r Region) Center() Point`</span>
- Uses: [`Point`](#type-point)

#### <a id="method-region-contains"></a><span class="api-method">Method</span> `Region.Contains`

- Signature: <span class="api-signature">`func (r Region) Contains(p Point) bool`</span>
- Uses: [`Point`](#type-point)

#### <a id="method-region-containsregion"></a><span class="api-method">Method</span> `Region.ContainsRegion`

- Signature: <span class="api-signature">`func (r Region) ContainsRegion(other Region) bool`</span>

#### <a id="method-region-exists"></a><span class="api-method">Method</span> `Region.Exists`

- Signature: <span class="api-signature">`func (r Region) Exists(source *Image, pattern *Pattern, timeout time.Duration) (Match, bool, error)`</span>
- Uses: [`Image`](#type-image), [`Pattern`](#type-pattern), [`Match`](#type-match)

#### <a id="method-region-find"></a><span class="api-method">Method</span> `Region.Find`

- Signature: <span class="api-signature">`func (r Region) Find(source *Image, pattern *Pattern) (Match, error)`</span>
- Uses: [`Image`](#type-image), [`Pattern`](#type-pattern), [`Match`](#type-match)

#### <a id="method-region-findall"></a><span class="api-method">Method</span> `Region.FindAll`

- Signature: <span class="api-signature">`func (r Region) FindAll(source *Image, pattern *Pattern) ([]Match, error)`</span>
- Uses: [`Image`](#type-image), [`Pattern`](#type-pattern), [`Match`](#type-match)

#### <a id="method-region-findallbycolumn"></a><span class="api-method">Method</span> `Region.FindAllByColumn`

- Signature: <span class="api-signature">`func (r Region) FindAllByColumn(source *Image, pattern *Pattern) ([]Match, error)`</span>
- Uses: [`Image`](#type-image), [`Pattern`](#type-pattern), [`Match`](#type-match)

#### <a id="method-region-findallbyrow"></a><span class="api-method">Method</span> `Region.FindAllByRow`

- Signature: <span class="api-signature">`func (r Region) FindAllByRow(source *Image, pattern *Pattern) ([]Match, error)`</span>
- Uses: [`Image`](#type-image), [`Pattern`](#type-pattern), [`Match`](#type-match)

#### <a id="method-region-findtext"></a><span class="api-method">Method</span> `Region.FindText`

- Signature: <span class="api-signature">`func (r Region) FindText(source *Image, query string, params OCRParams) ([]TextMatch, error)`</span>
- Uses: [`OCRParams`](#type-ocrparams), [`Image`](#type-image), [`TextMatch`](#type-textmatch)

#### <a id="method-region-grow"></a><span class="api-method">Method</span> `Region.Grow`

- Signature: <span class="api-signature">`func (r Region) Grow(dx, dy int) Region`</span>

#### <a id="method-region-has"></a><span class="api-method">Method</span> `Region.Has`

- Signature: <span class="api-signature">`func (r Region) Has(source *Image, pattern *Pattern, timeout time.Duration) (bool, error)`</span>
- Uses: [`Image`](#type-image), [`Pattern`](#type-pattern)

#### <a id="method-region-intersection"></a><span class="api-method">Method</span> `Region.Intersection`

- Signature: <span class="api-signature">`func (r Region) Intersection(other Region) Region`</span>

#### <a id="method-region-moveto"></a><span class="api-method">Method</span> `Region.MoveTo`

- Signature: <span class="api-signature">`func (r Region) MoveTo(x, y int) Region`</span>

#### <a id="method-region-offset"></a><span class="api-method">Method</span> `Region.Offset`

- Signature: <span class="api-signature">`func (r Region) Offset(dx, dy int) Region`</span>
- Uses: [`Offset`](#type-offset)

#### <a id="method-region-readtext"></a><span class="api-method">Method</span> `Region.ReadText`

- Signature: <span class="api-signature">`func (r Region) ReadText(source *Image, params OCRParams) (string, error)`</span>
- Uses: [`OCRParams`](#type-ocrparams), [`Image`](#type-image)

#### <a id="method-region-resetthrowexception"></a><span class="api-method">Method</span> `Region.ResetThrowException`

- Signature: <span class="api-signature">`func (r *Region) ResetThrowException()`</span>

#### <a id="method-region-setautowaittimeout"></a><span class="api-method">Method</span> `Region.SetAutoWaitTimeout`

- Signature: <span class="api-signature">`func (r *Region) SetAutoWaitTimeout(sec float64)`</span>

#### <a id="method-region-setobservescanrate"></a><span class="api-method">Method</span> `Region.SetObserveScanRate`

- Signature: <span class="api-signature">`func (r *Region) SetObserveScanRate(rate float64)`</span>

#### <a id="method-region-setsize"></a><span class="api-method">Method</span> `Region.SetSize`

- Signature: <span class="api-signature">`func (r Region) SetSize(w, h int) Region`</span>

#### <a id="method-region-setthrowexception"></a><span class="api-method">Method</span> `Region.SetThrowException`

- Signature: <span class="api-signature">`func (r *Region) SetThrowException(flag bool)`</span>

#### <a id="method-region-setwaitscanrate"></a><span class="api-method">Method</span> `Region.SetWaitScanRate`

- Signature: <span class="api-signature">`func (r *Region) SetWaitScanRate(rate float64)`</span>

#### <a id="method-region-union"></a><span class="api-method">Method</span> `Region.Union`

- Signature: <span class="api-signature">`func (r Region) Union(other Region) Region`</span>

#### <a id="method-region-wait"></a><span class="api-method">Method</span> `Region.Wait`

- Signature: <span class="api-signature">`func (r Region) Wait(source *Image, pattern *Pattern, timeout time.Duration) (Match, error)`</span>
- Uses: [`Image`](#type-image), [`Pattern`](#type-pattern), [`Match`](#type-match)

#### <a id="method-region-waitvanish"></a><span class="api-method">Method</span> `Region.WaitVanish`

- Signature: <span class="api-signature">`func (r Region) WaitVanish(source *Image, pattern *Pattern, timeout time.Duration) (bool, error)`</span>
- Uses: [`Image`](#type-image), [`Pattern`](#type-pattern)

## Raw Package Doc

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

type AppAPI interface {
	Open(name string, args []string, opts AppOptions) error
	Focus(name string, opts AppOptions) error
	Close(name string, opts AppOptions) error
	IsRunning(name string, opts AppOptions) (bool, error)
	ListWindows(name string, opts AppOptions) ([]Window, error)
}

type AppController struct {
	// Has unexported fields.
}

func NewAppController() *AppController

func (c *AppController) Close(name string, opts AppOptions) error

func (c *AppController) Focus(name string, opts AppOptions) error

func (c *AppController) IsRunning(name string, opts AppOptions) (bool, error)

func (c *AppController) ListWindows(name string, opts AppOptions) ([]Window, error)

func (c *AppController) Open(name string, args []string, opts AppOptions) error

func (c *AppController) SetBackend(backend core.App)

type AppOptions struct {
	Timeout time.Duration
}

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

type ObserveAPI interface {
	ObserveAppear(source *Image, region Region, pattern *Pattern, opts ObserveOptions) ([]ObserveEvent, error)
	ObserveVanish(source *Image, region Region, pattern *Pattern, opts ObserveOptions) ([]ObserveEvent, error)
	ObserveChange(source *Image, region Region, opts ObserveOptions) ([]ObserveEvent, error)
}

type ObserveEvent struct {
	Type      ObserveEventType
	Match     Match
	Timestamp time.Time
}

type ObserveEventType string

const (
	ObserveEventAppear ObserveEventType = "appear"
	ObserveEventVanish ObserveEventType = "vanish"
	ObserveEventChange ObserveEventType = "change"
)
type ObserveOptions struct {
	Interval time.Duration
	Timeout  time.Duration
}

type ObserverController struct {
	// Has unexported fields.
}

func NewObserverController() *ObserverController

func (c *ObserverController) ObserveAppear(source *Image, region Region, pattern *Pattern, opts ObserveOptions) ([]ObserveEvent, error)

func (c *ObserverController) ObserveChange(source *Image, region Region, opts ObserveOptions) ([]ObserveEvent, error)

func (c *ObserverController) ObserveVanish(source *Image, region Region, pattern *Pattern, opts ObserveOptions) ([]ObserveEvent, error)

func (c *ObserverController) SetBackend(backend core.Observer)

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

type Window struct {
	Title   string
	Bounds  Rect
	Focused bool
}

```
