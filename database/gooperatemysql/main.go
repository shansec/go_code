package main

import (
	"database/sql"
	"fmt"
)

var db *sql.DB

func initDB() (err error) {
	dns := "root:root@tcp(127.0.0.1:3306)/golang_code"
	db, err := sql.Open("mysql", dns)
	if err != nil {
		return err
	}
	// 尝试与数据库建立连接
	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}

func main() {
	// 初始化数据库
	err := initDB()
	if err != nil {
		fmt.Printf("数据库连接失败: %v", err)
		return
	}
}
