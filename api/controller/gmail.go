package controller

import (
	"log"
	"net/http"
	"net/smtp"

	"github.com/gin-gonic/gin"
)

func HandleSendGmail(c *gin.Context) {
	from := "shresthation@gmail.com"
	pass := "Dev@Sudip_1997g"
	to := "sudipstha08@gmail.com"

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: Hello there\n\n" +
		"Hello, This is the body of the mail"

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
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

// https://github.com/tangingw/go_smtp/blob/master/send_mail.go
// https://pkg.go.dev/google.golang.org/api/gmail/v1#hdr-Creating_a_client
// https://developers.google.com/gmail/api/quickstart/go