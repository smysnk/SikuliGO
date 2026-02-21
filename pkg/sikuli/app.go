package sikuli

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/sikulix/portgo/internal/app"
	"github.com/sikulix/portgo/internal/core"
)

type AppOptions struct {
	Timeout time.Duration
}

type Window struct {
	Title   string
	Bounds  Rect
	Focused bool
}

type AppController struct {
	backend core.App
}

var newAppBackend = func() core.App {
	return app.New()
}

func NewAppController() *AppController {
	return &AppController{
		backend: newAppBackend(),
	}
}

func (c *AppController) SetBackend(backend core.App) {
	if backend == nil {
		return
	}
	c.backend = backend
}

func (c *AppController) Open(name string, args []string, opts AppOptions) error {
	_, err := c.execute(core.AppActionOpen, name, args, opts)
	return err
}

func (c *AppController) Focus(name string, opts AppOptions) error {
	_, err := c.execute(core.AppActionFocus, name, nil, opts)
	return err
}

func (c *AppController) Close(name string, opts AppOptions) error {
	_, err := c.execute(core.AppActionClose, name, nil, opts)
	return err
}

func (c *AppController) IsRunning(name string, opts AppOptions) (bool, error) {
	res, err := c.execute(core.AppActionIsRunning, name, nil, opts)
	if err != nil {
		return false, err
	}
	return res.Running, nil
}

func (c *AppController) ListWindows(name string, opts AppOptions) ([]Window, error) {
	res, err := c.execute(core.AppActionListWindow, name, nil, opts)
	if err != nil {
		return nil, err
	}
	out := make([]Window, 0, len(res.Windows))
	for _, w := range res.Windows {
		out = append(out, Window{
			Title:   w.Title,
			Bounds:  NewRect(w.X, w.Y, w.W, w.H),
			Focused: w.Focused,
		})
	}
	return out, nil
}

func (c *AppController) execute(action core.AppAction, name string, args []string, opts AppOptions) (core.AppResult, error) {
	if c == nil || c.backend == nil {
		return core.AppResult{}, ErrBackendUnsupported
	}
	name = strings.TrimSpace(name)
	if name == "" {
		return core.AppResult{}, fmt.Errorf("%w: app name is empty", ErrInvalidTarget)
	}
	opts = normalizeAppOptions(opts)
	res, err := c.backend.Execute(core.AppRequest{
		Action:  action,
		Name:    name,
		Args:    args,
		Timeout: opts.Timeout,
	})
	if err != nil {
		return core.AppResult{}, mapAppError(err)
	}
	return res, nil
}

func normalizeAppOptions(in AppOptions) AppOptions {
	out := in
	if out.Timeout < 0 {
		out.Timeout = 0
	}
	return out
}

func mapAppError(err error) error {
	if err == nil {
		return nil
	}
	if errors.Is(err, core.ErrAppUnsupported) {
		return fmt.Errorf("%w: %v", ErrBackendUnsupported, err)
	}
	lower := strings.ToLower(err.Error())
	if strings.Contains(lower, "cannot be") ||
		strings.Contains(lower, "unsupported app action") {
		return fmt.Errorf("%w: %v", ErrInvalidTarget, err)
	}
	return err
}
