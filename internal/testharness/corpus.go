package testharness

import (
	"embed"
	"encoding/json"
	"fmt"
	"image"
	"image/color"
)

//go:embed testdata/golden_match_cases.json
var corpusFS embed.FS

type GoldenCase struct {
	Name         string          `json:"name"`
	Haystack     [][]int         `json:"haystack"`
	Needle       [][]int         `json:"needle"`
	Mask         [][]int         `json:"mask,omitempty"`
	Threshold    float64         `json:"threshold"`
	ResizeFactor float64         `json:"resize_factor"`
	MaxResults   int             `json:"max_results"`
	Expected     []ExpectedMatch `json:"expected"`
}

type ExpectedMatch struct {
	X        int     `json:"x"`
	Y        int     `json:"y"`
	W        int     `json:"w"`
	H        int     `json:"h"`
	ScoreMin float64 `json:"score_min"`
	ScoreMax float64 `json:"score_max"`
}

func LoadCorpus() ([]GoldenCase, error) {
	raw, err := corpusFS.ReadFile("testdata/golden_match_cases.json")
	if err != nil {
		return nil, fmt.Errorf("read corpus: %w", err)
	}
	var cases []GoldenCase
	if err := json.Unmarshal(raw, &cases); err != nil {
		return nil, fmt.Errorf("parse corpus: %w", err)
	}
	for i := range cases {
		if cases[i].ResizeFactor == 0 {
			cases[i].ResizeFactor = 1.0
		}
	}
	return cases, nil
}

func MatrixToGray(rows [][]int) (*image.Gray, error) {
	if len(rows) == 0 || len(rows[0]) == 0 {
		return nil, fmt.Errorf("matrix is empty")
	}
	w := len(rows[0])
	h := len(rows)
	img := image.NewGray(image.Rect(0, 0, w, h))
	for y := 0; y < h; y++ {
		if len(rows[y]) != w {
			return nil, fmt.Errorf("matrix row width mismatch")
		}
		for x := 0; x < w; x++ {
			v := rows[y][x]
			if v < 0 {
				v = 0
			}
			if v > 255 {
				v = 255
			}
			img.SetGray(x, y, color.Gray{Y: uint8(v)})
		}
	}
	return img, nil
}

