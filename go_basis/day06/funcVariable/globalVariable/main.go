package main

import "fmt"

// 定义初始化全局变量，该包中所有函数都可以使用，其中Age首字母大写，其它包也可以使用
/*
这里不能写成 username := "zhangsan" 和 Age  := 20
因为等价于：
    var username string
	username = "zhangsan"
但是第二句的赋值语句不能再函数体外，所以是错误的
*/
var username string = "zhangsan"
var Age int = 20

func test() {
	fmt.Println("username", username)
	fmt.Println("Age", Age)
}

func main() {
	test()
	fmt.Println("main username", username)
	fmt.Println("main Age", Age)
}
