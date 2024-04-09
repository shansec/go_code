package associate

import (
	"fmt"
	"gorm.io/gorm"
	"gorm_study/global"
)

type People struct {
	gorm.Model
	Name        string
	Age         uint
	Sex         uint
	CreditCards []CreditCard `gorm:"many2many:people_card;"`
}

type CreditCard struct {
	gorm.Model
	Number  string
	Color   string
	Peoples []People `gorm:"many2many:people_card;"`
}

func ManyToMany() {
	// 创建表
	//global.EXEC.AutoMigrate(&People{}, &CreditCard{})
	// 创建数据，以 people 为主
	//card1 := CreditCard{
	//	Number: "123456789",
	//	Color:  "red",
	//}
	//card2 := CreditCard{
	//	Number: "9876543121",
	//	Color:  "blue",
	//}
	//people := People{
	//	Name:        "李白",
	//	Age:         100,
	//	Sex:         0,
	//	CreditCards: []CreditCard{card1, card2},
	//}
	//global.EXEC.Model(&People{}).Create(&people)
	// 创建数据，以 card 为主
	//people1 := People{
	//	Name: "白居易",
	//	Age:  101,
	//	Sex:  0,
	//}
	//people2 := People{
	//	Name: "苏轼",
	//	Age:  102,
	//	Sex:  0,
	//}
	//card := CreditCard{
	//	Number:  "0000000000",
	//	Color:   "green",
	//	Peoples: []People{people1, people2},
	//}
	//global.EXEC.Model(&CreditCard{}).Create(&card)
	// 查找
	//var card4 []CreditCard
	//peo := People{Model: gorm.Model{
	//	ID: 1,
	//}}
	var people []People
	card := CreditCard{
		Model: gorm.Model{ID: 3},
	}
	//global.EXEC.Model(&peo).Preload("Peoples").Association("CreditCards").Find(&card4)
	global.EXEC.Model(&card).Preload("CreditCards").Where("Age > ?", 101).Association("Peoples").Find(&people)
	// 添加
	//cardd := CreditCard{Model: gorm.Model{
	//	ID: 7,
	//}}
	//card2 := CreditCard{
	//	Number: "98765431210000",
	//	Color:  "pink",
	//}
	//global.EXEC.Model(&peo).Association("CreditCards").Append(&card2)
	// 替换
	//global.EXEC.Model(&peo).Association("CreditCards").Replace(&cardd, &card2)
	// 删除
	//global.EXEC.Model(&peo).Association("CreditCards").Delete(&cardd)
	// 清除
	//global.EXEC.Model(&peo).Association("CreditCards").Clear()
	fmt.Println(people)
}
