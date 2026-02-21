package sikuli

import (
	"errors"
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/sikulix/portgo/internal/core"
	"github.com/sikulix/portgo/internal/cv"
	ocrbackend "github.com/sikulix/portgo/internal/ocr"
)

type Finder struct {
	source  *Image
	matcher core.Matcher
	ocr     core.OCR
	last    []Match
}

var newOCRBackend = func() core.OCR {
	return ocrbackend.New()
}

func NewFinder(source *Image) (*Finder, error) {
	if source == nil || source.Gray() == nil {
		return nil, fmt.Errorf("%w: source image is nil", ErrInvalidTarget)
	}
	return &Finder{
		source:  source,
		matcher: cv.NewNCCMatcher(),
		ocr:     newOCRBackend(),
		last:    nil,
	}, nil
}

func (f *Finder) SetMatcher(m core.Matcher) {
	if m == nil {
		return
	}
	f.matcher = m
}

func (f *Finder) SetOCRBackend(ocr core.OCR) {
	if ocr == nil {
		return
	}
	f.ocr = ocr
}

func (f *Finder) Find(pattern *Pattern) (Match, error) {
	req, err := f.buildRequest(pattern, 1)
	if err != nil {
		return Match{}, err
	}
	rawMatches, err := f.matcher.Find(req)
	if err != nil {
		return Match{}, err
	}
	if len(rawMatches) == 0 {
		f.last = nil
		return Match{}, ErrFindFailed
	}
	match := toMatch(rawMatches[0], pattern.Offset())
	match.Index = 0
	f.last = []Match{match}
	return match, nil
}

func (f *Finder) FindAll(pattern *Pattern) ([]Match, error) {
	req, err := f.buildRequest(pattern, 0)
	if err != nil {
		return nil, err
	}
	rawMatches, err := f.matcher.Find(req)
	if err != nil {
		return nil, err
	}
	matches := make([]Match, 0, len(rawMatches))
	for i, m := range rawMatches {
		out := toMatch(m, pattern.Offset())
		out.Index = i
		matches = append(matches, out)
	}
	f.last = matches
	return matches, nil
}

func (f *Finder) FindAllByRow(pattern *Pattern) ([]Match, error) {
	matches, err := f.FindAll(pattern)
	if err != nil {
		return nil, err
	}
	SortMatchesByRowColumn(matches)
	reindex(matches)
	f.last = matches
	return matches, nil
}

func (f *Finder) FindAllByColumn(pattern *Pattern) ([]Match, error) {
	matches, err := f.FindAll(pattern)
	if err != nil {
		return nil, err
	}
	SortMatchesByColumnRow(matches)
	reindex(matches)
	f.last = matches
	return matches, nil
}

// Exists returns the first match when present. Missing targets return (Match{}, false, nil).
func (f *Finder) Exists(pattern *Pattern) (Match, bool, error) {
	match, err := f.Find(pattern)
	if err != nil {
		if err == ErrFindFailed {
			return Match{}, false, nil
		}
		return Match{}, false, err
	}
	return match, true, nil
}

// Has reports whether the target exists and bubbles non-find errors.
func (f *Finder) Has(pattern *Pattern) (bool, error) {
	_, ok, err := f.Exists(pattern)
	return ok, err
}

func (f *Finder) Wait(pattern *Pattern, timeout time.Duration) (Match, error) {
	checkOnce := func() (Match, bool, error) {
		m, ok, err := f.Exists(pattern)
		if err != nil {
			return Match{}, false, err
		}
		return m, ok, nil
	}

	if timeout <= 0 {
		m, ok, err := checkOnce()
		if err != nil {
			return Match{}, err
		}
		if !ok {
			return Match{}, ErrTimeout
		}
		return m, nil
	}

	deadline := time.Now().Add(timeout)
	interval := finderWaitInterval()
	for {
		m, ok, err := checkOnce()
		if err != nil {
			return Match{}, err
		}
		if ok {
			return m, nil
		}
		if time.Now().After(deadline) {
			return Match{}, ErrTimeout
		}
		sleep := interval
		if remaining := time.Until(deadline); remaining < sleep {
			sleep = remaining
		}
		if sleep > 0 {
			time.Sleep(sleep)
		}
	}
}

func (f *Finder) WaitVanish(pattern *Pattern, timeout time.Duration) (bool, error) {
	checkOnce := func() (bool, error) {
		_, ok, err := f.Exists(pattern)
		if err != nil {
			return false, err
		}
		return !ok, nil
	}

	if timeout <= 0 {
		return checkOnce()
	}

	deadline := time.Now().Add(timeout)
	interval := finderWaitInterval()
	for {
		vanished, err := checkOnce()
		if err != nil {
			return false, err
		}
		if vanished {
			return true, nil
		}
		if time.Now().After(deadline) {
			return false, nil
		}
		sleep := interval
		if remaining := time.Until(deadline); remaining < sleep {
			sleep = remaining
		}
		if sleep > 0 {
			time.Sleep(sleep)
		}
	}
}

