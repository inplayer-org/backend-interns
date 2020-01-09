package main

import (
	"fmt"
	"sync"
)

type Account struct {
	mux     sync.Mutex
	balance int64
	status  bool
}

func main() {
	account := Account{balance: 10, status: true}
	fmt.Println(&account)
}
