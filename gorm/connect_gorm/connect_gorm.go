package connect_gorm

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"gorm_study/global"
)

func Connect() {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:               "root:root@tcp(127.0.0.1:3306)/gorm?charset=utf8&parseTime=True&loc=Local", // DSN data source name
		DefaultStringSize: 256,                                                                        // string 类型字段的默认长度
	}), &gorm.Config{
		SkipDefaultTransaction: false,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "s_",
			SingularTable: true,
		},
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		fmt.Println("数据库连接错误")
		return
	}
	global.DB = db
	global.EXEC = db.Debug()
}
