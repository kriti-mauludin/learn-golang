package main

import (
	"github.com/kriti-mauludin/learn-golang/section-2/5-oop/example-case/usecase"
)

func main() {
	senderUsecase := usecase.NewSmtpSenderUsecase()
	userUsecase := usecase.NewUserUsecase(senderUsecase)

	userUsecase.RegisterUser("km@example.com")
}
