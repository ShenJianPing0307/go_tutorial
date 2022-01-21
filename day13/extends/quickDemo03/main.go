package main

import "fmt"

type Student struct {
	Name string
	Age  int
}

type MiddleStudent struct {
	Student
}

type UgStudent struct {
	*Student
}

func main() {
	// 直接给匿名结构体初始化赋值
	ms := MiddleStudent{
		Student{Name: "kity", Age: 21},
	}

	ugs := UgStudent{
		&Student{Name: "jack", Age: 25},
	}

	fmt.Print(ms)
	fmt.Print(*ugs.Student)

}
