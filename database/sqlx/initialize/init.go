package initialize

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func InitDB() (err error) {
	dns := "root:root@tcp(127.0.0.1:3306)/golang_code?charset=utf8mb4&parseTime=True"
	DB, err = sqlx.Connect("mysql", dns)
	if err != nil {
		fmt.Printf("connect DB failed, err: %v\n", err)
		return
	}

	DB.SetMaxOpenConns(20)
	DB.SetMaxIdleConns(10)
	return
}
