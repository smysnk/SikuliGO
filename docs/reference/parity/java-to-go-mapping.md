# Java to Go API Mapping

This document is generated from `docs/reference/parity/java-to-go-seed.tsv` and source surfaces in `packages/api/pkg/sikuli/signatures.go` and `packages/api/proto/sikuli/v1/sikuli.proto`.

## Symbol Mapping

| Java/SikuliX Symbol | Go Surface | gRPC RPC | Node API | Python API | Status | Notes |
|---|---|---|---|---|---|---|
| `Pattern.similar(double)` | `(*Pattern).Similar(sim float64)` | `Find, FindOnScreen` | `Pattern().similar()` | `Pattern().similar()` | `parity-ready` | Similarity threshold semantics aligned. |
| `Pattern.exact()` | `(*Pattern).Exact()` | `Find, FindOnScreen` | `Pattern().exact()` | `Pattern().exact()` | `parity-ready` | Exact matching path preserved. |
| `Pattern.targetOffset(dx,dy)` | `(*Pattern).TargetOffset(dx,dy)` | `Find, FindOnScreen` | `Pattern().targetOffset(dx,dy)` | `Pattern().target_offset(dx,dy)` | `parity-ready` | Target-point offset supported. |
| `Pattern.resize(factor)` | `(*Pattern).Resize(factor)` | `Find, FindOnScreen` | `Pattern().resize(factor)` | `Pattern().resize(factor)` | `parity-ready` | Scale intent mapped into matcher request. |
| `Finder.find(Pattern)` | `(*Finder).Find(pattern), (*Finder).Iterate(pattern)` | `Find` | `Finder.find(...) via Region/Screen` | `Region.find(...)` | `parity-ready` | Go keeps explicit single-match returns and also exposes a no-throw compatibility iterator prep path. |
| `Finder.findAll(Pattern)` | `(*Finder).FindAll(pattern), (*Finder).IterateAll(pattern)` | `FindAll` | `Region.findAll(...)` | `Region.find_all(...)` | `parity-ready` | Deterministic ordering documented, with additive iterator preparation for SikuliX-style traversal. |
| `Finder.hasNext()` | `(*Finder).HasNext()` | `n/a` | `n/a` | `n/a` | `parity-ready` | Compatibility iterator exposes SikuliX-style traversal over the most recent match set. |
| `Finder.next()` | `(*Finder).Next()` | `n/a` | `n/a` | `n/a` | `parity-ready` | Compatibility iterator advances through the prepared match set without raising FindFailed on exhaustion. |
| `Finder.exists(Pattern, timeout)` | `(*Finder).Exists(pattern,timeout)` | `ExistsOnScreen` | `Region.exists(...)` | `Region.exists(...)` | `parity-ready` | Timeout polling semantics aligned. |
| `Finder.wait(Pattern, timeout)` | `(*Finder).Wait(pattern,timeout)` | `WaitOnScreen` | `Region.wait(...)` | `Region.wait(...)` | `parity-ready` | Wait semantics exposed across clients. |
| `Finder.waitVanish(Pattern, timeout)` | `(*Finder).WaitVanish(pattern,timeout)` | `WaitOnScreen + negative check` | `Region.waitVanish(...)` | `Region.wait_vanish(...)` | `partial` | Client-side vanish wrappers depend on repeated polling behavior. |
| `Region.findAnyList(List<Pattern>)` | `(*Finder).FindAnyList(patterns), (*Region).FindAnyList(source,patterns), (LiveRegion).FindAnyList(patterns), (Screen).FindAnyList(patterns)` | `n/a` | `n/a` | `n/a` | `partial` | API-level multi-target helper returns one best match per matched pattern from a single source image or screen capture; client wrappers land in the client parity plan. |
| `Region.findBestList(List<Pattern>)` | `(*Finder).FindBestList(patterns), (*Region).FindBestList(source,patterns), (LiveRegion).FindBestList(patterns), (Screen).FindBestList(patterns)` | `n/a` | `n/a` | `n/a` | `partial` | API-level best-of-list helper selects the highest-score match with deterministic tie-breaking; client wrappers land in the client parity plan. |
| `Region.waitAnyList(List<Pattern>, timeout)` | `(*Finder).WaitAnyList(patterns,timeout), (*Region).WaitAnyList(source,patterns,timeout), (LiveRegion).WaitAnyList(patterns,timeout), (Screen).WaitAnyList(patterns,timeout)` | `n/a` | `n/a` | `n/a` | `partial` | Additive Go wait helper polls one source image or one live capture per cycle until any pattern matches or ErrTimeout. |
| `Region.waitBestList(List<Pattern>, timeout)` | `(*Finder).WaitBestList(patterns,timeout), (*Region).WaitBestList(source,patterns,timeout), (LiveRegion).WaitBestList(patterns,timeout), (Screen).WaitBestList(patterns,timeout)` | `n/a` | `n/a` | `n/a` | `partial` | Additive Go wait helper returns the deterministic best match across the matched pattern list once any candidate appears. |
| `Region.find(Pattern)` | `(*Region).Find(source,pattern)` | `FindOnScreen` | `Region.find(pattern)` | `Region.find(pattern)` | `parity-ready` | Region-oriented search contract preserved. |
| `Region.click(Pattern)` | `Region + InputController` | `ClickOnScreen` | `Region.click(pattern)` | `Region.click(pattern)` | `parity-ready` | Server-side capture + click orchestration. |
| `Region.hover(Pattern)` | `Region + InputController.MoveMouse` | `FindOnScreen + MoveMouse` | `Region.hover(pattern)` | `Region.hover(pattern)` | `parity-ready` | Hover implemented as find target + move. |
| `Region.type(text)` | `InputController.TypeText` | `TypeText` | `Region.type(text)` | `Region.type_text(text)` | `parity-ready` | Text input mapped to backend input protocol. |
| `Region.readText()` | `(*Region).ReadText(source,params)` | `ReadText` | `Region.readText(...)` | `Region.read_text(...)` | `parity-ready` | OCR read flow supported. |
| `Region.findText(query)` | `(*Region).FindText(source,query,params)` | `FindText` | `Region.findText(...)` | `Region.find_text(...)` | `parity-ready` | OCR search flow supported. |
| `Region.collectWords()` | `(*Finder).CollectWords(params), (*Region).CollectWords(source,params), (LiveRegion).CollectWords(params), (Screen).CollectWords(params), (Match).CollectWords(params)` | `n/a` | `n/a` | `n/a` | `partial` | API-level OCR word collection is available across image-backed and live surfaces; client wrappers land in the client parity plan. |
| `Region.collectLines()` | `(*Finder).CollectLines(params), (*Region).CollectLines(source,params), (LiveRegion).CollectLines(params), (Screen).CollectLines(params), (Match).CollectLines(params)` | `n/a` | `n/a` | `n/a` | `partial` | API-level OCR line collection is available across image-backed and live surfaces with stable OCRLine results; client wrappers land in the client parity plan. |
| `Screen.start()/connect()` | `Sikuli auto/connect constructors` | `Channel + client bootstrap` | `Screen()/Screen.start()/Screen.connect()` | `Screen()/Screen.start()/Screen.connect()` | `parity-ready` | Client constructor patterns standardized. |
| `App.open(name,args)` | `(*AppController).Open(...)` | `OpenApp` | `Sikuli.openApp(...)` | `Sikuli.open_app(...)` | `parity-ready` | App lifecycle support mapped. |
| `App.focus(name)` | `(*AppController).Focus(...)` | `FocusApp` | `Sikuli.focusApp(...)` | `Sikuli.focus_app(...)` | `parity-ready` | Foreground focus support mapped. |
| `App.close(name)` | `(*AppController).Close(...)` | `CloseApp` | `Sikuli.closeApp(...)` | `Sikuli.close_app(...)` | `parity-ready` | Close app support mapped. |
| `App.isRunning(name)` | `(*AppController).IsRunning(...)` | `IsAppRunning` | `Sikuli.isAppRunning(...)` | `Sikuli.is_app_running(...)` | `parity-ready` | Running-state query mapped. |
| `App.window()` | `(*AppController).GetWindow(...), (*AppController).FindWindows(...), (*AppController).ListWindows(...)` | `GetWindow, FindWindows, ListWindows` | `Sikuli.listWindows(...)` | `Sikuli.list_windows(...)` | `partial` | API-level window selection helpers now exist; client wrappers still land in the client parity plan and platform metadata remains partially portable. |
| `App.focusedWindow()` | `(*AppController).FocusedWindow(...)` | `GetFocusedWindow` | `n/a` | `n/a` | `partial` | Focused window lookup is first-class in the API surface, with platform-specific metadata variance documented. |
| `App.allWindows()` | `(*AppController).ListWindows(...)` | `ListWindows` | `Sikuli.listWindows(...)` | `Sikuli.list_windows(...)` | `partial` | List-all behavior is stable at the API level; client parity wrappers still land in the client parity plan. |
| `Observe.onAppear` | `(*ObserverController).ObserveAppear(...)` | `ObserveAppear` | `Sikuli.observeAppear(...)` | `Sikuli.observe_appear(...)` | `parity-ready` | Polling observer path implemented. |
| `Observe.onVanish` | `(*ObserverController).ObserveVanish(...)` | `ObserveVanish` | `Sikuli.observeVanish(...)` | `Sikuli.observe_vanish(...)` | `parity-ready` | Vanish observer path implemented. |
| `Observe.onChange` | `(*ObserverController).ObserveChange(...)` | `ObserveChange` | `Sikuli.observeChange(...)` | `Sikuli.observe_change(...)` | `parity-ready` | Change observer path implemented. |
| `Region.keyDown()/keyUp()` | `InputController.KeyDown/KeyUp` | `KeyDown, KeyUp` | `Sikuli.keyDown/keyUp(keys)` | `Sikuli.key_down/key_up(keys)` | `parity-ready` | Stateful key transitions are exposed as dedicated API and protocol operations. |
| `Vision API features` | `internal/cv engine selections` | `Find + matcher_engine` | `engine option` | `engine option` | `partial` | Multiple OpenCV engines available; full SikuliX vision extensions not 1:1. |

