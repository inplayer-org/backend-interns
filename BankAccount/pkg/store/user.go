package store

import (
	"bankacc/pkg/entities"
	"database/sql"
)

type UserStore interface {
	Insert(FullName string, email string, PhoneNumber string, CreatedAt string, UpdatedAt string) (*entities.User, error)
	GetUserById(Id int) (*entities.User, error)
}

type UserModel struct {
	Db *sql.DB
}

func NewUserStoreModel(db *sql.DB) *UserModel {
	return &UserModel{
		Db: db,
	}
}

func (store *UserModel) Insert(FullName string, email string, PhoneNumber string, CreatedAt string, UpdatedAt string) (*entities.User, error) {

	user := entities.User{
		FullName:    FullName,
		Email:       email,
		PhoneNumber: PhoneNumber,
		CreatedAt:   CreatedAt,
		UpdatedAt:   UpdatedAt,
	}
	_, err := store.Db.Exec("INSERT INTO User (full_name, email, phone_number, created_at, updated_at) VALUES(?, ?, ?, ?, ?)", FullName, email, PhoneNumber, CreatedAt, UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (store *UserModel) GetUserById(Id int) (*[]entities.User, error) {
	var users []entities.User
	result, err := store.Db.Query("SELECT * FROM BankAccount.User WHERE id=?", Id)
	if err != nil {
		return nil, err
	}
	var user entities.User
	for result.Next() {
		err := result.Scan(&user.Id, &user.FullName, &user.Email, &user.PhoneNumber, &user.UpdatedAt, &user.CreatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return &users, nil
}
