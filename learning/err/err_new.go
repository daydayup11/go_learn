package main

import (
	"fmt"
	"github.com/pkg/errors"
	"os"
)

// FoundFile 表示找到文件时的错误类型
type FoundFile struct {
	Filename string
	Err      error
}

// 实现 Unwrap 方法
func (e *FoundFile) Unwrap() error {
	return e.Err
}

// 实现 error 接口的 Error 方法
func (e *FoundFile) Error() string {
	return fmt.Sprintf("%s: %v", e.Filename, e.Err)
}

var ErrNoPermission = fmt.Errorf("no permission to open file")
var ErrFileNotFound = fmt.Errorf("file not found")

func openFile() (*os.File, error) {
	file, err := os.Open("example.txt")
	if err != nil {
		return nil, &FoundFile{"example.txt", ErrFileNotFound}
	}
	return file, nil
}

func readFile() ([]byte, error) {
	file, err := openFile()
	if err != nil {
		return nil, errors.Wrap(err, "openFile 函数出错")
	}
	defer file.Close()

	return nil, nil
}

func main() {
	_, err := readFile()

	if errors.Is(err, ErrFileNotFound) {
		fmt.Printf("%+v", err)
	}
	//// 使用新的错误类型 FoundFile 进行处理
	//var foundFile *FoundFile
	//if ok := errors.As(err, &foundFile); ok {
	//	fmt.Printf("Encountered ErrFileNotFound for file: %V\n", err)
	//	// 处理文件不存在的情况
	//} else {
	//	// 处理其他错误
	//	fmt.Println("Encountered an unknown error:", err)
	//}
}
