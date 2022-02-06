package routes

import (
	"learning-go/api/httpd/handler"
	"learning-go/api/httpd/platform/newsfeed"

	"github.com/gin-gonic/gin"
)

func NewsFeedRoutes(route *gin.Engine) {
	feed := newsfeed.New()
	route.GET("/ping", handler.PingGet())
	route.GET("/newsfeed", handler.NewsfeedGet(feed))
	route.POST("/newsfeed", handler.NewsfeedPost(feed))
}
