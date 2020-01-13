package account

func (acc *Account) Close() (payout int64, ok bool) {
	var pay int64
	var close bool
	acc.Mux.Lock()
	defer acc.Mux.Unlock()
	{
		close = acc.Status
		if acc.Status {
			acc.Status = false
			pay = acc.BalanceAcc
			acc.BalanceAcc = 0

		}
	}
	return pay, close
}
