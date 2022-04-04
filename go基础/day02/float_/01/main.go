package main 

import "fmt"

func main() {
	// 浮点数在Go中的使用
	var a float32 = -0.369
	var b float64 = -0625

	fmt.Println("a=", a, "b=", b) //a= -0.369 b= -405

	// 尾数可能造成精度的丢失
	var c float32 = -203.50000006
	var d float64 = -203.50000006
	fmt.Println("c=", c, "d=", d) //c= -203.5 d= -203.50000006

	// Go中默认浮点型存储类型为float64
	var e = 0.236
	fmt.Printf("e的类型为%T", e) // e的类型为float64
}