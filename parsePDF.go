package main

import (
	"fmt"
	"github.com/lu4p/unipdf/v3/extractor"
	pdf "github.com/lu4p/unipdf/v3/model"
	"os"
	"path/filepath"
	"strings"
)

// 1.循环读取文件夹
func parsePdfFile(path string) {
	err := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			return nil
		}
		regexpss(path)
		return nil
	})
	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}
}

func regexpss(text string) {
	if !strings.HasSuffix(text, "pdf") {
		return
	}

	// 文件名作为标题，打印出来
	split := strings.Split(text, "｜")
	if len(split) < 2 {
		fmt.Println("无法分割，不处理：", text)
		return
	}

	suffix := strings.TrimSuffix(split[len(split)-1], ".pdf")
	fmt.Println(suffix) // 打印标题
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	parse(text)
}

func parse(fileName string) {
	f, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer f.Close()

	pdfReader, err := pdf.NewPdfReader(f)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	numPages, err := pdfReader.GetNumPages()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for i := 1; i <= numPages; i++ {
		page, _ := pdfReader.GetPage(i)
		ex, _ := extractor.New(page)
		text, _ := ex.ExtractText()
		fmt.Println(text)
	}
}
