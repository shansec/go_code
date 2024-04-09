package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

type user struct {
	Id   int
	Name string
	Age  int
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("connect DB failed, err: %v\n", err)
		return
	}
	queryRows()
}

func initDB() (err error) {
	dns := "root:root@tcp(127.0.0.1:3306)/golang_code?charset=utf8mb4&parseTime=True"
	db, err = sqlx.Connect("mysql", dns)
	if err != nil {
		fmt.Printf("connect DB failed, err: %v\n", err)
		return
	}

	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(10)
	return
}

// queryRow 查询单行
func queryRow() {
	sqlStr := "select id, name, age from user where id=?"
	var u user
	err := db.Get(&u, sqlStr, 1)
	if err != nil {
		fmt.Printf("get failed, err: %v\n", err)
		return
	}
	fmt.Printf("id: %d, name: %s, age: %d\n", u.Id, u.Name, u.Age)
}

func queryRows() {
	sqlStr := "select id, name, age from user where id > ?"
	var users []user
	err := db.Select(&users, sqlStr, 0)
	if err != nil {
		fmt.Printf("select failed err: %v\n", err)
		return
	}
	for i, u := range users {
		fmt.Printf("第%d位信息：id: %d, name: %s, age: %d\n", i+1, u.Id, u.Name, u.Age)
	}
}
