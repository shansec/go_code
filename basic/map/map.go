package main

import "fmt"

func main() {
	scoreMap := make(map[string]int)
	// 添加
	scoreMap["张三"] = 100
	scoreMap["李四"] = 90
	scoreMap["王二"] = 80

	// 删除
	delete(scoreMap, "张三")

	for key, value := range scoreMap {
		fmt.Println(key, value)
	}
}
