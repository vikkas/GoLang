package main

import (
	"server/service"
	"github.com/gin-gonic/gin"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func main() {
	//server := gin.Default()
	
	server := gin.New()
	//server.Use(gin.Recovery())
	//server.Use(gin.Logger())
	
	server.Use(gin.Recovery(), gin.Logger())
	
	server.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.FinaAll())
	})
	server.POST("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.Save(ctx))
	})
	server.Run(":8080")
}
