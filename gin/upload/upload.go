package upload

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

func FileUpload() {
	r := gin.Default()
	r.POST("/upload", func(ctx *gin.Context) {
		file, _ := ctx.FormFile("file")
		open, _ := file.Open()
		defer open.Close()
		create, _ := os.Create("./" + file.Filename)
		defer create.Close()
		io.Copy(create, open)
		//fmt.Println(file.Filename)
		ctx.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", file.Filename))
		ctx.File("./" + file.Filename)
	})
	r.Run(":8080")
}
