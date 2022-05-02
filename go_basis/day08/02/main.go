package main

import (
	"fmt"
	"time"
)

func main() {

	now := time.Now()

	// 通过now获取年月日、时分秒
	fmt.Printf("年=%v\n", now.Year())
	fmt.Printf("月=%v\n", now.Month())
	fmt.Printf("日=%v\n", now.Day())
	fmt.Printf("时=%v\n", now.Hour())
	fmt.Printf("分=%v\n", now.Minute())
	fmt.Printf("秒=%v\n", now.Second())

}

/*
年=2021
月=November
日=21
时=21
分=58
秒=45
*/
