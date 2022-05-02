package main

import "fmt"

func main() {
	// 声明map
	var a map[string]string

	// 在使用map前需要使用make分配空间
	a = make(map[string]string, 10)
	a["username"] = "张三"
	a["addr"] = "张家界"

	fmt.Println(a) // map[addr:张家界 username:张三]
}
