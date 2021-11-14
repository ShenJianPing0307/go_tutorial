package main

import "fmt"

func cal(n1 float64, n2 float64, operator byte) float64 {
	
	var res float64
	switch operator {

		case '+':
			res = n1 + n2
		case '-':
			res = n1 - n2
		case '*':
			res = n1 * n2
		case '/':
			res = n1 / n2
		default:
			fmt.Println("...")
	}
	return res

}

func main() {

	res := cal(3.25, 2.16, '+')
	fmt.Println("res=", res)
}