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
	// intChan<- 300
	fmt.Println("okook~")
	//当管道关闭后，读取数据是可以的
	n1 := <-intChan
	fmt.Println("n1=", n1)
}

func iterateChan() {
	//遍历管道
	intChan2 := make(chan int, 100)
	for i := 0; i < 100; i++ {
		//放入100个数据到管道
		intChan2 <- i * 2
	}

	//遍历管道不能使用普通的 for 循环
	// for i := 0; i < len(intChan2); i++ {

	// }

	//在遍历时，如果channel没有关闭，则会出现deadlock的错误
	//在遍历时，如果channel已经关闭，则会正常遍历数据，遍历完后，就会退出遍历
	close(intChan2)
	for v := range intChan2 {
		fmt.Println("v=", v)
	}
}

func main() {
	closeChan()
	iterateChan()
}