//go:build gosseract || gogosseract

package ocr

import "testing"

func TestParseHOCRWordsConformance(t *testing.T) {
	hocr := `
<div>
  <span class="ocrx_word" title="bbox 20 10 30 20; x_wconf 70">third</span>
  <span class="ocrx_word" title="bbox 5 5 15 15; x_wconf 95">first</span>
  <span class="ocrx_word" title="bbox 40 5 50 15; x_wconf 40">lowconf</span>
  <span class="ocrx_word" title="bbox 16 5 25 15; x_wconf 90">second</span>
</div>`

	words := parseHOCRWords(hocr, 0.6, 100, 200)
	if len(words) != 3 {
		t.Fatalf("expected 3 words after confidence filtering, got=%d", len(words))
	}

	if words[0].Text != "first" || words[0].X != 105 || words[0].Y != 205 {
		t.Fatalf("first word mismatch: %+v", words[0])
	}
	if words[1].Text != "second" || words[1].X != 116 || words[1].Y != 205 {
		t.Fatalf("second word mismatch: %+v", words[1])
	}
	if words[2].Text != "third" || words[2].X != 120 || words[2].Y != 210 {
		t.Fatalf("third word mismatch: %+v", words[2])
	}
}
