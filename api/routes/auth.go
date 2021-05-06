package routes

import (
	"learning-go/api/controller"
	service "learning-go/api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	loginService service.LoginService = service.NewLoginService()

	loginController controller.LoginController = controller.NewLoginController(loginService, jwtService)
)

// type AuthRoute struct {
// 	router     infrastructures.Server
// 	controller controller.LoginController
// }

// LOGIN ENDPOINTS
func AuthRoutes(route *gin.Engine) {
	route.POST("/login", func(ctx *gin.Context) {
		token := loginController.Login(ctx)
		if token != "" {
			ctx.JSON(http.StatusOK, gin.H{
				"token": token,
			})
		} else {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid username or password",
			})
		}
	})
}
