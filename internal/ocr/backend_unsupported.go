//go:build !gosseract && !gogosseract

package ocr

import (
	"fmt"

	"github.com/sikulix/portgo/internal/core"
)

type unsupportedBackend struct{}

func New() core.OCR {
	return &unsupportedBackend{}
}

func (b *unsupportedBackend) Read(req core.OCRRequest) (core.OCRResult, error) {
	if err := req.Validate(); err != nil {
		return core.OCRResult{}, err
	}
	return core.OCRResult{}, fmt.Errorf("%w: build with -tags gosseract", core.ErrOCRUnsupported)
}
