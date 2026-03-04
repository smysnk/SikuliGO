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
	"time"

	"github.com/smysnk/sikuligo/pkg/sikuli"
)

func screencaptureArgs(tmpPath string) []string {
	args := []string{"-x", "-t", "png"}
	display := strings.TrimSpace(os.Getenv("SIKULI_CAPTURE_DISPLAY"))
	if display == "" {
		display = strings.TrimSpace(os.Getenv("SIKULIGO_CAPTURE_DISPLAY"))
	}
	if display != "" {
		args = append(args, "-D", display)
	}
	args = append(args, tmpPath)
	return args
}

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

	args := screencaptureArgs(tmpPath)
	debugLogf("capture_screen.command.start name=%s args=%q", name, strings.Join(args, " "))
	cmd := exec.CommandContext(ctx, "screencapture", args...)
	start := time.Now()
	if out, err := cmd.CombinedOutput(); err != nil {
		debugLogf("capture_screen.command.error duration=%s err=%v output=%q", time.Since(start), err, strings.TrimSpace(string(out)))
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
	debugLogf("capture_screen.command.ok duration=%s tmp=%s", time.Since(start), tmpPath)

	buf, err := os.ReadFile(tmpPath)
	if err != nil {
		debugLogf("capture_screen.read.error tmp=%s err=%v", tmpPath, err)
		return nil, err
	}
	debugLogf("capture_screen.read.ok tmp=%s bytes=%d", tmpPath, len(buf))
	decoded, _, err := image.Decode(bytes.NewReader(buf))
	if err != nil {
		debugLogf("capture_screen.decode.error tmp=%s err=%v", tmpPath, err)
		return nil, err
	}
	b := decoded.Bounds()
	debugLogf("capture_screen.decode.ok width=%d height=%d", b.Dx(), b.Dy())
	out, err := sikuli.NewImageFromAny(name, decoded)
	if err != nil {
		debugLogf("capture_screen.image.error err=%v", err)
		return nil, err
	}
	debugLogf("capture_screen.image.ok name=%s width=%d height=%d", name, out.Width(), out.Height())
	return out, nil
}
