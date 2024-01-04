package main

import (
	"fmt"

	"time"
)

// 超时控制
func main() {
	select {
	case re := <-AsynService():
		fmt.Print("任务完成", re)
	case <-time.After(time.Millisecond * 3000):
		fmt.Print("超时啦")
		//default:
		//	fmt.Print("不能阻塞")
	}

}

// 服务
func service() string {
	time.Sleep(time.Millisecond * 3000)
	return "finish"
}

// 异步启动服务
func AsynService() chan string {
	rechan := make(chan string, 1)
	go func() {
		res := service()
		time.Sleep(time.Millisecond * 1000)
		rechan <- res
	}()
	return rechan
}
