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

type TransactionHistoryTestSuite struct {
	suite.Suite
	Transaction      entities.TransactionHistory
	Transactions     []entities.TransactionHistory
	TransactionStore store.TransactionHistoryModel
	Db               *sql.DB
}

func (suite *TransactionHistoryTestSuite) SetupTest() {
	var err error
	suite.Db = MySQLInit()
	transaction := store.NewTransactionHistoryStoreModel(suite.Db)
	suite.Transactions = []entities.TransactionHistory{
		{
			UserId:    1,
			AccountId: 1,
			Amount:    100,
			Action:    "Deposit",
		},
		{
			UserId:    1,
			AccountId: 1,
			Amount:    100,
			Action:    "Withdraw",
		},
		{
			UserId:    1,
			AccountId: 2,
			Amount:    44,
			Action:    "Deposit",
		},
		{
			UserId:    2,
			AccountId: 1,
			Amount:    741,
			Action:    "Withdraw",
		},
	}

	for i, current := range suite.Transactions {
		suite.Transaction, err = transaction.Insert(current.UserId, current.AccountId, current.Amount, current.Action)
		if err != nil {
			suite.T().Fatal("Unable to run InsertTransactionHistory store func")
		}
		suite.Transactions[i] = suite.Transaction
	}
}

func (suite *TransactionHistoryTestSuite) TestGetTransactionById() {
	store := store.NewTransactionHistoryStoreModel(suite.Db)
	var err error
	var transaction []entities.TransactionHistory
	var transactionsUserIDOne []entities.TransactionHistory
	var transactionsUserIDTwo []entities.TransactionHistory
	var transactionsUserIDFive []entities.TransactionHistory
	now := time.Now()

	transaction, err = store.GetTransactionsById(1)
	if err != nil {
		suite.T().Fatal("Unable to run GetTransactionsById store func")
	}
	for i, current := range suite.Transactions {
		current.CreatedAt = now
		if len(transaction) > i {
			transaction[i].CreatedAt = now
		}

		if current.UserId == 1 {
			transactionsUserIDOne = append(transactionsUserIDOne, current)
		}
	}
	suite.Equal(transactionsUserIDOne, transaction, "Users with ID 1 not equal")

	transaction, err = store.GetTransactionsById(2)
	if err != nil {
		suite.T().Fatal("Unable to run GetTransactionsById store func")
	}
	for i, current := range suite.Transactions {
		current.CreatedAt = now
		if len(transaction) > i {
			transaction[i].CreatedAt = now
		}
		if current.UserId == 2 {
			transactionsUserIDTwo = append(transactionsUserIDTwo, current)
		}
	}
	suite.Equal(transactionsUserIDTwo, transaction, "Users with ID 2 not equal")

	transaction, err = store.GetTransactionsById(5)
	if err != nil {
		suite.T().Fatal("Unable to run GetTransactionsById store func")
	}
	for i, current := range suite.Transactions {
		current.CreatedAt = now
		if len(transaction) > i {
			transaction[i].CreatedAt = now
		}
		if current.UserId == 5 {
			transactionsUserIDFive = append(transactionsUserIDFive, current)
		}
	}
	var emptyTransaction []entities.TransactionHistory
	suite.Equal(emptyTransaction, transaction, "GetTransactionById should return empty struct")
}

func (suite *TransactionHistoryTestSuite) TestGetTransactionByIdFromToDate() {
	store := store.NewTransactionHistoryStoreModel(suite.Db)
	var err error
	var transaction []entities.TransactionHistory
	var transactionsUserIDOne []entities.TransactionHistory
	var transactionsUserIDTwo []entities.TransactionHistory
	var transactionsUserIDFive []entities.TransactionHistory
	var now time.Time

	for _, current := range suite.Transactions {
		var k int
		for k = current.CreatedAt.Nanosecond(); k >= 10; k = k / 10 {
		}
		if k >= 5 {
			current.CreatedAt = suite.Transactions[0].CreatedAt.Add(time.Second * 1)
		}
		current.CreatedAt, _ = time.Parse("2006-01-02 15:04:05", current.CreatedAt.UTC().Format("2006-01-02 15:04:05"))
		now = current.CreatedAt
	}

	transaction, err = store.GetTransactionsByIdFromToDate(1, now.UTC(), now.UTC())
	if err != nil {
		suite.T().Fatal("Unable to run GetTransactionsById store func")
	}
	for i, current := range suite.Transactions {
		current.CreatedAt = now
		if len(transaction) > i {
			transaction[i].CreatedAt = now
		}

		if current.UserId == 1 {
			transactionsUserIDOne = append(transactionsUserIDOne, current)
		}
	}
	suite.Equal(transactionsUserIDOne, transaction, "Users with ID 1 not equal")

	transaction, err = store.GetTransactionsByIdFromToDate(2, now, now)
	if err != nil {
		suite.T().Fatal("Unable to run GetTransactionsById store func")
	}
	for i, current := range suite.Transactions {
		current.CreatedAt = now
		if len(transaction) > i {
			transaction[i].CreatedAt = now
		}
		if current.UserId == 2 {
			transactionsUserIDTwo = append(transactionsUserIDTwo, current)
		}
	}
	suite.Equal(transactionsUserIDTwo, transaction, "Users with ID 2 not equal")

	transaction, err = store.GetTransactionsByIdFromToDate(5, now, now)
	if err != nil {
		suite.T().Fatal("Unable to run GetTransactionsById store func")
	}
	for i, current := range suite.Transactions {
		current.CreatedAt = now
		if len(transaction) > i {
			transaction[i].CreatedAt = now
		}
		if current.UserId == 5 {
			transactionsUserIDFive = append(transactionsUserIDFive, current)
		}
	}
	suite.Equal(transactionsUserIDFive, transaction, "GetTransactionById should return empty struct")
}

func TestTransactionHistoryTestSuite(t *testing.T) {
	suite.Run(t, new(TransactionHistoryTestSuite))
}

func (suite *TransactionHistoryTestSuite) TearDownTest() {
	for i := 0; i < len(suite.Transactions); i++ {
		_, err := suite.Db.Exec("DELETE FROM TransactionHistory WHERE id=?", suite.Transactions[i].Id)
		if err != nil {
			suite.T().Fatal("Unable to run delete query")
		}
	}
}
