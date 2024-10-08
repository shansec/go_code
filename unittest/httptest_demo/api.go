package httptest_demo

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type ReqParam struct {
	X int `json:"x"`
}

type Result struct {
	Value int `json:"value"`
}

func GetResultByAPI(x, y int) int {
	p := &ReqParam{X: x}
	b, _ := json.Marshal(p)

	// 调用其它服务的 API
	resp, err := http.Post("http://localhost:8080/post", "application/json", bytes.NewBuffer(b))
	if err != nil {
		return -1
	}
	body, _ := ioutil.ReadAll(resp.Body)
	var ret Result
	if err := json.Unmarshal(body, &ret); err != nil {
		return -1
	}

	return ret.Value + y
}
