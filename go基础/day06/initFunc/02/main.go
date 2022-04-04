package main
import "fmt"

var age = test()

// 通过此函数可以看到是否调用变量
func test() int {
	fmt.Println("test...")
	return 12
}

func init() {
	fmt.Println("init...")
}

func main() {
	fmt.Println("main...")
}
/*
输出：

test...
init...
main...
*/
