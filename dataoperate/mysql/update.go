package main

import "fmt"

func Update() {
	res, err := DB.Exec("update person set username=? where user_id=?", "may", 1)
	if err != nil {
		fmt.Println("exec failed", err)
		return
	}
	row, err := res.RowsAffected()
	if err != nil {
		fmt.Println("rows failed", err)
	}
	fmt.Println("update success", row)
}
