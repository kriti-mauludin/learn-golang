package consumer

import (
	"context"
	"fmt"

	"github.com/kriti-mauludin/learn-golang/entity"
	"github.com/kriti-mauludin/learn-golang/internal/helper"
)

type SendEmailQueue struct {
	ctx context.Context
}

type SendEmail interface {
	ProcessSyncLog(payload map[string]interface{}) error
}

func NewSendEmail(
	ctx context.Context,
) SendEmail {
	return &SendEmailQueue{ctx}
}

func (l *SendEmailQueue) ProcessSyncLog(payload map[string]interface{}) error {
	var params entity.Log
	params.LoadFromMap(payload)

	helper.Dump(l.ctx)
	helper.Dump("Processing Send Email ...")

	fmt.Println("SYNC SUCCESS!")
	fmt.Println(params)

	return nil
}
