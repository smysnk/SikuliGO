package sikuli

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/sikulix/portgo/internal/core"
	inputbackend "github.com/sikulix/portgo/internal/input"
)

type MouseButton string

const (
	MouseButtonLeft   MouseButton = "left"
	MouseButtonRight  MouseButton = "right"
	MouseButtonMiddle MouseButton = "middle"
)

type InputOptions struct {
	Delay  time.Duration
	Button MouseButton
}

type InputController struct {
	backend core.Input
}

var newInputBackend = func() core.Input {
	return inputbackend.New()
}

func NewInputController() *InputController {
	return &InputController{
		backend: newInputBackend(),
	}
}

func (c *InputController) SetBackend(backend core.Input) {
	if backend == nil {
		return
	}
	c.backend = backend
}

func (c *InputController) MoveMouse(x, y int, opts InputOptions) error {
	opts = normalizeInputOptions(opts)
	return c.execute(core.InputRequest{
		Action: core.InputActionMouseMove,
		X:      x,
		Y:      y,
		Delay:  opts.Delay,
	})
}

func (c *InputController) Click(x, y int, opts InputOptions) error {
	opts = normalizeInputOptions(opts)
	return c.execute(core.InputRequest{
		Action: core.InputActionClick,
		X:      x,
		Y:      y,
		Button: string(opts.Button),
		Delay:  opts.Delay,
	})
}

func (c *InputController) TypeText(text string, opts InputOptions) error {
	opts = normalizeInputOptions(opts)
	text = strings.TrimSpace(text)
	if text == "" {
		return fmt.Errorf("%w: type text cannot be empty", ErrInvalidTarget)
	}
	return c.execute(core.InputRequest{
		Action: core.InputActionTypeText,
		Text:   text,
		Delay:  opts.Delay,
	})
}

func (c *InputController) Hotkey(keys ...string) error {
	if len(keys) == 0 {
		return fmt.Errorf("%w: hotkey requires at least one key", ErrInvalidTarget)
	}
	normalized := make([]string, 0, len(keys))
	for _, key := range keys {
		key = strings.TrimSpace(key)
		if key == "" {
			return fmt.Errorf("%w: hotkey keys cannot be empty", ErrInvalidTarget)
		}
		normalized = append(normalized, key)
	}
	return c.execute(core.InputRequest{
		Action: core.InputActionHotkey,
		Keys:   normalized,
	})
}

func normalizeInputOptions(in InputOptions) InputOptions {
	out := in
	if out.Delay < 0 {
		out.Delay = 0
	}
	switch out.Button {
	case MouseButtonLeft, MouseButtonRight, MouseButtonMiddle:
	default:
		out.Button = MouseButtonLeft
	}
	return out
}

func (c *InputController) execute(req core.InputRequest) error {
	if c == nil || c.backend == nil {
		return ErrBackendUnsupported
	}
	if err := c.backend.Execute(req); err != nil {
		if errors.Is(err, core.ErrInputUnsupported) {
			return fmt.Errorf("%w: %v", ErrBackendUnsupported, err)
		}
		if strings.Contains(strings.ToLower(err.Error()), "executable file not found") {
			return fmt.Errorf("%w: %v", ErrBackendUnsupported, err)
		}
		if strings.Contains(strings.ToLower(err.Error()), "cannot be") ||
			strings.Contains(strings.ToLower(err.Error()), "requires") ||
			strings.Contains(strings.ToLower(err.Error()), "unsupported input action") {
			return fmt.Errorf("%w: %v", ErrInvalidTarget, err)
		}
		return err
	}
	return nil
}
