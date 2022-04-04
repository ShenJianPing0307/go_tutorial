package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	// 打开一个文件
	content, err := ioutil.ReadFile("./1.txt")
	if err != nil {
		fmt.Println("文件打开出错...")
	}

	// 文件读操作
	// content是一个字节切片类型
	fmt.Println(string(content))

}
