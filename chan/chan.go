package main

import "fmt"

func main() {
	var closeChan chan byte
	closeChan = make(chan byte,1)
	// 关闭channel后读取channel打印出 0
	close(closeChan)
	// 不关闭channel读取死锁
	// fatal error: all goroutines are asleep - deadlock!
	// goroutine 1 [chan receive]:
	fmt.Println(<-closeChan)
}
