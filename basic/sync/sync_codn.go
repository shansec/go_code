package sync

//func main() {
//	co := sync.NewCond(&sync.Mutex{})
//
//	go func() {
//		co.L.Lock()
//		fmt.Println("lock1")
//		co.Wait()
//		co.L.Unlock()
//		fmt.Println("unlock1")
//	}()
//
//	go func() {
//		co.L.Lock()
//		fmt.Println("lock2")
//		co.Wait()
//		co.L.Unlock()
//		fmt.Println("unlock2")
//	}()
//
//	time.Sleep(3 * time.Second)
//	co.Broadcast()
//	time.Sleep(2 * time.Second)
//}
