package main

import (
	"errors"
	"fmt"
	"sync"
	"testing"
	"time"
)

// 可复用对象
type Reusable struct {
	name string
}

// 对象池
type ObjPool struct {
	bufChan chan *Reusable
}

// 创建对象池的方法
func NewObjPool(size int) *ObjPool {
	objPool := new(ObjPool)
	objPool.bufChan = make(chan *Reusable, size)
	for i := 0; i < size; i++ {
		objPool.bufChan <- &Reusable{
			name: fmt.Sprint(i),
		}
	}
	return objPool
}

// 获取对象的方法
// 需要加上超时控制和返回错误值
func getObj(pool *ObjPool) (*Reusable, error) {
	select {
	case res := <-pool.bufChan:
		return res, nil
	case <-time.After(time.Millisecond * 2):
		return nil, errors.New("获取对象超时")
	}
}

// 释放对象的方法
// 需要加上溢出控制
func (pool *ObjPool) release(obj *Reusable) error {
	select {
	case pool.bufChan <- obj:
		fmt.Println(obj.name, "被放回")
		return nil
	default:
		return errors.New("溢出")
	}
}

// 创建对象池、获取、释放测试
func TestPool(t *testing.T) {
	pool := NewObjPool(14)
	var wg sync.WaitGroup
	for i := 0; i < 15; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			obj, err := getObj(pool)
			if err != nil {
				t.Error(err)
			}
			fmt.Println(i, "协程获得了对象", obj.name)
			time.Sleep(time.Millisecond * 1)
			pool.release(obj)
		}(i)
	}
	wg.Wait()
}
