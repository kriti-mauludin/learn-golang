package main

import "fmt"

type Account struct {
	name     string
	username string
}

func (a *Account) SetName(name string) {
	a.name = name
}

func (a *Account) GetName() string {
	return a.name
}

func main() {
	account := &Account{username: "KM"}
	account.SetName("K M")
	fmt.Println("Account Name:", account.GetName())
	fmt.Println("Account Username:", account.username)
	fmt.Println("Account:", account)
}
