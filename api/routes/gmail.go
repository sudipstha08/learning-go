package routes

import (
	"learning-go/api/controller"

	"github.com/gin-gonic/gin"
)

func GmailRoutes(route *gin.Engine) {
	gmailRoutes := route.Group("/gmail")
	{
		gmailRoutes.POST("/send", controller.HandleSendGmail)
	}
}
