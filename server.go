package main

import (
	"io"
	"learning-go/controller"
	"learning-go/middlewares"
	"learning-go/service"
	"net/http"
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
	setupLogOutput()
	Router := gin.New()
	// Router.Use(gin.Recovery(), gin.Logger())
	Router.Static("/css", "./templates/css")
	Router.LoadHTMLGlob("templates/*.html")
	Router.Use(gin.Recovery(), middlewares.Logger(), middlewares.BasicAuth(), gindump.Dump())
	Router.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.FindAll())
	})

	Router.POST("/videos", func(ctx *gin.Context) {
		err := videoController.Save(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Error in adding videos": err.Error(),
			})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"message": "Video added successfully"})
		}
	})

	// LISTEN AND SERVE ON 127.0.0.1:8080
	Router.Run(":8080")
}
