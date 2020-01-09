package main

import (
	account "account/pkg/services"
	"fmt"
)

type Account account.Account

func main() {
	//Initializing an account and calling the methods
	account := account.Open(10)
	fmt.Println(account)
	fmt.Println(account.Balance())
	fmt.Println(account.Deposit(30))
	account.Close()
	fmt.Println(account)
}
