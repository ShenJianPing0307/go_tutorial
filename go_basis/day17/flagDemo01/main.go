package main

import (
	"fmt"
	"os"
)

func main() {
	// 获取命令行参数
	fmt.Printf("命令行参数个数为：%v \n", len(os.Args))

	// 打印每一个参数
	for i, v := range os.Args {

		fmt.Printf("第%v个参数是：%v \n", i, v)
	}

}
