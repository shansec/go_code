package templates

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

func InitTemplate() {
	r := gin.Default()
	//r.LoadHTMLGlob("templates/**/*")

	r.SetFuncMap(template.FuncMap{
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
	})

	r.LoadHTMLFiles("./index.tmpl")

	//r.GET("/posts/index", func(c *gin.Context) {
	//	c.HTML(http.StatusOK, "posts/index.html", gin.H{
	//		"title": "posts/index",
	//	})
	//})
	//
	//r.GET("/users/index", func(c *gin.Context) {
	//	c.HTML(http.StatusOK, "users/index.html", gin.H{
	//		"title": "users/index",
	//	})
	//})

	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", "<a href='#'>博客</a>")
	})

	r.Run(":8080")
}
