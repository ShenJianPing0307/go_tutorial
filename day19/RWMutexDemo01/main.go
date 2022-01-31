package main

import (
	"fmt"
	"sync"
)

var mapNum = make(map[int]int, 50)
var wg sync.WaitGroup
var rwlock sync.RWMutex // 声明一个读写锁变量

func write(i int) {
	defer wg.Done()
	rwlock.Lock()
	for j := 1; j <= i; j++ {
		res := 1
		res = res * i
		mapNum[i] = res

	}
	rwlock.Unlock()
}

func read() {
	defer wg.Done()
	rwlock.RLock()
	fmt.Printf("读取数据\n")
	for i, v := range mapNum {
		fmt.Printf("i=%d, n=%d \n", i, v)
	}
	fmt.Printf("读取结束\n")
	rwlock.RUnlock()
}

func main() {
	// 启动10个读，10个写的协程
	wg.Add(22)

	for i := 0; i < 20; i++ {
		go write(i)
	}

	for i := 0; i < 2; i++ {
		go read()
	}

	// 此时等待20个协程执行完毕
	wg.Wait()

}
