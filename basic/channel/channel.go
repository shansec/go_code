package main

import "fmt"

func recv(c chan int) {
	re := <-c
	fmt.Println("接收成功:", re)
}

func f2(ch chan int) {
	// for {
	// 	value, ok := <-ch
	// 	if !ok {
	// 		fmt.Println("通道已关闭")
	// 		break
	// 	}
	// 	fmt.Printf("v:%#v ok:%#v\n", value, ok)
	// }

	for v := range ch {
		fmt.Printf("v:%#v\n", v)
	}
}

func provider() chan int {
	ch := make(chan int, 2)
	go func() {
		for i := 0; i < 10; i++ {
			if i%2 == 1 {
				ch <- i
			}
		}
		close(ch)
	}()
	return ch
}

func customer(ch chan int) int {
	sum := 0
	for value := range ch {
		sum += value
	}
	return sum
}

func main() {
	// ch := make(chan int)
	// go recv(ch)
	// ch <- 10
	// fmt.Println("发送成功")

	// ch := make(chan int, 2)
	// ch <- 10
	// ch <- 20
	// close(ch)
	// f2(ch)

	// ch := provider()
	// res := customer(ch)
	// fmt.Println("总和：", res)

	ch := make(chan int, 1)
	for i := 1; i <= 10; i++ {
		select {
		case x := <-ch:
			fmt.Println(x)
		case ch <- i:
		}
	}
}
