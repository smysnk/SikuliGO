package core

import (
	"errors"
	"fmt"
	"image"
	"time"
)

var ErrObserveUnsupported = errors.New("observe backend unsupported")

type ObserveEventType string

const (
	ObserveEventAppear ObserveEventType = "appear"
	ObserveEventVanish ObserveEventType = "vanish"
	ObserveEventChange ObserveEventType = "change"
)

type ObserveRequest struct {
	Source   *image.Gray
	Region   image.Rectangle
	Pattern  *image.Gray
	Event    ObserveEventType
	Interval time.Duration
	Timeout  time.Duration
	Options  map[string]string
}

func (r ObserveRequest) Validate() error {
	if r.Source == nil {
		return fmt.Errorf("observe source cannot be nil")
	}
	if r.Region.Empty() {
		return fmt.Errorf("observe region cannot be empty")
	}
	if r.Interval < 0 {
		return fmt.Errorf("observe interval cannot be negative")
	}
	if r.Timeout < 0 {
		return fmt.Errorf("observe timeout cannot be negative")
	}
	switch r.Event {
	case ObserveEventAppear, ObserveEventVanish:
		if r.Pattern == nil {
			return fmt.Errorf("observe pattern cannot be nil for %q", r.Event)
		}
	case ObserveEventChange:
		// Pattern is optional for change observation.
	default:
		return fmt.Errorf("unsupported observe event %q", r.Event)
	}
	return nil
}

type ObserveEvent struct {
	Event     ObserveEventType
	X         int
	Y         int
	W         int
	H         int
	Score     float64
	Timestamp time.Time
}

type Observer interface {
	Observe(req ObserveRequest) ([]ObserveEvent, error)
}
