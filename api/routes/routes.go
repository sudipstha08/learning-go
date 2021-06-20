package routes

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetRoutes(Router *gin.Engine) {
	fmt.Println("-------- SETTING UP ROUTES 📦 .........")
	 
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
	SiteMapRoutes(Router)
	fmt.Println("-------- ROUTES SETUP COMPLETE 🚉 -------")
	fmt.Println("-------------------------------------")
}
