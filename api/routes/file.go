package routes

import (
	"net/http"

	"learning-go/api/controller"

	"github.com/gin-gonic/gin"
)

func FileRoutes(route *gin.Engine) {
	fileRoutes := route.Group("/")
	{
		fileRoutes.GET("/upload", func(c *gin.Context) {
			c.HTML(http.StatusOK, "select_file.html", gin.H{})
		})
		fileRoutes.POST("/upload", controller.Upload)
	}
	fileRoutes.StaticFS("/file", http.Dir("public"))
}
