package main

//import "fmt"
//
//func main() {
//	c := make(chan int)
//	var readc <-chan int = c
//	var writec chan<- int = c
//
//	go SetChan(writec)
//
//	getChane(readc)
//}
//
//func SetChan(write chan<- int) {
//	for i := 1; i < 10; i++ {
//		fmt.Printf("设置值：%d\n", i)
//		write <- i
//	}
//}
//
//func getChane(read <-chan int) {
//	for i := 1; i < 10; i++ {
//		fmt.Printf("从 getChan 函数中取出值：%d\n", <-read)
//	}
//}
