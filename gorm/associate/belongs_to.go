package associate

import (
	"fmt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name      string
	Age       int
	Hobby     []string
	CompanyID uint
	// 重写外键
	// Company Company `gorm:"foreignKey:CompanyRefer"`
	// 重写引用
	Company Company `gorm:"references:Code"`
}

type Company struct {
	ID   uint
	Code string
	Name string
}

func BelongsTo() {
	//global.DB.AutoMigrate(&User{})
	// 创建
	//company := Company{
	//	Code: "001",
	//	Name: "三只狼",
	//}
	//user := User{
	//	Name:    "苏轼",
	//	Age:     101,
	//	hobby:   []string{"写诗"},
	//	Company: company,
	//}
	//global.DB.Debug().Create(&user)
	// 关联查找
	//var company Company
	var user User
	//user := User{Model: gorm.Model{ID: 1}}
	// 启用关联
	//global.EXEC.Model(&user).Association("Company").Find(&company)
	//global.EXEC.Preload("Company").First(&user)
	// 添加关联
	//global.EXEC.Model(&user).Association("Company").Append(&Company{Name: "诗词大会", Code: "003"})
	// 替换关联
	//global.EXEC.Model(&user).Association("Company").Replace(&Company{ID: 1, Code: "001", Name: "三只狼"})
	// 删除关联
	//global.EXEC.Model(&user).Association("Company").Delete(&Company{ID: 1, Code: "001", Name: "三只狼"})
	// 清楚关联
	//global.EXEC.Model(&user).Association("Company").Clear()
	fmt.Println(user)
}
