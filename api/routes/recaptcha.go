package routes

import (
	"learning-go/api/controller"

	"github.com/gin-gonic/gin"
)

func RecaptchaRoutes(route *gin.Engine) {
	recaptchaRoutes := route.Group("")
	{
		recaptchaRoutes.POST("/verify-recaptcha", controller.VerifyReCaptcha)
	}
}
