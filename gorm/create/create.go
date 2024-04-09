package create

import (
	"github/May-cloud/study/gorm/global"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string
	Age  int
	Sex  int
}

func Create() {
	// 创建数据库表
	global.DB.AutoMigrate(&User{})
	//user := User{Name: "李白", Age: 100, Sex: 0}
	//users := [...]User{
	//	User{Name: "杜甫", Age: 100, Sex: 0},
	//	User{Name: "王维", Age: 100, Sex: 0},
	//	User{Name: "白居易", Age: 100, Sex: 0},
	//}
	// 制定要创建的字段
	//global.DB.Select("Name", "Age").Create(&user)
	// 忽略要创建的字段
	//global.DB.Omit("Name").Create(&user)
	// map 方式创建
	//global.DB.Model(&User{}).Create(map[string]interface{}{"Name": "李白", "Age": 100, "Sex": 0})
	global.DB.Model(&User{}).Create([]map[string]interface{}{
		{"Name": "杜甫", "Age": 101, "Sex": 0},
		{"Name": "王维", "Age": 102, "Sex": 0},
		{"Name": "白居易", "Age": 103, "Sex": 0},
	})
}
