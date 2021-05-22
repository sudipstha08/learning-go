package routes

import (
	"github.com/gin-gonic/gin"
)

func GetRoutes(Router *gin.Engine) {
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
	CookieRoutes(Router)
	ChatRoutes(Router)
	GmailRoutes(Router)
}
