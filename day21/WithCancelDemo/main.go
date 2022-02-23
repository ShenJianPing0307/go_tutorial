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
		case <-ctx.Done(): // 检查cancel是否到时间执行
			fmt.Println("获取信息结束！")
			return
		default:
			fmt.Println("获取信息！")
		}
	}
}

func main() {
	wg.Add(1)
	ctx, cancel := context.WithCancel(context.Background())
	go getInfo(ctx)
	time.Sleep(time.Second * 5)
	cancel() // 5s后执行cancel函数
	wg.Wait()

}

/*
获取信息！
获取信息！
获取信息！
获取信息！
获取信息结束！
*/
