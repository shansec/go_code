package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string
	Age  int
	Sex  bool
}

type Student struct {
	Class string
	User  User
}

func (u User) myName() {
	fmt.Println("我的名字是", u.Name)
}

func main() {
	u := User{
		Name: "may",
		Age:  23,
		Sex:  true,
	}
	stu := Student{
		Class: "三年级",
		User:  u,
	}
	check(&stu)
}

func check(inter interface{}) {
	t := reflect.TypeOf(inter)
	v := reflect.ValueOf(inter)
	for i := 0; i < t.NumField(); i++ {
		fmt.Println(v.Field(i))
	}
	//fmt.Println(v.FieldByName("Class"))
	//fmt.Println(v.FieldByName("User"))
	//fmt.Println(v.FieldByIndex([]int{1, 1}))
	v.FieldByName("Class").SetString("四年级三班")
	for i := 0; i < t.NumField(); i++ {
		fmt.Println(v.Field(i))
	}
}
