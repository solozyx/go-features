package main

import (
	"fmt"
	"net"
)
// TCP协议
// 基于连接稳定可靠,客户端 与 服务端通信,必须先建立连接
// TCP/IP(Transmission Control Protocol/Internet Protocol) 即传输控制协议/网间协议
// 是一种面向连接（连接导向）的、可靠的、基于字节流的传输层（Transport layer）通信协议
// 因为是面向连接的协议，数据像水流一样传输，会存在黏包问题。

// TCP服务端
// 一个TCP服务端可以同时连接很多个客户端，例如世界各地的用户使用自己电脑上的浏览器访问淘宝网。
// 因为Go语言中创建多个goroutine实现并发非常方便和高效
// 所以我们可以每建立一次链接就创建一个goroutine去处理

// TCP服务端程序处理流程:
// 1.监听端口,机器一共有 65535 个端口, 1024 往后的端口都能用
// 2.等待接收客户端请求服务端监听的端口,和客户端建立连接
// 3.创建goroutine处理连接,把处理每个客户端连接的请求封装为1个serve函数,1个server可以接收N多个client请求
// 使用goroutine实现并发的性能可以和nginx相当

func main() {
	// 1. 监听端口
	listener, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("tcp server 监听 :20000 端口失败, err :", err)
		return
	}
	// 程序退出时释放 20000 端口
	defer listener.Close()

	// 2. 接收客户端请求建立链接
	for {
		// 如果没有客户端连接就阻塞 一直在等待
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("tcp server 和 tcp client 连接失败，err :", err)
			continue
			// 连接没有建立,不要写 conn.Close()
		}
		// 3. 每个连接请求创建1个goroutine进行处理
		go serve(conn)
	}
}

// 单独处理连接的函数
func serve(conn net.Conn) {
	// 单独的 goroutine 结束之后关闭连接
	defer conn.Close()

	// tcp server 从连接中接收 tcp client 数据
	var buf [1024]byte
	// 把数据读取到buf中, n 表示读了多少个字节的数据, 1024个字节用不完
	// 也可能发的数据超过 1024个字节 也是只取 1024个字节的数据
	n, err := conn.Read(buf[:]) // bufio.NewReader(conn)
	// buf := make([]byte, 1024)
	if err != nil {
		fmt.Println("tcp server 接收 tcp client 发来的消息失败, err :", err)
		return
	}
	fmt.Println("tcp server 接收 tcp client 发来的消息: ", string(buf[:n]))
}