## Status Summary

- `parity-ready`: 25
- `partial`: 11
- `gap`: 0

## Go API Interface Surface

Extracted from `packages/api/pkg/sikuli/signatures.go`:

### `ImageAPI`

- `Name() string`
- `Width() int`
- `Height() int`
- `Gray() *image.Gray`
- `Clone() *Image`
- `Crop(rect Rect) (*Image, error)`

### `TargetPointProvider`

- `TargetPoint() Point`

### `PatternAPI`

- `Image() *Image`
- `Similar(sim float64) *Pattern`
- `Similarity() float64`
- `Exact() *Pattern`
- `TargetOffset(dx, dy int) *Pattern`
- `Offset() Point`
- `Resize(factor float64) *Pattern`
- `ResizeFactor() float64`
- `Mask() *image.Gray`

### `FinderAPI`

- `Find(pattern *Pattern) (Match, error)`
- `FindAll(pattern *Pattern) ([]Match, error)`
- `FindAllByRow(pattern *Pattern) ([]Match, error)`
- `FindAllByColumn(pattern *Pattern) ([]Match, error)`
- `FindAnyList(patterns []*Pattern) ([]Match, error)`
- `FindBestList(patterns []*Pattern) (Match, error)`
- `Iterate(pattern *Pattern) error`
- `IterateAll(pattern *Pattern) error`
- `Exists(pattern *Pattern) (Match, bool, error)`
- `Has(pattern *Pattern) (bool, error)`
- `Wait(pattern *Pattern, timeout time.Duration) (Match, error)`
- `WaitAnyList(patterns []*Pattern, timeout time.Duration) ([]Match, error)`
- `WaitBestList(patterns []*Pattern, timeout time.Duration) (Match, error)`
- `WaitVanish(pattern *Pattern, timeout time.Duration) (bool, error)`
- `CollectWords(params OCRParams) ([]OCRWord, error)`
- `CollectLines(params OCRParams) ([]OCRLine, error)`
- `HasNext() bool`
- `Next() (Match, bool)`
- `Reset()`
- `Destroy()`
- `ReadText(params OCRParams) (string, error)`
- `FindText(query string, params OCRParams) ([]TextMatch, error)`
- `LastMatches() []Match`

