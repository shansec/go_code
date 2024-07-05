package main

import (
	"gorm_study/associate"
	"gorm_study/gormconnect"
)

func main() {
	gormconnect.Connect()
	//create.Create()	// 创建
	//seek.Seek()		// 查询
	//update.Update()	// 更新
	//remove.Remove()		// 删除
	//native.Native() // 原生语句
	//associate.BelongsTo()
	//associate.HasOne()
	//associate.HasMany()
	associate.ManyToMany()
}
