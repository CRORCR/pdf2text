package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

const Regexp = "(<span data-slate-string=\"true\">)(.*?)(</span>)"

func parseHtmlFile(path string) {
	err := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			return nil
		}

		// 1、搜索并处理所有html文件
		if !strings.HasSuffix(f.Name(), ".html") {
			return nil
		}

		newFileName := paramsFileName(f.Name()) // 输出 文件名称
		fmt.Println("处理文件:", newFileName)

		// 2、创建一个txt文件，并把标题写进去
		newFile := createTxtFile(path)
		if newFile == nil {
			fmt.Println("创建新文件失败,err:", err)
			return nil
		}
		// 结束之后，关闭新文件
		defer newFile.Close()

		// 3、写入标题
		newFile.WriteString(newFileName)

		// 4、读取所有html文件内容
		text, err := paramsAndGetText(path)
		if err != nil {
			return nil
		}

		// 4、开始处理源文件，把html的内容处理之后，写入新文件
		regexpText(newFile, text)
		return nil
	})

	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}
}

// 文件名称处理：1.去掉html  2.去掉首尾空格
func paramsFileName(oldName string) string {
	oldName = strings.TrimSuffix(oldName, ".html")
	return strings.TrimSpace(oldName)
}

// 读取文件，并返回全文内容
func paramsAndGetText(filePath string) (string, error) {
	// 1、读取文件
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Println("read file fail", err)
		return "", err
	}
	defer f.Close()

	//fmt.Println("文件名称:", f.Name()) // 文件名称: /Users/pencil/go/src/inke_work/params_html/demo.html
	fd, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println("read to fd fail", err)
		return "", err
	}
	return string(fd), nil
}

// 创建一个txt结尾的文件，用来存储解析后的数据
func createTxtFile(oldName string) *os.File {
	newFileName := strings.TrimSuffix(oldName, ".html")
	newFileName = newFileName + ".txt"
	file, err := os.OpenFile(newFileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("写入文件有误", err)
	}
	return file
}

// 2.正则匹配
func regexpText(newFile *os.File, text string) {
	compile := regexp.MustCompile(Regexp)
	match := compile.FindAllStringSubmatch(text, -1) //findAll 匹配所有 参数-1表示所有
	// fmt.Println("match", match)

	for _, value := range match {
		if len(value) > 2 {
			//fmt.Println("---", len(value), "---", value)
			// 处理的时候，写入txt文件中
			newFile.Write([]byte("\r\n"))
			newFile.WriteString(value[2])
		}
	}
}
