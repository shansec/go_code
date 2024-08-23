package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type addParam struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type addResult struct {
	Code int `json:"code"`
	Data int `json:"data"`
}

func add(x, y int) int {
	return x + y
}

func AddHandler(w http.ResponseWriter, r *http.Request) {
	// 解析参数
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("解析参数失败")
		return
	}
	var param addParam
	json.Unmarshal(b, &param)
	ret := add(param.X, param.Y)
	respBytes, _ := json.Marshal(addResult{Code: 0, Data: ret})
	w.Write(respBytes)
}

func main() {
	http.HandleFunc("/add", AddHandler)
	log.Fatal(http.ListenAndServe(":9090", nil))
}
