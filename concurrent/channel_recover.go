package main
import (
	"fmt"
	"time"
)

func sayHello() {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		fmt.Println("hello,world")
	}
}

func testRecover() {
	// defer + recover
	defer func() {
		//捕获 testRecover 抛出的panic
		if err := recover(); err != nil {
			fmt.Println("testRecover() 发生错误", err)
		}
	}()
	//定义了一个map
	var myMap map[int]string
	myMap[0] = "golang" //error
}

func main() {
	go sayHello()
	go testRecover()
	for i := 0; i < 10; i++ {
		fmt.Println("main() ok=", i)
		time.Sleep(time.Second)
	}
}