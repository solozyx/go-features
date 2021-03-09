package main

/*
#include <stdio.h>
void foo(){
	printf("C func");
}
// import "C" 要和C语言代码相接
*/
import "C"

func main() {
	C.foo()
}
