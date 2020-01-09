package balance

func (acc *Account) Balance() (balance int64, ok bool) {
	return acc.balance, acc.status
}
