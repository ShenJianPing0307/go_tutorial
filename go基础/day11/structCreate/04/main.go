package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	// 方法四 var p1 *Person = &Person
	var p1 *Person = &Person{}
	/*
		p1是一个指针，访问字段的标准方式是(*p1).Name = "egg",
		但是go的设计者为了使用方便在底层进行了处理，所以可以这样写：
		p1.Name = "egg"，实际底层会对p1变成(*p1)
	*/
	p1.Name = "egg"
	fmt.Println(*p1)

}
