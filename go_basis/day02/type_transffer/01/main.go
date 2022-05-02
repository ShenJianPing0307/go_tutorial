package main

import "fmt"

func main() {
	// 数据类型默认值
	var a int
	var b float32
	var c float64
	var isStatus bool
	var username string

	// 打印出默认值
	fmt.Printf("a=%d, b=%v, c=%v, isStatus=%v, username=%v", a, b, c, isStatus, username) //a=0, b=0, c=0, isStatus=false, username=""
}
