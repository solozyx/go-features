package main
import (
	"fmt"
	"time"
)
/*
write Data
*/
func writeData(intChan chan int) {
	for i := 1; i <= 10; i++ {
		//放入数据
		intChan<- i
		fmt.Println("writeData ", i)
		time.Sleep(time.Second)
	}
	// 关闭channel 方便其readData协程读取 intChan不引起deadlock
	close(intChan)
}
/*
read data
*/
func readData(intChan chan int, exitChan chan bool) {
	for {
		v, ok := <-intChan
		if !ok {
			// 读取到push进channel的真实元素 ok为true
			// 未读取到push进channel的真实元素 即 pop掉channel的所有元素后 ok为false
			break
		}
		time.Sleep(time.Second)
		fmt.Printf("readData 读到数据=%v\n", v) 
	}
	// readData 读取完数据后，即任务完成
	exitChan <- true
	// 关闭channel 方便其main主协程读取 exitChan不引起deadlock
	close(exitChan)
}

func main() {
	intChan := make(chan int, 10)
	exitChan := make(chan bool, 1)
	go writeData(intChan)
	go readData(intChan, exitChan)

	// 终止程序控制
	// time.Sleep(time.Second * 50)
	// OR
	for {
		_, ok := <-exitChan
		if !ok {
			// exitChan 在 readData 协程关闭了 所以可以一直轮询遍历 不会deadlock
			fmt.Println("main 进来 了... ok = ",ok)
			break
		} else {
			// 只有在 readData协程执行了 exitChan <- true 操作
			// 这里的ok才为true
			// 该 else 分支只能执行 1次
			fmt.Println("main 进来 了...ok = ",ok)
		}
	}
}

/*
writeData  1
writeData  2
readData 读到数据=1
readData 读到数据=2
writeData  3
readData 读到数据=3
writeData  4
writeData  5
readData 读到数据=4
readData 读到数据=5
writeData  6
readData 读到数据=6
writeData  7
writeData  8
readData 读到数据=7
readData 读到数据=8
writeData  9
writeData  10
readData 读到数据=9
readData 读到数据=10
main 进来 了...ok =  true
main 进来 了... ok =  false

Process finished with exit code 0
*/
