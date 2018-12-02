package main

import (
	"fmt"
	"time"
)

/*
intChan 的值=0xc04207a000 intChan本身的地址=0xc042068018
channel len= 3 cap=3
num2= 211
channel len= 2 cap=3
num3= 50 num4= 98

Process finished with exit code 0
*/
func chan1() {
	//1. 创建一个可以存放3个int类型的管道
	var intChan chan int
	intChan = make(chan int, 3)
	// 2. channel是引用类型
	fmt.Printf("intChan 的值=%v intChan本身的地址=%p\n", intChan, &intChan)

	//3. 向管道写入数据
	intChan<- 10
	num := 211
	intChan<- num
	intChan<- 50
	//如果从channel取出数据后，可以继续放入
	<-intChan
	//注意, 当我们给管道写入数据时，不能超过其容量
	intChan<- 98
	//4. 管道的长度和cap(容量) 3, 3
	fmt.Printf("channel len= %v cap=%v \n", len(intChan), cap(intChan))

	//5. 从管道中读取数据
	var num2 int
	num2 = <-intChan
	fmt.Println("num2=", num2)
	// 2, 3
	fmt.Printf("channel len= %v cap=%v \n", len(intChan), cap(intChan))

	//6. 在没有使用协程的情况下，
	// 如果我们的管道数据已经全部取出，再取就会报告 deadlock
	num3 := <-intChan
	num4 := <-intChan

	// deadlock
	// num5 := <-intChan
	fmt.Println("num3=", num3, "num4=", num4/*, "num5=", num5*/)
}

func chanDeadlock1() {
	var intChan chan int
	fmt.Println(intChan) // <nil>
	fmt.Printf("%T\n",intChan) // chan int

	intChan = make(chan int)
	fmt.Println(intChan) // 0xc04203c060

	// fatal error: all goroutines are asleep - deadlock!
	// 阻塞式 执行这行代码后面的代码要等待，main goroutine向通道中写入数据
	intChan <- 100
	// 不能解除阻塞 阻塞发生在 intChan <- 100 根本走不了下面的代码
	<- intChan
	// 1个goroutine阻塞 不能自救 要其他goroutine来解除阻塞
	// 所以要把 <- intChan 放到另外的goroutine执行
}

func chanDeadlock2() {
	var intChan chan int
	intChan = make(chan int)
	// intChan <- 100 导致主协程阻塞
	intChan <- 100

	// 主协程阻塞了 需要另一个协程帮主协程解除阻塞
	// fatal error: all goroutines are asleep - deadlock!
	// 因为该子协程 启动晚了 上面 intChan <- 100 已经导致主协程 阻塞了 没启动
	go func() {
		<- intChan
	}()

	fmt.Println("main主协程over ...")
}

/*
子 goroutine 从通道中读取到main主协程传来的数据是:  100
main主协程over ...

Process finished with exit code 0
*/
func chanResolveDeadlock1() {
	var intChan chan int
	intChan = make(chan int)

	// 主协程阻塞了 需要另一个协程帮主协程解除阻塞
	go func() {
		// 阻塞式，从通道中读取数据
		// 子协程 data := <- intChan 也阻塞
		// 该子协程的阻塞由 主协程intChan <- 100 写channel解除
		data := <- intChan
		fmt.Println("子 goroutine 从通道中读取到main主协程传来的数据是: ",data)
	}()
	// 在 go func(){}() 中 intChan可以直接使用 goroutine执行1个匿名函数
	// 可以操作它上面定义的局部变量

	// intChan <- 100 导致主协程阻塞
	// 该主协程的阻塞 由 子协程 data := <- intChan 读channel解除
	intChan <- 100

	// 所以 主协程 子协程 批次解除对方的阻塞

	fmt.Println("main主协程over ...")
}

/*
main主协程over ...

Process finished with exit code 0
*/
func chanResolveDeadlock2() {
	var intChan chan int
	intChan = make(chan int)

	go func() {
		data := <- intChan
		time.Sleep(5*time.Second)
		fmt.Println("子 goroutine 从通道中读取到main主协程传来的数据是: ",data)
	}()
	intChan <- 100
	fmt.Println("main主协程over ...")
}

/*
子 goroutine 从通道中读取到main主协程传来的数据是:  100
main主协程over ...

Process finished with exit code 0
*/
func chanSubDoneMainExit() {
	var intChan chan int
	intChan = make(chan int)

	doneChan := make(chan bool)

	go func() {
		data := <- intChan
		time.Sleep(5*time.Second)
		fmt.Println("子 goroutine 从通道中读取到main主协程传来的数据是: ",data)
		doneChan <- true
	}()
	intChan <- 100
	<- doneChan
	fmt.Println("main主协程over ...")
}

/*
子goroutine执行 ..
main主协程 读取到数据:  hello
main主协程 执行时间:  5

Process finished with exit code 0
*/
func sleepChan() {
	ch1 := make(chan string)
	start := time.Now().Unix()
	go func() {
		fmt.Println("子goroutine执行 .. ")
		time.Sleep(5*time.Second)
		ch1 <- "hello"
	}()

	time.Sleep(2*time.Second)

	data := <- ch1
	end := time.Now().Unix()
	fmt.Println("main主协程 读取到数据: ",data)
	fmt.Println("main主协程 执行时间: ",end-start)
}

/*
main 主协程创建了通道ch1:  0xc04203c060
printLetter 子协程接收到通道ch1:  0xc04203c060
	1,A
printNum 子协程接收到通道ch1:  0xc04203c060
1
2
	2,B
main 主协程执行完毕..

Process finished with exit code 0
*/
func printTest(){
	ch1 :=make(chan bool)
	fmt.Println("main 主协程创建了通道ch1: ",ch1)
	go printNum(ch1) // 引用传递
	go printLetter(ch1)
	<- ch1
	<- ch1
	fmt.Println("main 主协程执行完毕..")
}

func printNum(ch1 chan bool){
	fmt.Println("printNum 子协程接收到通道ch1: ",ch1)
	for i := 1; i <= 2; i++ {
		fmt.Println(i)
		time.Sleep(10*time.Millisecond)
	}
	ch1<-true
}

func printLetter(ch1 chan bool){
	fmt.Println("printLetter 子协程接收到通道ch1: ",ch1)
	for i := 1; i<=2; i++ {
		fmt.Printf("\t%d,%c\n",i,64+i)
		time.Sleep(10*time.Millisecond)
	}
	ch1 <- true
}

func main() {
	// chan1()
	// chanDeadlock1()
	// chanDeadlock2()
	// chanResolveDeadlock1()
	// chanResolveDeadlock2()
	// chanSubDoneMainExit()
	// sleepChan()
	printTest()
}