package main
import (
	"fmt"
)

type Cat struct {
	Name string
	Age int
}

/*
newCat=main.Cat , newCat={小花猫 4}
newCat.Name=小花猫

Process finished with exit code 0
*/

func main() {
	// 定义一个存放任意数据类型的管道 3个数据
	// var allChan chan interface{}
	allChan := make(chan interface{}, 3)

	allChan<- 10
	allChan<- "tom jack"
	cat := Cat{"小花猫", 4}
	allChan<- cat

	// 获得管道中的第三个元素，则先将前2个推出
	<-allChan
	<-allChan
	newCat := <-allChan
	// 在运行层面能正确打印出数据
	// newCat=main.Cat , newCat={小花猫 4}
	fmt.Printf("newCat=%T , newCat=%v\n", newCat, newCat)

	// 下面的写法是错误的!编译不通过
	// 在编译的层面 认为 newCat 是空接口类型 接口类型里面是没有字段
	// 编译层面 =/= 运行层面
	// fmt.Printf("newCat.Name=%v", newCat.Name)

	//使用类型断言
	a := newCat.(Cat) 
	fmt.Printf("newCat.Name=%v", a.Name)
}