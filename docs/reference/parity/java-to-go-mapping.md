# Java to Go API Mapping

This document is generated from `docs/reference/parity/java-to-go-seed.tsv` and source surfaces in `packages/api/pkg/sikuli/signatures.go` and `packages/api/proto/sikuli/v1/sikuli.proto`.

## Symbol Mapping

| Java/SikuliX Symbol | Go Surface | gRPC RPC | Node API | Python API | Status | Notes |
|---|---|---|---|---|---|---|
| `Pattern.similar(double)` | `(*Pattern).Similar(sim float64)` | `Find, FindOnScreen` | `Pattern().similar()` | `Pattern().similar()` | `parity-ready` | Similarity threshold semantics aligned. |
| `Pattern.exact()` | `(*Pattern).Exact()` | `Find, FindOnScreen` | `Pattern().exact()` | `Pattern().exact()` | `parity-ready` | Exact matching path preserved. |
| `Pattern.targetOffset(dx,dy)` | `(*Pattern).TargetOffset(dx,dy)` | `Find, FindOnScreen` | `Pattern().targetOffset(dx,dy)` | `Pattern().target_offset(dx,dy)` | `parity-ready` | Target-point offset supported. |
| `Pattern.resize(factor)` | `(*Pattern).Resize(factor)` | `Find, FindOnScreen` | `Pattern().resize(factor)` | `Pattern().resize(factor)` | `parity-ready` | Scale intent mapped into matcher request. |
| `Finder.find(Pattern)` | `(*Finder).Find(pattern)` | `Find` | `Finder.find(...) via Region/Screen` | `Region.find(...)` | `parity-ready` | Primary single-target matching flow. |
| `Finder.findAll(Pattern)` | `(*Finder).FindAll(pattern)` | `FindAll` | `Region.findAll(...)` | `Region.find_all(...)` | `parity-ready` | Deterministic ordering documented. |
| `Finder.exists(Pattern, timeout)` | `(*Finder).Exists(pattern,timeout)` | `ExistsOnScreen` | `Region.exists(...)` | `Region.exists(...)` | `parity-ready` | Timeout polling semantics aligned. |
| `Finder.wait(Pattern, timeout)` | `(*Finder).Wait(pattern,timeout)` | `WaitOnScreen` | `Region.wait(...)` | `Region.wait(...)` | `parity-ready` | Wait semantics exposed across clients. |
| `Finder.waitVanish(Pattern, timeout)` | `(*Finder).WaitVanish(pattern,timeout)` | `WaitOnScreen + negative check` | `Region.waitVanish(...)` | `Region.wait_vanish(...)` | `partial` | Client-side vanish wrappers depend on repeated polling behavior. |
| `Region.find(Pattern)` | `(*Region).Find(source,pattern)` | `FindOnScreen` | `Region.find(pattern)` | `Region.find(pattern)` | `parity-ready` | Region-oriented search contract preserved. |
| `Region.click(Pattern)` | `Region + InputController` | `ClickOnScreen` | `Region.click(pattern)` | `Region.click(pattern)` | `parity-ready` | Server-side capture + click orchestration. |
| `Region.hover(Pattern)` | `Region + InputController.MoveMouse` | `FindOnScreen + MoveMouse` | `Region.hover(pattern)` | `Region.hover(pattern)` | `parity-ready` | Hover implemented as find target + move. |
| `Region.type(text)` | `InputController.TypeText` | `TypeText` | `Region.type(text)` | `Region.type_text(text)` | `parity-ready` | Text input mapped to backend input protocol. |
| `Region.readText()` | `(*Region).ReadText(source,params)` | `ReadText` | `Region.readText(...)` | `Region.read_text(...)` | `parity-ready` | OCR read flow supported. |
| `Region.findText(query)` | `(*Region).FindText(source,query,params)` | `FindText` | `Region.findText(...)` | `Region.find_text(...)` | `parity-ready` | OCR search flow supported. |
| `Screen.start()/connect()` | `Sikuli auto/connect constructors` | `Channel + client bootstrap` | `Screen()/Screen.start()/Screen.connect()` | `Screen()/Screen.start()/Screen.connect()` | `parity-ready` | Client constructor patterns standardized. |
| `App.open(name,args)` | `(*AppController).Open(...)` | `OpenApp` | `Sikuli.openApp(...)` | `Sikuli.open_app(...)` | `parity-ready` | App lifecycle support mapped. |
| `App.focus(name)` | `(*AppController).Focus(...)` | `FocusApp` | `Sikuli.focusApp(...)` | `Sikuli.focus_app(...)` | `parity-ready` | Foreground focus support mapped. |
| `App.close(name)` | `(*AppController).Close(...)` | `CloseApp` | `Sikuli.closeApp(...)` | `Sikuli.close_app(...)` | `parity-ready` | Close app support mapped. |
| `App.isRunning(name)` | `(*AppController).IsRunning(...)` | `IsAppRunning` | `Sikuli.isAppRunning(...)` | `Sikuli.is_app_running(...)` | `parity-ready` | Running-state query mapped. |
| `App.window()` | `(*AppController).ListWindows(...)` | `ListWindows` | `Sikuli.listWindows(...)` | `Sikuli.list_windows(...)` | `partial` | Window metadata fields not fully equivalent to Java window model. |
| `Observe.onAppear` | `(*ObserverController).ObserveAppear(...)` | `ObserveAppear` | `Sikuli.observeAppear(...)` | `Sikuli.observe_appear(...)` | `parity-ready` | Polling observer path implemented. |
| `Observe.onVanish` | `(*ObserverController).ObserveVanish(...)` | `ObserveVanish` | `Sikuli.observeVanish(...)` | `Sikuli.observe_vanish(...)` | `parity-ready` | Vanish observer path implemented. |
| `Observe.onChange` | `(*ObserverController).ObserveChange(...)` | `ObserveChange` | `Sikuli.observeChange(...)` | `Sikuli.observe_change(...)` | `parity-ready` | Change observer path implemented. |
| `Region.keyDown()/keyUp()` | `InputController.Hotkey(keys...)` | `Hotkey` | `Sikuli.hotkey(keys)` | `Sikuli.hotkey(keys)` | `gap` | Modifier key stateful keyDown/keyUp model not yet split as dedicated API. |
| `Vision API features` | `internal/cv engine selections` | `Find + matcher_engine` | `engine option` | `engine option` | `partial` | Multiple OpenCV engines available; full SikuliX vision extensions not 1:1. |

