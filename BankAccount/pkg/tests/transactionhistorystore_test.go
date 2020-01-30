package tests

import (
	"bankacc/pkg/entities"
	"bankacc/pkg/store"
	"database/sql"
	"fmt"
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
	Transaction      *entities.TransactionHistory
	Transactions     []entities.TransactionHistory
	TransactionsP    []*entities.TransactionHistory
	TransactionStore store.TransactionHistoryModel
	TransactionID    []int
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
			Amount:    65,
			Action:    "Deposit",
		},
		{
			UserId:    2,
			AccountId: 2,
			Amount:    44,
			Action:    "Deposit",
		},
	}

	for _, value := range suite.Transactions {
		suite.Transaction, err = transaction.Insert(value.UserId, value.AccountId, value.Amount, value.Action)
		if err != nil {
			suite.T().Fatal("Unable to run InsertTransactionHistory store func")
		}
		suite.TransactionsP = append(suite.TransactionsP, suite.Transaction)
		suite.TransactionID = []int{1, 2}
	}
}

func (suite *TransactionHistoryTestSuite) TestGetTransactionById() {
	store := store.NewTransactionHistoryStoreModel(suite.Db)
	var err error
	var transact []*entities.TransactionHistory
	//for _, value := range suite.TransactionID {
	transact, err = store.GetTransactionsById(1)
	if err != nil {
		suite.T().Fatal("Unable to run GetTransactionsById store func")
	}

	fmt.Println(transact)

	//}

	for _, value := range suite.TransactionsP {
		for i := range transact {
			if value.Id == transact[i].Id {
				suite.Equal(value.UserId, transact[i].UserId)
				suite.Equal(value.AccountId, transact[i].AccountId)
				suite.Equal(1, transact[i].Action)
				suite.Equal(value.Amount, transact[i].Amount)
			}
		}
	}
}

func (suite *TransactionHistoryTestSuite) TestGetTransactionByIdFromToDate() {
	store := store.NewTransactionHistoryStoreModel(suite.Db)
	var err error
	var transact []*entities.TransactionHistory
	for _, value := range suite.TransactionsP {
		var k int
		for k = value.CreatedAt.Nanosecond(); k >= 10; k = k / 10 {
		}
		if k >= 5 {
			value.CreatedAt = value.CreatedAt.Add(time.Second * 1)
		}
		value.CreatedAt, _ = time.Parse("2006-01-02 15:04:05", value.CreatedAt.UTC().Format("2006-01-02 15:04:05"))
		transact, err = store.GetTransactionsByIdFromToDate(value.UserId, value.CreatedAt.UTC(), value.CreatedAt.UTC())
		if err != nil {
			suite.T().Fatal("Unable to run GetTransactionsByIdFromDate store func")
		}
	}

	for _, value := range suite.TransactionsP {
		for i := range transact {
			if value.Id == transact[i].Id {
				suite.Equal(value.UserId, transact[i].UserId)
				suite.Equal(value.AccountId, transact[i].AccountId)
				suite.Equal(value.Action, transact[i].Action)
				suite.Equal(value.Amount, transact[i].Amount)
			}
		}
	}
}

func TestTransactionHistoryTestSuite(t *testing.T) {
	suite.Run(t, new(TransactionHistoryTestSuite))
}

func (suite *TransactionHistoryTestSuite) TearDownTest() {
	for i := 0; i < len(suite.TransactionsP); i++ {
		_, err := suite.Db.Exec("DELETE FROM TransactionHistory WHERE id=?", suite.TransactionsP[i].Id)
		if err != nil {
			suite.T().Fatal("Unable to run delete query")
		}
	}
}