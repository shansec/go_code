package main

import (
	"errors"
	"fmt"
)

func add1(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

func calc2(x, y int, op func(int, int) int) int {
	return op(x, y)
}

func do(s string) (func(int, int) int, error) {
	switch s {
	case "+":
		return add1, nil
	case "-":
		return sub, nil
	default:
		err := errors.New("无法识别的操作符")
		return nil, err
	}
}

func main() {
	result := calc2(20, 30, add1)
	fmt.Println(result)

	// 匿名函数
	add := func(x, y int) {
		fmt.Println(x + y)
	}
	add(50, 90)

	// defer 执行时机
	// 在Go语言的函数中return语句在底层并不是原子操作，它分为给返回值赋值和RET指令两步。
	//	而defer语句执行的时机就在返回值赋值操作后，RET指令执行前。
}
