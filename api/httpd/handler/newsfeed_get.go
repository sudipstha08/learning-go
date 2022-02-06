package handler

import (
	"learning-go/api/httpd/platform/newsfeed"
	"net/http"

	"github.com/gin-gonic/gin"
)

// func NewsfeedGet(feed *newsfeed.Repo) gin.HandlerFunc{
func NewsfeedGet(feed newsfeed.Getter) gin.HandlerFunc {
	// anynomus function
	return func(c *gin.Context) {
		results := feed.GetAll()
		c.JSON(http.StatusOK, results)
	}
}
