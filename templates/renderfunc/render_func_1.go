package renderfunc

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Name   string
	Gender string
	Age    int
}

func SayHello(w http.ResponseWriter, r *http.Request) {
	// 解析指定文件生成模板对象
	tmpl, err := template.ParseFiles("E:\\minor\\Go\\src\\golang_code\\web\\templates\\temfiles\\template_1.tmpl")
	if err != nil {
		fmt.Printf("create template failed err: %v\n", err)
		return
	}
	user := User{
		Name:   "小王子",
		Gender: "男",
		Age:    18,
	}
	// 利用给定数据渲染模板，并将结果写入w
	err = tmpl.Execute(w, user)
	if err != nil {
		fmt.Printf("render template failed err: %v\n", err)
		return
	}
}
