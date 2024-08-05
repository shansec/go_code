package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, err := http.Get("https://www.liwenzhou.com/")
	if err != nil {
		fmt.Printf("get failed, err: %v\n", err)
		return
	}
	defer resp.Body.Close() // 关闭回复的主题
	all, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("read failed, err: %v\n", err)
		return
	}
	fmt.Print(string(all))
}
