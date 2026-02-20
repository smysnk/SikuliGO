package sikuli

import (
	"fmt"
	"image"
	"image/color"
)

type Image struct {
	name string
	gray *image.Gray
}

func NewImageFromGray(name string, src *image.Gray) (*Image, error) {
	if src == nil {
		return nil, fmt.Errorf("%w: image is nil", ErrInvalidTarget)
	}
	cloned := image.NewGray(src.Bounds())
	copy(cloned.Pix, src.Pix)
	return &Image{name: name, gray: cloned}, nil
}

func NewImageFromAny(name string, src image.Image) (*Image, error) {
	if src == nil {
		return nil, fmt.Errorf("%w: image is nil", ErrInvalidTarget)
	}
	b := src.Bounds()
	g := image.NewGray(b)
	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			g.Set(x, y, color.GrayModel.Convert(src.At(x, y)))
		}
	}
	return &Image{name: name, gray: g}, nil
}

func NewImageFromMatrix(name string, rows [][]uint8) (*Image, error) {
	if len(rows) == 0 || len(rows[0]) == 0 {
		return nil, fmt.Errorf("%w: matrix is empty", ErrInvalidTarget)
	}
	w := len(rows[0])
	h := len(rows)
	g := image.NewGray(image.Rect(0, 0, w, h))
	for y := 0; y < h; y++ {
		if len(rows[y]) != w {
			return nil, fmt.Errorf("%w: matrix row width mismatch", ErrInvalidTarget)
		}
		for x := 0; x < w; x++ {
			g.SetGray(x, y, color.Gray{Y: rows[y][x]})
		}
	}
	return &Image{name: name, gray: g}, nil
}

func (i *Image) Name() string {
	return i.name
}

func (i *Image) Width() int {
	if i == nil || i.gray == nil {
		return 0
	}
	return i.gray.Bounds().Dx()
}

func (i *Image) Height() int {
	if i == nil || i.gray == nil {
		return 0
	}
	return i.gray.Bounds().Dy()
}

func (i *Image) Gray() *image.Gray {
	if i == nil {
		return nil
	}
	return i.gray
}

func (i *Image) Clone() *Image {
	if i == nil || i.gray == nil {
		return nil
	}
	copyGray := image.NewGray(i.gray.Bounds())
	copy(copyGray.Pix, i.gray.Pix)
	return &Image{name: i.name, gray: copyGray}
}

func (i *Image) Crop(rect Rect) (*Image, error) {
	if i == nil || i.gray == nil {
		return nil, fmt.Errorf("%w: image is nil", ErrInvalidTarget)
	}
	if rect.Empty() {
		return nil, fmt.Errorf("%w: crop rect is empty", ErrInvalidTarget)
	}
	srcBounds := i.gray.Bounds()
	crop := image.Rect(rect.X, rect.Y, rect.X+rect.W, rect.Y+rect.H).Intersect(srcBounds)
	if crop.Empty() {
		return nil, fmt.Errorf("%w: crop rect outside source bounds", ErrInvalidTarget)
	}

	out := image.NewGray(crop)
	for y := crop.Min.Y; y < crop.Max.Y; y++ {
		srcStart := i.gray.PixOffset(crop.Min.X, y)
		srcEnd := i.gray.PixOffset(crop.Max.X, y)
		dstStart := out.PixOffset(crop.Min.X, y)
		copy(out.Pix[dstStart:dstStart+crop.Dx()], i.gray.Pix[srcStart:srcEnd])
	}

	name := fmt.Sprintf("%s@crop[%d,%d %dx%d]", i.name, crop.Min.X, crop.Min.Y, crop.Dx(), crop.Dy())
	return &Image{name: name, gray: out}, nil
}
