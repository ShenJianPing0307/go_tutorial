package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	// 方式二/1
	var p1 Person = Person{
		"lily",
		20,
	}
	fmt.Println(p1)

	// 方式二/2
	p2 := Person{
		Name: "harry",
		Age:  30,
	}
	fmt.Println(p2)

	//方式二/3
	p3 := Person{}
	p3.Name = "alice"
	p3.Age = 36
	fmt.Print(p3)

}
