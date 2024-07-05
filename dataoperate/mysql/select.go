package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Person struct {
	UserId   int    `db:"user_id"`
	UserName string `db:"username"`
	Sex      string `db:"sex"`
	Email    string `db:"email"`
}

type Place struct {
	Country string `db:"country"`
	City    string `db:"city"`
	Tel     int    `db:"tel"`
}

func Select() {
	var person []Person
	err := DB.Select(&person, "select user_id, username, sex, email from person")
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}

	if len(person) != 0 {
		for key, value := range person {
			fmt.Println(fmt.Sprintf("第 %d 位人员信息", key+1), value)
		}
	} else {
		fmt.Println("select is null")
	}
}
