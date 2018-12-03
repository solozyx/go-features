package main

import (
	"time"
	"fmt"
)

/*
*time.Timer
2018-12-03 11:45:39.7049401 +0800 CST m=+0.004985101
<-chan time.Time
2018-12-03 11:45:42.7058003 +0800 CST m=+3.005845301

Process finished with exit code 0
*/
func newTimer(){
	// 创建计时器
	// 3秒之后时间到 就要等待这3秒
	timer1 := time.NewTimer(3 * time.Second)

	// *time.Timer
	fmt.Printf("%T\n",timer1)

	// 2018-12-03 11:45:39.7049401 +0800 CST m=+0.004985101
	fmt.Println(time.Now())

	// 定时器timer内部包含一个通道是 双向通道
	// 定时器自己向内部的双向通道写入 time.Time 结构体 写入1个时间 Duration之后的时间
	// 比如现在是 11:40:00 上面时间间隔3秒 将当前时间+设置的时间间隔 写入内部双向通道
	// 但是定时器给外部提供调用的是 单向通道
	time1 := <- timer1.C

	// <-chan time.Time
	fmt.Printf("%T\n",timer1.C)

	// 2018-12-03 11:45:42.7058003 +0800 CST m=+3.005845301
	fmt.Println(time1)
}

/*
2018-12-03 12:51:32.3719693 +0800 CST m=+0.003990601
2018-12-03 12:51:37.3803201 +0800 CST m=+5.012341301

Process finished with exit code 0
*/
func after(){
	// 使用After(),返回值 <- chan Time, 同T imer.C 一样
	ch1 := time.After(5 * time.Second)

	// 2018-12-03 12:51:32.3719693 +0800 CST m=+0.003990601
	fmt.Println(time.Now())

	time2 := <- ch1

	// 2018-12-03 12:51:37.3803201 +0800 CST m=+5.012341301
	fmt.Println(time2)
}

func main(){
	// newTimer()
	after()
}