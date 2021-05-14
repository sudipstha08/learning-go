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
	UserID := c.MustGet("UserID").(string)

	c.JSON(http.StatusOK, gin.H{
		"data":    books,
		"user_id": UserID,
	})
}
