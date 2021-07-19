package routes

import (
	"fmt"
	"io"
	"learning-go/api/middlewares"
	service "learning-go/api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleBucketFileUpload() gin.HandlerFunc {
	return func(c *gin.Context) {
		file, fileHeader, err := c.Request.FormFile("file")
		fmt.Println("file----------------------------------", file)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "No file is received",
			})
			return
		}

		// first 512 byte of file is supposed to contain file metadata like file header
		buff := make([]byte, 512)
		_, err = file.Read(buff)
		if err != nil {
			fmt.Println("err-------------", err.Error())
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}

		// detect the content-type of the file uploaded, allow only if its image file
		filetype := http.DetectContentType(buff)
		if filetype != "image/jpeg" && filetype != "image/png" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "The provided file format is not allowed. Please upload a JPEG or PNG image",
			})
			return
		}

		// seek the file cursor to start of the file
		_, err = file.Seek(0, io.SeekStart)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}

		// Upload the file to google cloud storage bucket
		storageDir := "images" // images directory inside the storage bucket - directory should be appended with filename before uploading
		fileHeader.Filename = storageDir + "/" + fileHeader.Filename

		storageService := service.NewGCPStorageService()

		url, err := storageService.UploadFile(c.Request, file, fileHeader)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
		}

		fmt.Println("url---------------------", url)
		c.JSON(http.StatusOK, map[string]string{
			"url": url.String(),
		})
	}
}

func BucketRoutes(route *gin.Engine) {
	bucketRoutes := route.Group("/bucket")
	bucketRoutes.Use(middlewares.CORSMiddleware())
	{
		bucketRoutes.POST("/upload", HandleBucketFileUpload())
	}
}
