package group

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func middleware1() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("我再方法前1")
		c.Next()
		fmt.Println("我再方法后1")
	}
}

func middleware2() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("我再方法前2")
		c.Next()
		fmt.Println("我再方法后2")
	}
}

func Group() {
	r := gin.Default()
	v1 := r.Group("v1").Use(middleware1()).Use(middleware2())
	v1.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"code": 200,
			"msg":  "请求到数据了",
		})
	})
	r.Run(":8080")
}
