package routes

import (
	"learning-go/controller"
	"learning-go/middlewares"
	repository "learning-go/repositories"
	service "learning-go/services"
	"net/http"

	"github.com/gin-gonic/gin"
)



var (
	videoRepository repository.VideoRepository = repository.NewVideoRepository()
	videoService    service.VideoService         = service.New(videoRepository)
	jwtService      service.JWTService           = service.NewJWTService()

	videoController controller.VideoController = controller.New(videoService)
)

func VideosRoute(route *gin.Engine) {
	// JWT Authorization Middleware applies to "/api" only.
	apiRoutes := route.Group("/api", middlewares.AuthorizeJWT())
	{
		apiRoutes.GET("/videos", func(ctx *gin.Context) {
			ctx.JSON(200, videoController.FindAll())
		})

		apiRoutes.POST("/videos", func(ctx *gin.Context) {
			err := videoController.Save(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"Error in adding videos": err.Error(),
				})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "Video added successfully"})
			}
		})

		apiRoutes.PUT("/videos/:id", func(ctx *gin.Context) {
			err := videoController.Update(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
			} else {
				ctx.JSON(http.StatusOK, gin.H{
					"message": "Success",
				})
			}
		})

		apiRoutes.DELETE("/videos/:id", func(ctx *gin.Context) {
			err := videoController.Delete(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
			} else {
				ctx.JSON(http.StatusOK, gin.H{
					"message": "Success",
				})
			}
		})
	}

	viewRoutes := route.Group("/view")
	{
		viewRoutes.GET("/videos", videoController.ShowAll)
	}
}
