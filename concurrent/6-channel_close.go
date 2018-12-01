package main

import (
	"fmt"
)

func closeChan() {
	intChan := make(chan int, 3)
	intChan <- 100
	intChan <- 200
	// close
	close(intChan)
	// channel关闭后 不能够再写入数到channel
	// panic:send on closed channel
	// intChan<- 300
	fmt.Println("okook~")

	//当管道关闭后，读取数据是可以的
	n1 := <-intChan
	fmt.Println("n1=", n1)
}

func iterateChan() {
	intChan2 := make(chan int, 100)
	for i := 0; i < 100; i++ {
		intChan2 <- i * 2
	}

	// 遍历管道不能使用普通的 for 循环
	// for i := 0; i < len(intChan2); i++ {
		// 这样遍历只能取出 50个元素 因为每 pop 1个元素 len会减去1
		// 容量遍历也不行，容量也不能代表管道有多少个数据
	// }

	// 在遍历时，如果channel没有关闭，则会出现deadlock的错误
	// 在遍历时，如果channel已经关闭，则会正常遍历数据，遍历完后，就会退出遍历
	close(intChan2)
	for v := range intChan2 {
		fmt.Println("v=", v)
	}
}

func main() {
	closeChan()
	iterateChan()
}