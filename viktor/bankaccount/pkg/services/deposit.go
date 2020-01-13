package account

func (acc *Account) Deposit(amount int64) (newBalance int64, ok bool) {
	acc.Mux.Lock()
	defer acc.Mux.Unlock()
	{
		if acc.Status {
			if (acc.BalanceAcc + amount) < 0 {
				return 0, false
			}
			acc.BalanceAcc += amount
		}
	}
	return acc.BalanceAcc, acc.Status
}
