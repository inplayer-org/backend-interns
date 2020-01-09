package closeaccount

func (acc *Account) Close() (payout int64, ok bool) {
	var pay int64
	var close bool
	acc.mux.Lock()
	acc.mux.Unlock()
	{
		close = acc.status
		if acc.status {
			acc.status = false
			pay = acc.balance
			acc.balance = 0
		}
	}
	return pay, close
}
