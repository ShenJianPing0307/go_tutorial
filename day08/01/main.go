package main

import (
	"fmt"
	"time"
)

func main() {

	now := time.Now()
	fmt.Printf("now=%v type=%T", now, now) // now=2021-11-21 21:55:00.8398932 +0800 CST m=+0.004958701 type=time.Time

}