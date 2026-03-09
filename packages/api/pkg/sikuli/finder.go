package sikuli

import (
	"errors"
	"fmt"
	"math"
	"sort"
	"strings"
	"time"

	"github.com/smysnk/sikuligo/internal/core"
	"github.com/smysnk/sikuligo/internal/cv"
	ocrbackend "github.com/smysnk/sikuligo/internal/ocr"
)

type Finder struct {
	source  *Image
	matcher core.Matcher
	ocr     core.OCR
	last    []Match
	nextIdx int
}

var newOCRBackend = func() core.OCR {
	return ocrbackend.New()
}

// NewFinder creates a search/OCR helper bound to a source image.
func NewFinder(source *Image) (*Finder, error) {
	if source == nil || source.Gray() == nil {
		return nil, fmt.Errorf("%w: source image is nil", ErrInvalidTarget)
	}
	return &Finder{
		source:  source,
		matcher: cv.NewDefaultMatcher(),
		ocr:     newOCRBackend(),
		last:    nil,
		nextIdx: 0,
	}, nil
}

// SetMatcher overrides the matcher backend used by this finder.
func (f *Finder) SetMatcher(m core.Matcher) {
	if m == nil {
		return
	}
	f.matcher = m
}

// SetOCRBackend overrides the OCR backend used by this finder.
func (f *Finder) SetOCRBackend(ocr core.OCR) {
	if ocr == nil {
		return
	}
	f.ocr = ocr
}

// Find returns the best match for the pattern.
func (f *Finder) Find(pattern *Pattern) (Match, error) {
	matches, err := f.searchMatches(pattern, 1)
	if err != nil {
		return Match{}, err
	}
	if len(matches) == 0 {
		f.clearTraversal()
		return Match{}, ErrFindFailed
	}
	f.setTraversal(matches)
	return matches[0], nil
}

// FindAll returns all candidate matches for the pattern.
func (f *Finder) FindAll(pattern *Pattern) ([]Match, error) {
	matches, err := f.searchMatches(pattern, 0)
	if err != nil {
		return nil, err
	}
	f.setTraversal(matches)
	return matches, nil
}

// Iterate prepares a compatibility iterator over the best match.
// Unlike Find, a miss does not return ErrFindFailed. Call HasNext to inspect presence.
func (f *Finder) Iterate(pattern *Pattern) error {
	matches, err := f.searchMatches(pattern, 1)
	if err != nil {
		return err
	}
	f.setTraversal(matches)
	return nil
}

// IterateAll prepares a compatibility iterator over all matches.
// Unlike Java SikuliX this additive Go surface keeps LastMatches available even after iteration.
func (f *Finder) IterateAll(pattern *Pattern) error {
	matches, err := f.searchMatches(pattern, 0)
	if err != nil {
		return err
	}
	f.setTraversal(matches)
	return nil
}

// Count returns the number of matches for the given pattern.
func (f *Finder) Count(pattern *Pattern) (int, error) {
	matches, err := f.FindAll(pattern)
	if err != nil {
		return 0, err
	}
	return len(matches), nil
}

func (f *Finder) FindAllByRow(pattern *Pattern) ([]Match, error) {
	matches, err := f.FindAll(pattern)
	if err != nil {
		return nil, err
	}
	SortMatchesByRowColumn(matches)
	reindex(matches)
	f.setTraversal(matches)
	return matches, nil
}

// FindAllByColumn returns all matches sorted left-to-right then top-to-bottom.
func (f *Finder) FindAllByColumn(pattern *Pattern) ([]Match, error) {
	matches, err := f.FindAll(pattern)
	if err != nil {
		return nil, err
	}
	SortMatchesByColumnRow(matches)
	reindex(matches)
	f.setTraversal(matches)
	return matches, nil
}

