package tests

import (
	"bankacc/pkg/entities"
	"bankacc/pkg/store"
	"database/sql"
	"log"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/suite"
)

func MySQLInit() *sql.DB {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "Password1!"
	dbName := "BankAccount?parseTime=true"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		log.Println(err)
	}
	return db
}

type TearDownTestSuite interface {
	TearDownTest()
}

type AccountTestSuite struct {
	suite.Suite
	Account      entities.Account
	Accounts     []entities.Account
	AccountStore store.AccountModel
	Db           *sql.DB
}

func (suite *AccountTestSuite) SetupTest() {
	var err error
	suite.Db = MySQLInit()
	account := store.NewAccountStoreModel(suite.Db)
	suite.Accounts = []entities.Account{
		{
			UserId:   3,
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

	for i, current := range suite.Accounts {
		suite.Account, err = account.InsertAccount(current.UserId, current.Balance, current.Currency)
		if err != nil {
			suite.T().Fatal("Unable to run InsertAccount store func")
		}
		suite.Accounts[i] = suite.Account
	}
}

func (suite *AccountTestSuite) TestGetAccountsById() {
	store := store.NewAccountStoreModel(suite.Db)
	var err error
	var account []entities.Account
	var accountsUserIDOne []entities.Account
	var accountsUserIDThree []entities.Account
	var accountsUserIDFive []entities.Account
	now := time.Now()

	account, err = store.GetAccountsByUserId(1)
	if err != nil {
		suite.T().Fatal("Unable to run GetAccountsByUserId")
	}
	for i, current := range suite.Accounts {
		current.CreatedAt = now
		current.UpdatedAt = now
		if len(account) > i {
			account[i].CreatedAt = now
			account[i].UpdatedAt = now
		}
		if current.UserId == 1 {
			accountsUserIDOne = append(accountsUserIDOne, current)
		}
	}
	suite.Equal(accountsUserIDOne, account)

	account, err = store.GetAccountsByUserId(3)
	if err != nil {
		suite.T().Fatal("Unable to run GetAccountsByUserId")
	}
	for i, current := range suite.Accounts {
		current.CreatedAt = now
		current.UpdatedAt = now
		if len(account) > i {
			account[i].CreatedAt = now
			account[i].UpdatedAt = now
		}
		if current.UserId == 3 {
			accountsUserIDThree = append(accountsUserIDThree, current)
		}
	}
	suite.Equal(accountsUserIDThree, account)

	account, err = store.GetAccountsByUserId(5)
	if err != nil {
		suite.T().Fatal("Unable to run GetAccountsByUserId")
	}
	for i, current := range suite.Accounts {
		current.CreatedAt = now
		current.UpdatedAt = now
		if len(account) > i {
			account[i].CreatedAt = now
			account[i].UpdatedAt = now
		}
		if current.UserId == 5 {
			accountsUserIDFive = append(accountsUserIDFive, current)
		}
	}
	suite.Equal(accountsUserIDFive, account)
}

func (suite *AccountTestSuite) TestUpdateAccount() {
	now := time.Now()
	store := store.NewAccountStoreModel(suite.Db)
	var err error
	var account entities.Account
	time.Sleep(2 * time.Second)
	for _, current := range suite.Accounts {
		account, err = store.UpdateAccount(current.Id, current.UserId, current.Balance, current.Currency)
		if err != nil {
			suite.T().Fatal("Unable to run UpdateAccount store func")
		}
		current.CreatedAt = now
		current.UpdatedAt = now
		account.CreatedAt = now
		account.UpdatedAt = now
		suite.Equal(current, account)
	}
}

func (suite *AccountTestSuite) TestCloseAccount() {
	now := time.Now()
	store := store.NewAccountStoreModel(suite.Db)
	var err error
	var account entities.Account
	for _, current := range suite.Accounts {
		account, err = store.CloseAccount(current.Id, current.UserId)
		if err != nil {
			suite.T().Fatal("Unable to run CloseAccount store func")
		}
		current.Status = false
		current.Currency = ""
		current.Balance = 0
		current.CreatedAt = now
		current.UpdatedAt = now
		account.CreatedAt = now
		account.UpdatedAt = now
		suite.Equal(current, account)
	}
}

func TestAccountTestSuite(t *testing.T) {
	suite.Run(t, new(AccountTestSuite))
}

func (suite *AccountTestSuite) TearDownTest() {
	for _, current := range suite.Accounts {
		_, err := suite.Db.Exec("DELETE FROM Account WHERE id=?", current.Id)
		if err != nil {
			suite.T().Fatal("Unable to run delete query")
		}
	}
}
