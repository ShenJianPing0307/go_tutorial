package main

import (
	"fmt"
	"reflect"
)

type User struct {
	NickName string
	Age      int
}

func transType(v interface{}) {
	// 使用interface{}接收任意类型的反射
	// 1、将interface{}转成reflect.Value类型
	rVal := reflect.ValueOf(v)
	// 2、将reflect.Value转成interface{}
	iVal := rVal.Interface()
	// 3、通过断言将interface{}转成User类型
	user := iVal.(User)

	fmt.Printf("原类型为%T", user) //原类型为main.User

}

func main() {
	user := User{
		NickName: "lily",
		Age:      20,
	}

	transType(user)
}
