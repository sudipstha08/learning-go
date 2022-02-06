package routes

import (
	"learning-go/api/controller"

	"github.com/gin-gonic/gin"
)

func CookieRoutes(route *gin.Engine) {
	cookieRoutes := route.Group("/cookie")
	{
		cookieRoutes.GET("", controller.HandleGetCookie())
	}
}
