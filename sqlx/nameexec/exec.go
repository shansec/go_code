package nameexec

import (
	"fmt"
	"sqlx_study/initialize"
)

// InsertByNameExec 通过 NameExec 实现增加
func InsertByNameExec() {
	sqlStr := "insert into user (name, age) VALUES (:name, :age)"
	_, err := initialize.DB.NamedExec(sqlStr,
		map[string]interface{}{
			"name": "曹操",
			"age":  98,
		},
	)
	if err != nil {
		fmt.Printf("NameExec failed err: %v\n", err)
		return
	}
	return
}

// UpdateByNameExec 通过 NameExec 实现更新
func UpdateByNameExec() {
	sqlStr := "update user set age = :age where id = :id"
	_, err := initialize.DB.NamedExec(sqlStr, map[string]interface{}{"age": 999, "id": 7})
	if err != nil {
		fmt.Printf("update data failed err: %v\n", err)
		return
	}
	return
}
