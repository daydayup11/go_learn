package main

import (
	"github.com/pkg/errors"
)

//func main() {
//	err := originalFunction()
//	if err != nil {
//		fmt.Printf("Error: %+v\n", err)
//	}
//}

func originalFunction() error {
	err := someFunction()
	if err != nil {
		// 使用 errors.Wrap 包装错误并添加堆栈信息
		return errors.Wrap(err, "originalFunction 发生错误")
		//return fmt.Errorf(" %w", err)
	}
	return nil
}

func someFunction() error {
	// 模拟一个出错的情况
	return errors.New("这是一个模拟的错误")
}
