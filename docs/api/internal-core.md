# API: `internal/core`

[Back to API Index](./)

## Full Package Doc

```text
package core // import "github.com/sikulix/portgo/internal/core"


VARIABLES

var ErrInputUnsupported = errors.New("input backend unsupported")
var ErrOCRUnsupported = errors.New("ocr backend unsupported")

FUNCTIONS

func ResizeGrayNearest(src *image.Gray, factor float64) *image.Gray

TYPES

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

type SearchRequest struct {
	Haystack     *image.Gray
	Needle       *image.Gray
	Mask         *image.Gray
	Threshold    float64
	ResizeFactor float64
	MaxResults   int
}

func (r SearchRequest) Validate() error

```
