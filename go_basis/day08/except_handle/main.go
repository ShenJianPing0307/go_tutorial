package main

import (
	"errors"
	_ "fmt"
)

func cal(number int) (err error) {
	if number != 5 {
		return errors.New("输入错误...")
	} else {
		return nil
	}
}

func test() {
	err := cal(10)
	if err != nil {
		panic(err)
	}

}

func main() {
	test()
}

/*
err= runtime error: integer divide by zero
*/
