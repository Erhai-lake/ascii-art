package export

import (
	"ascii-art/internal/models"
	"log"
	"os"
	"strings"
)

// GenerateTXTContent 生成TXT内容
func GenerateTXTContent(ascii [][]models.CharPixel) string {
	var contentBuilder strings.Builder
	for y, row := range ascii {
		for _, pixel := range row {
			contentBuilder.WriteString(pixel.Char)
		}
		if y < len(ascii)-1 {
			contentBuilder.WriteString("\n")
		}
	}
	return contentBuilder.String()
}

// TXT 导出TXT文件
func TXT(content string, filename string) {
	outFile, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer func(outFile *os.File) {
		err := outFile.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(outFile)
	if _, err := outFile.WriteString(content); err != nil {
		log.Fatal(err)
	}
}
