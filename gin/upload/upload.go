package upload

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FileUpload() {
	r := gin.Default()
	r.POST("/upload", func(ctx *gin.Context) {
		// 上传单个文件，对文件的格式/大小没有限制
		//file, err := ctx.FormFile("file")
		//if err != nil {
		//	ctx.String(500, "上传图片出错")
		//}
		//open, _ := file.Open()
		//defer open.Close()
		//create, _ := os.Create("./" + file.Filename)
		//defer create.Close()
		//io.Copy(create, open)
		//ctx.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", file.Filename))

		// SaveUploadedFile 和上面代码实现的效果是一样的
		//ctx.SaveUploadedFile(file, file.Filename)
		//ctx.File("./" + file.Filename)

		// 上传指定类型的文件
		_, headers, err := ctx.Request.FormFile("file")
		if err != nil {
			ctx.String(500, fmt.Sprintf("Err when try to get file: %v", err))
		}
		if headers.Size > 1024*1024*2 {
			ctx.String(500, "文件大小超过 2M")
			return
		}
		if headers.Header.Get("Content-Type") != "image/png" {
			ctx.String(500, "只允许上传 png 图片")
			return
		}
		ctx.SaveUploadedFile(headers, "./upload/"+headers.Filename)
		ctx.File("./upload/" + headers.Filename)
	})
	r.NoRoute(func(ctx *gin.Context) {
		ctx.String(http.StatusNotFound, "404 Page Not Found")
	})
	r.Run(":8080")
}
