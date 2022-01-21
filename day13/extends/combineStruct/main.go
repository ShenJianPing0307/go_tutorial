package main

type Student struct {
	Name string
	Age  int
}

type MiddleStudent struct {
	s Student // 有名结构体，组合关系
}

func main() {
	var ms MiddleStudent
	// 必须带上有名结构体的名字s
	ms.s.Name = "Alice"
	ms.s.Age = 25
}