func (f *Finder) LastMatches() []Match {
	if len(f.last) == 0 {
		return nil
	}
	out := make([]Match, len(f.last))
	copy(out, f.last)
	return out
}

func (f *Finder) ReadText(params OCRParams) (string, error) {
	result, err := f.readOCR(params)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(result.Text), nil
}

func (f *Finder) FindText(query string, params OCRParams) ([]TextMatch, error) {
	query = strings.TrimSpace(query)
	if query == "" {
		return nil, fmt.Errorf("%w: text query is empty", ErrInvalidTarget)
	}

	opts := normalizeOCRParams(params)
	result, err := f.readOCR(opts)
	if err != nil {
		return nil, err
	}

	matches := make([]TextMatch, 0)
	for _, word := range result.Words {
		if !containsText(word.Text, query, opts.CaseSensitive) {
			continue
		}
		matches = append(matches, TextMatch{
			Rect:       NewRect(word.X, word.Y, word.W, word.H),
			Text:       word.Text,
			Confidence: word.Confidence,
		})
	}
	if len(matches) == 0 && containsText(result.Text, query, opts.CaseSensitive) && f.source != nil && f.source.Gray() != nil {
		b := f.source.Gray().Bounds()
		matches = append(matches, TextMatch{
			Rect:       NewRect(b.Min.X, b.Min.Y, b.Dx(), b.Dy()),
			Text:       strings.TrimSpace(result.Text),
			Confidence: 0,
		})
	}
	if len(matches) == 0 {
		return nil, ErrFindFailed
	}
	sort.Slice(matches, func(i, j int) bool {
		if matches[i].Y == matches[j].Y {
			return matches[i].X < matches[j].X
		}
		return matches[i].Y < matches[j].Y
	})
	for i := range matches {
		matches[i].Index = i
	}
	return matches, nil
}

func (f *Finder) buildRequest(pattern *Pattern, maxResults int) (core.SearchRequest, error) {
	if f == nil || f.source == nil || f.source.Gray() == nil {
		return core.SearchRequest{}, fmt.Errorf("%w: source image is nil", ErrInvalidTarget)
	}
	if pattern == nil || pattern.Image() == nil || pattern.Image().Gray() == nil {
		return core.SearchRequest{}, fmt.Errorf("%w: pattern image is nil", ErrInvalidTarget)
	}
	req := core.SearchRequest{
		Haystack:     f.source.Gray(),
		Needle:       pattern.Image().Gray(),
		Mask:         pattern.Mask(),
		Threshold:    pattern.Similarity(),
		ResizeFactor: pattern.ResizeFactor(),
		MaxResults:   maxResults,
	}
	return req, nil
}

func (f *Finder) readOCR(params OCRParams) (core.OCRResult, error) {
	if f == nil || f.source == nil || f.source.Gray() == nil {
		return core.OCRResult{}, fmt.Errorf("%w: source image is nil", ErrInvalidTarget)
	}
	if f.ocr == nil {
		return core.OCRResult{}, ErrBackendUnsupported
	}
	opts := normalizeOCRParams(params)
	result, err := f.ocr.Read(core.OCRRequest{
		Image:            f.source.Gray(),
		Language:         opts.Language,
		TrainingDataPath: opts.TrainingDataPath,
		MinConfidence:    opts.MinConfidence,
		Timeout:          opts.Timeout,
	})
	if err != nil {
		if errors.Is(err, core.ErrOCRUnsupported) {
			return core.OCRResult{}, fmt.Errorf("%w: %v", ErrBackendUnsupported, err)
		}
		return core.OCRResult{}, err
	}
	return result, nil
}

func finderWaitInterval() time.Duration {
	rate := DefaultWaitScanRate
	if s := GetSettings(); s.WaitScanRate > 0 {
		rate = s.WaitScanRate
	}
	interval := time.Duration(float64(time.Second) / rate)
	if interval < time.Millisecond {
		return time.Millisecond
	}
	return interval
}

func toMatch(candidate core.MatchCandidate, off Point) Match {
	return NewMatch(candidate.X, candidate.Y, candidate.W, candidate.H, candidate.Score, off)
}

// SortMatchesByRowColumn keeps parity with Java helper behavior for "by row".
func SortMatchesByRowColumn(matches []Match) {
	sort.Slice(matches, func(i, j int) bool {
		if matches[i].Y == matches[j].Y {
			return matches[i].X < matches[j].X
		}
		return matches[i].Y < matches[j].Y
	})
}

// SortMatchesByColumnRow keeps parity with Java helper behavior for "by column".
func SortMatchesByColumnRow(matches []Match) {
	sort.Slice(matches, func(i, j int) bool {
		if matches[i].X == matches[j].X {
			return matches[i].Y < matches[j].Y
		}
		return matches[i].X < matches[j].X
	})
}

func reindex(matches []Match) {
	for i := range matches {
		matches[i].Index = i
	}
}
