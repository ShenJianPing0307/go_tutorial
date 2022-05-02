package main

import "fmt"

func JudgeType(items ...interface{}) {
	for index, value := range items {
		switch value.(type) {
		case bool:
			fmt.Printf("第%v个参数时bool类型，值为%v\n", index, value)
		case float32:
			fmt.Printf("第%v个参数时float32类型，值为%v\n", index, value)
		case float64:
			fmt.Printf("第%v个参数时float64类型，值为%v\n", index, value)
		case int:
			fmt.Printf("第%v个参数时int类型，值为%v\n", index, value)
		case string:
			fmt.Printf("第%v个参数时string类型，值为%v\n", index, value)
		default:
			fmt.Printf("第%v个参数类型不确定，值为%v\n", index, value)
		}

	}
}

func main() {
	var a int = 10
	var b string = "happy"
	var c float32 = 3.21
	JudgeType(a, b, c)
}
