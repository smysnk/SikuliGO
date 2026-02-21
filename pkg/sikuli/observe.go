package sikuli

import (
	"errors"
	"fmt"
	"image"
	"strings"
	"time"

	"github.com/sikulix/portgo/internal/core"
	observebackend "github.com/sikulix/portgo/internal/observe"
)

type ObserveOptions struct {
	Interval time.Duration
	Timeout  time.Duration
}

type ObserveEventType string

const (
	ObserveEventAppear ObserveEventType = "appear"
	ObserveEventVanish ObserveEventType = "vanish"
	ObserveEventChange ObserveEventType = "change"
)

type ObserveEvent struct {
	Type      ObserveEventType
	Match     Match
	Timestamp time.Time
}

type ObserverController struct {
	backend core.Observer
}

var newObserveBackend = func() core.Observer {
	return observebackend.New()
}

func NewObserverController() *ObserverController {
	return &ObserverController{
		backend: newObserveBackend(),
	}
}

func (c *ObserverController) SetBackend(backend core.Observer) {
	if backend == nil {
		return
	}
	c.backend = backend
}

func (c *ObserverController) ObserveAppear(source *Image, region Region, pattern *Pattern, opts ObserveOptions) ([]ObserveEvent, error) {
	return c.observe(source, region, pattern, ObserveEventAppear, opts)
}

func (c *ObserverController) ObserveVanish(source *Image, region Region, pattern *Pattern, opts ObserveOptions) ([]ObserveEvent, error) {
	return c.observe(source, region, pattern, ObserveEventVanish, opts)
}

func (c *ObserverController) ObserveChange(source *Image, region Region, opts ObserveOptions) ([]ObserveEvent, error) {
	return c.observe(source, region, nil, ObserveEventChange, opts)
}

func (c *ObserverController) observe(source *Image, region Region, pattern *Pattern, event ObserveEventType, opts ObserveOptions) ([]ObserveEvent, error) {
	if c == nil || c.backend == nil {
		return nil, ErrBackendUnsupported
	}
	req, err := buildObserveRequest(source, region, pattern, event, opts)
	if err != nil {
		return nil, err
	}
	out, err := c.backend.Observe(req)
	if err != nil {
		return nil, mapObserveError(err)
	}
	events := make([]ObserveEvent, 0, len(out))
	for _, e := range out {
		events = append(events, ObserveEvent{
			Type:      ObserveEventType(e.Event),
			Match:     NewMatch(e.X, e.Y, e.W, e.H, e.Score, NewPoint(0, 0)),
			Timestamp: e.Timestamp,
		})
	}
	return events, nil
}

func buildObserveRequest(source *Image, region Region, pattern *Pattern, event ObserveEventType, opts ObserveOptions) (core.ObserveRequest, error) {
	if source == nil || source.Gray() == nil {
		return core.ObserveRequest{}, fmt.Errorf("%w: source image is nil", ErrInvalidTarget)
	}
	if region.Empty() {
		return core.ObserveRequest{}, fmt.Errorf("%w: region is empty", ErrInvalidTarget)
	}
	sourceBounds := source.Gray().Bounds()
	observeRect := image.Rect(region.X, region.Y, region.X+region.W, region.Y+region.H).Intersect(sourceBounds)
	if observeRect.Empty() {
		return core.ObserveRequest{}, fmt.Errorf("%w: region does not intersect source bounds", ErrInvalidTarget)
	}

	var eventKind core.ObserveEventType
	switch event {
	case ObserveEventAppear:
		eventKind = core.ObserveEventAppear
	case ObserveEventVanish:
		eventKind = core.ObserveEventVanish
	case ObserveEventChange:
		eventKind = core.ObserveEventChange
	default:
		return core.ObserveRequest{}, fmt.Errorf("%w: unsupported observe event %q", ErrInvalidTarget, event)
	}

	opts = normalizeObserveOptions(opts)
	req := core.ObserveRequest{
		Source:   source.Gray(),
		Region:   observeRect,
		Event:    eventKind,
		Interval: opts.Interval,
		Timeout:  opts.Timeout,
	}

	if eventKind == core.ObserveEventAppear || eventKind == core.ObserveEventVanish {
		if pattern == nil || pattern.Image() == nil || pattern.Image().Gray() == nil {
			return core.ObserveRequest{}, fmt.Errorf("%w: observe pattern image is nil", ErrInvalidTarget)
		}
		req.Pattern = pattern.Image().Gray()
	}

	return req, nil
}

func normalizeObserveOptions(in ObserveOptions) ObserveOptions {
	out := in
	if out.Timeout < 0 {
		out.Timeout = 0
	}
	if out.Interval <= 0 {
		out.Interval = defaultObserveInterval()
	}
	return out
}

func defaultObserveInterval() time.Duration {
	rate := DefaultObserveScanRate
	if s := GetSettings(); s.ObserveScanRate > 0 {
		rate = s.ObserveScanRate
	}
	interval := time.Duration(float64(time.Second) / rate)
	if interval < time.Millisecond {
		return time.Millisecond
	}
	return interval
}

func mapObserveError(err error) error {
	if err == nil {
		return nil
	}
	if errors.Is(err, core.ErrObserveUnsupported) {
		return fmt.Errorf("%w: %v", ErrBackendUnsupported, err)
	}
	lower := strings.ToLower(err.Error())
	if strings.Contains(lower, "cannot be") ||
		strings.Contains(lower, "requires") ||
		strings.Contains(lower, "unsupported observe event") {
		return fmt.Errorf("%w: %v", ErrInvalidTarget, err)
	}
	return err
}
