package mailer

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/mail"
	"net/smtp"
	"os"
	"strings"
)

type SvcInterface interface {
	Send(obj MailObject) error
	Fire(message string, from, to mail.Address) error
}

type svcImplementation struct {
}

func NewHandler() SvcInterface {
	return svcImplementation{}
}

func (s svcImplementation) Send(obj MailObject) (err error) {
	log.Println("sending mail to:", obj.Email)

	from := mail.Address{Name: "", Address: os.Getenv("SMTP_SENDER")}
	to := mail.Address{Name: "", Address: strings.TrimSpace(obj.Email)}

	// Setup headers
	headers := make(map[string]string)
	headers["From"] = from.String()
	headers["To"] = to.String()
	headers["Subject"] = obj.Subject

	message := ""
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";"
	message += "\r\n" + obj.ParsedBody

	if err := s.Fire(message, from, to); err != nil {
		log.Println("Err on sending mail:", err)
		return err
	}

	log.Println("mail sent to:", obj.Email)
	return nil
}

func (s svcImplementation) Fire(message string, from, to mail.Address) error {
	host := os.Getenv("SMTP_HOST")
	servername := fmt.Sprintf("%s:%s", host, os.Getenv("SMTP_PORT"))

	auth := smtp.PlainAuth("", os.Getenv("SMTP_USERNAME"), os.Getenv("SMTP_PASSWORD"), host)

	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         host,
	}

	conn, err := tls.Dial("tcp", servername, tlsconfig)
	if err != nil {
		log.Println("Err on dialing tcp:", err)
		return err
	}

	c, err := smtp.NewClient(conn, host)
	if err != nil {
		log.Println("Err on setting new smtp client:", err)
		return err
	}

	if err = c.Auth(auth); err != nil {
		log.Println("Err on authenticate smtp client:", err)
		return err
	}

	if err = c.Mail(from.Address); err != nil {
		log.Println("Err on setting mail sender:", err)
		return err
	}

	if err = c.Rcpt(to.Address); err != nil {
		log.Println("Err on setting recipient:", err)
		return err
	}

	w, err := c.Data()
	if err != nil {
		log.Println("Err on get smtp client writer data:", err)
		return err
	}

	_, err = w.Write([]byte(message))
	if err != nil {
		log.Println("Err on write writer:", err)
		return err
	}

	err = w.Close()
	if err != nil {
		log.Println("Err on closing writer:", err)
		return err
	}

	c.Quit()
	return nil
}
