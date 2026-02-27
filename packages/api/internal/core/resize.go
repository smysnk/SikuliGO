package core

import "image"

func ResizeGrayNearest(src *image.Gray, factor float64) *image.Gray {
	if src == nil {
		return nil
	}
	if factor == 1.0 {
		dst := image.NewGray(src.Bounds())
		copy(dst.Pix, src.Pix)
		return dst
	}
	srcB := src.Bounds()
	srcW := srcB.Dx()
	srcH := srcB.Dy()
	if srcW == 0 || srcH == 0 {
		return image.NewGray(image.Rect(0, 0, 0, 0))
	}
	dstW := int(float64(srcW)*factor + 0.5)
	dstH := int(float64(srcH)*factor + 0.5)
	if dstW < 1 {
		dstW = 1
	}
	if dstH < 1 {
		dstH = 1
	}
	dst := image.NewGray(image.Rect(0, 0, dstW, dstH))
	for y := 0; y < dstH; y++ {
		srcY := int(float64(y) / factor)
		if srcY >= srcH {
			srcY = srcH - 1
		}
		for x := 0; x < dstW; x++ {
			srcX := int(float64(x) / factor)
			if srcX >= srcW {
				srcX = srcW - 1
			}
			dst.SetGray(x, y, src.GrayAt(srcB.Min.X+srcX, srcB.Min.Y+srcY))
		}
	}
	return dst
}

