package sikuli

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/smysnk/sikuligo/internal/app"
	"github.com/smysnk/sikuligo/internal/core"
)

type AppOptions struct {
	Timeout time.Duration
}

type Window struct {
	ID      string
	App     string
	PID     int
	Title   string
	Bounds  Rect
	Focused bool
}

type WindowQuery struct {
	ID            string
	TitleExact    string
	TitleContains string
	FocusedOnly   bool
	Index         int
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
	return windowsFromCore(res.Windows), nil
}

func (c *AppController) FindWindows(name string, query WindowQuery, opts AppOptions) ([]Window, error) {
	windows, err := c.ListWindows(name, opts)
	if err != nil {
		return nil, err
	}
	return filterWindows(windows, query), nil
}

func (c *AppController) GetWindow(name string, query WindowQuery, opts AppOptions) (Window, bool, error) {
	windows, err := c.FindWindows(name, query, opts)
	if err != nil {
		return Window{}, false, err
	}
	if len(windows) == 0 {
		return Window{}, false, nil
	}
	index := normalizeWindowQuery(query).Index
	if index < 0 || index >= len(windows) {
		return Window{}, false, nil
	}
	return windows[index], true, nil
}

func (c *AppController) FocusedWindow(name string, opts AppOptions) (Window, bool, error) {
	return c.GetWindow(name, WindowQuery{FocusedOnly: true}, opts)
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

func windowsFromCore(in []core.WindowInfo) []Window {
	out := make([]Window, 0, len(in))
	for _, w := range in {
		out = append(out, Window{
			ID:      w.ID,
			App:     w.App,
			PID:     w.PID,
			Title:   w.Title,
			Bounds:  NewRect(w.X, w.Y, w.W, w.H),
			Focused: w.Focused,
		})
	}
	return out
}

func normalizeWindowQuery(in WindowQuery) WindowQuery {
	out := in
	out.ID = strings.TrimSpace(out.ID)
	out.TitleExact = strings.TrimSpace(out.TitleExact)
	out.TitleContains = strings.TrimSpace(out.TitleContains)
	if out.Index < 0 {
		out.Index = 0
	}
	return out
}

func filterWindows(windows []Window, query WindowQuery) []Window {
	query = normalizeWindowQuery(query)
	matches := make([]Window, 0, len(windows))
	for _, window := range windows {
		if query.ID != "" && window.ID != query.ID {
			continue
		}
		if query.FocusedOnly && !window.Focused {
			continue
		}
		if query.TitleExact != "" && window.Title != query.TitleExact {
			continue
		}
		if query.TitleContains != "" && !strings.Contains(strings.ToLower(window.Title), strings.ToLower(query.TitleContains)) {
			continue
		}
		matches = append(matches, window)
	}
	return matches
}
