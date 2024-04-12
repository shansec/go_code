package assertion

//import "fmt"
//
//type User struct {
//	Name string
//	Age  int
//	Sex  bool
//}
//
//type Student struct {
//	User
//}
//
//func (u User) myName() {
//	fmt.Println("我的名字是", u.Name)
//}
//
//func main() {
//	u := User{
//		Name: "may",
//		Age:  23,
//		Sex:  false,
//	}
//	check(u)
//}
//
//func check(v interface{}) {
//	// 断言
//	//v.(User).myName()
//	switch v.(type) {
//	case User:
//		fmt.Println("我是 User")
//	case Student:
//		fmt.Println("我是 Student")
//	}
//}
