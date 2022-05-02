package main

import "fmt"

func test1() {
	// 关闭文件资源
	f := openfile("filePath")
	defer f.close()
	// 操作文件代码
	// ...
}

func test2() {
	// 关闭数据库资源
	connect := openDatabase("connect path")
	defer connect.close()
	// 操作数据库代码
	// ...

}

func main() {

}
