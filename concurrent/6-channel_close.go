package main

import (
	"fmt"
	"time"
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

func sendData(ch1 chan int){
	for i := 1; i <= 5; i++ {
		ch1 <- i
	}
	fmt.Println("sendData 子协程发送方,写入数据完毕..")
	// 发送方:关闭通道,用于通知接受方没有数据继续发送了
	close(ch1)
}

/*
main主协程 接收方 读取到数据:  1 true
main主协程 接收方 读取到数据:  2 true
main主协程 接收方 读取到数据:  3 true
main主协程 接收方 读取到数据:  4 true
main主协程 接收方 读取到数据:  5 true
sendData 子协程发送方,写入数据完毕..
main主协程 接收方 读取完毕,发送方把通道关闭了.. ok =  false
main主协程 接收方 读取到数据:  0 false
main主协程 接收方 读取完毕,发送方把通道关闭了.. ok =  false
main主协程 接收方 读取到数据:  0 false
main主协程 接收方 读取完毕,发送方把通道关闭了.. ok =  false
main主协程 接收方 读取到数据:  0 false
main主协程 接收方 读取完毕,发送方把通道关闭了.. ok =  false
main主协程 接收方 读取到数据:  0 false

Process finished with exit code 2
*/
func closeChan1(){
	ch1 := make(chan int)
	go sendData(ch1)
	// main主协程 接收方 读取数据
	for{
		time.Sleep(1*time.Second)
		data ,ok:= <- ch1
		if !ok{
			fmt.Println("main主协程 接收方 读取完毕,发送方把通道关闭了.. ok = ",ok)
		}
		fmt.Println("main主协程 接收方 读取到数据: ",data,ok)
	}
}

/*
main主协程 接收方 读取到数据:  1 true
main主协程 接收方 读取到数据:  2 true
main主协程 接收方 读取到数据:  3 true
main主协程 接收方 读取到数据:  4 true
main主协程 接收方 读取到数据:  5 true
sendData 子协程发送方,写入数据完毕..
main主协程 接收方 读取完毕,发送方把通道关闭了.. ok =  false

Process finished with exit code 0
*/
func closeChan2(){
	ch1 := make(chan int)
	go sendData(ch1)
	// main主协程 接收方 读取数据
	for{
		time.Sleep(1*time.Second)
		data ,ok:= <- ch1
		if !ok{
			fmt.Println("main主协程 接收方 读取完毕,发送方把通道关闭了.. ok = ",ok)
			break
		}
		fmt.Println("main主协程 接收方 读取到数据: ",data,ok)
	}
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

func sendDataString(ch1 chan  string){
	for i := 1; i <= 10; i++ {
		ch1 <- fmt.Sprint("数据",i)
	}
	close(ch1)
}

func iterateChan1(){
	ch1 := make(chan string)
	go sendDataString(ch1)
	for value := range ch1 { // 停止条件：通道关闭，显示的调用close()
		fmt.Println("main主协程 从通道中读取数据: ",value)
	}
}

func main() {
	// closeChan()
	// closeChan1()
	// closeChan2()
	// iterateChan()
	iterateChan1()
}