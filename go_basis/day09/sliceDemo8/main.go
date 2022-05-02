package main

import "fmt"

func main() {

	str8_1 := "abcdef"

	// 通过 string-->[]byte方式修改值，但是不能处理中文字符，否则乱码，因为byte是按字节处理，而一个汉字三个字节
	s8_1 := []byte(str8_1)
	s8_1[0] = 'h'
	str8_2 := string(s8_1)
	fmt.Println(str8_2) // hbcdef

	// string-->[]rune, 可以处理中文问题
	s8_3 := []byte(str8_1)
	s8_3[0] = 'f'
	s8_4 := string(s8_3)
	fmt.Println(s8_4) // fbcdef

}
