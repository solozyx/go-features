package main
import (
	"fmt"
	"time"
)

type Cat struct {
	Name string
	Age int
}

/*
非缓冲通道
*/
func noBufferChannel(){
	ch1 :=  make(chan int)
	// 非缓冲通道:  0 0
	fmt.Println("非缓冲通道: ", len(ch1), cap(ch1))

	go func() {
		time.Sleep(3*time.Second)
		<- ch1 // 即时阻塞
	}()

	ch1 <- 100 // 即时阻塞
	fmt.Println("noBufferChannel 写完了..")
}

/*
缓冲通道 缓冲区满了才会阻塞
*/
func bufferChannel(){
	// 缓冲大小 5
	ch2 := make(chan int, 5)
	// 缓冲通道:  0 5
	fmt.Println("缓冲通道: ", len(ch2), cap(ch2))

	go func() {
		data1 := <- ch2
		fmt.Println("子协程获取数据: ", data1)
	}()

	ch2 <- 1
	// 1 5
	fmt.Println(len(ch2), cap(ch2))

	ch2 <- 2
	ch2 <- 3
	// 3 5
	fmt.Println(len(ch2), cap(ch2))

	ch2 <- 4
	ch2 <- 5
	// 5 5
	fmt.Println(len(ch2), cap(ch2))

	ch2 <- 6 // 非即时阻塞
	fmt.Println("main协程 over..")
}

func sendData2(ch3 chan string){
	for i := 1; i <= 100; i++ {
		ch3 <- fmt.Sprint("数据：",i)
		fmt.Println("已经写出数据：",i)
	}
	close(ch3)
}

/*
chan string
已经写出数据： 1
已经写出数据： 2
已经写出数据： 3
已经写出数据： 4
已经写出数据： 5
	读取数据:  数据：1
已经写出数据： 6
	读取数据:  数据：2
...
	读取数据:  数据：100
	读取数据:
读取完毕..

Process finished with exit code 0
*/
func bufferChanne2(){
	ch3 := make(chan string,5)
	// chan string
	fmt.Printf("%T\n",ch3)

	go sendData2(ch3)

	for {
		time.Sleep(100 * time.Millisecond)
		data,ok := <- ch3
		fmt.Println("\t读取数据: ",data)
		if !ok{
			fmt.Println("读取完毕..")
			break
		}
	}
}

/*
newCat=main.Cat , newCat={小花猫 4}
newCat.Name=小花猫

Process finished with exit code 0
*/
func interfaceChan(){
	// 定义一个存放任意数据类型的管道 3个数据
	// var allChan chan interface{}
	allChan := make(chan interface{}, 3)

	allChan<- 10
	allChan<- "tom jack"
	cat := Cat{"小花猫", 4}
	allChan<- cat

	// 获得管道中的第三个元素，则先将前2个推出
	<-allChan
	<-allChan
	newCat := <-allChan
	// 在运行层面能正确打印出数据
	// newCat=main.Cat , newCat={小花猫 4}
	fmt.Printf("newCat=%T , newCat=%v\n", newCat, newCat)

	// 下面的写法是错误的!编译不通过
	// 在编译的层面 认为 newCat 是空接口类型 接口类型里面是没有字段
	// 编译层面 =/= 运行层面
	// fmt.Printf("newCat.Name=%v", newCat.Name)

	//使用类型断言
	a := newCat.(Cat)
	fmt.Printf("newCat.Name=%v", a.Name)
}

func main() {
	// noBufferChannel()
	// bufferChannel()
	bufferChanne2()
}