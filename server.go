package main
import "github.com/gin-gonic/gin"

func main() {
	Router := gin.Default()
	Router.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "ok!!!",
		})
	})
	Router.Run(":8080") // listen and serve on 0.0.0.0:8080
}