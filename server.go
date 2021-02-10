package main
import "github.com/gin-gonic/gin"

func main() {
	Router := gin.Default()
	Router.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ok!!!",
		})
	})
	Router.Run(":5000") // listen and serve on 0.0.0.0:5000
}