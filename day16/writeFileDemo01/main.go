package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	filePath := "./2.txt"
	// 只写模式打开文件、如果文件不存在就创建
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("文件错误...")
		return
	}
	// 函数执行完毕关闭文件
	defer file.Close()

	// 向文件写入内容
	str := "hello world!"

	// 使用带缓存的写入
	writer := bufio.NewWriter(file)
	writer.WriteString(str)

	// writer带缓存的，所以在调用WriterString方法时，其实
	// 时先写入到缓存中，所以需要调用Flush方法，将缓存中的数据
	// 真正写入到文件，否则文件中没有数据
	writer.Flush()

}
