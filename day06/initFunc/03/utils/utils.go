package utils

import "fmt"

var Age int

// Age 全局变量，在init函数中进行初始化
func init() {
	fmt.Println("utils包init...")
	Age = 20
}