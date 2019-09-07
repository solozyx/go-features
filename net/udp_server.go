package main

import (
	"fmt"
	"net"
)

// UDP server端
func process(listener *net.UDPConn) {
	// 释放监听端口,否则占据端口其他程序无法使用
	defer listener.Close()

	// 通过 listener 收发数据
	for {
		// udp server 接收 udp client 消息
		var buf [1024]byte
		// 数组转为切片 buf[:]
		// int 读取数据量
		// udpAddr 因为udp协议不是基于连接的,任何客户端都能给服务端发数据,区分客户端 方便返回数据给客户端
		n, udpAddr, err := listener.ReadFromUDP(buf[:])
		if err != nil {
			fmt.Println("udp server 接收 udp client 消息失败：err: ", err)
			return
		}
		fmt.Printf("udp server 接收到来自 %v 的消息: %v\n", udpAddr, string(buf[:n]))

		// udp server 回复消息给 udp client
		n, err = listener.WriteToUDP([]byte("i have received " + string(buf[:n])), udpAddr)
		if err != nil {
			fmt.Println("udp server 回复消息给 udp client 失败，err: ", err)
			return
		}
	}
}

func main() {
	listener, err := net.ListenUDP("udp", &net.UDPAddr{
		// IP:   net.IPv4(0,0,0,0),
		IP:   net.ParseIP("127.0.0.1"),
		Port: 30000,
		// Zone 在 ipv6使用,ipv4用不到
	})
	if err != nil {
		fmt.Println("启动 udp server 失败，err:", err)
		return
	}

	process(listener)
}

/*
go run udp_server.go
udp server 接收到来自 127.0.0.1:50112 的消息: i am udp client
udp server 接收到来自 127.0.0.1:50202 的消息: i am udp client
*/