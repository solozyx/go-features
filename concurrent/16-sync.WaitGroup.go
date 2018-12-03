package main

import (
	"sync"
	"math/rand"
	"time"
	"fmt"
)

func printNum1(wg *sync.WaitGroup) {
	for i := 1; i <= 100; i++ {
		fmt.Println("子g1 i = ", i)
		time.Sleep(time.Duration(rand.Intn(1000)))
	}
	// WaitGroup同步等待组内置计数器 减1
	wg.Done()
}

func printNum2(wg *sync.WaitGroup) {
	for j := 1; j <= 100; j++ {
		fmt.Println("\t子g2 j = ", j)
		time.Sleep(time.Duration(rand.Intn(1000)))
	}
	// WaitGroup同步等待组内置计数器 减1
	wg.Done()
}

func wg(){
	var wg sync.WaitGroup
	fmt.Printf("sync.WaitGroup类型: %T\n",wg)
	fmt.Println(wg)

	// 设置同步等待组WaitGroup内置的计数器值
	// 要执行2个goroutine 就设置 2
	wg.Add(2)

	// error
	// 可能有1个子协程执行不完 程序退出 不崩溃
	// 如果2个子协程都执行完 WaitGroup的内置counter值最终是 -1 panic 恐慌错误
	// wg.Add(1)

	// fatal error: all goroutines are asleep - deadlock!
	// wg.Add(3)

	// 设置种子
	rand.Seed(time.Now().UnixNano())

	go printNum1(&wg)
	go printNum2(&wg)

	// main协程执行Wait(), main协程进入阻塞状态
	// WaitGroup底层内置计数器为0时,main协程才能解除阻塞
	wg.Wait()

	// time.Sleep(1 * time.Second)
	fmt.Println("main协程解除阻塞结束了...")
}

func main()  {
	wg()
}