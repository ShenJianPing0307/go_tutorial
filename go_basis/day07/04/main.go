package main

import (
	"fmt"
	"strconv"
)

func main() {

	str := strconv.Itoa(123)
	fmt.Printf("str值为%v,str类型为%T", str, str) // str值为123,str类型为string

}
