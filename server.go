package main

import (
	"fmt"
	"io"
	"learning-go/controller"
	"learning-go/middlewares"
	"learning-go/service"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	// gindump as alias
	// gindump "github.com/tpkeeper/gin-dump" // DEBUGGER
)

var (
	videoService service.VideoService = service.New()
	loginService service.LoginService = service.NewLoginService()
	jwtService   service.JWTService   = service.NewJWTService()

	videoController controller.VideoController = controller.New(videoService)
	loginController controller.LoginController = controller.NewLoginController(loginService, jwtService)
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
	// Router.Use(gin.Recovery(), middlewares.Logger(), middlewares.BasicAuth(), gindump.Dump())
	Router.Use(gin.Recovery(), gin.Logger())

	Router.POST("/login", func(ctx *gin.Context) {
		fmt.Println("Helllooo")
		token := loginController.Login(ctx)
		if token != "" {
			ctx.JSON(http.StatusOK, gin.H{
				"token": token,
			})
		} else {
			ctx.JSON(http.StatusUnauthorized, nil)
		}
	})

	// JWT Authorization Middleware applies to "/api" only.
	apiRoutes := Router.Group("/api", middlewares.AuthorizeJWT())
	{
		apiRoutes.GET("/videos", func(ctx *gin.Context) {
			ctx.JSON(200, videoController.FindAll())
		})

		apiRoutes.POST("/videos", func(ctx *gin.Context) {
			err := videoController.Save(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"Error in adding videos": err.Error(),
				})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "Video added successfully"})
			}
		})
	}

	viewRoutes := Router.Group("/view")
	{
		viewRoutes.GET("/videos", videoController.ShowAll)
	}

	// LISTEN AND SERVE ON 127.0.0.1:8080
	Router.Run(":8080")
}
