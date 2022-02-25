package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:9000")
	if err != nil {
		return
	}

	reader := bufio.NewReader(os.Stdin)

	// 进行通讯
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("客户端读取终端输入出错了...")
		}
		// 将line数据发送到服务器
		_, err = conn.Write([]byte(line))
		if err != nil {
			fmt.Println("客户端发送数据出错了...")
		}

	}

}
