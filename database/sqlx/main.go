package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jmoiron/sqlx"
	"sqlx_study/initialize"
	"sqlx_study/namequery"
)

func main() {
	err := initialize.InitDB()
	if err != nil {
		fmt.Printf("connect DB failed, err: %v\n", err)
		return
	}
	// baseoperate.DeleteData()
	// nameexec.InsertByNameExec()
	//nameexec.UpdateByNameExec()
	namequery.NamedQuery()
}
