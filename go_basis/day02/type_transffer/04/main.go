package main

import (
	"fmt"
	"strconv"
)

func main() {

	var a1 int = 10
	var a2 float64 = 10.236
	var a3 bool = false

	// FormatInt函数第一个参数是int64类型的参数，所以需要转换
	b1 := strconv.FormatInt(int64(a1), 10)
	fmt.Printf("b1类型%T b1=%q\n", b1, b1)
	// 'f'表示格式，10表示小数位保留10位，64表示这个小数是float64类型
	b2 := strconv.FormatFloat(a2, 'f', 10, 64)
	fmt.Printf("b2类型%T b2=%q\n", b2, b2)
	b3 := strconv.FormatBool(a3)
	fmt.Printf("b3类型%T b3=%q\n", b3, b3)
	/*
		  输出
			b1类型string b1="10"
			b2类型string b2="10.2360000000"
			b3类型string b3="false"
	*/
}
