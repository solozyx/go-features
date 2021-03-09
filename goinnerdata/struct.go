package main

import "fmt"

/*
0x63d578
0x63d578
true
true
0x63d578
0x63d578
true
true
*/

func main() {
	a := struct{}{}
	b := struct{}{}
	fmt.Printf("%p \r\n",&a)
	fmt.Printf("%p \r\n",&b)
	fmt.Println(a==b)
	fmt.Println(&a==&b)

	var c struct{}
	var d struct{}
	fmt.Printf("%p \r\n",&c)
	fmt.Printf("%p \r\n",&d)
	fmt.Println(c==d)
	fmt.Println(&c==&d)
}
