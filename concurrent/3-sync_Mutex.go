// 需求：现在要计算 1-200 的各个数的阶乘，并且把各个数的阶乘放入到map中
// 最后显示出来。要求使用goroutine完成

// 思路
// 1. 编写一个函数，来计算各个数的阶乘，并放入到 map中.
// 2. 我们启动的协程多个，统计的将结果放入到 map中
// 3. map 应该做成一个全局的.

package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	myMap = make(map[int]int, 10)
	// 声明一个全局的互斥锁
	// lock 是一个全局的互斥锁，
	// sync 是包: synchornized 同步
	// Mutex : 是互斥
	lock sync.Mutex
)

/*
test1 函数就是计算 n!, 让将这个结果放入到 myMap
*/
func test1(n int) {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	//这里我们将 res 放入到myMap

	//加锁
	lock.Lock()
	// fatal error : concurrent map writes
	// 我的机器是 4个逻辑CPU，同一个时刻有4个协程的向 myMap 写入数据
	// 同时读数据可以，但是同时写数据不行
	myMap[n] = res

	//解锁
	lock.Unlock()
}

func main() {
	// 开启多个协程完成这个任务 [200个]
	for i := 1; i <= 20; i++ {
		go test1(i)
	}

	// 主线程休眠10秒钟 防止协程在任务执行完前就退出
	time.Sleep(time.Second * 10)

	// 输出结果,遍历这个结果
	lock.Lock()

	for i, v := range myMap {
		fmt.Printf("map[%d]=%d\n", i, v)
	}

	lock.Unlock()
}