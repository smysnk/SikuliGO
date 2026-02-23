//go:build !darwin

package grpcv1

import (
	"fmt"

	"github.com/sikulix/portgo/pkg/sikuli"
)

func captureScreenImage(name string) (*sikuli.Image, error) {
	return nil, fmt.Errorf("%w: screen capture backend unsupported on this platform", sikuli.ErrBackendUnsupported)
}
