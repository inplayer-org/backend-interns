package account

import "sync"

type (
	Account interface {
		Close() (payout int64, ok bool)
		Balance() (balance int64, ok bool)
		Deposit(amount int64) (newBalance int64, ok bool)
	}
)

type accountData struct {
	balance int64
	open    bool
	mux     sync.Mutex
}

func Open(initialDeposit int64) Account {

	if initialDeposit < 0 {
		return nil
	}

	var acc Account
	acc = &accountData{
		balance: initialDeposit,
		open:    true,
	}

	return acc
}

func (acc *accountData) Close() (payout int64, ok bool) {
	acc.mux.Lock()
	defer acc.mux.Unlock()

	if !acc.open {
		return 0, false
	}

	acc.open = false
	return acc.balance, true
}

func (acc *accountData) Balance() (balance int64, ok bool) {

	acc.mux.Lock()
	defer acc.mux.Unlock()

	if !acc.open {
		return 0, false
	}

	return acc.balance, true
}

func (acc *accountData) Deposit(amount int64) (int64, bool) {

	acc.mux.Lock()
	defer acc.mux.Unlock()

	if !acc.open {
		return 0, false
	}

	if acc.balance+amount < 0 {
		return acc.balance, false
	}

	acc.balance += amount
	return acc.balance, true
}
