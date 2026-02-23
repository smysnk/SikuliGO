//go:build darwin

package grpcv1

import (
	"bytes"
	"fmt"
	"image"
	_ "image/png"
	"os"
	"os/exec"
	"strings"

	"github.com/sikulix/portgo/pkg/sikuli"
)

func captureScreenImage(name string) (*sikuli.Image, error) {
	tmp, err := os.CreateTemp("", "sikuligrpc-screen-*.png")
	if err != nil {
		return nil, err
	}
	tmpPath := tmp.Name()
	_ = tmp.Close()
	defer func() { _ = os.Remove(tmpPath) }()

	cmd := exec.Command("screencapture", "-x", "-t", "png", tmpPath)
	if out, err := cmd.CombinedOutput(); err != nil {
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
