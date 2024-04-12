package _struct

import "fmt"

type May struct {
	name  string
	age   int
	sex   bool
	hobby []string
}

func (may *May) Song() {
	fmt.Printf("%v真帅！！！", may.name)
}

//func main() {
//var may May
//may.name = "may"
//may.age = 12
//may.hobby = []string{"读书"}
//may.Song()
//fmt.Println(may)

//may := May{
//	name:  "may",
//	age:   0,
//	sex:   false,
//	hobby: []string{"读书"},
//}
//fmt.Println(may)

//var may = new(May)
//may.name = "may"
//may.age = 12
//may.hobby = []string{"读书"}
//fmt.Println(may)

//var may *May
//mayP := &may
//may.name = "may"
//(*mayP).name = "true"
//
//fmt.Println(mayP)
//}
