package main

import (
	"fmt"
	"math"
)

func main() {
	// 十进制
	var a int = 10
	fmt.Printf("%d \n", a)
	fmt.Printf("%b \n", a)

	// 八进制
	var b int = 077
	fmt.Printf("%o \n", b)

	// 十六进制
	var c int = 0xff
	fmt.Printf("%x \n", c)
	fmt.Printf("%X \n", c)

	// 浮点型
	fmt.Printf("%f\n", math.Pi)
	fmt.Printf("%.2f\n", math.Pi)

	// 字符串
	str := "hello"
	str1 := "self"
	fmt.Print(str, str1)

	str2 := `第一行
第二行
第三行`
	fmt.Println(str2)
	fmt.Println(len(str2))
	const name, age = "may", 23
	fmt.Println(fmt.Sprintf("%s is %d years old.\n", name, age))

	// byte类型 和 rune类型
	s := "helloShansec"
	for i := 0; i < len(s); i++ { // byte
		fmt.Printf("%v(%c)", s[i], s[i])
	}
}
