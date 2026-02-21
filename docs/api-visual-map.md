# API Visual Map

<style>
  .api-class { color: #0f766e; font-weight: 700; }
  .api-method { color: #7c3aed; font-weight: 700; }
  .api-protocol { color: #b45309; font-weight: 700; }
  .api-value { color: #1d4ed8; font-weight: 700; }
  .api-link { color: #be123c; font-weight: 700; }
  .api-note { color: #374151; }
</style>

This page augments the generated API reference with color-coded type/method groupings and cross-links where types are used.

## Legend

- <span class="api-class">Class/Controller</span>
- <span class="api-method">Method</span>
- <span class="api-protocol">Protocol/Interface</span>
- <span class="api-value">Value/Object Type</span>
- <span class="api-link">Cross-reference link</span>

## Navigation

- [Public API package (`pkg/sikuli`)](https://smysnk.github.io/SikuliGO/api/pkg-sikuli)
- [Core protocol package (`internal/core`)](https://smysnk.github.io/SikuliGO/api/internal-core)
- [Matcher backend (`internal/cv`)](https://smysnk.github.io/SikuliGO/api/internal-cv)
- [OCR backend (`internal/ocr`)](https://smysnk.github.io/SikuliGO/api/internal-ocr)
- [Input backend (`internal/input`)](https://smysnk.github.io/SikuliGO/api/internal-input)
- [Observe backend (`internal/observe`)](https://smysnk.github.io/SikuliGO/api/internal-observe)
- [App backend (`internal/app`)](https://smysnk.github.io/SikuliGO/api/internal-app)

## Type Index

| Type | Reference |
|---|---|
| <span class="api-class">Finder</span> | [`pkg/sikuli.Finder`](https://smysnk.github.io/SikuliGO/api/pkg-sikuli/#finder) |
| <span class="api-class">Region</span> | [`pkg/sikuli.Region`](https://smysnk.github.io/SikuliGO/api/pkg-sikuli/#region) |
| <span class="api-class">InputController</span> | [`pkg/sikuli.InputController`](https://smysnk.github.io/SikuliGO/api/pkg-sikuli/#inputcontroller) |
| <span class="api-class">ObserverController</span> | [`pkg/sikuli.ObserverController`](https://smysnk.github.io/SikuliGO/api/pkg-sikuli/#observercontroller) |
| <span class="api-class">AppController</span> | [`pkg/sikuli.AppController`](https://smysnk.github.io/SikuliGO/api/pkg-sikuli/#appcontroller) |
| <span class="api-value">Pattern</span> | [`pkg/sikuli.Pattern`](https://smysnk.github.io/SikuliGO/api/pkg-sikuli/#pattern) |
| <span class="api-value">Match</span> | [`pkg/sikuli.Match`](https://smysnk.github.io/SikuliGO/api/pkg-sikuli/#match) |
| <span class="api-value">OCRParams</span> | [`pkg/sikuli.OCRParams`](https://smysnk.github.io/SikuliGO/api/pkg-sikuli/#ocrparams) |
| <span class="api-value">ObserveEvent</span> | [`pkg/sikuli.ObserveEvent`](https://smysnk.github.io/SikuliGO/api/pkg-sikuli/#observeevent) |
| <span class="api-value">Window</span> | [`pkg/sikuli.Window`](https://smysnk.github.io/SikuliGO/api/pkg-sikuli/#window) |

## Public Types

### <span class="api-class">Finder</span>

<span class="api-method">Methods</span>:
- `Find(pattern *[Pattern](#pattern)) ([Match](#match), error)`
- `FindAll(pattern *[Pattern](#pattern)) ([][Match](#match), error)`
- `FindAllByRow(pattern *[Pattern](#pattern)) ([][Match](#match), error)`
- `FindAllByColumn(pattern *[Pattern](#pattern)) ([][Match](#match), error)`
- `Exists(pattern *[Pattern](#pattern)) ([Match](#match), bool, error)`
- `Has(pattern *[Pattern](#pattern)) (bool, error)`
- `Wait(pattern *[Pattern](#pattern), timeout time.Duration) ([Match](#match), error)`
- `WaitVanish(pattern *[Pattern](#pattern), timeout time.Duration) (bool, error)`
- `ReadText(params [OCRParams](#ocrparams)) (string, error)`
- `FindText(query string, params [OCRParams](#ocrparams)) ([][TextMatch](#textmatch), error)`

<span class="api-note">Uses types</span>:
- [`Pattern`](#pattern), [`Match`](#match), [`OCRParams`](#ocrparams), [`TextMatch`](#textmatch)

### <span class="api-class">Region</span>

<span class="api-method">Methods</span>:
- Geometry: `Center`, `Grow`, `Offset`, `MoveTo`, `SetSize`, `Contains`, `ContainsRegion`, `Union`, `Intersection`
- Image search: `Find`, `Exists`, `Has`, `Wait`, `WaitVanish`, `FindAll`, `FindAllByRow`, `FindAllByColumn`
- OCR/text: `ReadText`, `FindText`

<span class="api-note">Uses types</span>:
- [`Rect`](#rect), [`Point`](#point), [`Pattern`](#pattern), [`Match`](#match), [`OCRParams`](#ocrparams), [`TextMatch`](#textmatch)

### <span class="api-class">InputController</span>

<span class="api-method">Methods</span>:
- `MoveMouse(x, y int, opts [InputOptions](#inputoptions)) error`
- `Click(x, y int, opts [InputOptions](#inputoptions)) error`
- `TypeText(text string, opts [InputOptions](#inputoptions)) error`
- `Hotkey(keys ...string) error`

<span class="api-note">Uses types</span>:
- [`InputOptions`](#inputoptions), [`MouseButton`](#mousebutton)

### <span class="api-class">ObserverController</span>

<span class="api-method">Methods</span>:
- `ObserveAppear(source *[Image](#image), region [Region](#region), pattern *[Pattern](#pattern), opts [ObserveOptions](#observeoptions)) ([][ObserveEvent](#observeevent), error)`
- `ObserveVanish(source *[Image](#image), region [Region](#region), pattern *[Pattern](#pattern), opts [ObserveOptions](#observeoptions)) ([][ObserveEvent](#observeevent), error)`
- `ObserveChange(source *[Image](#image), region [Region](#region), opts [ObserveOptions](#observeoptions)) ([][ObserveEvent](#observeevent), error)`

<span class="api-note">Uses types</span>:
- [`Image`](#image), [`Region`](#region), [`Pattern`](#pattern), [`ObserveOptions`](#observeoptions), [`ObserveEvent`](#observeevent)

### <span class="api-class">AppController</span>

<span class="api-method">Methods</span>:
- `Open(name string, args []string, opts [AppOptions](#appoptions)) error`
- `Focus(name string, opts [AppOptions](#appoptions)) error`
- `Close(name string, opts [AppOptions](#appoptions)) error`
- `IsRunning(name string, opts [AppOptions](#appoptions)) (bool, error)`
- `ListWindows(name string, opts [AppOptions](#appoptions)) ([][Window](#window), error)`

<span class="api-note">Uses types</span>:
- [`AppOptions`](#appoptions), [`Window`](#window)

## Value/Object Types

### <span class="api-value">Point</span>
- Coordinate pair used by [`Rect`](#rect), [`Region`](#region), [`Match`](#match), [`Location`](#location), [`Offset`](#offset)

### <span class="api-value">Rect</span>
- Geometry primitive used by [`Region`](#region), [`Match`](#match), [`Window`](#window)

### <span class="api-value">Location</span>
- User-facing coordinate wrapper; converts to [`Point`](#point)

### <span class="api-value">Offset</span>
- Target offset wrapper; converts to [`Point`](#point)

### <span class="api-value">Image</span>
- Source container used by [`Pattern`](#pattern), [`Finder`](#finder), [`Region`](#region), [`ObserverController`](#observercontroller)

### <span class="api-value">Pattern</span>
- Search target definition used by [`Finder`](#finder), [`Region`](#region), [`ObserverController`](#observercontroller)

### <span class="api-value">Match</span>
- Image match result used by [`Finder`](#finder), [`Region`](#region), [`ObserveEvent`](#observeevent)

### <span class="api-value">TextMatch</span>
- OCR text match result used by [`Finder`](#finder), [`Region`](#region)

### <span class="api-value">OCRParams</span>
- OCR options used by [`Finder`](#finder), [`Region`](#region)

### <span class="api-value">InputOptions</span>
- Input action options used by [`InputController`](#inputcontroller)

### <span class="api-value">MouseButton</span>
- Mouse button enum used by [`InputOptions`](#inputoptions)

### <span class="api-value">ObserveOptions</span>
- Observe action options used by [`ObserverController`](#observercontroller)

### <span class="api-value">ObserveEvent</span>
- Observe payload type returned by [`ObserverController`](#observercontroller)

### <span class="api-value">AppOptions</span>
- App action options used by [`AppController`](#appcontroller)

### <span class="api-value">Window</span>
- Window payload returned by [`AppController`](#appcontroller)

## Usage Matrix

Legend: ✅ strongly used in method signatures or primary flows.

| Type | Finder | Region | InputController | ObserverController | AppController |
|---|---|---|---|---|---|
| [Pattern](#pattern) | ✅ | ✅ |  | ✅ |  |
| [Match](#match) | ✅ | ✅ |  |  |  |
| [OCRParams](#ocrparams) | ✅ | ✅ |  |  |  |
| [TextMatch](#textmatch) | ✅ | ✅ |  |  |  |
| [InputOptions](#inputoptions) |  |  | ✅ |  |  |
| [ObserveOptions](#observeoptions) |  |  |  | ✅ |  |
| [ObserveEvent](#observeevent) |  |  |  | ✅ |  |
| [AppOptions](#appoptions) |  |  |  |  | ✅ |
| [Window](#window) |  |  |  |  | ✅ |

## Protocol Map

- <span class="api-protocol">`core.Matcher`</span> consumed by [`Finder`](#finder) ([internal API](https://smysnk.github.io/SikuliGO/api/internal-core/#matcher))
- <span class="api-protocol">`core.OCR`</span> consumed by [`Finder`](#finder) ([internal API](https://smysnk.github.io/SikuliGO/api/internal-core/#ocr))
- <span class="api-protocol">`core.Input`</span> consumed by [`InputController`](#inputcontroller) ([internal API](https://smysnk.github.io/SikuliGO/api/internal-core/#input))
- <span class="api-protocol">`core.Observer`</span> consumed by [`ObserverController`](#observercontroller) ([internal API](https://smysnk.github.io/SikuliGO/api/internal-core/#observer))
- <span class="api-protocol">`core.App`</span> consumed by [`AppController`](#appcontroller) ([internal API](https://smysnk.github.io/SikuliGO/api/internal-core/#app))
