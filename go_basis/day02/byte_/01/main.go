package main

import "fmt"

func main() {

	// 使用byte来存储字符类型，值使用单引号('value')来表示，直接输出byte值就是输出对应的字符码值
	var a = '0'
	var b = 'h'
	fmt.Println("a=", a, "b=", b) //a= 48 b= 104

	// 如果输出对应的字符，需要格式化输出
	fmt.Printf("a=%c b=%c\n", a, b) //a=0 b=h

	// byte类型的表数范围是0~255，如果存储的字符码值大于这个，可以使用int类型
	//var c byte = '张' //constant 24352 overflows byte
	var c int = '张'
	fmt.Println("c=", c) //c= 24352

	// 如果给一个码值，按照格式化输出可以输出一个unicode字符
	var d int = 24352
	fmt.Printf("d=%c\n", d) //d=张

	// 字符类型可以运算
	var e = 'm' + 10
	fmt.Println("e=", e) //e= 119

}
