package store

import (
	"bankacc/pkg/entities"
	"database/sql"
	"fmt"
	"time"
)

type TransactionHistoryStore interface {
	Insert(userId int, accountId int, amount float64, action string) (entities.TransactionHistory, error)
	GetTransactionsById(id int) ([]entities.TransactionHistory, error)
	GetTransactionsByIdFromToDate(id int, fromDate time.Time, toDate time.Time) ([]entities.TransactionHistory, error)
}

type TransactionHistoryModel struct {
	Db *sql.DB
}

func NewTransactionHistoryStoreModel(db *sql.DB) *TransactionHistoryModel {
	return &TransactionHistoryModel{
		Db: db,
	}
}

func (store *TransactionHistoryModel) Insert(userId int, accountId int, amount float64, action string) (entities.TransactionHistory, error) {
	now := time.Now()

	result, err := store.Db.Exec("INSERT INTO TransactionHistory(user_id, account_id, amount, action, created_at) VALUES(?, ?, ?, ?, ?)", userId, accountId, amount, action, now)
	if err != nil {
		fmt.Println(err)
	}
	res, err := result.LastInsertId()
	if err != nil {
		fmt.Println(err)
	}
	transaction := entities.TransactionHistory{
		Id:        int(res),
		UserId:    userId,
		AccountId: accountId,
		Amount:    amount,
		Action:    action,
		CreatedAt: now,
	}
	return transaction, nil
}

func (store *TransactionHistoryModel) GetTransactionsById(id int) ([]entities.TransactionHistory, error) {
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
	return transactions, nil
}

func (store *TransactionHistoryModel) GetTransactionsByIdFromToDate(id int, fromDate time.Time, toDate time.Time) ([]entities.TransactionHistory, error) {
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
	return transactions, nil
}
