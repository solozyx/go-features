package main
import (
	"fmt"
	"time"
)
func main() {
	intChannel := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intChannel <- i
	}
	stringChannel := make(chan string, 5)
	for i := 0; i < 5; i++ {
		stringChannel <- "hello" + fmt.Sprintf("%d", i)
	}

	//传统的方法在遍历管道时，如果不关闭会阻塞而导致 deadlock
	//问题，在实际开发中，可能我们不好确定什么关闭该管道.
	//可以使用select 方式可以解决

	// label:
	for {
		select {
			// 注意:
			// 这里,如果intChannel一直没有关闭,不会一直阻塞而deadlock
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

				// 关键字 break到 select，跳不出for循环
				// break

				// 使用标签跳出for循环
				// break label
		}
	}
}