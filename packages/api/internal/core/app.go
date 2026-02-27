package core

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

var ErrAppUnsupported = errors.New("app backend unsupported")

type AppAction string

const (
	AppActionOpen       AppAction = "open"
	AppActionFocus      AppAction = "focus"
	AppActionClose      AppAction = "close"
	AppActionIsRunning  AppAction = "is_running"
	AppActionListWindow AppAction = "list_windows"
)

type AppRequest struct {
	Action  AppAction
	Name    string
	Args    []string
	Timeout time.Duration
	Options map[string]string
}

func (r AppRequest) Validate() error {
	if strings.TrimSpace(string(r.Action)) == "" {
		return fmt.Errorf("app action cannot be empty")
	}
	if strings.TrimSpace(r.Name) == "" {
		return fmt.Errorf("app name cannot be empty")
	}
	if r.Timeout < 0 {
		return fmt.Errorf("app timeout cannot be negative")
	}
	switch r.Action {
	case AppActionOpen, AppActionFocus, AppActionClose, AppActionIsRunning, AppActionListWindow:
		return nil
	default:
		return fmt.Errorf("unsupported app action %q", r.Action)
	}
}

type WindowInfo struct {
	Title   string
	X       int
	Y       int
	W       int
	H       int
	Focused bool
}

type AppResult struct {
	Running bool
	PID     int
	Windows []WindowInfo
}

type App interface {
	Execute(req AppRequest) (AppResult, error)
}
