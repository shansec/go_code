package main

import (
	"fmt"
	"sync"
)

var (
	x  int64
	wg sync.WaitGroup // 等待组
	m  sync.Mutex     // 互斥锁
)

func add() {
	for i := 0; i < 5000; i++ {
		m.Lock()
		x = x + 1
		m.Unlock()
	}
	wg.Done()
}

func main() {
	wg.Add(2)

	go add()
	go add()

	wg.Wait()
	fmt.Println(x)
}
