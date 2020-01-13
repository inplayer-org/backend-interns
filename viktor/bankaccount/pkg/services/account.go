package account

import (
	"sync"
)

type Account struct {
	Mux        sync.Mutex
	BalanceAcc int64
	Status     bool
}
