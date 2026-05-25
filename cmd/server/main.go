package main

import (
	"ascii-art/internal/export"
	"ascii-art/internal/imageManipulation"
	"ascii-art/internal/models"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// ThumbnailMaxSize 缩略图最大尺寸
const ThumbnailMaxSize = 200

// main 主函数
func main() {
	if len(os.Args) < 3 {
		printUsage()
		return
	}

	imgPath := os.Args[1]
	scale := parseScale(os.Args[2])
	if scale == 0 {
		return
	}

	format := "html"
	if len(os.Args) >= 4 {
		format = strings.ToLower(os.Args[3])
	}

	if format == "gray" {
		baseName := getBaseName(imgPath)
		imageManipulation.ConvertToGrayAndSave(imgPath, baseName+"_gray.png")
		return
	}

	if format != "html" && format != "txt" {
		fmt.Println("格式必须是 html, txt, 或 gray")
		return
	}

	ascii := imageManipulation.ProcessImage(imgPath, scale)
	baseName := getBaseName(imgPath)

	exportResult(format, imgPath, ascii, baseName)
}

// printUsage 打印说明书
func printUsage() {
	fmt.Println("用法: go run main.go <图片路径> <缩放比例> [格式: html|txt|gray]")
	fmt.Println("缩放比例: 大于0的数字，小于1缩小，大于1放大")
	fmt.Println("示例: go run main.go test.png 0.5")
	fmt.Println("示例: go run main.go test.png 2.0")
	fmt.Println("示例: go run main.go test.png 0.5 txt")
	fmt.Println("示例: go run main.go test.png 0.5 gray")
}

// parseScale 解析缩放比例
func parseScale(s string) float64 {
	scale, err := strconv.ParseFloat(s, 64)
	if err != nil || scale <= 0 {
		fmt.Println("缩放比例必须是大于 0 的数字")
		return 0
	}
	return scale
}

// getBaseName 获取图片基础名
func getBaseName(imgPath string) string {
	return strings.TrimSuffix(filepath.Base(imgPath), filepath.Ext(imgPath))
}

// exportResult 导出结果
func exportResult(format string, imgPath string, ascii [][]models.CharPixel, baseName string) {
	switch format {
	case "html":
		content := export.GenerateHTMLContent(ascii)
		thumbnail := imageManipulation.GenerateThumbnailBase64(imgPath, ThumbnailMaxSize)
		export.HTML(content, thumbnail, baseName+".html")
		fmt.Println("生成完成:", baseName+".html")
	case "txt":
		content := export.GenerateTXTContent(ascii)
		export.TXT(content, baseName+".txt")
		fmt.Println("生成完成:", baseName+".txt")
	}
}
