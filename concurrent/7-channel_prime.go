package main
import (
	"fmt"
	"time"
)

/*
向 intChan放入 1-80000个数
管道是引用类型,和main函数的是同一个
*/
func putNum(intChan chan int) {
	for i := 1; i <= 80000; i++ {    
		intChan <- i
	}
	close(intChan)
}

/*
从 intChan取出数据，并判断是否为素数,如果是，就放入到primeChan
*/
func primeNum(intChan chan int, primeChan chan int, exitChan chan bool) {
	var flag bool
	for {
		// time.Sleep(time.Millisecond * 10) //10毫秒
		num, ok := <- intChan
		if !ok {
			//intChan 取不到..
			break
		}
		// 假设是素数
		flag = true
		// 判断num是不是素数
		for i := 2; i < num; i++ {
			if num % i == 0 {
				// 说明该num不是素数
				flag = false
				break
			}
		}
		if flag {
			// 将这个数就放入到primeChan
			primeChan <- num
		}
	}

	fmt.Println("有一个 primeNum 协程因为取不到数据，退出")
	// 这里我们还不能关闭 primeChan
	// 向 exitChan 写入true
	exitChan <- true
}

func main() {
	intChan := make(chan int , 1000)
	// 放入素数结果
	primeChan := make(chan int, 20000)
	// 标识退出的管道 8个true 表示求出了所有素数 退出程序
	exitChan := make(chan bool, 8)

	start := time.Now().Unix()
	
	// 开启一个协程，向intChan放入1-80000个数
	go putNum(intChan)
	// 开启8个协程，从 intChan取出数据，并判断是否为素数,如果是，就放入到primeChan
	for i := 0; i < 8; i++ {
		go primeNum(intChan, primeChan, exitChan)
	}

	//这里我们主线程，进行处理
	go func(){
		for i := 0; i < 8; i++ {
			<- exitChan
		}
		end := time.Now().Unix()
		fmt.Println("使用协程耗时=", end - start)
		// 当我们从exitChan 取出了8个结果，就可以放心的关闭 prprimeChan
		close(primeChan)
	}()

	//遍历 primeChan 把结果取出
	for {
		res, ok := <-primeChan
		if !ok{
			break
		}
		// 将结果输出
		fmt.Printf("素数=%d\n", res)
	}
	fmt.Println("main线程退出")
}