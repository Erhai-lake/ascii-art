package imageManipulation

import (
	"image"

	"github.com/nfnt/resize"
)

// ResizeImage 缩放图片, 保持纵横比
func ResizeImage(img image.Image, scale float64) image.Image {
	srcBounds := img.Bounds()
	width := uint(float64(srcBounds.Dx()) * scale)
	height := uint(float64(srcBounds.Dy()) * scale * 0.6)
	if height == 0 {
		height = 1
	}
	resized := resize.Resize(width, height, img, resize.Lanczos3)
	return resized
}
