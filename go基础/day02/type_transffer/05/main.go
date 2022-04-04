package main

import (
	"fmt"
	"strconv"
)

func main() {
	var str1 string = "123"
	var n1 int64
	// strconv.ParseInt返回两个值(i int64, err error)
	// 将第二个err忽视使用 _
	n1, _ = strconv.ParseInt(str1, 10, 64)
	fmt.Printf("n1的类型位%T n1=%v\n", n1, n1) //n1的类型位int64 n1=123

	var str2 string = "123.56"
	var f1 float64
	f1, _ = strconv.ParseFloat(str2, 64)
	fmt.Printf("f1的类型位%T f1=%v\n", f1, f1) //f1的类型位float64 f1=123.56

	var str3 string = "false"	
	var b1 bool
	b1, _ = strconv.ParseBool(str3)
	fmt.Printf("b1的类型位%T f1=%v\n", b1, b1) //b1的类型位bool f1=false

}
