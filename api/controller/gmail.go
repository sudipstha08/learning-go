package controller

import (
	"io/ioutil"
	service "learning-go/api/services"
	"learning-go/utils"
	"log"
	"net/http"
	"net/smtp"
	"os"

	"github.com/gin-gonic/gin"
)

func SentryTest(c *gin.Context) {
	c.JSON(http.StatusUnauthorized, gin.H{
		"error": "Error occured",
	})
	utils.SendMsgToSentry(c, "Error sentry -----")
}

func HandleSendGmail(c *gin.Context) {
	from := os.Getenv("GMAIL_MAIL")
	pass := os.Getenv("GMAIL_PASSWORD")
	to := "testmail@mailinator.com"

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: Hello there\n\n" +
		"Hello, This is the body of the mail"

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
		utils.SendMsgToSentry(c, err.Error())
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err,
		})
		return
	}

	log.Print("sent, visit http://foobarbazz.mailinator.com")
	c.JSON(http.StatusOK, gin.H{
		"success": "The mail is sent. Please check your mail" + to,
	})
}

func HandleSendMail(c *gin.Context) {
	sender := service.NewSender(os.Getenv("GMAIL"), os.Getenv("GMAIL_PASSWORD"))
	// sender := service.Sender{os.Getenv("GMAIL_MAIL"), os.Getenv("GMAIL_PASSWORD")}

	//The receiver needs to be in slice as the receive supports multiple receiver
	Receiver := []string{"johncena7@mailinator.com", "larrypage@mailinator.com"}

	Subject := "Testing HTLML Email from golang"
	currentDir, err := os.Getwd()
	if err != nil {
		utils.SendMsgToSentry(c, err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	message, err := ioutil.ReadFile(currentDir + "/templates/email.html")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Error occured",
		})
		utils.SendMsgToSentry(c, err.Error())
		log.Fatal("Error finding the file :", err)
		return
	}
	bodyMessage := sender.WriteHTMLEmail(Receiver, Subject, string(message))

	err = sender.SendMail(Receiver, Subject, bodyMessage)
	if err != nil {
		utils.SendMsgToSentry(c, err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": "Mail sent successfully",
	})
}

// https://pkg.go.dev/google.golang.org/api/gmail/v1#hdr-Creating_a_client
// https://developers.google.com/gmail/api/quickstart/go
