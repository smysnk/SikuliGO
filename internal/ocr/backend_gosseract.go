//go:build gosseract

package ocr

import (
	"bytes"
	"fmt"
	"html"
	"image/png"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/otiai10/gosseract/v2"
	"github.com/sikulix/portgo/internal/core"
)

var (
	wordSpanPattern = regexp.MustCompile(`(?is)<span[^>]*class=["'][^"']*ocrx_word[^"']*["'][^>]*title=["'][^"']*bbox\s+(\d+)\s+(\d+)\s+(\d+)\s+(\d+)(?:;[^"']*x_wconf\s+(\d+))?[^"']*["'][^>]*>(.*?)</span>`)
	tagPattern      = regexp.MustCompile(`(?s)<[^>]*>`)
	spacePattern    = regexp.MustCompile(`\s+`)
	tessdataEnvMu   sync.Mutex
)

type gosseractBackend struct{}

func New() core.OCR {
	return &gosseractBackend{}
}

func (b *gosseractBackend) Read(req core.OCRRequest) (core.OCRResult, error) {
	if err := req.Validate(); err != nil {
		return core.OCRResult{}, err
	}
	if req.Timeout <= 0 {
		return readWithGosseract(req)
	}

	type result struct {
		out core.OCRResult
		err error
	}
	done := make(chan result, 1)
	go func() {
		out, err := readWithGosseract(req)
		done <- result{out: out, err: err}
	}()

	select {
	case got := <-done:
		return got.out, got.err
	case <-time.After(req.Timeout):
		return core.OCRResult{}, fmt.Errorf("ocr timed out after %s", req.Timeout)
	}
}

func readWithGosseract(req core.OCRRequest) (core.OCRResult, error) {
	return withTessdataPrefix(req.TrainingDataPath, func() (core.OCRResult, error) {
		client := gosseract.NewClient()
		defer client.Close()

		if err := client.SetLanguage(req.Language); err != nil {
			return core.OCRResult{}, err
		}

		var imageBuf bytes.Buffer
		if err := png.Encode(&imageBuf, req.Image); err != nil {
			return core.OCRResult{}, err
		}
		if err := client.SetImageFromBytes(imageBuf.Bytes()); err != nil {
			return core.OCRResult{}, err
		}

		text, err := client.Text()
		if err != nil {
			return core.OCRResult{}, err
		}

		hocr, err := client.HOCRText()
		if err != nil {
			// hOCR may be unavailable in some environments, while plain text can still work.
			hocr = ""
		}

		base := req.Image.Bounds().Min
		words := parseHOCRWords(hocr, req.MinConfidence, base.X, base.Y)

		return core.OCRResult{
			Text:  strings.TrimSpace(text),
			Words: words,
		}, nil
	})
}

func withTessdataPrefix(prefix string, fn func() (core.OCRResult, error)) (core.OCRResult, error) {
	if strings.TrimSpace(prefix) == "" {
		return fn()
	}

	tessdataEnvMu.Lock()
	defer tessdataEnvMu.Unlock()

	prev, hadPrev := os.LookupEnv("TESSDATA_PREFIX")
	if err := os.Setenv("TESSDATA_PREFIX", prefix); err != nil {
		return core.OCRResult{}, err
	}
	defer func() {
		if hadPrev {
			_ = os.Setenv("TESSDATA_PREFIX", prev)
			return
		}
		_ = os.Unsetenv("TESSDATA_PREFIX")
	}()

	return fn()
}

func parseHOCRWords(hocr string, minConfidence float64, baseX, baseY int) []core.OCRWord {
	if strings.TrimSpace(hocr) == "" {
		return nil
	}

	found := wordSpanPattern.FindAllStringSubmatch(hocr, -1)
	if len(found) == 0 {
		return nil
	}

	out := make([]core.OCRWord, 0, len(found))
	for _, m := range found {
		if len(m) < 7 {
			continue
		}
		x0, err0 := strconv.Atoi(m[1])
		y0, err1 := strconv.Atoi(m[2])
		x1, err2 := strconv.Atoi(m[3])
		y1, err3 := strconv.Atoi(m[4])
		if err0 != nil || err1 != nil || err2 != nil || err3 != nil || x1 <= x0 || y1 <= y0 {
			continue
		}

		conf := 1.0
		if len(m) > 5 && strings.TrimSpace(m[5]) != "" {
			if v, err := strconv.Atoi(m[5]); err == nil {
				conf = float64(v) / 100.0
			}
		}
		if conf < minConfidence {
			continue
		}

		raw := strings.TrimSpace(m[6])
		clean := strings.TrimSpace(spacePattern.ReplaceAllString(html.UnescapeString(tagPattern.ReplaceAllString(raw, " ")), " "))
		if clean == "" {
			continue
		}

		out = append(out, core.OCRWord{
			Text:       clean,
			X:          baseX + x0,
			Y:          baseY + y0,
			W:          x1 - x0,
			H:          y1 - y0,
			Confidence: conf,
		})
	}

	sort.Slice(out, func(i, j int) bool {
		if out[i].Y == out[j].Y {
			return out[i].X < out[j].X
		}
		return out[i].Y < out[j].Y
	})
	return out
}
