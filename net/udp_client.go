package main

import (
	"fmt"
	"net"
)

// UDP client
func main() {
	// TODO : NOTICE 官方库返回 Conn 不恰当 , udp协议没有建立连接
	conn, err := net.Dial("udp", "127.0.0.1:30000")
	if err != nil {
		fmt.Println("udp client 连接 udp server 失败，err: ", err)
		return
	}
	defer conn.Close()

	n, err := conn.Write([]byte("i am udp client"))
	if err != nil {
		fmt.Println("udp client 发送消息给 udp server 失败，err: ", err)
		return
	}

	// udp client 接收 udp server 消息
	var buf [1024]byte
	n, err = conn.Read(buf[:])
	if err != nil {
		fmt.Println("udp client 读取 udp server 消息失败,err: ", err)
		return
	}
	fmt.Println("收到回复: ", string(buf[:n]))
}

/*
go run udp_client.go
收到回复:  i have received i am udp client
*/