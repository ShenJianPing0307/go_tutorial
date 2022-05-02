package main

import (
	"fmt"
	"time"
)

func main() {

	now := time.Now()

	// 通过now获取年月日、时分秒
	date := fmt.Sprintf("年月日%d-%d-%d %d:%d:%d", now.Year(),
		now.Month(), now.Day(), now.Hour(),
		now.Minute(), now.Second())

	fmt.Printf("date=%v", date) // date=年月日2021-11-21 22:7:43

}
