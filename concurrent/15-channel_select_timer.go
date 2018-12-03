package main

import (
	"time"
	"fmt"
)

func selectTimer(){
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		// time.Sleep(2 * time.Second)
		// time.Sleep(4 * time.Second)
		data := <-ch1
		fmt.Println("子goroutine读取ch1数据: ", data)
	}()

	select {
	case ch1 <- 100: // 一直阻塞
		fmt.Println("ch1中写数据..")
		time.Sleep(1 * time.Second)

	case ch2 <- 200: // 一直阻塞
		fmt.Println("ch2中写数据..")

	// time.After() 返回装时间 time.Time 结构体的通道
	case <- time.After(3 * time.Second): // 阻塞3秒
		fmt.Println("定时器到时间了,向通道投递当前时间..")
		// go func(){}() 子协程睡眠 4秒的时候 deadlock 说明  case ch1 <- 100: 没执行 ch1 <- 100
		// ch1 := make(chan int) 创建的通道cap是0 没其他人读 根本写不进去
		// fmt.Println(<-ch1) // fatal error: all goroutines are asleep - deadlock!

	// default:
	//	fmt.Println("default..")
	}
}

func main() {
	selectTimer()
}