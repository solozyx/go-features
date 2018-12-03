package main

import (
	"fmt"
	"time"
)

func select0(){
	intChannel := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intChannel <- i
	}

	stringChannel := make(chan string, 5)
	for i := 0; i < 5; i++ {
		stringChannel <- "hello" + fmt.Sprintf("%d", i)
	}

	// 传统的方法在遍历管道时，如果不关闭会阻塞而导致 deadlock
	// 在实际开发中，可能我们不好确定什么关闭该管道
	// 可以使用select 方式可以解决

	// label:
	for {
		select {
		// 注意 这里,如果intChannel一直没有关闭,不会一直阻塞而deadlock
		// 会自动到下一个case匹配
		case v := <- intChannel :
			fmt.Printf("从intChannel读取的数据%d\n", v)
			time.Sleep(time.Second)
		case v := <- stringChannel :
			fmt.Printf("从stringChannel读取的数据%s\n", v)
			time.Sleep(time.Second)
		default :
			fmt.Printf("intChannel和stringChannel都取不到数据 \n")
			time.Sleep(time.Second)
			return

			// 关键字 break 到 select，跳不出for循环
			// break

			// 使用标签跳出for循环
			// break label
		}
	}
}

/*
fatal error: all goroutines are asleep - deadlock!

ch1 和 ch2 没有写入数据操作 在select中等待 <-ch1 和 <-ch2
都不能读出数据
没有default 则deadlock
*/
func select1(){
	ch1 := make(chan int)
	ch2 := make(chan int)
	select {
	case data := <-ch1:
		fmt.Println("ch1中读取数据了: ", data)
	case data := <-ch2:
		fmt.Println("ch2中读取数据了: ", data)
	}
}

/*
第一次执行 ch2中读取数据了:  200
第二次执行 ch1中读取数据了:  100
第三次执行 ch1中读取数据了:  100
...
*/
func select2(){
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		time.Sleep(5 * time.Second)
		ch1 <- 100
	}()
	go func() {
		time.Sleep(5 * time.Second)
		ch2 <- 200
	}()
	select {
	case data := <-ch1:
		fmt.Println("ch1中读取数据了: ", data)
		fmt.Println("再从 ch2中读取数据: ", <- ch2)
	case data := <-ch2:
		fmt.Println("ch2中读取数据了: ", data)
		fmt.Println("再从 ch1中读取数据: ", <- ch1)
	}
}

/*
执行了default ...

Process finished with exit code 0
*/
func select3(){
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		time.Sleep(5 * time.Second)
		ch1 <- 100
	}()

	go func() {
		time.Sleep(5 * time.Second)
		ch2 <- 200
	}()

	select {
	case data := <-ch1:
		fmt.Println("ch1中读取数据了: ", data)
	case data := <-ch2:
		fmt.Println("ch2中读取数据了: ", data)
	default:
		fmt.Println("执行了default ... ")
	}
}


func main() {
	// select0()
	// select1()
	// select2()
	select3()
}