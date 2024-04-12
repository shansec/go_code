package sync

//func main() {
//	wg := &sync.WaitGroup{}
//	wg.Add(2)
//	go func() {
//		time.Sleep(3 * time.Second)
//		wg.Done()
//		fmt.Println("打掉血量减一")
//	}()
//
//	go func() {
//		time.Sleep(net * time.Second)
//		wg.Done()
//		fmt.Println("打掉血量减一")
//	}()
//
//	wg.Wait()
//	fmt.Println("血量为零，GAME OVER")
//}
