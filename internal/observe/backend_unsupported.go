package observe

import (
	"fmt"

	"github.com/sikulix/portgo/internal/core"
)

type unsupportedBackend struct{}

func New() core.Observer {
	return &unsupportedBackend{}
}

func (b *unsupportedBackend) Observe(req core.ObserveRequest) ([]core.ObserveEvent, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	return nil, fmt.Errorf("%w: no observe backend configured", core.ErrObserveUnsupported)
}
