package main

import (
	"fmt"
	"reflect"
)

func transType(v interface{}) {
	// 使用interface{}接收任意类型的反射
	// 1、将interface{}转成reflect.Value类型
	rVal := reflect.ValueOf(v)
	// 2、将reflect.Value转成reflect.Type类型
	rType := rVal.Type()
	// 3、调用reflect.Type类型中的方法，获取类别
	kind := rType.Kind()
	fmt.Print(kind)

}

func main() {
	var num = 10

	transType(num)
}
