package tests

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"bankacc/pkg/entities"
)

func (suite *AccountTestSuite) TestUpdateAccount() {
	store := store.NewAccountStoreModel(suite.Db)
	var err error
	var account *entities.Account
	var accounts []*entities.Account
	//time.Sleep(2 * time.Second)
	for _, value := range suite.AccountsP {
		account, err = store.UpdateAccount(value.Id, value.UserId, value.Balance, value.Currency)
		if err != nil {
			suite.T().Fatal("Unable to run UpdateAccount store func")
		}
		accounts = append(accounts, account)
	}
	for _, value := range suite.AccountsP {
		for i := range accounts {
			if value.Id == accounts[i].Id {
				suite.Equal(value.UserId, accounts[i].UserId)
				suite.Equal(value.Balance, accounts[i].Balance)
				suite.Equal(value.Currency, accounts[i].Currency)
			}
		}
	}
}
func (suite *AccountTestSuite) TestCloseAccount() {
	store := store.NewAccountStoreModel(suite.Db)
	var err error
	var account *entities.Account
	var accounts []*entities.Account
	for _, value := range suite.AccountsP {
		account, err = store.CloseAccount(value.Id, value.UserId)
		if err != nil {
			suite.T().Fatal("Unable to run CloseAccount store func")
		}
		accounts = append(accounts, account)
	}
	for _, value := range suite.AccountsP {
		for i := range accounts {
			if value.Id == accounts[i].Id {
				suite.Equal(value.UserId, accounts[i].UserId)
				suite.Equal(false, accounts[i].Status)
			}
		}
	}
}
func TestAccountTestSuite(t *testing.T) {
	suite.Run(t, new(AccountTestSuite))
}
func (suite *AccountTestSuite) TearDownTest() {
	for i := 0; i < len(suite.AccountsP); i++ {
		_, err := suite.Db.Exec("DELETE FROM Account WHERE id=?", suite.AccountsP[i].Id)
		if err != nil {
			suite.T().Fatal("Unable to run delete query")
		}
	}
}
