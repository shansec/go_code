package main

//
//import "fmt"
//
//type Animal interface {
//	Eat()
//	Run()
//}
//
//type Cat struct {
//	Name string
//	Sex  bool
//}
//
//func (c *Cat) Eat() {
//	fmt.Println(c.Name, "开始吃")
//}
//
//func (c *Cat) Run() {
//	fmt.Println(c.Name, "开始跑")
//}
//
//type Dog struct {
//	Name string
//}
//
//func (d *Dog) Eat() {
//	fmt.Println(d.Name, "开始吃")
//}
//
//func (d *Dog) Run() {
//	fmt.Println(d.Name, "开始跑")
//}
//
//func main() {
//	var a Animal
//
//	c := Cat{
//		Name: "Tom",
//		Sex:  false,
//	}
//
//	a = &c
//	a.Run()
//	a.Eat()
//}
//
//type ordered interface {
//	int | bool
//}
//
//func MyFunc[T ordered](a T) {
//	fmt.Println(a)
//}