## Status Summary

- `parity-ready`: 22
- `partial`: 3
- `gap`: 1

## Go API Interface Surface

Extracted from `packages/api/pkg/sikuli/signatures.go`:

### `ImageAPI`

- `Name() string`
- `Width() int`
- `Height() int`
- `Gray() *image.Gray`
- `Clone() *Image`
- `Crop(rect Rect) (*Image, error)`

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
- `Exists(pattern *Pattern) (Match, bool, error)`
- `Has(pattern *Pattern) (bool, error)`
- `Wait(pattern *Pattern, timeout time.Duration) (Match, error)`
- `WaitVanish(pattern *Pattern, timeout time.Duration) (bool, error)`
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
- `ReadText(source *Image, params OCRParams) (string, error)`
- `FindText(source *Image, query string, params OCRParams) ([]TextMatch, error)`

### `InputAPI`

- `MoveMouse(x, y int, opts InputOptions) error`
- `Click(x, y int, opts InputOptions) error`
- `TypeText(text string, opts InputOptions) error`
- `Hotkey(keys ...string) error`

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


## gRPC Surface

Extracted from `packages/api/proto/sikuli/v1/sikuli.proto`:

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
- `rpc Hotkey(HotkeyRequest) returns (ActionResponse);`
- `rpc ObserveAppear(ObserveRequest) returns (ObserveResponse);`
- `rpc ObserveVanish(ObserveRequest) returns (ObserveResponse);`
- `rpc ObserveChange(ObserveChangeRequest) returns (ObserveResponse);`
- `rpc OpenApp(AppActionRequest) returns (ActionResponse);`
- `rpc FocusApp(AppActionRequest) returns (ActionResponse);`
- `rpc CloseApp(AppActionRequest) returns (ActionResponse);`
- `rpc IsAppRunning(AppActionRequest) returns (IsAppRunningResponse);`
- `rpc ListWindows(AppActionRequest) returns (ListWindowsResponse);`

## Maintenance

- Update the seed file when parity mappings change.
- Run `./scripts/generate-parity-docs.sh` after updates.
- CI verifies this file is up to date.
