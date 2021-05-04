package routes

import "github.com/gin-gonic/gin"

import (
	"learning-go/controller"

)

func RecaptchaRoutes(route *gin.Engine) {
	recaptchaRoutes := route.Group("")
	{
		recaptchaRoutes.POST("/verify-recaptcha", controller.VerifyReCaptcha)
	}
}
