package main

import (
	"fmt"
	"strconv"
)

func main() {
	// string to int
	//var str string = "100"
	//i1, err := strconv.Atoi(str)
	//if err != nil {
	//	fmt.Println("can not convert string to int")
	//} else {
	//	fmt.Printf("type: %T, value: %#v\n", i1, i1)
	//}

	// int to string
	//var str int = 101
	//i2 := strconv.Itoa(str)
	//fmt.Printf("type: %T, value: %#v\n", i2, i2)

	// Parse*
	// ParseBool	string to bool
	// ParseInt		string to int 与 Atoi 的区别是 ParseInt 可以控制转换的进制
	// ParseFloat	string to float
	// ParseUint	string to uint

	// Format*
	// FormatBool bool to string
	// FormatInt(i int64, base int) bool to int 返回i的base进制的字符串表示。base 必须在2到36之间，结果中会使用小写字母’a’到’z’表示大于10的数字。
	// FormatFloat float to string
	// FormatUint uint to string

	// Append*
	var bytes []byte = []byte{'2'}
	var bool1 bool = true
	fmt.Printf("type: %T value: %v\n", strconv.AppendBool(bytes, bool1), strconv.AppendBool(bytes, bool1))
}
