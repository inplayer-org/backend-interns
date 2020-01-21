package store

import (
	"bankacc/pkg/config"
	"bankacc/pkg/entities"
	"database/sql"
)

type TransactionHistoryStore interface {
	Insert(UserId int, AccountId int, Amount float64, Action string, Created string) (*entities.TransactionHistory, error)
}

type TransactionHistoryModel struct {
	Db *sql.DB
}

func NewTransactionHistoryStoreModel(db *sql.DB) *TransactionHistoryModel {
	return &TransactionHistoryModel{
		Db: db,
	}
}

func (store *TransactionHistoryModel) Insert(UserId int, AccountId int, Amount float64, Action string, Created string) (*entities.TransactionHistory, error) {
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
			Created:   Created,
		}
		_, err := TransactionModel.Db.Exec("INSERT INTO TransactionHistory(user_id, account_id, amount, action, created_at) VALUES(?, ?, ?, ?, ?)", UserId, AccountId, Amount, Action, Created)

		if err != nil {
			return nil, err
		}
		return &transaction, nil
	}
}
