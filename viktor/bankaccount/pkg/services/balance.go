package account

func (acc *Account) Balance() (balance int64, ok bool) {
	return acc.BalanceAcc, acc.Status
}
