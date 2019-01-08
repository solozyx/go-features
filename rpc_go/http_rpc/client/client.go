package main 

import(
	"fmt"
	"log"
	"net/rpc"
	"os"
)

// 客户端RPC调用服务端服务 客户端和服务端注册的RPC类型要一致 
type Args struct{
	A , B int 
}

type Quotient struct {
	Quo int 
	Rem int 
}

// path\client>go run client.go 127.0.0.1
// Math.Multiply 17 * 8 = 136
// Math.Quotient 17 / 8 quo = 2 rem = 1
func main() {
	// 调用服务端RPC服务
	// 命令行参数取到远程服务地址 
	if len(os.Args) != 2 {
		fmt.Println("Usage : ",os.Args[0], "server")
		os.Exit(1)
	}
	// 命令行第0个参数是程序本身 
	serverAddr := os.Args[1]
	client,err := rpc.DialHTTP("tcp",serverAddr+":1234")
	if err != nil{	
		log.Fatal("dialing err : ",err)
	}
	
	// 实例化参数类型 
	args := &Args{17,8}
	
	var reply int 
	if err = client.Call("Math.Multiply",args,&reply); err != nil {
		log.Fatal("Math.Multiply err : ",err)
	}
	fmt.Printf("Math.Multiply %d * %d = %d \n",args.A,args.B,reply)

	var quotient Quotient
	if err = client.Call("Math.Devide",args,&quotient); err != nil {
		log.Fatal("Math.Quotient err : ",err)
	}
	fmt.Printf("Math.Quotient %d / %d quo = %d rem = %d \n",
		args.A,args.B,quotient.Quo,quotient.Rem)
}