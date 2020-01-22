package store

import (
	"bankacc/pkg/entities"
	"database/sql"
)

type TransactionHistoryStore interface {
	Insert(userId int, accountId int, amount float64, action string, createdAt string) (*entities.TransactionHistory, error)
	GetTransactionsById(id int) (*[]entities.TransactionHistory, error)
	GetTransactionsByIdFromToDate(id int, fromDate string, toDate string) (*[]entities.TransactionHistory, error)
}

type TransactionHistoryModel struct {
	Db *sql.DB
}

func NewTransactionHistoryStoreModel(db *sql.DB) *TransactionHistoryModel {
	return &TransactionHistoryModel{
		Db: db,
	}
}

func (store *TransactionHistoryModel) Insert(userId int, accountId int, amount float64, action string, createdAt string) (*entities.TransactionHistory, error) {
	transaction := entities.TransactionHistory{
		UserId:    userId,
		AccountId: accountId,
		Amount:    amount,
		Action:    action,
		CreatedAt: createdAt,
	}
	_, err := store.Db.Exec("INSERT INTO TransactionHistory(user_id, account_id, amount, action, created_at) VALUES(?, ?, ?, ?, ?)", userId, accountId, amount, action, createdAt)

	if err != nil {
		return nil, err
	}
	return &transaction, nil
}

func (store *TransactionHistoryModel) GetTransactionsById(id int) (*[]entities.TransactionHistory, error) {
	var transactions []entities.TransactionHistory
	result, err := store.Db.Query("SELECT * FROM TransactionHistory WHERE user_id = ?", id)
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

func (store *TransactionHistoryModel) GetTransactionsByIdFromToDate(id int, fromDate string, toDate string) (*[]entities.TransactionHistory, error) {
	var transactions []entities.TransactionHistory
	result, err := store.Db.Query("SELECT * FROM TransactionHistory WHERE user_id = ? and created_at BETWEEN ? and ?", id, fromDate, toDate)
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
