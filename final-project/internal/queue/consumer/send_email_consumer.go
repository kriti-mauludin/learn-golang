package consumer

import (
	"context"
	"fmt"
	"log"

	"github.com/kriti-mauludin/final-project/entity"
	"github.com/kriti-mauludin/final-project/internal/mailer"
)

type SendEmailQueue struct {
	ctx context.Context
}

type SendEmail interface {
	ProcessSendNotif(payload map[string]interface{}) error
}

func NewSendEmail(
	ctx context.Context,
) SendEmail {
	return &SendEmailQueue{ctx}
}

func (l *SendEmailQueue) ProcessSendNotif(payload map[string]interface{}) error {
	var params entity.SendEmailReq
	params.LoadFromMap(payload)

	mailObj, err := entity.FillMailObjLogin(params)
	if err != nil {
		log.Println("Err FillMailObj:", err)
		return err
	}
	err = mailer.NewHandler().Send(mailObj)
	if err != nil {
		log.Println("Err send mail:", err)
	}

	fmt.Println("SEND NOTIF SUCCESS!")

	return nil
}