### `RegionAPI`

- `Center() Point`
- `Grow(dx, dy int) Region`
- `Offset(dx, dy int) Region`
- `MoveTo(x, y int) Region`
- `SetSize(w, h int) Region`
- `Contains(p Point) bool`
- `ContainsRegion(other Region) bool`
- `Union(other Region) Region`
- `Intersection(other Region) Region`
- `Find(source *Image, pattern *Pattern) (Match, error)`
- `Exists(source *Image, pattern *Pattern, timeout time.Duration) (Match, bool, error)`
- `Has(source *Image, pattern *Pattern, timeout time.Duration) (bool, error)`
- `Wait(source *Image, pattern *Pattern, timeout time.Duration) (Match, error)`
- `WaitVanish(source *Image, pattern *Pattern, timeout time.Duration) (bool, error)`
- `FindAll(source *Image, pattern *Pattern) ([]Match, error)`
- `FindAllByRow(source *Image, pattern *Pattern) ([]Match, error)`
- `FindAllByColumn(source *Image, pattern *Pattern) ([]Match, error)`
- `FindAnyList(source *Image, patterns []*Pattern) ([]Match, error)`
- `FindBestList(source *Image, patterns []*Pattern) (Match, error)`
- `WaitAnyList(source *Image, patterns []*Pattern, timeout time.Duration) ([]Match, error)`
- `WaitBestList(source *Image, patterns []*Pattern, timeout time.Duration) (Match, error)`
- `ReadText(source *Image, params OCRParams) (string, error)`
- `FindText(source *Image, query string, params OCRParams) ([]TextMatch, error)`
- `CollectWords(source *Image, params OCRParams) ([]OCRWord, error)`
- `CollectLines(source *Image, params OCRParams) ([]OCRLine, error)`

