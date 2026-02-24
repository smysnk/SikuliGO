//go:build !darwin

package grpcv1

import (
	"context"
	"fmt"

	"github.com/smysnk/sikuligo/pkg/sikuli"
)

func captureScreenImage(_ context.Context, name string) (*sikuli.Image, error) {
	return nil, fmt.Errorf("%w: screen capture backend unsupported on this platform", sikuli.ErrBackendUnsupported)
}
