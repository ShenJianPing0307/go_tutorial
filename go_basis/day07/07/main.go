package main

import (
	"fmt"
	"strconv"
)

func main() {

	str1 := strconv.FormatInt(123, 2)
	fmt.Printf("123对应的二进制为%v\n", str1) // str值为123,str类型为string

	str2 := strconv.FormatInt(123, 16)
	fmt.Printf("123对应的十六进制为%v", str2) // str值为123,str类型为string

}

/*
123对应的二进制为1111011
123对应的十六进制为7b
*/
