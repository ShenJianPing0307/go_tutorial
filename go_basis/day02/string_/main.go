package main

import "fmt"

func main() {
	// 1、string基本使用
	var a = "hello world!"
	fmt.Println(a)

	// 2、Go中的字符串是不可变的
	// a[0] = 'k' // cannot assign to a[0] (strings are immutable)
	// fmt.Println(a)

	// 3、字符串中转义字符使用
	b := "aaa\nbbb"
	fmt.Println(b)

	// 4、反引号``，以字符串原生形式输出
	c := `
	  func main() {

	  }
	`
	fmt.Println(c)

	// 5、字符串拼接
	d := "hello" + "world"
	fmt.Println(d)

	// 6、字符串多行处理
	// 字符串分行，保留 + 到上一行
	e := "hello" +
		"world"
	fmt.Println(e)
}
