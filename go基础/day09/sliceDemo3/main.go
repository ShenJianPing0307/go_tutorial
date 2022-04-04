package main

import "fmt"

func main() {

	// 定义一个切片，直接指定具体数组，类似make
	var s3 []string = []string{"zhangsan", "lisi", "wangwu"} 

	fmt.Println("s3 =", s3) // s3 = [zhangsan lisi wangwu]
	fmt.Println("len =", len(s3)) // len = 3
	fmt.Println("cap =", cap(s3)) // cap = 3

}