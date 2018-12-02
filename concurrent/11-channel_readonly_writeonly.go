package main

import (
	"fmt"
	"time"
)

func doubleDirectionChan(ch1 chan string){
	ch1 <- "我是小明"

	time.Sleep(2*time.Second)
	data := <-ch1
	fmt.Println("子协程 接收到main回应: ",data)
}

func doubleTest(){
	// 双向通道
	ch1 := make(chan string)
	go doubleDirectionChan(ch1)

	time.Sleep(3 * time.Second)
	data := <- ch1
	fmt.Println("main主协程 接受到子协程打招呼: ",data)

	ch1 <- "你要上学么?"
}

/*
只有写入数据
*/
func writeDataOnly(ch1 chan <- string){
	// 只能写入
	// invalid operation: <-ch1 (receive from send-only type chan<- string)
	// <- ch1
}

/*
只有读取数据
*/
func readDataOnly(ch1 <- chan string){
	<- ch1
	// invalid operation: ch1 <- "hello" (send to receive-only type <-chan string)
	// ch1 <- "hello"
}

func singleTest(){
	// 双向通道
	ch1 := make(chan string)
	writeDataOnly(ch1)
	readDataOnly(ch1)

}

func writeOnlyReadOnlyChan(){
	// 管道可以声明为只读或者只写
	// 在默认情况下，管道是双向 可读可写
	// var chan1 chan int

	// 2.声明为只写
	var chan2 chan<- int
	chan2 = make(chan int, 3)
	chan2<- 20
	// num := <-chan2 //error
	fmt.Println("chan2=", chan2)

	// 3.声明为只读
	var chan3 <-chan int
	num2 := <-chan3
	// chan3<- 30 //err
	fmt.Println("num2", num2)
}

func main() {
	// writeOnlyReadOnlyChan()
	// doubleTest()
	singleTest()
}