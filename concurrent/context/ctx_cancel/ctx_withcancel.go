package main

import (
	"context"
	"fmt"
	"time"
)

// 起子goroutine干活 用context.WithCancel控制它结束
// 干活的goroutine要使用 <-ctx.Done 外部有人取消它

func main()  {
	test()
	time.Sleep(time.Hour)
}

func test() {
	ctx, cancel := context.WithCancel(context.Background())
	// 调用cancel取消函数 会触发 case <-ctx.Done():
	defer cancel()

	intChan := gen(ctx)
	for n := range intChan {
		fmt.Println(n)
		if n == 5 {
			// n==5 是遍历管道结束 但是调用 gen 启动的5个goroutine没有结束
			break
		}
	}
}

func gen(ctx context.Context) <-chan int {
	dst := make(chan int)
	n := 1
	go func() {
		for {
			select {
			case <-ctx.Done():
				// 触发了 <-ctx.Done() 对应的goroutine 返回结束
				// ctx 控制子goroutine的生命周期
				// 后台goroutine如何退出 可以用context控制
				fmt.Println("gen func goroutien exited")
				return // returning not to leak the goroutine
			case dst <- n:
				n++
			}
		}
	}()
	return dst
}