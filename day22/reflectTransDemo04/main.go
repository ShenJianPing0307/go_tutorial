package main

import (
	"fmt"
	"reflect"
)

func transType(v interface{}) {
	// 使用interface{}接收任意类型的反射
	// 1、将interface{}转成reflect.Value类型
	rVal := reflect.ValueOf(v)
	// 2、将reflect.Value调用对应的方法修改值
	rVal.Elem().SetInt(20)

}

func main() {
	var num = 10
	transType(&num)
	fmt.Println(num) // 20
}
