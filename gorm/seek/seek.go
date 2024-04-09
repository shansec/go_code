package seek

import (
	"fmt"
	"gorm.io/gorm"
	"gorm_study/global"
)

type User struct {
	gorm.Model
	Name string
	Age  int
	Sex  int
}

func Seek() {
	var user User
	//var users []User
	// 获取第一条数据
	//global.DB.First(&user)
	// 获取第二条数据
	//global.DB.Last(&user)
	// 根据主键检索
	//global.DB.First(&user, 10)
	//global.DB.Find(&users, []int{10, 11, 12})
	// where 条件
	//global.DB.Where("name = ?", "杜甫").First(&user)
	//global.DB.Where("name IN ?", []string{"李白", "杜甫"}).Find(&users)
	//global.DB.Where(&User{Name: "杜甫"}).First(&user)
	//global.DB.Where(map[string]interface{}{"Age": 100}).Find(&users)
	// 指定结构体查询字段
	//global.DB.Where(&User{Name: "杜甫", Age: 101}, "name", "age").Find(&users)
	// 内联条件
	//global.DB.Find(&users, "name LIKE ?", "%李白%")
	//global.DB.Find(&users, User{Age: 100})
	// Or 条件
	//global.DB.Where("name = ?", "杜甫").Or("name = ?", "王维").Find(&users)
	// 选择特定字段
	global.DB.Debug().Select("name").Where("name = ?", "杜甫").First(&user)
	fmt.Println(user)
}
