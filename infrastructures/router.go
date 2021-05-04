package infrastructures

import (
	"learning-go/routes"
	"os"

	"github.com/gin-contrib/cors"
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
	Router.Use(cors.Default())

	// ROUTES
	routes.ArticlesRoutes(Router)
	routes.VideosRoute(Router)
	routes.AuthRoutes(Router)
	routes.NewsFeedRoutes(Router)
	routes.FileRoutes(Router)
	routes.CsvRoutes(Router)
	routes.RecaptchaRoutes(Router)

	// LISTEN AND SERVE ON 127.0.0.1:5000
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}
	Router.Run(":" + port)
}
