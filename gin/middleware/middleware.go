package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		fmt.Println("中间件开始执行了")
		c.Set("request", "中间件")
		c.Next()
		status := c.Writer.Status()
		fmt.Println("中间件执行完毕", status)
		t2 := time.Since(t)
		fmt.Println("time", t2)
	}
}

func Enter() {
	r := gin.Default()

	r.Use(middleware())

	{
		r.GET("/middleware", func(ctx *gin.Context) {
			req, _ := ctx.Get("request")
			fmt.Println("request", req)
			ctx.JSON(200, gin.H{"request": req})
		})
	}
	r.Run(":8080")
}
