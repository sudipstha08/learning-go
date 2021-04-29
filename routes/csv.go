package routes

import "github.com/gin-gonic/gin"

import (
	"learning-go/controller"

)

func CsvRoutes(route *gin.Engine) {
	csvRoutes := route.Group("csv")
	{
		csvRoutes.GET("/download", controller.HandleCsvDownload)
	}
}
