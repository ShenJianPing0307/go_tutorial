package main

import "fmt"

func sum(n1 int, n2 int) int {
	// 当执行当defer会被延迟执行，先执行defer后面的语句
	// defer的语句会被压入到栈中，按照先如入后出的方式出栈
	// 当sum函数执行完毕后会执行defer
	defer fmt.Println("sum n1=", n1) // 第三步
	defer fmt.Println("sum n2=", n2) // 第二步

	res := n1 + n2
	fmt.Println("sum res=", res) // 第一步
	return res
}

func main() {

	res := sum(5, 10)
	fmt.Println("main res=", res) // 第四步

}

/*
输出：
sum res= 15
sum n2= 10
sum n1= 5
main res= 15
*/
