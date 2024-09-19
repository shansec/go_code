package sql_demo

import "database/sql"

func recordStats(db *sql.DB, userID, productID int64) (err error) {
	// 开启事务
	// 操作 views 和 product_viewers 两张表
	tx, err := db.Begin()
	if err != nil {
		return
	}
	defer func() {
		switch err {
		case nil:
			err = tx.Commit()
		default:
			tx.Rollback()
		}
	}()

	if _, err := tx.Exec("UPDATE products SET views = views + 1"); err != nil {
		return err
	}

	_, err = tx.Exec("INSERT INTO product_viewers (user_id, product_id) VALUES (?, ?)", userID, productID)
	if err != nil {
		return
	}
	return
}
