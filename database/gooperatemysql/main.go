package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

type user struct {
	id   int
	age  int
	name string
}

func initDB() (err error) {
	dns := "root:root@tcp(127.0.0.1:3306)/golang_code?charset=utf8mb4&parseTime=True"
	db, err = sql.Open("mysql", dns)
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
	deleteRow()
}

// queryRow 查询单行
func queryRow() {
	sqlStr := "select id, name, age from user where id = ?"
	var u user
	// 非常重要：确保QueryRow之后调用Scan方法，否则持有的数据库链接不会被释放
	err := db.QueryRow(sqlStr, 1).Scan(&u.id, &u.name, &u.age)
	if err != nil {
		fmt.Printf("scan failed: %v\n", err)
		return
	}
	fmt.Printf("id: %d, name: %s, age: %d", u.id, u.name, u.age)
}

// queryRows 查询多行
func queryRows() {
	sqlStr := "select id, name, age from user where id > ?"
	rows, err := db.Query(sqlStr, 1)
	if err != nil {
		fmt.Printf("query failed err: %v\n", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var u user
		err := rows.Scan(&u.id, &u.name, &u.age)
		if err != nil {
			fmt.Printf("Scan failed err: %v\n", err)
			return
		}
		fmt.Printf("id: %d, name: %s, age: %d\n", u.id, u.name, u.age)
	}
}

// insertRow 插入数据
func insertRow() {
	sqlStr := "insert into user(name, age) values (?, ?)"
	exec, err := db.Exec(sqlStr, "马超", 100)
	if err != nil {
		fmt.Printf("insert failed err: %v\n", err)
		return
	}
	// 新插入数据的 id
	id, err := exec.LastInsertId()
	if err != nil {
		fmt.Printf("get lastID failed err: %v\n", err)
		return
	}
	fmt.Printf("insert data sucess, lastID is %d\n", id)
}

// updateRow 更新数据
func updateRow() {
	sqlStr := "update user set age = ? where id = ?"
	exec, err := db.Exec(sqlStr, 99, 2)
	if err != nil {
		fmt.Printf("update failed err: %v\n", err)
		return
	}
	affected, err := exec.RowsAffected()
	if err != nil {
		fmt.Printf("get affected row failed err: %v\n", err)
		return
	}
	fmt.Printf("update data sucess, affected row  %d\n", affected)
}

// deleteRow 删除数据
func deleteRow() {
	sqlStr := "delete from user where id = ?"
	ret, err := db.Exec(sqlStr, 3)
	if err != nil {
		fmt.Printf("delete failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("delete success, affected rows:%d\n", n)
}
