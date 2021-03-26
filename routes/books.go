package routes

import (
	"learning-go/controller"

	"github.com/gin-gonic/gin"
)

func BooksRoutes(route *gin.Engine) {
	route.GET("/books", controller.FindBooks)
}
