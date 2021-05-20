package routes

// WebSockets basically offer us duplex communication from a non-trusted
// source to a server that we own across a TCP socket connection.
import (
	"learning-go/api/controller"
	"learning-go/api/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ChatRoutes(route *gin.Engine) {
	pool := helpers.NewPool()
	go pool.Start()
	chatRoutes := route.Group("/ws")

	chatRoutes.GET("", gin.WrapF(func(w http.ResponseWriter, r *http.Request) {
		controller.ServeWs(pool, w, r)
	}))
	// chatRoutes.GET("", gin.WrapF(controller.ServeWs))
}
