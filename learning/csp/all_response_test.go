package main

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

// 开十个go程去执行任务，当所有go程返回结果就返回给用户
// 通过channel接受完数据才返回
// 注意防止协程泄露这里需要使用buffer channel
func AllResponse() string {
	numRunner := 10
	cn := make(chan string, numRunner)
	for i := 0; i < numRunner; i++ {
		go func(i int) {
			ret := task(i)
			cn <- ret
		}(i)
	}
	//return <-cn
	var res string
	for i := 0; i < numRunner; i++ {
		res = res + <-cn
	}
	return res
}
func task2(id int) string {
	time.Sleep(time.Millisecond * 100)
	//都会调用这个方法，但是只有第一个返回的go程的值才会被取用
	return fmt.Sprintf("result from %d", id)
}

func TestAll(t *testing.T) {
	t.Log("Before:", runtime.NumGoroutine())
	t.Log(AllResponse())
	time.Sleep(time.Second * 3)
	t.Log("After:", runtime.NumGoroutine())
}
