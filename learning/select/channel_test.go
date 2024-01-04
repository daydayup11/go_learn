package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var wg = sync.WaitGroup{}

func TestChannel(t *testing.T) {
	cn := make(chan int, 1)
	wg.Add(1)
	producer(cn)
	wg.Add(1)
	consumer(cn)
	wg.Wait()
	fmt.Print("任务完成")
}

// 1.消息生产者向通道发送消息，阻塞式
func producer(cn chan int) {
	//好嘛，不异步就会死锁了
	go func() {
		for i := 0; i < 10; i++ {
			cn <- i
		}
		//不关闭就死锁了，接受者会一直等待
		close(cn)
		fmt.Println("发送完毕")
		wg.Done()
	}()

}

// 2.消息消费者从通道获取消费消息
func consumer(cn chan int) {
	//不异步就阻塞主线程了
	go func() {
		for {
			if data, ok := <-cn; ok {
				fmt.Println("消费完了一个消息：", data)
			} else {
				break
			}
		}
		fmt.Print("消费完毕")
		wg.Done()
	}()

}

// 测试任务取消
func TestCancel(t *testing.T) {
	cn := make(chan struct{})
	for i := 1; i < 6; i++ {
		go func(cn chan struct{}) {
			for {
				if isCanceled(cn) {
					break
				} else {
					time.Sleep(time.Millisecond * 1000)
				}
			}
			fmt.Print("任务取消")
		}(cn)
	}
	cancel2(cn)

}

// 1.普通向cancel通道发送取消通知，这种做法需要事先知道有多少个正在执行的任务
func cancel1(cn chan struct{}) {
	cn <- struct{}{}
}

// 2.向采取close方法关闭所有任务
func cancel2(cn chan struct{}) {
	close(cn)
}

// 3.select判断任务是否取消
func isCanceled(cn chan struct{}) bool {
	select {
	case <-cn:
		fmt.Println("任务取消")
		return true
	default:
		fmt.Print("任务不取消，继续执行")
		return false
	}
}
