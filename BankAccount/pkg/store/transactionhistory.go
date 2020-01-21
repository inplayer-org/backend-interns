package store

import (
	"bankacc/pkg/config"
	"bankacc/pkg/entities"
	"database/sql"
)

type TransactionHistoryStore interface {
	Insert(UserId int, AccountId int, Amount float64, Action string, CreatedAt string) (*entities.TransactionHistory, error)
	GetTransactionsById(Id int) (*[]entities.TransactionHistory, error)
	GetTransactionsByIdFromToDate(Id int, FromDate string, ToDate string) (*[]entities.TransactionHistory, error)
}

type TransactionHistoryModel struct {
	Db *sql.DB
}

func NewTransactionHistoryStoreModel(db *sql.DB) *TransactionHistoryModel {
	return &TransactionHistoryModel{
		Db: db,
	}
}

func (store *TransactionHistoryModel) Insert(UserId int, AccountId int, Amount float64, Action string, CreatedAt string) (*entities.TransactionHistory, error) {
	db, err := config.GetMySQLDB()
	if err != nil {
		return nil, err
	} else {
		TransactionModel := TransactionHistoryModel{
			Db: db,
		}

		transaction := entities.TransactionHistory{
			UserId:    UserId,
			AccountId: AccountId,
			Amount:    Amount,
			Action:    Action,
			CreatedAt: CreatedAt,
		}
		_, err := TransactionModel.Db.Exec("INSERT INTO TransactionHistory(user_id, account_id, amount, action, created_at) VALUES(?, ?, ?, ?, ?)", UserId, AccountId, Amount, Action, CreatedAt)

		if err != nil {
			return nil, err
		}
		return &transaction, nil
	}
}

func (store *TransactionHistoryModel) GetTransactionsById(Id int) (*[]entities.TransactionHistory, error) {
	var transactions []entities.TransactionHistory
	db, err := config.GetMySQLDB()
	if err != nil {
		return nil, err
	} else {
		TransactionModel := TransactionHistoryModel{
			Db: db,
		}
		result, err := TransactionModel.Db.Query("SELECT * FROM BankAccount.TransactionHistory WHERE user_id = ?", Id)
		if err != nil {
			return nil, err
		}
		var transaction entities.TransactionHistory
		for result.Next() {
			err := result.Scan(&transaction.Id, &transaction.UserId, &transaction.AccountId, &transaction.Amount, &transaction.Action, &transaction.CreatedAt)
			if err != nil {
				return nil, err
			}
			transactions = append(transactions, transaction)
		}
		return &transactions, nil
	}
}

func (store *TransactionHistoryModel) GetTransactionsByIdFromToDate(Id int, FromDate string, ToDate string) (*[]entities.TransactionHistory, error) {
	var transactions []entities.TransactionHistory
	db, err := config.GetMySQLDB()
	if err != nil {
		return nil, err
	} else {
		TransactionModel := TransactionHistoryModel{
			Db: db,
		}
		result, err := TransactionModel.Db.Query("SELECT * FROM BankAccount.TransactionHistory WHERE user_id = ? and created_at between ? and ?", Id, FromDate, ToDate)
		if err != nil {
			return nil, err
		}
		var transaction entities.TransactionHistory
		for result.Next() {
			err := result.Scan(&transaction.Id, &transaction.UserId, &transaction.AccountId, &transaction.Amount, &transaction.Action, &transaction.CreatedAt)
			if err != nil {
				return nil, err
			}
			transactions = append(transactions, transaction)
		}
		return &transactions, nil
	}
}
