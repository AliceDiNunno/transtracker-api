package V1

import "github.com/gin-gonic/gin"

func ProductsRoute(context *gin.RouterGroup) {
	context.GET("/", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"hello": "world",
		})
	})
}