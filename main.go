package main

import "github.com/gin-gonic/gin"

func main() {

	ginServer := gin.Default()
	ginServer.GET("/hello", func(context *gin.Context) {
		context.JSON(200, gin.H{"msg": "这是git上传的测试"})
	})

	ginServer.Run(":8888")

}
