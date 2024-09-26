package main

import "github.com/gin-gonic/gin"

func main() {

	ginServer := gin.Default()
	ginServer.GET("/hello", func(context *gin.Context) {
		context.JSON(200, gin.H{"msg": "一个月之内写一个go的后端项目"})
	})

	ginServer.Run(":8888")

}
