package bind

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type PostParams struct {
	Name string `json:"name" form:"name"`
	Age  int    `json:"age" form:"age" binding:"required,mustBig"`
	Sex  bool   `json:"sex" form:"sex"`
}

func mustBig(fl validator.FieldLevel) bool {
	if fl.Field().Interface().(int) < 18 {
		return false
	}
	return true
}

func Bind() {
	r := gin.Default()

	// shouldBind* 系列方法和 Bind* 系列方法区别
	//  Bind* 数据绑定失败，这些方法会自动终止请求处理流程，并向客户端返回一个错误响应
	//	ShouldBind* 方法在数据绑定失败时不会自动终止请求处理流程或向客户端返回错误。相反，它们仅返回一个错误值，允许开发者自行决定如何处理这个错误。
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("mustBig", mustBig)
	}
	r.POST("/testBind", func(ctx *gin.Context) {
		var param PostParams
		err := ctx.ShouldBindQuery(&param)
		if err != nil {
			ctx.JSON(200, gin.H{
				"code": "报错了",
				"msg":  gin.H{},
			})
			return
		} else {
			ctx.JSON(200, gin.H{
				"code": "请求到数据",
				"msg":  param,
			})
		}
	})
	r.Run(":8080")
}
