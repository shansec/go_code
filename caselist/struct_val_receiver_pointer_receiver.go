package main

import "fmt"

type student struct {
	name string
	age  uint
}

func (s student) setName(name string) {
	s.name = name
}

func (s *student) setAge(age uint) {
	s.age = age
}

func main() {
	stu := student{name: "赵云", age: 101}
	fmt.Printf("student modification before: %+v\n", stu)
	stu.setName("张飞")
	fmt.Printf("student modification after(val): %+v\n", stu)
	stu.setAge(102)
	fmt.Printf("student modification after(pointer): %+v\n", stu)
}
