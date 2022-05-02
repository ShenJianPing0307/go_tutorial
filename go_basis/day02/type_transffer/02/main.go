package main

import "fmt"

func main() {
	var i int32 = 10

	// i int32-->float
	var j float32 = float32(i)

	// 高精度-->低精度
	var k int8 = int8(i)

	// 低精度-->高精度
	var m int64 = int64(i)

	fmt.Printf("i=%v j=%v k=%v m=%v", i, j, k, m) //i=10 j=10 k=10 m=10

}
