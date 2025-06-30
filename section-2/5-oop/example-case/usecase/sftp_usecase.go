package usecase

import "fmt"

type SmtpSenderUsecase struct{}

func NewSmtpSenderUsecase() *SmtpSenderUsecase {
	return &SmtpSenderUsecase{}
}

type ISmtpSenderUsecase interface {
	SendEmail(to string, body string)
}

func (s SmtpSenderUsecase) SendEmail(to string, body string) {
	fmt.Printf("Sending email to: %s\nmessage: %s\n", to, body)
}
