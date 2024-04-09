package model

import "database/sql/driver"

type User struct {
	Id   int
	Name string
	Age  int
}

type UserIn struct {
	Name string `db:"name"`
	Age  int    `db:"age"`
}

func (u UserIn) Value() (driver.Value, error) {
	return []interface{}{u.Name, u.Age}, nil
}
