package routes

import (
	"learning-go/api/controller"

	"github.com/gin-gonic/gin"
)

func CsvRoutes(route *gin.Engine) {
	csvRoutes := route.Group("csv")
	{
		csvRoutes.GET("/download", controller.HandleCsvDownload)
	}
}
