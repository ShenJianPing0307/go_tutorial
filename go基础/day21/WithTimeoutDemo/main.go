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
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	go getInfo(ctx)
	defer cancel() // 超时之后释放资源
	wg.Wait()

}

/*
获取信息！
获取信息！
获取信息！
获取信息！
获取信息结束！
*/
