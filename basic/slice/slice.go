package main

import "fmt"

func main() {
	//type slice struct {
	//	array unsafe.Pointer
	//	len   int
	//	cap   int
	//}

	// 声明切片类型
	// var name []T
	var a []string
	var b = []int{}
	var c = []bool{true, false}
	var d = []bool{false, true}

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	arr := [...]int{55, 56, 57, 58, 59}
	arrB := arr[1:4]
	fmt.Println(arrB)
	fmt.Printf("type of b:%T\n", b)

	// 切片再切片
	sliceA := [...]string{"北京", "上海", "杭州", "河南", "成都", "湖南"}
	fmt.Printf("a:%v type:%T len:%d cap:%d\n", sliceA, sliceA, len(sliceA), cap(sliceA))
	sliceB := sliceA[1:3]
	fmt.Printf("a:%v type:%T len:%d cap:%d\n", sliceB, sliceB, len(sliceB), cap(sliceB))
	sliceC := sliceB[1:5]
	fmt.Printf("a:%v type:%T len:%d cap:%d\n", sliceC, sliceC, len(sliceC), cap(sliceC))

	// 使用make构造切片
	sliceMake := make([]int, 2, 10)
	fmt.Printf("a:%v type:%T len:%d cap:%d\n", sliceMake, sliceMake, len(sliceMake), cap(sliceMake))

	// 切片的赋值拷贝
	//  切片的赋值拷贝共享底层数组，对一个切片的修改会影响另一个切片的内容
	sliceCopy := make([]int, 3)
	sliceCopy2 := sliceCopy
	sliceCopy[0] = 100
	fmt.Println(sliceCopy)
	fmt.Println(sliceCopy2)

	// append
	var numSlice []int
	for i := 0; i < 10; i++ {
		numSlice = append(numSlice, i)
		fmt.Printf("%v len:%d cap:%d ptr:%p\n", numSlice, len(numSlice), cap(numSlice), numSlice)
	}

	// 删除元素
	// 要从切片a中删除索引为index的元素，操作方法是a = append(a[:index], a[index+1:]...)
}
