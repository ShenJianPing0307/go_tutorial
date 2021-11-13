package main

import "fmt"

func main() {

	var a1 int = 10
	var a2 float64 = 10.236
	var a3 bool = false
	var a4 byte = 'k'

	// fmt.Sprintf转换
	b1 := fmt.Sprintf("%d", a1)
	fmt.Printf("b1类型%T b1=%q\n", b1, b1)
	b2 := fmt.Sprintf("%f", a2)
	fmt.Printf("b2类型%T b2=%q\n", b2, b2)
	b3 := fmt.Sprintf("%t", a3)
	fmt.Printf("b3类型%T b3=%q\n", b3, b3)
	b4 := fmt.Sprintf("%c", a4)
	fmt.Printf("b4类型%T b4=%q\n", b4, b4)
	/*
	  输出
	  b1类型string b1="10"
	  b2类型string b2="10.236000"
	  b3类型string b3="false"    
      b4类型string b4="k"        
      b5类型string b5="hello"  
	*/
}