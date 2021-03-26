package controller

import (
	// "learning-go/infrastructures"
	"learning-go/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FindBooks(c *gin.Context) {
	var books []models.Book
	// infrastructures.DB.Find(&books)
	c.JSON(http.StatusOK, gin.H{
		"data": books,
	})
}
