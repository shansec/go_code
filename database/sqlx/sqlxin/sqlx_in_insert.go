package sqlxin

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"sqlx_study/initialize"
	"sqlx_study/model"
	"strings"
)

// BatchInsertUsersTrouble 自行拼接完成插入
func BatchInsertUsersTrouble(users []*model.UserIn) {
	valueStrings := make([]string, 0, len(users))
	valueArgs := make([]interface{}, 0, len(users)*2)
	for _, u := range users {
		valueStrings = append(valueStrings, "(?, ?)")
		valueArgs = append(valueArgs, u.Name)
		valueArgs = append(valueArgs, u.Age)
	}

	stmt := fmt.Sprintf("insert into user (name, age) values %s",
		strings.Join(valueStrings, ","))

	exec, err := initialize.DB.Exec(stmt, valueArgs...)
	if err != nil {
		fmt.Printf("batchinsert failed err: %v\n", err)
		return
	}

	affected, err := exec.RowsAffected()
	if err != nil {
		fmt.Printf("get affected failed err: %v\n", err)
		return
	}

	fmt.Printf("batchinsert success, affected is %d\n", affected)
}

// BatchInsertUsersBySqlxIn 通过 sqlx.in 实现批量插入
func BatchInsertUsersBySqlxIn(users []interface{}) error {
	query, args, _ := sqlx.In(
		"INSERT INTO user (name, age) VALUES (?), (?), (?)",
		users..., // 如果 arg 实现了 driver.Valuer, sqlx.In 会通过调用 Value()来展开它
	)
	fmt.Println(query)
	fmt.Println(args)
	_, err := initialize.DB.Exec(query, args...)
	return err
}

// QueryByIds 通过 id 进行批量查询
func QueryByIds(ids []int) {
	query, args, err := sqlx.In("select name, age from user where id in (?)", ids)
	if err != nil {
		fmt.Printf("sqlx.in failed err: %v\n", err)
		return
	}
	fmt.Printf("query %s\n", query)
	fmt.Printf("args %v\n", args)
	var users []model.UserIn
	err = initialize.DB.Select(&users, query, args...)
	if err != nil {
		fmt.Printf("DB Select failed err: %v\n", err)
		return
	}
	for index, user := range users {
		fmt.Printf("第%d位历史人物信息：name=%s, age=%d\n", index+1, user.Name, user.Age)
	}
}
