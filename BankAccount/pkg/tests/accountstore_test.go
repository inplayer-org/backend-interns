package tests

import (
	"bankacc/pkg/entities"
	"bankacc/pkg/store"
	"database/sql"
	"log"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/suite"
)

type TearDownTestSuite interface {
	TearDownTest()
}

type AccountTestSuite struct {
	suite.Suite
	Account      *entities.Account
	Accounts     []entities.Account
	AccountsP    []*entities.Account
	AccountStore store.AccountModel
	Db           *sql.DB
}

func (suite *AccountTestSuite) SetupTest() {
	var err error
	suite.Db = MySQLInit()
	account := store.NewAccountStoreModel(suite.Db)
	suite.Accounts = []entities.Account{
		{
			UserId:   2,
			Balance:  550,
			Currency: "Euros",
		},
		{
			UserId:   3,
			Balance:  420,
			Currency: "Euros",
		},
		{
			UserId:   1,
			Balance:  330,
			Currency: "$",
		},
		{
			UserId:   3,
			Balance:  650,
			Currency: "Dinar",
		},
	}

	for _, value := range suite.Accounts {
		suite.Account, err = account.InsertAccount(value.UserId, value.Balance, value.Currency)
		if err != nil {
			suite.T().Fatal("Unable to run InsertAccount store func")
		}
		suite.AccountsP = append(suite.AccountsP, suite.Account)
	}
}

func (suite *AccountTestSuite) TestGetAccountsById() {
	store := store.NewAccountStoreModel(suite.Db)
	var err error
	var accounts []*entities.Account
	for _, value := range suite.AccountsP {
		accounts, err = store.GetAccountsByUserId(value.UserId)
		if err != nil {
			suite.T().Fatal("Unable to run GetAccountsById store func")
		}
	}
	for _, value := range suite.AccountsP {
		for i := range accounts {
			if value.Id == accounts[i].Id {
				suite.Equal(value.UserId, accounts[i].UserId)
				suite.Equal(value.Balance, accounts[i].Balance)
				suite.Equal(value.Currency, accounts[i].Currency)
				suite.Equal(value.Status, accounts[i].Status)
			}
		}
	}
}

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
				suite.Equal(value.Status, accounts[i].Status)
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
