package main

//import (
//	"fmt"
//	"sync"
//	"time"
//)
//
//func main() {
//	l := &sync.RWMutex{}
//	go lockFunc(l)
//	go lockFunc(l)
//	go lockFunc(l)
//	go readLockFunc(l)
//	go readLockFunc(l)
//	go readLockFunc(l)
//	go readLockFunc(l)
//	for {
//	}
//}
//
//func lockFunc(lock *sync.RWMutex) {
//	lock.Lock()
//	fmt.Println("等等互斥锁....")
//	time.Sleep(1 * time.Second)
//	lock.Unlock()
//}
//
//func readLockFunc(lock *sync.RWMutex) {
//	lock.RLock() // 读锁在读取的时候不会阻塞其它锁，但是会排斥写锁，写锁相反
//	fmt.Println("等等读写互斥锁....")
//	time.Sleep(1 * time.Second)
//	lock.RUnlock()
//}
