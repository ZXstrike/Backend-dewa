package mail

import (
	"log"
	"zxsttm/server/config"

	"gopkg.in/gomail.v2"
)

func SendMail(to []string, cc []string, subject, message string) error {
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", config.Config.SMTP.Email)
	mailer.SetHeader("To", to...)
	mailer.SetAddressHeader("Cc", config.Config.SMTP.Email, "Admin")
	mailer.SetHeader("Subject", subject)
	mailer.SetBody("text/html", "Hello, this is your OTP <b>"+message+"</b>")

	dialer := &gomail.Dialer{
		Host: config.Config.SMTP.Host,
		Port: 25,
	}

	err := dialer.DialAndSend(mailer)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println("Mail sent!")
	return nil
}
