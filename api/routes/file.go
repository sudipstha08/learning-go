package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"learning-go/api/controller"

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
