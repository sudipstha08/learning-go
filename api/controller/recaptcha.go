package controller

import (
	"encoding/json"
	"log"
	"os"

	"io/ioutil"
	"net/http"
	"net/url"
	"learning-go/models"

	"github.com/gin-gonic/gin"
)

const recaptchaServerName = "https://www.google.com/recaptcha/api/siteverify"

//VerifyReCaptcha verifies the recaptcha
func VerifyReCaptcha(ctx *gin.Context) {
	secretKey := os.Getenv("RECAPTCHA_SECRET_KEY")
	log.Println(secretKey)
	var input models.RecaptchaRequest

	if err := ctx.ShouldBind(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request: " + err.Error(),
		})
		return
	}
	resp, err := http.PostForm(recaptchaServerName,
		url.Values{"secret": {secretKey}, "response": {input.Response}})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var responseData models.RecaptchaResponse
	if err := json.Unmarshal(body, &responseData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, responseData)
}
