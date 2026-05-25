package imageManipulation

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

func ConvertToGrayAndSave(imgPath, outputPath string) {
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
	bounds := img.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()
	grayImg := image.NewGray(image.Rect(0, 0, width, height))
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			c := img.At(x, y)
			grayValue := RGBToGray(c)
			grayImg.SetGray(x, y, color.Gray{Y: uint8(grayValue)})
		}
	}
	outFile, err := os.Create(outputPath)
	if err != nil {
		log.Fatal(err)
	}
	defer func(outFile *os.File) {
		_ = outFile.Close()
	}(outFile)
	if err := png.Encode(outFile, grayImg); err != nil {
		log.Fatal(err)
	}
	log.Printf("灰度图片已保存到: %s\n", outputPath)
}
