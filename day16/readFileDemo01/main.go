package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {

	// 打开一个文件
	file, err := os.Open("./1.txt")
	if err != nil {
		fmt.Println("文件打开出错...")
	}

	// 延时关闭文件, 函数执行完毕自动关闭文件
	defer file.Close()

	// 文件读操作
	// 返回的是一个*Reader类型的缓冲区, 默认的缓存大小4096
	reader := bufio.NewReader(file)
	// 循环读取文件内容
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			// io.EOF表示文件末尾，就结束读取
			break
		}
		// 输出每一行内容
		fmt.Println(str)
	}

}
