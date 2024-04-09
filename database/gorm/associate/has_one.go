package associate

//type People struct {
//	gorm.Model
//	CreditCard CreditCard
//}
//
//type CreditCard struct {
//	gorm.Model
//	Number   string
//	PeopleID uint
//}

func HasOne() {
	// 创建表
	//global.EXEC.AutoMigrate(&CreditCard{}, &People{})
	//card := CreditCard{Number: "6259650871772098"}
	//people := People{CreditCard: card}
	// 增加
	//global.EXEC.Create(&people)
	//var people People
	//var card CreditCard
	// 查找
	//global.EXEC.Model(&People{Model: gorm.Model{ID: 1}}).Preload("CreditCard", func(db *gorm.DB) *gorm.DB {
	//	return db.Select("number")
	//}).First(&people)
	//global.EXEC.Model(&people).Association("CreditCard").Find(&card)
	// 添加
	//global.EXEC.Model(&people).Association("CreditCard").Append(&CreditCard{Number: "6259650871771111"})
	// 替换
	//global.EXEC.Model(&people).Association("CreditCard").Replace(&CreditCard{Model: gorm.Model{
	//	ID: 2,
	//}})
	// 删除
	//global.EXEC.Model(&people).Association("CreditCard").Delete(&CreditCard{Model: gorm.Model{
	//	ID: 1,
	//}})
	// 清除
	//global.EXEC.Model(&people).Association("CreditCard").Clear()
	//fmt.Println(people)
}
