package account

func Open(initialDeposit int64) *Account {

	if initialDeposit < 0 {
		return nil
	}
	return &Account{BalanceAcc: initialDeposit, Status: true}
}
