# API: `internal/core`

[Back to API Index](./)

## Full Package Doc

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
