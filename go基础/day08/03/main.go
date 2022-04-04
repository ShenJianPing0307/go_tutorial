package main

import (
	"fmt"
	"time"
)

func main() {

	now := time.Now()

	// 通过now获取年月日、时分秒
	fmt.Printf("年月日%d-%d-%d %d:%d:%d", now.Year(),
	now.Month(),now.Day(),now.Hour(),
	now.Minute(),now.Second()) // 年月日2021-11-21 22:4:53

}
