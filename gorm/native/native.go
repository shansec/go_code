package native

import (
	"fmt"
	"github/May-cloud/study/gorm/global"
)

type User struct {
	Name string
	Age  int
	Sex  int
}

func Native() {
	var users []User
	//global.DB.Debug().Raw("Select name, age, sex from may_users WHERE id in ?", []int{11, 12, 13}).Scan(&users)
	global.DB.Debug().Exec("drop table may_users")
	fmt.Println(users)
}
