package routes

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func TestRoutes(route *gin.Engine) {
	recaptchaRoutes := route.Group("/test")
	signature := ""

	recaptchaRoutes.Use(func(c *gin.Context) {
		signature += "A"
		// Use Next, the internal will to perform in front of the route has been registered middleware
		fmt.Println("First middleware", signature)
		// Next should be used only inside middleware.
		// The operation before next() is executed before the handler;
		// The operation after next () is executed after the handler;
		c.Next()
		signature += "W"
		fmt.Println("After First middleware", signature)
	})

	recaptchaRoutes.Use(func(c *gin.Context) {
		signature += "B"
		fmt.Println("Second Middleware", signature)
		c.Next()
		fmt.Println("After second middleware", signature)
	})

	{
		recaptchaRoutes.GET("", func(c *gin.Context) {
			fmt.Println("Route is executed", signature)
		})
	}
}
