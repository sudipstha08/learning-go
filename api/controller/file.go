package controller

import (
	"fmt"
	"io"
	"learning-go/utils"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// upload controller
func Upload(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		utils.SendMsgToSentry(c, err.Error())
		c.String(http.StatusBadRequest, fmt.Sprintf("file err : %s", err.Error()))
		return
	}
	filename := header.Filename
	out, err := os.Create("public/" + filename)
	if err != nil {
		utils.SendMsgToSentry(c, err.Error())
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
