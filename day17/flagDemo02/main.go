package main

import (
	"flag"
	"fmt"
)

func main() {
	// 定义用于接收命令行指定变量
	var user string
	var pwd string
	var host string
	var port int

	// &user用来接收命令行输入的参数值, 即-u 后的值
	// "u" 就是 -u 参数
	// "" 默认值
	// "用户名默认为空" 用于说明

	flag.StringVar(&user, "u", "", "用户名默认为空")
	flag.StringVar(&pwd, "p", "", "密码默认为空")
	flag.StringVar(&host, "h", "localhost", "主机名默认为localhost")
	flag.IntVar(&port, "P", 3306, "端口号默认为3306")

	// 接收值后需要进行解析
	flag.Parse()

	// 输出
	fmt.Printf("user=%v, pwd=%v, host=%v, port=%v", user, pwd, host, port)

}
