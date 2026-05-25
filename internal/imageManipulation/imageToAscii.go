package imageManipulation

import (
	"ascii-art/internal/models"
	"image"
)

// 字符集
var asciiChars = []rune("@%#*+=-:. ")

// ImageToAscii 图片转换为ASCII字符矩阵
func ImageToAscii(img image.Image) [][]models.CharPixel {
	bounds := img.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()
	result := make([][]models.CharPixel, height)
	for y := 0; y < height; y++ {
		row := make([]models.CharPixel, width)
		for x := 0; x < width; x++ {
			c := img.At(x+bounds.Min.X, y+bounds.Min.Y)
			gray := RGBToGray(c)
			v := gray / 255.0
			v = v * v
			idx := int((1.0 - v) * float64(len(asciiChars)-1))
			if idx < 0 {
				idx = 0
			}
			if idx >= len(asciiChars) {
				idx = len(asciiChars) - 1
			}
			row[x] = models.CharPixel{
				Char: string(asciiChars[idx]),
				Gray: uint8(gray),
			}
		}
		result[y] = row
	}
	return result
}