// Exists returns the first match when present. Missing targets return (Match{}, false, nil).
func (f *Finder) Exists(pattern *Pattern) (Match, bool, error) {
	return SearchExists(func() (Match, error) {
		return f.Find(pattern)
	}, 0, finderWaitInterval())
}

// Has reports whether the target exists and bubbles non-find errors.
func (f *Finder) Has(pattern *Pattern) (bool, error) {
	_, ok, err := f.Exists(pattern)
	return ok, err
}

func (f *Finder) Wait(pattern *Pattern, timeout time.Duration) (Match, error) {
	return SearchWait(func() (Match, error) {
		return f.Find(pattern)
	}, timeout, finderWaitInterval())
}

// WaitVanish blocks until the pattern disappears or timeout expires.
func (f *Finder) WaitVanish(pattern *Pattern, timeout time.Duration) (bool, error) {
	return SearchWaitVanish(func() (Match, error) {
		return f.Find(pattern)
	}, timeout, finderWaitInterval())
}

// HasNext reports whether the compatibility iterator has another match available.
func (f *Finder) HasNext() bool {
	return f != nil && f.nextIdx >= 0 && f.nextIdx < len(f.last)
}

// Next returns the next compatibility-iterator match and advances the iterator.
// It returns false when the iterator is empty or exhausted.
func (f *Finder) Next() (Match, bool) {
	if !f.HasNext() {
		return Match{}, false
	}
	match := f.last[f.nextIdx]
	f.nextIdx++
	return match, true
}

// Reset rewinds the compatibility iterator to the start of the most recent match set.
func (f *Finder) Reset() {
	if f == nil {
		return
	}
	f.nextIdx = 0
}

// Destroy clears the compatibility iterator state and last-match cache.
func (f *Finder) Destroy() {
	f.clearTraversal()
}

// LastMatches returns a copy of the full most recent match set.
// It does not shrink as the compatibility iterator advances.
func (f *Finder) LastMatches() []Match {
	if len(f.last) == 0 {
		return nil
	}
	out := make([]Match, len(f.last))
	copy(out, f.last)
	return out
}

func (f *Finder) searchMatches(pattern *Pattern, maxResults int) ([]Match, error) {
	req, err := f.buildRequest(pattern, maxResults)
	if err != nil {
		return nil, err
	}
	rawMatches, err := f.matcher.Find(req)
	if err != nil {
		if errors.Is(err, core.ErrMatcherUnsupported) {
			return nil, ErrBackendUnsupported
		}
		return nil, err
	}
	matches := make([]Match, 0, len(rawMatches))
	for i, m := range rawMatches {
		out := toMatch(m, pattern.Offset())
		out.Index = i
		matches = append(matches, out)
	}
	return matches, nil
}

func (f *Finder) setTraversal(matches []Match) {
	if f == nil {
		return
	}
	if len(matches) == 0 {
		f.clearTraversal()
		return
	}
	f.last = matches
	f.nextIdx = 0
}

func (f *Finder) clearTraversal() {
	if f == nil {
		return
	}
	f.last = nil
	f.nextIdx = 0
}

// ReadText runs OCR and returns normalized text.
func (f *Finder) ReadText(params OCRParams) (string, error) {
	result, err := f.readOCR(params)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(result.Text), nil
}

// CollectWords runs OCR and returns word-level geometry for the source image.
func (f *Finder) CollectWords(params OCRParams) ([]OCRWord, error) {
	result, err := f.readOCR(params)
	if err != nil {
		return nil, err
	}
	return wordsFromOCRResult(result), nil
}

// CollectLines runs OCR and groups word-level geometry into line-level results.
func (f *Finder) CollectLines(params OCRParams) ([]OCRLine, error) {
	result, err := f.readOCR(params)
	if err != nil {
		return nil, err
	}
	return linesFromOCRResult(result, f.sourceBounds()), nil
}

