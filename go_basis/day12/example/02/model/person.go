package model

import "fmt"

type person struct {
	Name   string
	age    int
	salary int
}

// 结构体时小写变量名，提供工厂方法
func CreatePerson(name string) *person {

	return &person{
		Name: name,
	}
}

// 提供年龄赋值工厂方法
func (p *person) SetAge(age int) {
	if age >= 18 {
		p.age = age
	} else {
		fmt.Println("年龄输入不合法")
	}
}

// 提供获取年龄的方法
func (p *person) GetAge() int {
	return p.age
}

// 提供薪资赋值工厂方法
func (p *person) SetSalary(salary int) {
	p.salary = salary
}

// 提供获取薪资的方法
func (p *person) GetSalary() int {
	return p.salary
}
