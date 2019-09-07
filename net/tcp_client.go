package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

// TCP 客户端
// 1.建立与服务端的链接
// 2.进行数据收发
// 3.关闭链接

func main() {
	// 1. tcp client 和 tcp server 拨号,和 tcp server 建立连接
	conn, err := net.Dial("tcp", "127.0.0.1:20000")
	// 不能在这里写 defer conn.Close() 当创建连接出错,连接没创建出来执行关闭引发panic
	if err != nil {
		fmt.Println("tcp client 连接 tcp server 失败, err ：", err)
		return
	}
	// 创建连接成功后,保证对正常的连接执行关闭操作
	defer conn.Close()

	// 2. tcp client 向 tcp server 发消息

	// 打开文件句柄 和 打开网络socket连接都是1个IO,文件IO 网络IO ,都可以通过IO读写数据
	// fmt.Fprintln(conn, "i am tcp client")

	// 获取用户输入
	// var input string
	// 在终端一直等用户输入回车输入结束
	// fmt.Scanln(&input)
	// fmt.Scanf("%s\n", &input)
	// 这2种写法不能识别 终端输入的 ?

	// 从标准输入获取输入
	reader := bufio.NewReader(os.Stdin)
	// 读取内容直到遇到 \n 结束
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("tcp client 获取用户标准输入失败, err :", err)
		return
	}
	// 向连接 网络IO 写字节数据
	_, err = conn.Write([]byte(input))
	if err != nil {
		fmt.Println("tcp client 向 tcp server 发送消息失败, err :", err)
		return
	}
}
