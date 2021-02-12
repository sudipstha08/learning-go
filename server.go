package main

import (
	"go_crash_course/controller"
	"go_crash_course/service"

	"github.com/gin-gonic/gin"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func main() {
	Router := gin.Default()
	// Router.GET("/test", func(ctx *gin.Context) {
	// 	ctx.JSON(200, gin.H{
	// 		"message": "ok!!!",
	// 	})
	// })
	Router.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.FindAll())
	})

	Router.POST("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.Save(ctx))
	})
	// LISTEN AND SERVE ON 127.0.0.1:8080
	Router.Run(":8080") 
}
