package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var m map[int]int
	var p uintptr
	fmt.Println(unsafe.Sizeof(m), unsafe.Sizeof(p)) // 8 8 (linux/amd64)


}
