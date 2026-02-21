package app

import (
	"fmt"

	"github.com/sikulix/portgo/internal/core"
)

type unsupportedBackend struct{}

func New() core.App {
	return &unsupportedBackend{}
}

func (b *unsupportedBackend) Execute(req core.AppRequest) (core.AppResult, error) {
	if err := req.Validate(); err != nil {
		return core.AppResult{}, err
	}
	return core.AppResult{}, fmt.Errorf("%w: no app backend configured", core.ErrAppUnsupported)
}
