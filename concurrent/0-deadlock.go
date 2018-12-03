package main

import (
	"time"
	"fmt"
)

func main(){
	a := make(chan int)
	fmt.Println("a 指针存储的地址 = ",a)
	go func(intChan chan int) {
		fmt.Println("intChan 指针存储的地址 = ",intChan)
		intChan <- 1
		fmt.Println("子协程阻塞 执行不到这里...")
		intChan <- 2
	}(a)
	time.Sleep(5 * time.Second)
	fmt.Println("main协程 退出时间 ",time.Now().Unix())
}