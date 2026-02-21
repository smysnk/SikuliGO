package observe

import (
	"fmt"
	"image"
	"strconv"
	"time"

	"github.com/sikulix/portgo/internal/core"
	"github.com/sikulix/portgo/internal/cv"
)

const (
	defaultObserveThreshold = 0.7
	minObserveInterval      = time.Millisecond
)

type pollingBackend struct {
	matcher core.Matcher
	now     func() time.Time
	sleep   func(time.Duration)
	// onPoll is a test hook invoked before each sample.
	onPoll func(iteration int)
}

func New() core.Observer {
	return newPollingBackend(nil)
}

func newPollingBackend(m core.Matcher) *pollingBackend {
	if m == nil {
		m = cv.NewNCCMatcher()
	}
	return &pollingBackend{
		matcher: m,
		now:     time.Now,
		sleep:   time.Sleep,
	}
}

func (b *pollingBackend) Observe(req core.ObserveRequest) ([]core.ObserveEvent, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	if b == nil || b.matcher == nil {
		return nil, fmt.Errorf("%w: matcher unavailable", core.ErrObserveUnsupported)
	}

	region := req.Region.Intersect(req.Source.Bounds())
	if region.Empty() {
		return nil, fmt.Errorf("observe region does not intersect source bounds")
	}

	interval := req.Interval
	if interval <= 0 {
		interval = minObserveInterval
	}

	threshold := parseClampedFloat(req.Options, "threshold", defaultObserveThreshold)
	changeThreshold := parseClampedFloat(req.Options, "change_threshold", 0)

	now := b.now
	sleep := b.sleep
	if now == nil {
		now = time.Now
	}
	if sleep == nil {
		sleep = time.Sleep
	}

	start := now()
	deadline := start.Add(req.Timeout)

	var previous *image.Gray
	var lastMatch *core.MatchCandidate

	for i := 0; ; i++ {
		if b.onPoll != nil {
			b.onPoll(i)
		}

		frame, err := cloneGrayRegion(req.Source, region)
		if err != nil {
			return nil, err
		}
		timestamp := now()

		switch req.Event {
		case core.ObserveEventAppear:
			match, found, err := b.findFirst(frame, req.Pattern, threshold)
			if err != nil {
				return nil, err
			}
			if found {
				return []core.ObserveEvent{{
					Event:     core.ObserveEventAppear,
					X:         match.X,
					Y:         match.Y,
					W:         match.W,
					H:         match.H,
					Score:     match.Score,
					Timestamp: timestamp,
				}}, nil
			}
		case core.ObserveEventVanish:
			match, found, err := b.findFirst(frame, req.Pattern, threshold)
			if err != nil {
				return nil, err
			}
			if found {
				cpy := match
				lastMatch = &cpy
			}
			if !found {
				ev := core.ObserveEvent{
					Event:     core.ObserveEventVanish,
					X:         region.Min.X,
					Y:         region.Min.Y,
					W:         region.Dx(),
					H:         region.Dy(),
					Score:     1,
					Timestamp: timestamp,
				}
				if lastMatch != nil {
					ev.X = lastMatch.X
					ev.Y = lastMatch.Y
					ev.W = lastMatch.W
					ev.H = lastMatch.H
				}
				return []core.ObserveEvent{ev}, nil
			}
		case core.ObserveEventChange:
			if previous != nil {
				score := normalizedDiff(previous, frame)
				if score > changeThreshold {
					return []core.ObserveEvent{{
						Event:     core.ObserveEventChange,
						X:         region.Min.X,
						Y:         region.Min.Y,
						W:         region.Dx(),
						H:         region.Dy(),
						Score:     score,
						Timestamp: timestamp,
					}}, nil
				}
			}
			previous = frame
		default:
			return nil, fmt.Errorf("unsupported observe event %q", req.Event)
		}

		if req.Timeout <= 0 {
			return nil, nil
		}
		if !timestamp.Before(deadline) {
			return nil, nil
		}

		wait := interval
		if remaining := deadline.Sub(timestamp); remaining < wait {
			wait = remaining
		}
		if wait > 0 {
			sleep(wait)
		}
	}
}

func (b *pollingBackend) findFirst(frame, pattern *image.Gray, threshold float64) (core.MatchCandidate, bool, error) {
	matches, err := b.matcher.Find(core.SearchRequest{
		Haystack:     frame,
		Needle:       pattern,
		Threshold:    threshold,
		ResizeFactor: 1,
		MaxResults:   1,
	})
	if err != nil {
		return core.MatchCandidate{}, false, err
	}
	if len(matches) == 0 {
		return core.MatchCandidate{}, false, nil
	}
	return matches[0], true, nil
}

func cloneGrayRegion(src *image.Gray, region image.Rectangle) (*image.Gray, error) {
	if src == nil {
		return nil, fmt.Errorf("source image is nil")
	}
	crop := region.Intersect(src.Bounds())
	if crop.Empty() {
		return nil, fmt.Errorf("region does not intersect source bounds")
	}

	dst := image.NewGray(crop)
	for y := crop.Min.Y; y < crop.Max.Y; y++ {
		srcStart := src.PixOffset(crop.Min.X, y)
		srcEnd := src.PixOffset(crop.Max.X, y)
		dstStart := dst.PixOffset(crop.Min.X, y)
		copy(dst.Pix[dstStart:dstStart+crop.Dx()], src.Pix[srcStart:srcEnd])
	}
	return dst, nil
}

func parseClampedFloat(options map[string]string, key string, fallback float64) float64 {
	if options == nil {
		return fallback
	}
	raw, ok := options[key]
	if !ok {
		return fallback
	}
	parsed, err := strconv.ParseFloat(raw, 64)
	if err != nil {
		return fallback
	}
	if parsed < 0 {
		return 0
	}
	if parsed > 1 {
		return 1
	}
	return parsed
}

func normalizedDiff(a, b *image.Gray) float64 {
	if a == nil || b == nil {
		return 0
	}
	ab := a.Bounds()
	bb := b.Bounds()
	if ab.Dx() != bb.Dx() || ab.Dy() != bb.Dy() {
		return 1
	}
	if ab.Dx() == 0 || ab.Dy() == 0 {
		return 0
	}
	var total int
	var count int
	for y := 0; y < ab.Dy(); y++ {
		for x := 0; x < ab.Dx(); x++ {
			av := int(a.GrayAt(ab.Min.X+x, ab.Min.Y+y).Y)
			bv := int(b.GrayAt(bb.Min.X+x, bb.Min.Y+y).Y)
			diff := av - bv
			if diff < 0 {
				diff = -diff
			}
			total += diff
			count++
		}
	}
	if count == 0 {
		return 0
	}
	return float64(total) / float64(count*255)
}
