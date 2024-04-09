package main

import "fmt"

func main() {
	var c calculation
	c = add
	fmt.Println(c(2, 5))
}

// 函数的参数说明
func intSum(x, y int) int {
	return x + y
}

// 可变参数
//
//	可变参数只能作为函数参数的最后一个
func intSum2(x ...int) int {
	fmt.Println(x)
	sum := 0
	for _, value := range x {
		sum = sum + value
	}
	return sum
}

// 多返回值
func calc(x, y int) (int, int) {
	sum := x + y
	sub := x - y
	return sum, sub
}

// 返回值命名
func calc1(x, y int) (sum, sub int) {
	sum = x + y
	sub = x - y
	return
}

// 函数类型与变量
type calculation func(int, int) int

func add(x, y int) int {
	return x + y
}
