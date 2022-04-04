package main

import "fmt"

// 全局变量的使用
var a = 1
var b = 2

// 也可以一次声明多个变量
var (
	c = 3
	d = 4
)

func main() {

	fmt.Println("a=", a, "b=", b, "c=", c, "d=", d) //a= 1 b= 2 c= 3 d= 4
}