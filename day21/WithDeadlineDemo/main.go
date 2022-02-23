package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func getInfo(ctx context.Context) {
	defer wg.Done()
	for {
		time.Sleep(time.Second)
		select {
		case <-ctx.Done(): // 到时间会自动调用cancel函数，结束协程函数
			fmt.Println("获取信息结束！")
			return
		default:
			fmt.Println("获取信息！")
		}
	}
}

func main() {
	wg.Add(1)
	d := time.Now().Add(time.Second * 5)
	ctx, _ := context.WithDeadline(context.Background(), d) // 如果没有返回cancel函数也是可行的，只要在规定的d时间内就自动结束传入ctx的函数
	go getInfo(ctx)
	wg.Wait()

}

/*
获取信息！
获取信息！
获取信息！
获取信息！
获取信息结束！
*/
