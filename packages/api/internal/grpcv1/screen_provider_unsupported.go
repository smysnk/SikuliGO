//go:build !darwin

package grpcv1

import (
	"context"
	"fmt"

	"github.com/smysnk/sikuligo/pkg/sikuli"
)

func listScreens(context.Context) ([]sikuli.Screen, error) {
	return nil, fmt.Errorf("%w: screen enumeration backend unsupported on this platform", sikuli.ErrBackendUnsupported)
}
