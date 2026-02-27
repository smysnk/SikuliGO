package core

import (
	"errors"
	"fmt"
	"image"
	"strings"
	"time"
)

var ErrOCRUnsupported = errors.New("ocr backend unsupported")

type OCRRequest struct {
	Image            *image.Gray
	Language         string
	TrainingDataPath string
	MinConfidence    float64
	Timeout          time.Duration
}

func (r OCRRequest) Validate() error {
	if r.Image == nil {
		return fmt.Errorf("ocr image cannot be nil")
	}
	if strings.TrimSpace(r.Language) == "" {
		return fmt.Errorf("ocr language cannot be empty")
	}
	if r.MinConfidence < 0 || r.MinConfidence > 1 {
		return fmt.Errorf("ocr minimum confidence must be in [0,1]")
	}
	if r.Timeout < 0 {
		return fmt.Errorf("ocr timeout cannot be negative")
	}
	return nil
}

type OCRWord struct {
	Text       string
	X          int
	Y          int
	W          int
	H          int
	Confidence float64
}

type OCRResult struct {
	Text  string
	Words []OCRWord
}

type OCR interface {
	Read(req OCRRequest) (OCRResult, error)
}