### `LiveRegionAPI`

- `Bounds() Region`
- `Center() Point`
- `TargetPoint() Point`
- `Grow(dx, dy int) LiveRegion`
- `Offset(dx, dy int) LiveRegion`
- `MoveTo(x, y int) LiveRegion`
- `SetSize(w, h int) LiveRegion`
- `WithMatcherEngine(engine MatcherEngine) LiveRegion`
- `Capture() (*Image, error)`
- `Find(pattern *Pattern) (Match, error)`
- `FindAnyList(patterns []*Pattern) ([]Match, error)`
- `FindBestList(patterns []*Pattern) (Match, error)`
- `Exists(pattern *Pattern, timeout time.Duration) (Match, bool, error)`
- `Has(pattern *Pattern, timeout time.Duration) (bool, error)`
- `Wait(pattern *Pattern, timeout time.Duration) (Match, error)`
- `WaitAnyList(patterns []*Pattern, timeout time.Duration) ([]Match, error)`
- `WaitBestList(patterns []*Pattern, timeout time.Duration) (Match, error)`
- `WaitVanish(pattern *Pattern, timeout time.Duration) (bool, error)`
- `ReadText(params OCRParams) (string, error)`
- `FindText(query string, params OCRParams) ([]TextMatch, error)`
- `CollectWords(params OCRParams) ([]OCRWord, error)`
- `CollectLines(params OCRParams) ([]OCRLine, error)`
- `Hover(opts InputOptions) error`
- `Click(opts InputOptions) error`
- `RightClick(opts InputOptions) error`
- `DoubleClick(opts InputOptions) error`
- `MouseDown(opts InputOptions) error`
- `MouseUp(opts InputOptions) error`
- `TypeText(text string, opts InputOptions) error`
- `Paste(text string, opts InputOptions) error`
- `DragDrop(target TargetPointProvider, opts InputOptions) error`
- `Wheel(direction WheelDirection, steps int, opts InputOptions) error`
- `KeyDown(keys ...string) error`
- `KeyUp(keys ...string) error`

### `MatchAPI`

