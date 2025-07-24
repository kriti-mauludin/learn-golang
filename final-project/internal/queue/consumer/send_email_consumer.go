package consumer

import (
	"context"
	"fmt"

	"github.com/kriti-mauludin/try-consumer-rabbitmq/entity"
	"github.com/kriti-mauludin/try-consumer-rabbitmq/internal/helper"
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

	helper.Dump("Processing Notif Success add new todolist ...")
	fmt.Println(params)

	fmt.Println("SEND NOTIF SUCCESS!")

	return nil
}
