package main

import "fmt"

type Cell struct {
	length float64
	width  float64
	height float64
}

func (cell *Cell) GetVolume() float64 {

	return cell.length * cell.width * cell.height

}

func main() {
	// 创建一个立方体结构体变量
	var cell Cell
	// 为每个字段赋值
	fmt.Print("请输入立方体的长：")
	fmt.Scanln(&cell.length)
	fmt.Print("请输入立方体的宽：")
	fmt.Scanln(&cell.width)
	fmt.Print("请输入立方体的高：")
	fmt.Scanln(&cell.height)

	// 调用结构体变量的方法
	res := cell.GetVolume()
	fmt.Println(res)

}
