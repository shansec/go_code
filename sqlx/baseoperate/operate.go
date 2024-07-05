package baseoperate

import (
	"fmt"
	"sqlx_study/initialize"
	"sqlx_study/model"
)

// QueryRow 查询单行
func QueryRow() {
	sqlStr := "select id, name, age from user where id=?"
	var u model.User
	err := initialize.DB.Get(&u, sqlStr, 1)
	if err != nil {
		fmt.Printf("get failed, err: %v\n", err)
		return
	}
	fmt.Printf("id: %d, name: %s, age: %d\n", u.Id, u.Name, u.Age)
}

// QueryRows 查询多行
func QueryRows() {
	sqlStr := "select id, name, age from user where id > ?"
	var users []model.User
	err := initialize.DB.Select(&users, sqlStr, 0)
	if err != nil {
		fmt.Printf("select failed err: %v\n", err)
		return
	}
	for i, u := range users {
		fmt.Printf("第%d位信息：id: %d, name: %s, age: %d\n", i+1, u.Id, u.Name, u.Age)
	}
}

// InsertData 插入数据
func InsertData() {
	sqlStr := "insert into user(name, age) values (?, ?)"
	row, err := initialize.DB.Exec(sqlStr, "周仓", 101)
	if err != nil {
		fmt.Printf("insert failed err: %v\n", err)
		return
	}
	lastId, err := row.LastInsertId()
	if err != nil {
		fmt.Printf("get lastId failed err: %v\n", err)
		return
	}
	fmt.Printf("insert success, the lastId is %d\n", lastId)
}

// UpdateData 更新数据
func UpdateData() {
	sqlStr := "update user set age = ? where id = ?"
	row, err := initialize.DB.Exec(sqlStr, 99, 6)
	if err != nil {
		fmt.Printf("update data failed err: %v\n", err)
		return
	}
	affected, err := row.RowsAffected()
	if err != nil {
		fmt.Printf("get affected failed err: %v\n", err)
		return
	}
	fmt.Printf("update success, the affected is %d\n", affected)
}

// DeleteData 删除数据
func DeleteData() {
	sqlStr := "delete from user  where id = ?"
	row, err := initialize.DB.Exec(sqlStr, 6)
	if err != nil {
		fmt.Printf("delete data failed err: %v\n", err)
		return
	}
	affected, err := row.RowsAffected()
	if err != nil {
		fmt.Printf("get affected failed err: %v\n", err)
		return
	}
	fmt.Printf("delete success, the affected is %d\n", affected)
}
