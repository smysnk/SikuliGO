# API: `internal/core`

[Back to API Index](./)

<style>
  .api-type { color: #0f766e; font-weight: 700; }
  .api-func { color: #1d4ed8; font-weight: 700; }
  .api-method { color: #7c3aed; font-weight: 700; }
  .api-signature { font-family: ui-monospace, SFMono-Regular, Menlo, Consolas, "Liberation Mono", monospace; }
</style>

Legend: <span class="api-type">Type</span>, <span class="api-func">Function</span>, <span class="api-method">Method</span>

Package: `package core // import "github.com/sikulix/portgo/internal/core"`

## Symbol Index

### Types

- <span class="api-type">[`App`](#type-app)</span>
- <span class="api-type">[`AppAction`](#type-appaction)</span>
- <span class="api-type">[`AppRequest`](#type-apprequest)</span>
- <span class="api-type">[`AppResult`](#type-appresult)</span>
- <span class="api-type">[`Input`](#type-input)</span>
- <span class="api-type">[`InputAction`](#type-inputaction)</span>
- <span class="api-type">[`InputRequest`](#type-inputrequest)</span>
- <span class="api-type">[`MatchCandidate`](#type-matchcandidate)</span>
- <span class="api-type">[`Matcher`](#type-matcher)</span>
- <span class="api-type">[`OCR`](#type-ocr)</span>
- <span class="api-type">[`OCRRequest`](#type-ocrrequest)</span>
- <span class="api-type">[`OCRResult`](#type-ocrresult)</span>
- <span class="api-type">[`OCRWord`](#type-ocrword)</span>
- <span class="api-type">[`ObserveEvent`](#type-observeevent)</span>
- <span class="api-type">[`ObserveEventType`](#type-observeeventtype)</span>
- <span class="api-type">[`ObserveRequest`](#type-observerequest)</span>
- <span class="api-type">[`Observer`](#type-observer)</span>
- <span class="api-type">[`SearchRequest`](#type-searchrequest)</span>
- <span class="api-type">[`WindowInfo`](#type-windowinfo)</span>

### Functions

- <span class="api-func">[`ResizeGrayNearest`](#func-resizegraynearest)</span>

### Methods

- <span class="api-method">[`AppRequest.Validate`](#method-apprequest-validate)</span>
- <span class="api-method">[`InputRequest.Validate`](#method-inputrequest-validate)</span>
- <span class="api-method">[`OCRRequest.Validate`](#method-ocrrequest-validate)</span>
- <span class="api-method">[`ObserveRequest.Validate`](#method-observerequest-validate)</span>
- <span class="api-method">[`SearchRequest.Validate`](#method-searchrequest-validate)</span>

## Declarations

### Types

#### <a id="type-app"></a><span class="api-type">Type</span> `App`

- Signature: <span class="api-signature">`type App interface {`</span>

#### <a id="type-appaction"></a><span class="api-type">Type</span> `AppAction`

- Signature: <span class="api-signature">`type AppAction string`</span>

#### <a id="type-apprequest"></a><span class="api-type">Type</span> `AppRequest`

- Signature: <span class="api-signature">`type AppRequest struct {`</span>

#### <a id="type-appresult"></a><span class="api-type">Type</span> `AppResult`

- Signature: <span class="api-signature">`type AppResult struct {`</span>

#### <a id="type-input"></a><span class="api-type">Type</span> `Input`

- Signature: <span class="api-signature">`type Input interface {`</span>

#### <a id="type-inputaction"></a><span class="api-type">Type</span> `InputAction`

- Signature: <span class="api-signature">`type InputAction string`</span>

#### <a id="type-inputrequest"></a><span class="api-type">Type</span> `InputRequest`

- Signature: <span class="api-signature">`type InputRequest struct {`</span>

#### <a id="type-matchcandidate"></a><span class="api-type">Type</span> `MatchCandidate`

- Signature: <span class="api-signature">`type MatchCandidate struct {`</span>

#### <a id="type-matcher"></a><span class="api-type">Type</span> `Matcher`

- Signature: <span class="api-signature">`type Matcher interface {`</span>

#### <a id="type-ocr"></a><span class="api-type">Type</span> `OCR`

- Signature: <span class="api-signature">`type OCR interface {`</span>

#### <a id="type-ocrrequest"></a><span class="api-type">Type</span> `OCRRequest`

- Signature: <span class="api-signature">`type OCRRequest struct {`</span>

#### <a id="type-ocrresult"></a><span class="api-type">Type</span> `OCRResult`

- Signature: <span class="api-signature">`type OCRResult struct {`</span>

#### <a id="type-ocrword"></a><span class="api-type">Type</span> `OCRWord`

- Signature: <span class="api-signature">`type OCRWord struct {`</span>

#### <a id="type-observeevent"></a><span class="api-type">Type</span> `ObserveEvent`

- Signature: <span class="api-signature">`type ObserveEvent struct {`</span>

#### <a id="type-observeeventtype"></a><span class="api-type">Type</span> `ObserveEventType`

- Signature: <span class="api-signature">`type ObserveEventType string`</span>

#### <a id="type-observerequest"></a><span class="api-type">Type</span> `ObserveRequest`

- Signature: <span class="api-signature">`type ObserveRequest struct {`</span>

#### <a id="type-observer"></a><span class="api-type">Type</span> `Observer`

- Signature: <span class="api-signature">`type Observer interface {`</span>

#### <a id="type-searchrequest"></a><span class="api-type">Type</span> `SearchRequest`

- Signature: <span class="api-signature">`type SearchRequest struct {`</span>

#### <a id="type-windowinfo"></a><span class="api-type">Type</span> `WindowInfo`

- Signature: <span class="api-signature">`type WindowInfo struct {`</span>

### Functions

#### <a id="func-resizegraynearest"></a><span class="api-func">Function</span> `ResizeGrayNearest`

- Signature: <span class="api-signature">`func ResizeGrayNearest(src *image.Gray, factor float64) *image.Gray`</span>

### Methods

#### <a id="method-apprequest-validate"></a><span class="api-method">Method</span> `AppRequest.Validate`

- Signature: <span class="api-signature">`func (r AppRequest) Validate() error`</span>

#### <a id="method-inputrequest-validate"></a><span class="api-method">Method</span> `InputRequest.Validate`

- Signature: <span class="api-signature">`func (r InputRequest) Validate() error`</span>

#### <a id="method-ocrrequest-validate"></a><span class="api-method">Method</span> `OCRRequest.Validate`

- Signature: <span class="api-signature">`func (r OCRRequest) Validate() error`</span>

#### <a id="method-observerequest-validate"></a><span class="api-method">Method</span> `ObserveRequest.Validate`

- Signature: <span class="api-signature">`func (r ObserveRequest) Validate() error`</span>

#### <a id="method-searchrequest-validate"></a><span class="api-method">Method</span> `SearchRequest.Validate`

- Signature: <span class="api-signature">`func (r SearchRequest) Validate() error`</span>

## Raw Package Doc

```text
package core // import "github.com/sikulix/portgo/internal/core"


VARIABLES

var ErrAppUnsupported = errors.New("app backend unsupported")
var ErrInputUnsupported = errors.New("input backend unsupported")
var ErrOCRUnsupported = errors.New("ocr backend unsupported")
var ErrObserveUnsupported = errors.New("observe backend unsupported")

FUNCTIONS

func ResizeGrayNearest(src *image.Gray, factor float64) *image.Gray

TYPES

type App interface {
	Execute(req AppRequest) (AppResult, error)
}

type AppAction string

const (
	AppActionOpen       AppAction = "open"
	AppActionFocus      AppAction = "focus"
	AppActionClose      AppAction = "close"
	AppActionIsRunning  AppAction = "is_running"
	AppActionListWindow AppAction = "list_windows"
)
type AppRequest struct {
	Action  AppAction
	Name    string
	Args    []string
	Timeout time.Duration
	Options map[string]string
}

func (r AppRequest) Validate() error

type AppResult struct {
	Running bool
	PID     int
	Windows []WindowInfo
}

type Input interface {
	Execute(req InputRequest) error
}

type InputAction string

const (
	InputActionMouseMove InputAction = "mouse_move"
	InputActionClick     InputAction = "click"
	InputActionTypeText  InputAction = "type_text"
	InputActionHotkey    InputAction = "hotkey"
)
type InputRequest struct {
	Action  InputAction
	X       int
	Y       int
	Button  string
	Text    string
	Keys    []string
	Delay   time.Duration
	Options map[string]string
}

func (r InputRequest) Validate() error

type MatchCandidate struct {
	X     int
	Y     int
	W     int
	H     int
	Score float64
}

type Matcher interface {
	Find(req SearchRequest) ([]MatchCandidate, error)
}

type OCR interface {
	Read(req OCRRequest) (OCRResult, error)
}

type OCRRequest struct {
	Image            *image.Gray
	Language         string
	TrainingDataPath string
	MinConfidence    float64
	Timeout          time.Duration
}

func (r OCRRequest) Validate() error

type OCRResult struct {
	Text  string
	Words []OCRWord
}

type OCRWord struct {
	Text       string
	X          int
	Y          int
	W          int
	H          int
	Confidence float64
}

type ObserveEvent struct {
	Event     ObserveEventType
	X         int
	Y         int
	W         int
	H         int
	Score     float64
	Timestamp time.Time
}

type ObserveEventType string

const (
	ObserveEventAppear ObserveEventType = "appear"
	ObserveEventVanish ObserveEventType = "vanish"
	ObserveEventChange ObserveEventType = "change"
)
type ObserveRequest struct {
	Source   *image.Gray
	Region   image.Rectangle
	Pattern  *image.Gray
	Event    ObserveEventType
	Interval time.Duration
	Timeout  time.Duration
	Options  map[string]string
}

func (r ObserveRequest) Validate() error

type Observer interface {
	Observe(req ObserveRequest) ([]ObserveEvent, error)
}

type SearchRequest struct {
	Haystack     *image.Gray
	Needle       *image.Gray
	Mask         *image.Gray
	Threshold    float64
	ResizeFactor float64
	MaxResults   int
}

func (r SearchRequest) Validate() error

type WindowInfo struct {
	Title   string
	X       int
	Y       int
	W       int
	H       int
	Focused bool
}

```
