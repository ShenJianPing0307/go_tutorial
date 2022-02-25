package main

import (
	"fmt"
	"net"
)

func processReq(conn net.Conn) {
	// 延时关闭conn
	defer conn.Close()
	// 与服务端进行通讯
	for {
		// 创建一个切片，盛放消息
		buf := make([]byte, 1024)
		// 1、等待客户端通过conn发送消息
		// 2、如果没有消息，协程就阻塞在这里
		m, err := conn.Read(buf)
		if err != nil {
			fmt.Println("客户端退出")
			return
		}
		// 3、显示客户端发送的内容
		fmt.Print(string(buf[:m]))
	}
}

func main() {
	fmt.Println("服务器开始监听，等待客户端的连接：")
	ls, err := net.Listen("tcp", "0.0.0.0:9000")

	if err != nil {
		return
	}
	// 延时关闭listen
	defer ls.Close()

	// 等待客户端连接
	for {
		fmt.Println("等待客户端来连接：")
		conn, err := ls.Accept()
		if err != nil {
			return
		} else {
			fmt.Printf("con=%v, 客户端ip=%v\n", conn, conn.RemoteAddr().String())
		}

		// 每一个客户端启动一个协程处理，防止过长时间相应
		go processReq(conn) //conn接口类型，是引用

	}
}
