package core

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

var ErrInputUnsupported = errors.New("input backend unsupported")

type InputAction string

const (
	InputActionMouseMove InputAction = "mouse_move"
	InputActionClick     InputAction = "click"
	InputActionTypeText  InputAction = "type_text"
	InputActionHotkey    InputAction = "hotkey"
)

type InputRequest struct {
	Action  InputAction
	X       int
	Y       int
	Button  string
	Text    string
	Keys    []string
	Delay   time.Duration
	Options map[string]string
}

func (r InputRequest) Validate() error {
	if strings.TrimSpace(string(r.Action)) == "" {
		return fmt.Errorf("input action cannot be empty")
	}
	if r.Delay < 0 {
		return fmt.Errorf("input delay cannot be negative")
	}
	switch r.Action {
	case InputActionMouseMove:
		return nil
	case InputActionClick:
		if strings.TrimSpace(r.Button) == "" {
			return fmt.Errorf("click button cannot be empty")
		}
		return nil
	case InputActionTypeText:
		if strings.TrimSpace(r.Text) == "" {
			return fmt.Errorf("type text cannot be empty")
		}
		return nil
	case InputActionHotkey:
		if len(r.Keys) == 0 {
			return fmt.Errorf("hotkey requires at least one key")
		}
		for _, k := range r.Keys {
			if strings.TrimSpace(k) == "" {
				return fmt.Errorf("hotkey keys cannot be empty")
			}
		}
		return nil
	default:
		return fmt.Errorf("unsupported input action %q", r.Action)
	}
}

type Input interface {
	Execute(req InputRequest) error
}
