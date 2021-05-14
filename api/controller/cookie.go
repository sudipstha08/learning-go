package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func HandleGetCookie() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Cookie("gin_cookie")

		if err != nil {
			cookie = "NotSet"
			c.SetCookie("gin_cookie", "test", 3600, "/", "localhost", false, true)
		}
		fmt.Printf("Cookie value: %s \n", cookie)
	}

}
