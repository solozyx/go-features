package main

import "fmt"

func main() {
	var b = 10
	var a = func() int {
		b = 20
		changeB(b)
		return 1
	}()
	// a= 1
	// b= 20
	fmt.Println("a=",a)
	fmt.Println("b=",b)
}

func changeB(x int)  {
	x = 30
}