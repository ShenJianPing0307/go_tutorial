package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// 接收两个路径参数
func CopyFile(dstFileName string, srcFileName string) (written int64, err error) {
	// 打开源文件
	srcFile, err := os.Open(srcFileName)
	if err != nil {
		fmt.Printf("open srcFile err %v", err)
		return
	}
	// 延时关闭文件
	defer srcFile.Close()
	// 创建一个读文件的缓冲区
	reader := bufio.NewReader(srcFile)

	// 打开目标地址
	dstFile, err := os.OpenFile(dstFileName, os.O_CREATE|os.O_WRONLY, 0666)

	if err != nil {
		fmt.Printf("open dstFile err %v", err)
		return
	}
	// 创建一个写文件的缓冲区
	writer := bufio.NewWriter(dstFile)
	// 延时关闭文件
	defer dstFile.Close()

	return io.Copy(writer, reader)

}

func main() {
	srcFileName := "d:/1.jpg"
	dstFileName := "e:/1.jpg"
	_, err := CopyFile(dstFileName, srcFileName)
	if err == nil {
		fmt.Println("拷贝完成")
	} else {
		fmt.Printf("拷贝错误 %v", err)
	}
}
