package sikuli

import (
	"strings"
	"time"
)

const DefaultOCRLanguage = "eng"

type OCRParams struct {
	Language         string
	TrainingDataPath string
	MinConfidence    float64
	Timeout          time.Duration
	CaseSensitive    bool
}

type TextMatch struct {
	Rect
	Text       string
	Confidence float64
	Index      int
}

func normalizeOCRParams(in OCRParams) OCRParams {
	out := in
	if strings.TrimSpace(out.Language) == "" {
		out.Language = DefaultOCRLanguage
	}
	if out.MinConfidence < 0 {
		out.MinConfidence = 0
	}
	if out.MinConfidence > 1 {
		out.MinConfidence = 1
	}
	if out.Timeout < 0 {
		out.Timeout = 0
	}
	return out
}

func containsText(haystack, needle string, caseSensitive bool) bool {
	if caseSensitive {
		return strings.Contains(haystack, needle)
	}
	return strings.Contains(strings.ToLower(haystack), strings.ToLower(needle))
}
