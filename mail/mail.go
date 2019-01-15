package mail

import (
	"log"
	"net/smtp"
	"os"
)

const recipient = "rozhkov@mc.asu.ru"

var (
	login = os.Getenv("MAIL_LOGIN")
	pass  = os.Getenv("MAIL_PASS")
)

// Check verifies working of ASU's mail in mail.asu.ru
func Check() []byte {
	auth := smtp.PlainAuth("", login, pass, "mail.asu.ru")

	to := []string{recipient}
	msg := []byte("To: " + recipient + "\r\n" +
		"Subject: Mail is working!\r\n" +
		"\r\n" +
		"There is nothing interesting.\r\n")

	err := smtp.SendMail("mail.asu.ru:25", auth, "Checker@asu.ru", to, msg)
	if err != nil {
		log.Println("unable to send email. ", err)
		return []byte("false")
	}

	log.Println("Почта на mail.asu.ru успешно работает!")
	return []byte("true")
}