// FindText runs OCR and returns word-level matches for the query string.
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
	words := wordsFromOCRResult(result)
	matches := make([]TextMatch, 0)
	for _, word := range words {
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

func (f *Finder) sourceBounds() Rect {
	if f == nil || f.source == nil || f.source.Gray() == nil {
		return Rect{}
	}
	b := f.source.Gray().Bounds()
	return NewRect(b.Min.X, b.Min.Y, b.Dx(), b.Dy())
}

func wordsFromOCRResult(result core.OCRResult) []OCRWord {
	if len(result.Words) == 0 {
		return nil
	}
	words := make([]OCRWord, 0, len(result.Words))
	for i, word := range result.Words {
		words = append(words, OCRWord{
			Rect:       NewRect(word.X, word.Y, word.W, word.H),
			Text:       word.Text,
			Confidence: word.Confidence,
			Index:      i,
		})
	}
	sort.Slice(words, func(i, j int) bool {
		if words[i].Y == words[j].Y {
			return words[i].X < words[j].X
		}
		return words[i].Y < words[j].Y
	})
	for i := range words {
		words[i].Index = i
	}
	return words
}

func linesFromOCRResult(result core.OCRResult, fallback Rect) []OCRLine {
	words := wordsFromOCRResult(result)
	if len(words) == 0 {
		text := strings.TrimSpace(result.Text)
		if text == "" || fallback.Empty() {
			return nil
		}
		parts := strings.Split(text, "\n")
		lines := make([]OCRLine, 0, len(parts))
		for _, part := range parts {
			part = strings.TrimSpace(part)
			if part == "" {
				continue
			}
			lines = append(lines, OCRLine{Rect: fallback, Text: part})
		}
		for i := range lines {
			lines[i].Index = i
		}
		return lines
	}

	type lineBuilder struct {
		words       []OCRWord
		centerYMean float64
		maxHeight   int
	}
	builders := make([]lineBuilder, 0, len(words))
	for _, word := range words {
		centerY := float64(word.Y) + float64(word.H)/2.0
		if len(builders) == 0 {
			builders = append(builders, lineBuilder{words: []OCRWord{word}, centerYMean: centerY, maxHeight: word.H})
			continue
		}
		last := &builders[len(builders)-1]
		limit := float64(max(last.maxHeight, word.H)) * 0.6
		if math.Abs(centerY-last.centerYMean) <= limit {
			last.words = append(last.words, word)
			last.centerYMean = ((last.centerYMean * float64(len(last.words)-1)) + centerY) / float64(len(last.words))
			if word.H > last.maxHeight {
				last.maxHeight = word.H
			}
			continue
		}
		builders = append(builders, lineBuilder{words: []OCRWord{word}, centerYMean: centerY, maxHeight: word.H})
	}

	lines := make([]OCRLine, 0, len(builders))
	for _, builder := range builders {
		sort.Slice(builder.words, func(i, j int) bool { return builder.words[i].X < builder.words[j].X })
		bounds := builder.words[0].Rect
		parts := make([]string, 0, len(builder.words))
		confSum := 0.0
		lineWords := make([]OCRWord, len(builder.words))
		copy(lineWords, builder.words)
		for _, word := range builder.words {
			bounds = unionRect(bounds, word.Rect)
			parts = append(parts, strings.TrimSpace(word.Text))
			confSum += word.Confidence
		}
		line := OCRLine{
			Rect:       bounds,
			Text:       strings.TrimSpace(strings.Join(parts, " ")),
			Confidence: confSum / float64(len(builder.words)),
			Words:      lineWords,
		}
		lines = append(lines, line)
	}
	for i := range lines {
		lines[i].Index = i
	}
	return lines
}

func unionRect(a, b Rect) Rect {
	left := min(a.X, b.X)
	top := min(a.Y, b.Y)
	right := max(a.X+a.W, b.X+b.W)
	bottom := max(a.Y+a.H, b.Y+b.H)
	return NewRect(left, top, right-left, bottom-top)
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
