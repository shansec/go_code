package belongsto

import (
	"gorm_study/global"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string
}

type Profile struct {
	gorm.Model
	UserID int
	User   User
	Name   string
}

func BelongToCreate() {
	global.DB.AutoMigrate(&User{}, &Profile{})
}

func BelongToOperate() {
	profile := Profile{
		User: User{Name: "张三"},
		Name: "设置",
	}
	global.EXEC.Create(&profile)
}
