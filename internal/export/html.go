package export

import (
	"ascii-art/internal/models"
	_ "embed"
	"log"
	"os"
	"strings"
	"text/template"
)

//go:embed template.html
var HTMLTemplate string

// HTMLData 导出HTML数据结构体
type HTMLData struct {
	Content   string
	Thumbnail string
}

// GenerateHTMLContent 生成HTML内容
func GenerateHTMLContent(ascii [][]models.CharPixel) string {
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

// HTML 导出HTML文件
func HTML(content string, thumbnail string, filename string) {
	t, err := template.New("ascii").Parse(HTMLTemplate)
	if err != nil {
		log.Fatal("解析模板失败: ", err)
	}
	outFile, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer func(outFile *os.File) {
		_ = outFile.Close()
	}(outFile)
	data := HTMLData{
		Content:   content,
		Thumbnail: thumbnail,
	}
	if err := t.Execute(outFile, data); err != nil {
		log.Fatal(err)
	}
}
