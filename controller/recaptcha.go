package controller

import (
	"encoding/json"
	"log"
	"os"

	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/gin-gonic/gin"
)

//RecaptchaResponse structure.
type RecaptchaResponse struct {
	Success     bool      `json:"success"`
	Score       float64   `json:"score"`
	Action      string    `json:"action"`
	ChallengeTS time.Time `json:"challenge_ts"`
	Hostname    string    `json:"hostname"`
	ErrorCodes  []string  `json:"error-codes"`
}

//RecaptchaRequest structure
type RecaptchaRequest struct {
	Response string `json:"response" binding:"required"`
}

const recaptchaServerName = "https://www.google.com/recaptcha/api/siteverify"

//VerifyReCaptcha verifies the recaptcha
func VerifyReCaptcha(ctx *gin.Context) {
	secretKey := os.Getenv("RECAPTCHA_SECRET_KEY")
	log.Println(secretKey)
	var input RecaptchaRequest

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

	var responseData RecaptchaResponse
	if err := json.Unmarshal(body, &responseData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, responseData)
}
