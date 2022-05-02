package service

import (
	"go_tutorial/day15/crm/model"
)

type CustomerService struct {
	// 声明一个切片，用于存放customer
	CustomerSlice []model.Customer

	// 定义一个用于保存customer自增Id的变量，下一个Id就是Id+1
	CustomerNum int
}

// 向切片中添加customer
func (cs *CustomerService) Add(customer model.Customer) bool {
	cs.CustomerNum++
	customer.Id = cs.CustomerNum
	cs.CustomerSlice = append(cs.CustomerSlice, customer)
	return true
}

// 查询切片中的customer
func (cs *CustomerService) List() []model.Customer {

	return cs.CustomerSlice
}

// 查询切片中的编号
func (cs *CustomerService) FindById(id int) int {

	for index, value := range cs.CustomerSlice {
		if value.Id == id {
			return index
		}
	}
	return -1
}

// 删除customer
func (cs *CustomerService) Delete(id int) bool {
	index := cs.FindById(id)
	if index == -1 {
		// 说明不存在这个customer
		return false
	}
	cs.CustomerSlice = append(cs.CustomerSlice[:index], cs.CustomerSlice[index+1:]...)
	return true
}

// 获取一个customer
func (cs *CustomerService) GetCustomer(id int) (model.Customer, bool) {
	index := cs.FindById(id)
	if index != -1 {
		return cs.CustomerSlice[index], true
	}
	return model.Customer{}, false
}

// 修改customer
func (cs *CustomerService) Update(id int, name string, gender string, age int, phone string, email string) bool {
	index := cs.FindById(id)
	if index == -1 {
		// 说明不存在这个customer
		return false
	}
	cs.CustomerSlice[index].Name = name
	cs.CustomerSlice[index].Gender = gender
	cs.CustomerSlice[index].Age = age
	cs.CustomerSlice[index].Phone = phone
	cs.CustomerSlice[index].Email = email
	return true
}
