package main

import (
	"fmt"
	"time"
)

func main() {

	now := time.Now()

	// 通过now获取年月日、时分秒
	fmt.Printf(now.Format("2006-01-02 15:04:05")) // 2021-11-21 22:10:11
	fmt.Printf(now.Format("2006-01-02")) // 2021-11-21
	fmt.Printf(now.Format("15:04:05")) // 22:10:11

}
