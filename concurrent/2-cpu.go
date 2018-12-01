package main

import (
	"runtime"
	"fmt"
)

/*
在主goroutine中 没有init()函数则不执行 有 则执行顺序
init() -> main()
*/
func init(){
	// 获取机器逻辑CPU核心数量
	cpuNum := runtime.NumCPU()
	fmt.Println("cpuNum=", cpuNum)
	// 可以自己设置使用多个cpu
	// go程序的执行在 go1.8之前默认使用1个核心
	// 在go1.8之后默认使用多核心 不设置也可以
	//
	// 理论上取值范围 [1-256]
	//
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {

	fmt.Println("ok")
}