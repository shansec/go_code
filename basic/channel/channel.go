package main

import "fmt"

func main() {
	c1 := make(chan int, 5)
	//go func() {
	//	for i := 0; i < 10; i++ {
	//		c1 <- 1
	//	}
	//}()
	//for i := 0; i < 10; i++ {
	//	fmt.Println(<-c1)
	//}
	//var readChan <-chan int = c1
	//var writeChan chan<- int = c1
	//
	//writeChan <- 2
	//fmt.Println(<-readChan)
	c1 <- 1
	c1 <- 2
	c1 <- 3
	close(c1)
	c1 <- 4
	c1 <- 5

	fmt.Println(<-c1)
	fmt.Println(<-c1)
	fmt.Println(<-c1)
	fmt.Println(<-c1)
	fmt.Println(<-c1)
}
