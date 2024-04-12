package main

import "fmt"

func main() {
	// 算术运算符
	// 关系运算符
	// 逻辑运算符
	// 位运算符
	//	A	B	A&B	A|B	A^B
	//	0	0	0	0	0
	//	0	1	0	1	1
	//	1	1	1	1	0
	//	1	0	0	1	1
	A := 60
	B := 12
	fmt.Println(A | B)
	fmt.Println(A & B)
	fmt.Println(A ^ B)
	// 赋值运算符
	// 运算符优先级
	//	优先级	运算符
	//	7	~ ! ++ --
	//	net	* / % << >> & &^
	//	5	+ - ^
	//	4	== != < <= >= >
	//	3	<-
	//	2	&&
	//	1	||
}
