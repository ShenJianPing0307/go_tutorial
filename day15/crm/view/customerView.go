package main

import (
	"fmt"
	"go_tutorial/day15/crm/model"
	"go_tutorial/day15/crm/service"
)

type customerView struct {
	key             string //获取客户端的输入
	forloop         bool   // 判断 是否跳出for循环
	serviceCustomer service.CustomerService
}

// 主页面中添加customer，调用customerService中的Add方法
func (cv *customerView) add() {
	fmt.Println("姓名：")
	name := " "
	fmt.Scanln(&name)
	fmt.Println("性别")
	gender := " "
	fmt.Scanln(&gender)
	fmt.Println("年龄：")
	age := 0
	fmt.Scanln(&age)
	fmt.Println("电话：")
	phone := " "
	fmt.Scanln(&phone)
	fmt.Println("邮箱：")
	email := " "
	fmt.Scanln(&email)

	customer := model.CreateCustomer(name, gender, age, phone, email)
	if cv.serviceCustomer.Add(customer) {
		fmt.Println("---------添加完成---------")
	} else {
		fmt.Println("---------添加失败---------")
	}

}

// 主页面中查看customer，调用customerService中的List方法
func (cv *customerView) list() {
	fmt.Println("---------客户列表---------")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱")
	customers := cv.serviceCustomer.List()
	for i := 0; i < len(customers); i++ {
		fmt.Println(customers[i].ShowInfo())
	}
	fmt.Println("-------客户列表完成-------")
}

// 主页面中更新customer，调用customerService中的Update方法
func (cv *customerView) update() {
	fmt.Println("---------修改客户---------")
	fmt.Println("请输入待修改的客户编号：")
	Id := -1
	fmt.Scanln(&Id)
	// 这里需要判断是否存在，如果存在获取当前customer
	customer, flag := cv.serviceCustomer.GetCustomer(Id)
	if !flag {
		return
	}
	name := customer.Name
	fmt.Println("姓名：")

	fmt.Scanln(&name)
	gender := customer.Gender
	fmt.Println("性别：")

	fmt.Scanln(&gender)
	age := customer.Age
	fmt.Println("年龄：")

	fmt.Scanln(&age)
	phone := customer.Phone
	fmt.Println("电话：")
	fmt.Scanln(&phone)
	email := customer.Email
	fmt.Println("邮箱：")
	fmt.Scanln(&email)
	flag1 := cv.serviceCustomer.Update(Id, name, gender, age, phone, email)
	if flag1 {
		fmt.Println("---------修改完成---------")
	} else {
		fmt.Println("---------修改失败---------")
	}

}

// 主页面中删除customer，调用customerService中的Delete方法
func (cv *customerView) delete() {
	fmt.Println("---------删除客户---------")
	fmt.Println("请输入待删除的客户编号：")
	Id := -1
	fmt.Scanln(&Id)
	cv.serviceCustomer.Delete(Id)
	fmt.Println("---------删除完成---------")

}

// 主页面退出循环，通过修改forloop状态，退出循环
func (cv *customerView) exit() {
	fmt.Println("请确认是否退出(Y/N)：")
	for {
		fmt.Scanln(&cv.key)
		if cv.key == "Y" || cv.key == "y" || cv.key == "N" || cv.key == "n" {
			break
		}
		fmt.Println("你的输入有误，请重新输入(Y/N)：")
	}
	if cv.key == "Y" || cv.key == "y" {
		cv.forloop = false
	}

}

// 主页面
func (cv *customerView) mainMenu() {
	for {
		fmt.Println("---------客户关系管理系统---------")
		fmt.Println("            1 添加客户")
		fmt.Println("            2 修改客户")
		fmt.Println("            3 删除客户")
		fmt.Println("            4 客户列表")
		fmt.Println("            5 退   出")
		fmt.Println("            请选择(1-5):")

		// 获取用户输入
		fmt.Scanln(&cv.key)

		switch cv.key {

		case "1":
			cv.add()
		case "2":
			cv.update()
		case "3":
			cv.delete()
		case "4":
			cv.list()
		case "5":
			cv.exit()
		default:
			fmt.Println("你的输入有误，请重新输入...")
		}

		if !cv.forloop {
			break
		}

	}

}

// 主方法，初始化变量
func main() {
	cs := service.CustomerService{CustomerNum: 0}
	cv := customerView{
		key:             " ",
		forloop:         true,
		serviceCustomer: cs,
	}

	cv.mainMenu()
}
