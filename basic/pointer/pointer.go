package main

import "fmt"

func main() {
	// 指针地址
	// 指针类型
	// 指针取值 &（取地址） *（地址取值）
	pointerA := 100
	pointerB := &pointerA
	fmt.Printf("pointerA:%d pointer:%p\n", pointerA, &pointerA)
	fmt.Printf("pointerB:%p type:%T\n", pointerB, pointerB)

	pointerValueA := 10
	pointerValueB := &pointerValueA
	fmt.Printf("type of pointerValueB:%T\n", pointerB)
	pointerValueC := *pointerValueB
	fmt.Printf("type of c:%T\n", pointerValueC)
	fmt.Printf("value of c:%v\n", pointerValueC)

	// 指向指针的指针
	var a int
	var ptr *int
	var pptr **int

	ptr = &a
	pptr = &ptr
	fmt.Printf("变量 a = %d\n", a)
	fmt.Printf("指针ptr= %v\n", ptr)
	fmt.Printf("取指针变量ptr的值 *ptr= %d\n", *ptr)
	fmt.Printf("指向指针的指针变量 **pptr = %d\n", **pptr)
	fmt.Printf("变量的内存地址 a = %v\n", ptr)
	fmt.Printf("指针ptr取的内存地址为 &ptr = %v\n", &ptr)
	fmt.Printf("第一个指针变量的内存地址 *ptr = %v\n", pptr)

	// make 和 new 的区别
	//	二者都是用来申请内存的。
	//	new很少用,一般用来给基本类型申请内存的。并且内存对应的值为类型零值，返回的是指向类型的指针。比如string int返回的是对应类型的指针（*string *int）
	//	make只用于slice、map以及channel的初始化，返回的还是这三个引用类型本身；
}
