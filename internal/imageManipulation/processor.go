package imageManipulation

import (
	"ascii-art/internal/models"
	"bytes"
	"encoding/base64"
	"image"
	_ "image/jpeg"
	"image/png"
	_ "image/png"
	"log"
	"os"

	"github.com/nfnt/resize"
)

// ProcessImage 处理图片，返回ASCII字符矩阵
func ProcessImage(imgPath string, scale float64) [][]models.CharPixel {
	file, err := os.Open(imgPath)
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)
	img, _, err := image.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	resized := ResizeImage(img, scale)
	ascii := ImageToAscii(resized)
	return ascii
}

// GenerateThumbnailBase64 生成图片缩略图的Base64编码
func GenerateThumbnailBase64(imgPath string, maxSize int) string {
	file, err := os.Open(imgPath)
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)
	img, _, err := image.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	width := img.Bounds().Dx()
	height := img.Bounds().Dy()
	aspectRatio := float64(width) / float64(height)
	var newWidth, newHeight uint
	if width > height {
		newWidth = uint(maxSize)
		newHeight = uint(float64(maxSize) / aspectRatio)
	} else {
		newHeight = uint(maxSize)
		newWidth = uint(float64(maxSize) * aspectRatio)
	}
	thumbnail := resize.Resize(newWidth, newHeight, img, resize.Lanczos3)
	var buf bytes.Buffer
	err = png.Encode(&buf, thumbnail)
	if err != nil {
		log.Fatal(err)
	}
	return "data:image/png;base64," + base64.StdEncoding.EncodeToString(buf.Bytes())
}
