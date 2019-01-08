package main 

import(
	"errors"
	"fmt"
	"net"
	"os"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type Args struct{
	A , B int
}

type Math bool 

func (m *Math)Multiply(args *Args,reply *int) error {
	*reply = args.A * args.B 
	return nil 
}

type Quotient struct {
	Quo int
	Rem int
}

func (m *Math)Devide(args *Args,quotient *Quotient) error {
	if args.B == 0 {
		return errors.New("devide by zero error")
	}
	quotient.Quo = args.A / args.B 
	quotient.Rem = args.A % args.B 
	return nil 
}

func main(){
	math := new(Math)
	rpc.Register(math)
	tcpAddr,err := net.ResolveTCPAddr("tcp",":1234")
	if err != nil {
		fmt.Printf("tcp server rpc service start err = %v\n",err)
		os.Exit(2)
	}
	listener,err := net.ListenTCP("tcp",tcpAddr)
	if err != nil {
		fmt.Printf("tcp server rpc service listener err = %v\n",err)
		os.Exit(2)
	}
	for{
		conn,err := listener.Accept()
		if err != nil{
			fmt.Printf("conn err = %v\n",err)
			continue
		}
		// rpc.ServeConn(conn) -> jsonrpc.ServeConn(conn)
		jsonrpc.ServeConn(conn)
	}
}