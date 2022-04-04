package main
import (
	"fmt"
	"unsafe"
)

func main() {
	var a = 10

	// 查看默认类型
	fmt.Printf("a的类型为%T\n", a) //a的类型为int

	// 查看a变量占用的字节的大小
	fmt.Printf("a占用的字节大小%d", unsafe.Sizeof(a)) //a占用的字节大小8
}