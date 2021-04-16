package routes

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func upload(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("file err : %s", err.Error()))
		return
	}
	filename := header.Filename
	out, err := os.Create("public/" + filename)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		log.Fatal(err)
	}
	filepath := "http://localhost:8080/file/upload" + filename
	c.JSON(http.StatusOK, gin.H{"filepath": filepath})
}

func FileRoutes(route *gin.Engine) {
	fileRoutes := route.Group("/")
	{
		fileRoutes.GET("/upload", func(c *gin.Context) {
			c.HTML(http.StatusOK, "select_file.html", gin.H{})
		})
		fileRoutes.POST("/upload", upload)
	}
	fileRoutes.StaticFS("/file", http.Dir("public"))
}
