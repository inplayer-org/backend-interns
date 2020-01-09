package deposit

func (acc *Account) Deposit(amount int64) (newBalance int64, ok bool) {
	acc.mux.Lock()
	defer acc.mux.Unlock()
	{
		if acc.status {
			if (acc.balance + amount) < 0 {
				return 0, false
			}
			acc.balance += amount
		}
	}
	return acc.balance, acc.status
}
