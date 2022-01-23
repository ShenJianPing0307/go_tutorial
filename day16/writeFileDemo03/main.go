package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {

	filePath := "./2.txt"
	// 覆盖已经存在文件中的内容
	file, err := os.OpenFile(filePath, os.O_RDONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("文件错误...")
		return
	}
	// 函数执行完毕关闭文件
	defer file.Close()

	// 读文件
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			fmt.Println("已经读到文件末尾了...")
			break
		}
		fmt.Print(str)
	}

	// 向文件中写内容
	str := "hello iveboy"

	// 使用带缓存的写入
	writer := bufio.NewWriter(file)
	writer.WriteString(str)

	// writer带缓存的，所以在调用WriterString方法时，其实
	// 时先写入到缓存中，所以需要调用Flush方法，将缓存中的数据
	// 真正写入到文件，否则文件中没有数据
	writer.Flush()

}
