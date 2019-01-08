package main 

import(
	"errors"
	"fmt"
	"net/http"
	// 标准库rpc包 Go自带rpc框架 
	"net/rpc"
)

type Args struct{
	// 使用gob格式传输数据 
	// 公开成员才能被识别到 在数据转换的过程中使用反射 
	// 私有成员是非导出的 反射不到 导致数据转换失败 
	// 
	// gob是Golang包自带的一个数据结构序列化的编码/解码工具
	// 编码使用Encoder,解码使用Decoder
	// 一种典型的应用场景就是RPC(remote procedure calls)
	A , B int 
}

// 返回值 RPC类型 
// bool int 都无所谓 只是要求有这么1个类型
// 给类型添加提供RPC服务的方法 
type Math bool 

// 返回值类型 必须是指针 指向int的指针 
func (m *Math)Multiply(args *Args,reply *int) error {
	*reply = args.A * args.B 
	return nil 
}

// 定义 商 类型
type Quotient struct {
	Quo int // 商 
	Rem int // 余数 
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
	rpc.HandleHTTP()

	if err := http.ListenAndServe(":1234",nil); err != nil{
		fmt.Println(err.Error())
	}
}