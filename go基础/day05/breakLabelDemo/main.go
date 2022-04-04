package main
import "fmt"

func main() {
	label1:
	for i := 0; i < 10; i++ {
		if i > 5 {
			break label1
		}
		fmt.Println("i=", i)
	}
}
/* 执行结果
i= 0
i= 1
i= 2
i= 3
i= 4
i= 5
*/