package middlewares

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func SetMiddlewareAuthentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		// Set example variable
		c.Set("UserID", "12345")
		// before request
		fmt.Println("---------Before next-------", time.Since(t).Minutes())
		c.Next()
		// after request
		latency := time.Since(t).Nanoseconds()
		fmt.Println("---------After next-------", latency)
		// access the status we are sending
		status := c.Writer.Status()
		fmt.Println("------STATUS--------", status)
	}
}
