//go:build !darwin

package input

import (
	"fmt"

	"github.com/sikulix/portgo/internal/core"
)

type unsupportedBackend struct{}

func New() core.Input {
	return &unsupportedBackend{}
}

func (b *unsupportedBackend) Execute(req core.InputRequest) error {
	if err := req.Validate(); err != nil {
		return err
	}
	return fmt.Errorf("%w: no input backend configured", core.ErrInputUnsupported)
}