- `Bounds() Region`
- `Region() Region`
- `Center() Point`
- `TargetPoint() Point`
- `Live() bool`
- `Capture() (*Image, error)`
- `Find(pattern *Pattern) (Match, error)`
- `FindAnyList(patterns []*Pattern) ([]Match, error)`
- `FindBestList(patterns []*Pattern) (Match, error)`
- `Exists(pattern *Pattern, timeout time.Duration) (Match, bool, error)`
- `Has(pattern *Pattern, timeout time.Duration) (bool, error)`
- `Wait(pattern *Pattern, timeout time.Duration) (Match, error)`
- `WaitAnyList(patterns []*Pattern, timeout time.Duration) ([]Match, error)`
- `WaitBestList(patterns []*Pattern, timeout time.Duration) (Match, error)`
- `WaitVanish(pattern *Pattern, timeout time.Duration) (bool, error)`
- `ReadText(params OCRParams) (string, error)`
- `FindText(query string, params OCRParams) ([]TextMatch, error)`
- `CollectWords(params OCRParams) ([]OCRWord, error)`
- `CollectLines(params OCRParams) ([]OCRLine, error)`
- `MoveMouse(opts InputOptions) error`
- `Hover(opts InputOptions) error`
- `Click(opts InputOptions) error`
- `RightClick(opts InputOptions) error`
- `DoubleClick(opts InputOptions) error`
- `MouseDown(opts InputOptions) error`
- `MouseUp(opts InputOptions) error`
- `TypeText(text string, opts InputOptions) error`
- `Paste(text string, opts InputOptions) error`
- `DragDrop(target TargetPointProvider, opts InputOptions) error`
- `Wheel(direction WheelDirection, steps int, opts InputOptions) error`
- `KeyDown(keys ...string) error`
- `KeyUp(keys ...string) error`

### `ScreenAPI`

- `Live() bool`
- `TargetPoint() Point`
- `FullRegion() LiveRegion`
- `Region(x, y, w, h int) LiveRegion`
- `RegionRect(rect Rect) LiveRegion`
- `Capture() (*Image, error)`
- `Find(pattern *Pattern) (Match, error)`
- `FindAnyList(patterns []*Pattern) ([]Match, error)`
- `FindBestList(patterns []*Pattern) (Match, error)`
- `Exists(pattern *Pattern, timeout time.Duration) (Match, bool, error)`
- `Has(pattern *Pattern, timeout time.Duration) (bool, error)`
- `Wait(pattern *Pattern, timeout time.Duration) (Match, error)`
- `WaitAnyList(patterns []*Pattern, timeout time.Duration) ([]Match, error)`
- `WaitBestList(patterns []*Pattern, timeout time.Duration) (Match, error)`
- `WaitVanish(pattern *Pattern, timeout time.Duration) (bool, error)`
- `ReadText(params OCRParams) (string, error)`
- `FindText(query string, params OCRParams) ([]TextMatch, error)`
- `CollectWords(params OCRParams) ([]OCRWord, error)`
- `CollectLines(params OCRParams) ([]OCRLine, error)`
- `Hover(opts InputOptions) error`
- `Click(opts InputOptions) error`
- `RightClick(opts InputOptions) error`
- `DoubleClick(opts InputOptions) error`
- `MouseDown(opts InputOptions) error`
- `MouseUp(opts InputOptions) error`
- `TypeText(text string, opts InputOptions) error`
- `Paste(text string, opts InputOptions) error`
- `DragDrop(target TargetPointProvider, opts InputOptions) error`
- `Wheel(direction WheelDirection, steps int, opts InputOptions) error`
- `KeyDown(keys ...string) error`
- `KeyUp(keys ...string) error`

### `RuntimeAPI`

- `Address() string`
- `Close() error`
- `Screens() ([]Screen, error)`
- `PrimaryScreen() (Screen, error)`
- `Screen(id int) (Screen, error)`
- `Capture() (*Image, error)`
- `CaptureRegion(region Region) (*Image, error)`
- `Region(region Region) LiveRegion`

### `InputAPI`

- `MoveMouse(x, y int, opts InputOptions) error`
- `Hover(x, y int, opts InputOptions) error`
- `Click(x, y int, opts InputOptions) error`
- `RightClick(x, y int, opts InputOptions) error`
- `DoubleClick(x, y int, opts InputOptions) error`
- `MouseDown(x, y int, opts InputOptions) error`
- `MouseUp(x, y int, opts InputOptions) error`
- `TypeText(text string, opts InputOptions) error`
- `Paste(text string, opts InputOptions) error`
- `Hotkey(keys ...string) error`
- `KeyDown(keys ...string) error`
- `KeyUp(keys ...string) error`
- `Wheel(x, y int, direction WheelDirection, steps int, opts InputOptions) error`
- `DragDrop(fromX, fromY, toX, toY int, opts InputOptions) error`

