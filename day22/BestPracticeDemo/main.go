package main

import (
	"fmt"
	"reflect"
)

type User struct {
	UserName string `json:"username"`
	Age      int    `json:"age"`
}

func (u User) GetUser() {
	fmt.Println("获取用户信息")
}

func (u User) GetAge(age int, num int) int {
	lastAge := age + num
	return lastAge
}

func TestRefStru(v interface{}) {
	// 对结构体变量进行反射
	// 获取reflect.Type类型
	rType := reflect.TypeOf(v)
	// 获取reflect.Value类型
	rVal := reflect.ValueOf(v)

	// 获取v对应的类别
	kd := rVal.Kind()

	// 如果传入的非struct就退出
	if kd != reflect.Struct {
		return
	}

	// 获取结构体中有多少个字段
	FieldNum := rVal.NumField()

	// 对字段的tag进行反射，获取tag标签值
	for i := 0; i < FieldNum; i++ {
		// 获取struct标签，通过reflect.Type来获取tag标签值
		tagVal := rType.Field(i).Tag.Get("json")
		// 如果字段存在就打印出来
		if tagVal != "" {
			fmt.Printf("tag:%v \n", tagVal)
		}
	}

	// 获取结构体中有多少个方法
	MethodNum := rVal.NumMethod()
	fmt.Println("方法个数为:", MethodNum)

	// 方法排序是按照ASCII排序，与在结构体中的顺序无关
	// 该方法无参数的调用
	rVal.Method(1).Call(nil) //调用第一个方法

	// 含参数的调用，注意Call方法传入的的[]reflect.Value类型的参数,返回值是[]reflect.Value类型
	var parames []reflect.Value
	parames = append(parames, reflect.ValueOf(20))
	parames = append(parames, reflect.ValueOf(5))
	res := rVal.Method(0).Call(parames)
	fmt.Println("年龄", res[0].Int())
}

func main() {
	// 创建一个User实例
	u := User{
		UserName: "lily",
		Age:      20,
	}
	TestRefStru(u)
}
