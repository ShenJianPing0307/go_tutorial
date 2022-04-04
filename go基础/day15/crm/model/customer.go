package model

import "fmt"

// 声明一个customer结构体
type Customer struct {
	Id     int
	Name   string
	Gender string
	Age    int
	Phone  string
	Email  string
}

// 定义一个结构体变量创建的工厂方法，其中Id让其自动生成, 生成的实例中id默认为0
func CreateCustomer(name string, gender string, age int, phone string, email string) Customer {
	return Customer{
		Name:   name,
		Gender: gender,
		Age:    age,
		Phone:  phone,
		Email:  email,
	}
}

// 显示customer信息
func (customer *Customer) ShowInfo() string {
	res := fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v\t", customer.Id,
		customer.Name, customer.Gender, customer.Age, customer.Phone, customer.Email)
	return res

}
