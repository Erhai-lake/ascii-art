# ASCII Art Generator - ASCII艺术生成器

这是一个简单的 ASCII 艺术生成器, 为集成到我的API站点中做技术测试.

## 功能特性

- 将图片转换为 ASCII 艺术
- 支持多种导出格式: HTML/TXT/灰度图
- 可自定义缩放比例

## 项目结构

```
ascii-art/
├── cmd/
│   └── server/
│       └── main.go          # 主程序入口
├── internal/
│   ├── export/              # 导出功能
│   │   ├── html.go
│   │   ├── template.html
│   │   └── txt.go
│   ├── imageManipulation/   # 图像处理
│   │   ├── grayExport.go
│   │   ├── imageToAscii.go
│   │   ├── processor.go
│   │   ├── resizeImage.go
│   │   └── rgbToGray.go
│   └── models/              # 数据模型
│       └── charPixel.go
├── go.mod
└── go.sum
```

## 使用方法

```bash
# 基本用法
go run cmd/server/main.go <图片路径> <缩放比例>

# 导出为 HTML (默认)
go run cmd/server/main.go test.png 0.5

# 导出为 TXT
go run cmd/server/main.go test.png 0.5 txt

# 导出为灰度图
go run cmd/server/main.go test.png 0.5 gray
```

## 技术栈

- Go 1.26.1
- github.com/nfnt/resize
