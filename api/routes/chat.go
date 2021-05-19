package routes

// WebSockets basically offer us duplex communication from a non-trusted
// source to a server that we own across a TCP socket connection.
import (
	"learning-go/api/controller"

	"github.com/gin-gonic/gin"
)

func ChatRoutes(route *gin.Engine) {
	chatRoutes := route.Group("/ws")

	chatRoutes.GET("", controller.ServeWs)
	// chatRoutes.GET("", gin.WrapF(controller.ServeWs))
}
