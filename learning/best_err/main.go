package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

// 自定义错误类型
// 暴露具体的错误类型，需要导包强耦合，容易循环依赖问题
func errorsType() {
	_, err := os.Open("/path/to/nonexistent/file")
	if err != nil {
		if pErr, ok := err.(*os.PathError); ok {
			fmt.Printf("Operation: %s\n", pErr.Op)
			fmt.Printf("Path: %s\n", pErr.Path)
			fmt.Printf("Error: %v\n", pErr.Err)
		} else {
			fmt.Println("Non-path error:", err)
		}
	}
	fmt.Println(err)
}

// 不透明错误
// 通过判断行为而不是错误的类型来实现弱耦合
func opaqueError() {
	// 模拟网络连接超时的情况
	_, err := net.DialTimeout("tcp", "example.com:80", 1*time.Second)
	if err != nil {
		if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
			fmt.Println("Timeout error:", netErr)
		} else {
			fmt.Println("Non-timeout error:", err)
		}
	}
}

//func AuthenticateRequest(r *http.Request) error {
//	err := authenticate(r.User)
//	if err != nil {
//		return err
//	}
//	return nil
//	//return authenticate(r.User)
//}
