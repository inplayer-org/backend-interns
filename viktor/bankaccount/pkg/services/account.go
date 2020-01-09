package account

import (
	"sync"
)

type Account struct {
	Mux        sync.Mutex
	Balanceacc int64
	Status     bool
}
