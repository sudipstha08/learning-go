package routes

import (
	// "learning-go/routes"
	"os"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Gin *gin.Engine
}

func GetRoutes() {
	Router := gin.Default()
	Router.Static("/css", "./templates/css")
	Router.LoadHTMLGlob("templates/*.html")
	// Router.Use(gin.Recovery(), gin.Logger())
	// Router.Use(gin.Recovery(), middlewares.Logger(), middlewares.BasicAuth(), gindump.Dump())
	Router.Use(gin.Recovery(), gin.Logger())

	// ROUTES
	ArticlesRoutes(Router)
	VideosRoute(Router)
	AuthRoutes(Router)
	NewsFeedRoutes(Router)
	FileRoutes(Router)

	// routes.BooksRoutes(Router)

	// LISTEN AND SERVE ON 127.0.0.1:5000
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}
	Router.Run(":" + port)
}
