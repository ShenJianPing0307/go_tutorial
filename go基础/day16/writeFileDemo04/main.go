package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	// 1、将已经存在的文件a.txt中的内容读取到内存
	// 2、将内存中的内容写入到文件b.txt

	filePath1 := "./a.txt"
	filePath2 := "./b.txt"

	// 读取a.txt文件中的内容
	data, err := ioutil.ReadFile(filePath1)
	if err != nil {
		fmt.Println("文件读取错误...")
		return
	}

	// 打开写入的文件，如果文件不存在就会创建，然后写入数据
	err = ioutil.WriteFile(filePath2, data, 0666)
	if err != nil {
		fmt.Println("打开写入的文件错误...")
	}

}
