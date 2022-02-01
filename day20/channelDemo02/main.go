package main

import "fmt"

type User struct {
	name string
	age  int
}

func main() {
	var userChan chan User = make(chan User, 2)
	userChan <- User{
		name: "bili",
		age:  20,
	}
	userChan <- User{
		name: "harry",
		age:  23,
	}

	// 输出结构体变量属性
	var user1 User = <-userChan
	fmt.Printf("name=%v, age=%v", user1.name, user1.age)

}
