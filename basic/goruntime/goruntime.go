package main

import (
	"fmt"
	"sync"
)

// 声明全局等待组变量
var wg sync.WaitGroup

func main() {
	for i := 0; i < 10; i++ {
		// 登记1个goroutine
		wg.Add(1)
		go hello(i)
	}
	// 阻塞等待登记的goroutine完成
	wg.Wait()
}

func hello(i int) {
	fmt.Println("hello", i)
	// 告知 go 当前的 goruntime 结束
	defer wg.Done()
}
