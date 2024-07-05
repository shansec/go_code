package update

import (
	"gorm.io/gorm"
	"gorm_study/global"
)

type User struct {
	gorm.Model
	Name string
	Age  int
	Sex  int
}

func Update() {
	// 保存所有字段，没有primary key 时，save会执行创建操作
	//var user User
	//user.ID = 13
	//user.Name = "李商隐"
	//user.Age = 105
	//user.Sex = 1
	// 更新单个列
	//global.DB.Debug().Model(&user).Update("age", 100)
	// 更新多个列
	//global.DB.Debug().Model(&User{}).Where("name LIKE ?", "%李白%").Updates(User{Sex: 1})
	//global.DB.Debug().Model(&User{}).Where("name LIKE ?", "%李白%").Updates(map[string]interface{}{"Age": 100, "Sex": 0})
	// 更新制定字段
	//global.DB.Debug().Model(&User{}).Select("Age").Where("name = ?", "李白2").Updates(User{Age: 30})
	// 批量更新
	global.DB.Debug().Model(&User{}).Where("id in ?", []int{10, 11, 12, 13}).Updates(map[string]interface{}{"Sex": 1})
}
