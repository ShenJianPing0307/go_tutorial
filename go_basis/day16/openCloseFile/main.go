package main

import (
	"fmt"
	"os"
)

func main() {
	// 打开该目录下的1.txt文件
	file, err := os.Open("./1.txt")
	if err != nil {
		fmt.Println("文件打开出错...")
	}
	// 输出文件，它是一个指针
	fmt.Println(file) // &{0xc00007c780}

	// 关闭文件
	err = file.Close()
	if err != nil {
		fmt.Println("文件关闭出错...")
	}

}