### `ObserveAPI`

- `ObserveAppear(source *Image, region Region, pattern *Pattern, opts ObserveOptions) ([]ObserveEvent, error)`
- `ObserveVanish(source *Image, region Region, pattern *Pattern, opts ObserveOptions) ([]ObserveEvent, error)`
- `ObserveChange(source *Image, region Region, opts ObserveOptions) ([]ObserveEvent, error)`

### `AppAPI`

- `Open(name string, args []string, opts AppOptions) error`
- `Focus(name string, opts AppOptions) error`
- `Close(name string, opts AppOptions) error`
- `IsRunning(name string, opts AppOptions) (bool, error)`
- `ListWindows(name string, opts AppOptions) ([]Window, error)`
- `FindWindows(name string, query WindowQuery, opts AppOptions) ([]Window, error)`
- `GetWindow(name string, query WindowQuery, opts AppOptions) (Window, bool, error)`
- `FocusedWindow(name string, opts AppOptions) (Window, bool, error)`


## gRPC Surface

Extracted from `packages/api/proto/sikuli/v1/sikuli.proto`:

- `rpc ListScreens(ListScreensRequest) returns (ListScreensResponse);`
- `rpc GetPrimaryScreen(GetPrimaryScreenRequest) returns (GetPrimaryScreenResponse);`
- `rpc CaptureScreen(CaptureScreenRequest) returns (CaptureScreenResponse);`
- `rpc Find(FindRequest) returns (FindResponse);`
- `rpc FindAll(FindRequest) returns (FindAllResponse);`
- `rpc FindOnScreen(FindOnScreenRequest) returns (FindResponse);`
- `rpc ExistsOnScreen(ExistsOnScreenRequest) returns (ExistsOnScreenResponse);`
- `rpc WaitOnScreen(WaitOnScreenRequest) returns (FindResponse);`
- `rpc ClickOnScreen(ClickOnScreenRequest) returns (FindResponse);`
- `rpc ReadText(ReadTextRequest) returns (ReadTextResponse);`
- `rpc FindText(FindTextRequest) returns (FindTextResponse);`
- `rpc MoveMouse(MoveMouseRequest) returns (ActionResponse);`
- `rpc Click(ClickRequest) returns (ActionResponse);`
- `rpc TypeText(TypeTextRequest) returns (ActionResponse);`
- `rpc PasteText(TypeTextRequest) returns (ActionResponse);`
- `rpc Hotkey(HotkeyRequest) returns (ActionResponse);`
- `rpc MouseDown(ClickRequest) returns (ActionResponse);`
- `rpc MouseUp(ClickRequest) returns (ActionResponse);`
- `rpc KeyDown(HotkeyRequest) returns (ActionResponse);`
- `rpc KeyUp(HotkeyRequest) returns (ActionResponse);`
- `rpc ScrollWheel(ScrollWheelRequest) returns (ActionResponse);`
- `rpc ObserveAppear(ObserveRequest) returns (ObserveResponse);`
- `rpc ObserveVanish(ObserveRequest) returns (ObserveResponse);`
- `rpc ObserveChange(ObserveChangeRequest) returns (ObserveResponse);`
- `rpc OpenApp(AppActionRequest) returns (ActionResponse);`
- `rpc FocusApp(AppActionRequest) returns (ActionResponse);`
- `rpc CloseApp(AppActionRequest) returns (ActionResponse);`
- `rpc IsAppRunning(AppActionRequest) returns (IsAppRunningResponse);`
- `rpc ListWindows(AppActionRequest) returns (ListWindowsResponse);`
- `rpc FindWindows(WindowQueryRequest) returns (ListWindowsResponse);`
- `rpc GetWindow(WindowQueryRequest) returns (GetWindowResponse);`
- `rpc GetFocusedWindow(AppActionRequest) returns (GetWindowResponse);`

## Maintenance

- Update the seed file when parity mappings change.
- Run `./scripts/generate-parity-docs.sh` after updates.
- CI verifies this file is up to date.
