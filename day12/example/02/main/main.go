package main

import (
	"fmt"
	"go_tutorial/day12/example/02/model"
)

func main() {

	// 创建一个person的结构体变量
	p := model.CreatePerson("ariry")
	p.SetAge(30)
	p.SetSalary(5000)
	fmt.Println("姓名：", p.Name, "年龄：", p.GetAge(), "薪资：", p.GetSalary())

}
