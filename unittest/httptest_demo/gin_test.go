package httptest_demo

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_helloHandler(t *testing.T) {
	tests := []struct {
		name  string
		param string
		want  string
	}{
		{"base case", `{"name": "五月"}`, "hello 五月"},
		{"bad case", "", "we need a name"},
	}

	r := setRouter()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// mock 一个 HTTP 请求
			req := httptest.NewRequest("POST", "/hello", strings.NewReader(tt.param))

			// mock 一个响应记录器
			w := httptest.NewRecorder()

			r.ServeHTTP(w, req)

			// 检验状态码是否符合预期
			assert.Equal(t, http.StatusOK, w.Code)

			// 解析并检测响应内容是否符合预期
			var resp map[string]string
			err := json.Unmarshal([]byte(w.Body.String()), &resp)
			assert.Nil(t, err)
			assert.Equal(t, tt.want, resp["msg"])
		})
	}
}
