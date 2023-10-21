package main

import (
	"fmt"
	"os"

	"github.com/lu4p/unipdf/v3/extractor"
	pdf "github.com/lu4p/unipdf/v3/model"
)

func main() {
	f, err := os.Open("/Volumes/G盘/test/lcq.pdf")
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
