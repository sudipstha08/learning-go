package routes

import (
	"github.com/gin-gonic/gin"
	"learning-go/api/controller"

)

func SiteMapRoutes(route *gin.Engine) {
	routes := route.Group("/sitemap")
	{
		routes.GET("", controller.GetSiteMap)
	}
}
