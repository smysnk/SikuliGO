//go:build gogosseract

package ocr

import (
	"bytes"
	"context"
	"fmt"
	"html"
	"image/png"
	"reflect"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/sikulix/portgo/internal/core"
	"github.com/danlock/gogosseract"
)

var (
	wordSpanPattern = regexp.MustCompile(`(?is)<span[^>]*class=["'][^"']*ocrx_word[^"']*["'][^>]*title=["'][^"']*bbox\s+(\d+)\s+(\d+)\s+(\d+)\s+(\d+)(?:;[^"']*x_wconf\s+(\d+))?[^"']*["'][^>]*>(.*?)</span>`)
	tagPattern      = regexp.MustCompile(`(?s)<[^>]*>`)
	spacePattern    = regexp.MustCompile(`\s+`)
)

type gogosseractBackend struct{}

func New() core.OCR {
	return &gogosseractBackend{}
}

func (b *gogosseractBackend) Read(req core.OCRRequest) (core.OCRResult, error) {
	if err := req.Validate(); err != nil {
		return core.OCRResult{}, err
	}

	ctx := context.Background()
	cancel := func() {}
	if req.Timeout > 0 {
		ctx, cancel = context.WithTimeout(ctx, req.Timeout)
	}
	defer cancel()

	cfg := gogosseract.Config{}
	setStructStringField(&cfg, "Language", req.Language)
	if req.TrainingDataPath != "" {
		setStructStringField(&cfg, "TrainingData", req.TrainingDataPath)
		setStructStringField(&cfg, "TrainingDataPath", req.TrainingDataPath)
	}

	engine, err := gogosseract.New(ctx, cfg)
	if err != nil {
		return core.OCRResult{}, err
	}
	defer callClose(ctx, engine)

	var imageBuf bytes.Buffer
	if err := png.Encode(&imageBuf, req.Image); err != nil {
		return core.OCRResult{}, err
	}

	if err := callMethodError(engine, "LoadImage", ctx, bytes.NewReader(imageBuf.Bytes())); err != nil {
		return core.OCRResult{}, err
	}
	text, err := callMethodString(engine, "GetText", ctx)
	if err != nil {
		return core.OCRResult{}, err
	}

	hocr, err := callMethodString(engine, "GetHOCR", ctx)
	if err != nil {
		// Some backends may not expose hOCR; full-text OCR still remains available.
		hocr = ""
	}
	base := req.Image.Bounds().Min
	words := parseHOCRWords(hocr, req.MinConfidence, base.X, base.Y)

	return core.OCRResult{
		Text:  strings.TrimSpace(text),
		Words: words,
	}, nil
}

func setStructStringField(target any, fieldName, value string) {
	if strings.TrimSpace(value) == "" {
		return
	}
	v := reflect.ValueOf(target)
	if v.Kind() != reflect.Pointer || v.IsNil() {
		return
	}
	elem := v.Elem()
	if elem.Kind() != reflect.Struct {
		return
	}
	field := elem.FieldByName(fieldName)
	if !field.IsValid() || !field.CanSet() || field.Kind() != reflect.String {
		return
	}
	field.SetString(value)
}

func callClose(ctx context.Context, engine any) {
	_ = callMethodError(engine, "Close", ctx)
}

func callMethodError(target any, method string, args ...any) error {
	out, err := callMethod(target, method, args...)
	if err != nil {
		return err
	}
	if len(out) == 0 {
		return nil
	}
	last := out[len(out)-1]
	if !last.IsValid() || isNilValue(last) {
		return nil
	}
	e, ok := last.Interface().(error)
	if !ok {
		return nil
	}
	return e
}

func callMethodString(target any, method string, args ...any) (string, error) {
	out, err := callMethod(target, method, args...)
	if err != nil {
		return "", err
	}
	if len(out) == 0 {
		return "", nil
	}
	if len(out) == 1 {
		if s, ok := out[0].Interface().(string); ok {
			return s, nil
		}
		if out[0].Type().Implements(reflect.TypeOf((*error)(nil)).Elem()) {
			if isNilValue(out[0]) {
				return "", nil
			}
			return "", out[0].Interface().(error)
		}
		return "", fmt.Errorf("%s returned unsupported output shape", method)
	}

	first := out[0]
	last := out[len(out)-1]

	if last.Type().Implements(reflect.TypeOf((*error)(nil)).Elem()) && !isNilValue(last) {
		return "", last.Interface().(error)
	}
	if s, ok := first.Interface().(string); ok {
		return s, nil
	}
	return "", fmt.Errorf("%s did not return string output", method)
}

func isNilValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Pointer, reflect.Slice:
		return v.IsNil()
	default:
		return false
	}
}

func callMethod(target any, method string, args ...any) ([]reflect.Value, error) {
	v := reflect.ValueOf(target)
	m := v.MethodByName(method)
	if !m.IsValid() {
		return nil, fmt.Errorf("missing method %q", method)
	}
	mt := m.Type()
	in := make([]reflect.Value, 0, mt.NumIn())
	for i := 0; i < mt.NumIn(); i++ {
		paramType := mt.In(i)
		if i < len(args) {
			argVal := reflect.ValueOf(args[i])
			if !argVal.IsValid() {
				in = append(in, reflect.Zero(paramType))
				continue
			}
			if argVal.Type().AssignableTo(paramType) {
				in = append(in, argVal)
				continue
			}
			if argVal.Type().ConvertibleTo(paramType) {
				in = append(in, argVal.Convert(paramType))
				continue
			}
			if paramType.Kind() == reflect.Interface && argVal.Type().Implements(paramType) {
				in = append(in, argVal)
				continue
			}
			return nil, fmt.Errorf("method %q argument %d has incompatible type", method, i)
		}
		in = append(in, reflect.Zero(paramType))
	}
	return m.Call(in), nil
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
