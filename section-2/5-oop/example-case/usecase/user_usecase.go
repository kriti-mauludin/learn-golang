package usecase

import "fmt"

//like class user in oop
type UserUsecase struct {
	sender ISmtpSenderUsecase
}

//constructor
func NewUserUsecase(sender ISmtpSenderUsecase) *UserUsecase {
	return &UserUsecase{
		sender: sender,
	}
}

//abstract interface UserUsecase like (polymorphism)
type IUserUsecase interface {
	RegisterUser(email string)
}

//implementation of IUserUsecase
func (u *UserUsecase) RegisterUser(email string) {
	fmt.Println("Registering successfully with email:", email)
	u.sender.SendEmail(email, "Welcome to our service!")
}
