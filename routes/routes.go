package routes

import (
	"learning-go/infrastructures"
	"os"

	"github.com/gin-gonic/gin"
)

type Router struct {
	Route infrastructures.Server
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
	CsvRoutes(Router)

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}
	Router.Run(":" + port)
}
