package main

import "fmt"

/*
intChan 的值=0xc04207a000 intChan本身的地址=0xc042068018
channel len= 3 cap=3
num2= 211
channel len= 2 cap=3
num3= 50 num4= 98

Process finished with exit code 0
*/

func main() {
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