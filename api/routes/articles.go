package routes

import (
	"learning-go/api/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
)

// type ArticlesRoutes struct {
// 	controller controller.ArticlesController
// 	router infrastructures.Router
// }

func HandleGetArticles() gin.HandlerFunc {
	return func(c *gin.Context) {
		ReqID := c.Writer.Header().Get("X-Request-Id")
		c.JSON(http.StatusOK, gin.H{
			"data": ReqID,
		})
	}
}

func ArticlesRoutes(route *gin.Engine) {
	articlesRoutes := route.Group("/articles")
	articlesRoutes.Use(middlewares.RequestIdMiddleware())
	{
		articlesRoutes.GET("", HandleGetArticles())
	}
}
