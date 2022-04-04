package main

import "fmt"

func main() {
	// 省略var，使用 := 的方式，表示进行声明赋值，可以将 : 理解成变量声明
	i, j := 1, "aliy"
	
	fmt.Println("i=", i, "j=", j) //i= 1 j= aliy
}