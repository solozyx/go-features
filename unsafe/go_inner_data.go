package main

import (
	"fmt"
	"unsafe"
)

/*
切片 size= 24
string size= 16
struct{} size= 0
interface{} size= 16
map size= 8
channel size= 8
*/

func main() {
	var s1 []int
	size1 := unsafe.Sizeof(s1)
	fmt.Println("[]int size=",size1)

	var s2 []string
	size2 := unsafe.Sizeof(s2)
	fmt.Println("[]string size=",size2)

	var s3 []byte
	size3 := unsafe.Sizeof(s3)
	fmt.Println("[]byte size=",size3)

	var s4 string
	size4 := unsafe.Sizeof(s4)
	fmt.Println("string size=",size4)

	var s5 struct{}
	size5 := unsafe.Sizeof(s5)
	fmt.Println("struct{} size=",size5)

	var s6 interface{}
	size6 := unsafe.Sizeof(s6)
	fmt.Println("interface{} size=",size6)

	var s7 map[int]int
	size7 := unsafe.Sizeof(s7)
	fmt.Println("map[int]int{} size=",size7)

	var s8 map[string]struct{}
	size8 := unsafe.Sizeof(s8)
	fmt.Println("map[string]struct{} size=",size8)

	s9 := make(map[string]struct{},10)
	size9 := unsafe.Sizeof(s9)
	fmt.Println("make(map[string]struct{},10) size=",size9)

	s10 := make(chan int)
	size10 := unsafe.Sizeof(s10)
	fmt.Println("make(chan int) size=",size10)

	s11 := make(chan int,10)
	size11 := unsafe.Sizeof(s11)
	fmt.Println("make(chan int,10) size=",size11)

	s12 := make(chan interface{},10)
	size12 := unsafe.Sizeof(s12)
	fmt.Println("make(chan interface{},10) size=",size12)
}
