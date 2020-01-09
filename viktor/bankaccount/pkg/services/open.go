package account

// func newAccount(initilaVal int64) (*Account, error) {
// 	if initilaVal < 0 {
// 		return nil, fmt.Errorf("Initial value must be greater than 0")
// 	}
// 	return &Account{balance: initilaVal, status: true}, nil
// }

func Open(initialDeposit int64) *Account {
	// if account, err := newAccount(initialDeposit); err == nil {
	// 	return account
	// }
	if initialDeposit < 0 {
		return nil
	}
	return &Account{Balanceacc: initialDeposit, Status: true}

	//return nil
}
