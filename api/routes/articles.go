package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ArticlesRoutes(route *gin.Engine) {
	articlesRoutes := route.Group("/articles")
	{
		articlesRoutes.GET("/", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "Success",
			})
		})
	}
}
