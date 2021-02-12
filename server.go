package main

import (
	"io"
	"learning-go/controller"
	"learning-go/middlewares"
	"learning-go/service"
	"os"

	"github.com/gin-gonic/gin"
	// gindump as alias
	gindump "github.com/tpkeeper/gin-dump" // DEBUGGER
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

// Creates a log file
func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {
	// Router := gin.Default()
	// Router.GET("/test", func(ctx *gin.Context) {
	// 	ctx.JSON(200, gin.H{
	// 		"message": "ok!!!",
	// 	})
	// })
	setupLogOutput()
	Router := gin.New()
	// Router.Use(gin.Recovery(), gin.Logger())
	Router.Use(gin.Recovery(), middlewares.Logger(), middlewares.BasicAuth(), gindump.Dump())
	Router.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.FindAll())
	})

	Router.POST("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.Save(ctx))
	})
	// LISTEN AND SERVE ON 127.0.0.1:8080
	Router.Run(":8080")
}
