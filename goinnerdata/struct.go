package main

import "fmt"

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
