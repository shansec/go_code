package basic

import "github.com/gin-gonic/gin"

func Basic() {
	r := gin.Default()
	r.GET("/path/:id", func(c *gin.Context) {
		id := c.Param("id")
		name := c.Query("name")
		age := c.DefaultQuery("age", "未知")
		c.JSON(200, gin.H{
			"id":   id,
			"name": name,
			"age":  age,
		})
	})
	r.POST("/path", func(c *gin.Context) {
		user := c.DefaultPostForm("user", "admin")
		pwd := c.PostForm("password")
		c.JSON(200, gin.H{
			"user":     user,
			"password": pwd,
		})
	})
	r.DELETE("/path/:id", func(c *gin.Context) {
		id := c.Param("id")
		name := c.PostForm("name")
		c.JSON(200, gin.H{
			"id":   id,
			"name": name,
		})
	})
	r.PUT("/path", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"success": true,
		})
	})
	r.Run(":1010")
}
