package main

import "fmt"

func main() {
	// 数组属性：
	//	固定长度，数组声明后长度便不能修改
	//	只能存储一种特定类型元素的序列
	// 声明方式：
	//	直接声明 var arr [3]int
	//	make	arr := make([]int, 3)
	//	字面量	arr := [3]int{1, 2, 3}
	// 	自动识别长度	arr  := [...]int{1, 2, 3}
	//	new 	arr := new([10]int)
	var testArray [3]int                        //数组会初始化为int类型的零值
	var numArray = [3]int{1, 2}                 //使用指定的初始值完成初始化
	var cityArray = [3]string{"北京", "上海", "深圳"} //使用指定的初始值完成初始化
	fmt.Println(testArray)                      //[0 0 0]
	fmt.Println(numArray)                       //[1 2 0]
	fmt.Println(cityArray)                      //[北京 上海 深圳]
	// 以使用指定索引值的方式来初始化数组
	a := [...]int{1: 1, 3: 5}
	fmt.Println(a)                  // [0 1 0 5]
	fmt.Printf("type of a:%T\n", a) //type of a:[4]int
}
