package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func Insert() {
	r, err := DB.Exec("insert into person(username, sex, email)values(?, ?, ?)", "stu002", "man", "stu02@qq.com")
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
	id, err := r.LastInsertId()
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}

	fmt.Println("insert success:", id)
}
