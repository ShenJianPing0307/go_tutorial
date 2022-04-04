package main

import "fmt"

type Point struct {
	x int
	y int
}

type Rect struct {
	leftUp, rightDown Point
}

type Rect1 struct {
	leftUp, rightDown *Point
}

func main() {

	r1 := Rect{
		Point{1, 2},
		Point{3, 4},
	}

	r2 := Rect1{
		&Point{1, 2},
		&Point{3, 4},
	}

	// r1有4个int，在内存中是连续分布的
	fmt.Printf("r1.leftUp.x地址为：%p r1.leftUp.y地址为：%p r1.rightDown.x地址为：%p r1.rightDown.y地址：%p \n",
		&r1.leftUp.x, &r1.leftUp.y, &r1.rightDown.x, &r1.rightDown.y)
	/*
		r1.leftUp.x地址为：0xc00000e200
		r1.leftUp.y地址为：0xc00000e208
		r1.rightDown.x地址为：0xc00000e210
		r1.rightDown.y地址: 0xc00000e218
	*/

	// r2有2个*Point类型，它们本身的地址也是连续的
	fmt.Printf("r2.leftUp地址为：%p r2.rightDown地址为：%p \n", &r2.leftUp, &r2.rightDown)
	/*
		r2.leftUp地址为：0xc000050240
		r2.rightDown地址为：0xc000050248
	*/

	// 但是2个*Point类型指向的地址不一定是连续的
	fmt.Printf("r2.leftUp地址为：%p r2.rightDown地址为：%p", r2.leftUp, r2.rightDown)
	/*
		r2.leftUp地址为：0xc0000140b0
		r2.rightDown地址为：0xc0000140c0
	*/

}
