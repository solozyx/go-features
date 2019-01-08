package main 

import(
	"fmt"
	"log"
	"os"
	"net/rpc/jsonrpc"
)

type Args struct{
	A , B int 
}

type Quotient struct {
	Quo int 
	Rem int 
}


func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage : ",os.Args[0], "server")
		os.Exit(1)
	}
	serverAddr := os.Args[1]

	// rpc.Dial() -> jsonrpc.Dial()
	client,err := jsonrpc.Dial("tcp",serverAddr+":1234")
	if err != nil{	
		log.Fatal("dialing err : ",err)
	}

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