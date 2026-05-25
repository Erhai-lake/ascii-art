package imageManipulation

import "image/color"

// RGBToGray 人眼感知亮度模型
func RGBToGray(c color.Color) float64 {
	r, g, b, _ := c.RGBA()
	R := float64(r >> 8)
	G := float64(g >> 8)
	B := float64(b >> 8)
	return 0.2126*R + 0.7152*G + 0.0722*B
}
