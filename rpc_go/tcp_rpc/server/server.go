package main 

import(
	"errors"
	"fmt"
	"net"
	"net/rpc"
	"os"
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

// http_rpc -> tcp_rpc 是降级 因为http基于tcp
// 对于client只是改了服务名称 server不听接收请求
// 阻塞式 1对1 RPC 不使用goroutine处理并发请求
// 这里实现 来1个请求处理完 然后顺序处理下1个请求
func main(){
	math := new(Math)
	rpc.Register(math)
	// 启动tcp监听
	tcpAddr,err := net.ResolveTCPAddr("tcp",":1234")
	if err != nil {
		fmt.Printf("tcp server rpc service start err = %v\n",err)
		os.Exit(2)
	}
	// 获取监听器对象
	listener,err := net.ListenTCP("tcp",tcpAddr)
	if err != nil {
		fmt.Printf("tcp server rpc service listener err = %v\n",err)
		os.Exit(2)
	}
	// 轮询 开启监听
	for{
		//
		conn,err := listener.Accept()
		if err != nil{
			fmt.Printf("conn err = %v\n",err)
			// 直接忽略本次连接
			continue
		}
		rpc.ServeConn(conn)
	}
}