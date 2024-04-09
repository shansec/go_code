package associate

import (
	"gorm.io/gorm"
)

type Pen struct {
	gorm.Model
	Written   string
	PenHolder string
	Refill    []Refill
}

type Refill struct {
	gorm.Model
	Color    string
	Material string
	PenID    uint
}

func HasMany() {
	// 创建表
	//global.EXEC.AutoMigrate(&Pen{}, &Refill{})

}
