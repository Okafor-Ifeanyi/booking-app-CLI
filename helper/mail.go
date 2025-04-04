package helper

import (
	"crypto/tls"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gopkg.in/gomail.v2"
)


func SendMail(reciever string) {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get environment variable
	username := os.Getenv("SMTP_USER")
	password := os.Getenv("SMTP_PASS")
	host := os.Getenv("SMTP_HOST")
	

	mailConfig := gomail.NewMessage()
	mailConfig.SetHeader("From", username)
	mailConfig.SetHeader("To", reciever)
	// mailConfig.SetAddressHeader("Cc", "dan@example.com", "Dan")
	mailConfig.SetHeader("Subject", "Backend Dev Role!")
	mailConfig.SetBody("text/html", "Hello <b>Testing GO</b> and <i>Testing GO</i>!")
	// mailConfig.Attach("/home/Alex/lolcat.jpg")

	fmt.Println("Sending email...", username, password)

	d := gomail.NewDialer(host, 587, username, password)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(mailConfig); err != nil {
		panic(err)
	}
}