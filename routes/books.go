package routes

import (
	"learning-go/controller"
	"learning-go/middlewares"

	"github.com/gin-gonic/gin"
)

func BooksRoutes(route *gin.Engine) {
	route.Use(middlewares.SetMiddlewareAuthentication())

	route.GET("/books", controller.FindBooks)
}
