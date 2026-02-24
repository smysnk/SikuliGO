//go:build darwin

package grpcv1

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"image"
	_ "image/png"
	"os"
	"os/exec"
	"strings"

	"github.com/smysnk/sikuligo/pkg/sikuli"
)

func captureScreenImage(ctx context.Context, name string) (*sikuli.Image, error) {
	if ctx == nil {
		ctx = context.Background()
	}
	tmp, err := os.CreateTemp("", "sikuligrpc-screen-*.png")
	if err != nil {
		return nil, err
	}
	tmpPath := tmp.Name()
	_ = tmp.Close()
	defer func() { _ = os.Remove(tmpPath) }()

	cmd := exec.CommandContext(ctx, "screencapture", "-x", "-t", "png", tmpPath)
	if out, err := cmd.CombinedOutput(); err != nil {
		if errors.Is(ctx.Err(), context.DeadlineExceeded) {
			return nil, fmt.Errorf("%w: screencapture canceled by request deadline", sikuli.ErrTimeout)
		}
		if errors.Is(ctx.Err(), context.Canceled) {
			return nil, fmt.Errorf("%w: screencapture canceled", sikuli.ErrTimeout)
		}
		msg := strings.TrimSpace(string(out))
		if msg == "" {
			msg = err.Error()
		}
		return nil, fmt.Errorf("%w: screencapture failed: %s", sikuli.ErrBackendUnsupported, msg)
	}

	buf, err := os.ReadFile(tmpPath)
	if err != nil {
		return nil, err
	}
	decoded, _, err := image.Decode(bytes.NewReader(buf))
	if err != nil {
		return nil, err
	}
	return sikuli.NewImageFromAny(name, decoded)
}
