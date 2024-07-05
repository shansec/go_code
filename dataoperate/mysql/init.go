package main

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func Init() {
	database, err := sqlx.Open("mysql", "root:root@tcp(127.0.0.1:3306)/gorm")
	if err != nil {
		fmt.Println("open mysql failed,", err)
		return
	}
	DB = database
}
