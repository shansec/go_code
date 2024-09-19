package unittesthttp

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ReqParam struct {
	X int `json:"x"`
}

func UnitTestHttp() {
	r := gin.Default()

	r.POST("/post", func(c *gin.Context) {
		var req ReqParam
		if err := c.ShouldBindJSON(&req); err != nil {
			fmt.Printf("接受参数失败，err: %v\n", err)
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"x": req.X + 10,
		})
	})
	r.Run(":8080")
}
