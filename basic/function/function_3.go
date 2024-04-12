package main

import (
	"fmt"
	"strings"
)

// 闭包
func returnNum() func() (int, int) {
	return func() (int, int) {
		return 0, 1
	}
}

func makeSuffix(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func main() {
	q := returnNum()
	a, b := q()
	fmt.Println(a, b)

	f2 := makeSuffix(".jpg")
	fmt.Println("文件名处理后=", f2("winter"))
	fmt.Println("文件名处理后=", f2("bird.jpg"))
}
