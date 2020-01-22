package store

import (
	"database/sql"
	"time"

	"bankacc/pkg/entities"
)

type AccountStore interface {
	InsertAccount(userId int, balance float64, currency string, createdAt time.Time, updatedAt time.Time) (*entities.Account, error)
	GetAccountsByUserId(userId int) (*[]entities.Account, error)
	UpdateAccount(id int, userId int, balance float64, currency string, updatedAt time.Time) (*entities.User, error)
	CloseAccount(id int, userId int, updatedAt time.Time) (*entities.Account, error)
}

type AccountModel struct {
	Db *sql.DB
}

func NewAccountStoreModel(db *sql.DB) *AccountModel {
	return &AccountModel{
		Db: db,
	}
}

func (store *AccountModel) InsertAccount(userId int, balance float64, currency string, createdAt time.Time, updatedAt time.Time) (*entities.Account, error) {
	createdAt = time.Now()
	updatedAt = time.Now()
	account := entities.Account{
		UserId:    userId,
		Balance:   balance,
		Currency:  currency,
		Status:    true,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
	_, err := store.Db.Exec("INSERT INTO Account (user_id, balance, currency, status, created_at, updated_at) VALUES(?, ?, ?, 1, ?, ?)", userId, balance, currency, createdAt, updatedAt)
	if err != nil {
		return nil, err
	}
	return &account, nil
}

func (store *AccountModel) GetAccountsByUserId(userId int) (*[]entities.Account, error) {
	var accounts []entities.Account
	result, err := store.Db.Query("SELECT * FROM BankAccount.Account WHERE user_id=?", userId)
	if err != nil {
		return nil, err
	}
	var account entities.Account
	for result.Next() {
		err := result.Scan(&account.Id, &account.UserId, &account.Balance, &account.Currency, &account.Status, &account.CreatedAt, &account.UpdatedAt)
		if err != nil {
			return nil, err
		}
		accounts = append(accounts, account)
	}
	return &accounts, nil
}

func (store *AccountModel) UpdateAccount(id int, userId int, balance float64, currency string, updatedAt time.Time) (*entities.User, error) {
	updatedAt = time.Now()
	_, err := store.Db.Exec("UPDATE BankAccount.Account SET balance = ?, currency =?, updated_at = ? WHERE user_id = ? AND id =?", balance, currency, updatedAt, userId, id)
	if err != nil {
		return nil, err
	}
	return nil, err
}

func (store *AccountModel) CloseAccount(id int, userId int, updatedAt time.Time) (*entities.Account, error) {
		updatedAt = time.Now()
	_, err := store.Db.Exec("UPDATE BankAccount.Account SET status = 0, updated_at = ? WHERE user_id =  ? AND id = ?", updatedAt, userId, id)
	if err != nil {
		return nil, err
	}
	return nil, err
}
