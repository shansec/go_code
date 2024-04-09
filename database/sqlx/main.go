package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jmoiron/sqlx"
	"sqlx_study/initialize"
	"sqlx_study/model"
	"sqlx_study/sqlxin"
)

func main() {
	err := initialize.InitDB()
	if err != nil {
		fmt.Printf("connect DB failed, err: %v\n", err)
		return
	}
	// baseoperate.DeleteData()
	// nameexec.InsertByNameExec()
	// nameexec.UpdateByNameExec()
	// namequery.NamedQuery()
	users := make([]interface{}, 3)
	user1 := model.UserIn{
		Name: "赵云",
		Age:  103,
	}
	users[0] = user1
	user2 := model.UserIn{
		Name: "马超",
		Age:  104,
	}
	users[1] = user2
	user3 := model.UserIn{
		Name: "诸葛亮",
		Age:  105,
	}
	users[2] = user3
	//sqlxin.BatchInsertUsersTrouble(users)
	//err = sqlxin.BatchInsertUsersBySqlxIn(users)
	sqlxin.QueryByIds([]int{2, 4, 5})
}
