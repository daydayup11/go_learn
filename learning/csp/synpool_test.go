package main

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
)

// 测试syn.Pool
// 创建、获取、放入、GC、获取
func TestSyn(t *testing.T) {
	pool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("create new obj")
			return 100
		},
	}
	get := pool.Get().(int)
	fmt.Println(get)
	pool.Put(3)
	runtime.GC()
	//go 1.13版本之后对回收策略改成上一个GC周期被使用过不回收
	runtime.GC()
	obj := pool.Get().(int)
	fmt.Println(obj)
}

func TestSyncPoolInMultiGoroutine(t *testing.T) {
	pool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("Create a new object.")
			return 10
		},
	}

	pool.Put(100)
	pool.Put(100)
	pool.Put(100)

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			fmt.Println(pool.Get())
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println(pool.Get())
}
