package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main(){
	TestSyncPool()
	fmt.Println("--------------------")
	TestSyncPoolInMultiGoroutine()
}

func TestSyncPool() {
	pool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("Create a new object")
			return 100
		},
	}

	v := pool.Get().(int)
	fmt.Println(v)
	pool.Put(3)
	// GC 会清除 sync.Pool 中缓存的对象
	// 在 go 1.13 版本 sync.Pool 支持跨 GC
	runtime.GC()
	v1, _ := pool.Get().(int)
	fmt.Println(v1)
	v2, _ := pool.Get().(int)
	fmt.Println(v2)
}

func TestSyncPoolInMultiGoroutine() {
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
		wg.Wait()
	}
}