package infrastructures

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Server struct {
	Gin *gin.Engine
}

func GetRoutes() Server {
	Router := gin.Default()
	Router.Use(cors.Default())

	return Server{
		Gin: Router,
	}
}
