package main

import (
	"fmt"
	"os"
)

func main() {
	// 在函数内部，可以使用更简略的 := 方式声明并初始化变量。
	if len(os.Args) != 0 {
		fmt.Println(os.Args[0])
	}
	fmt.Println(os.Args[1])
}
