package store

import (
	"bankacc/pkg/entities"
	"database/sql"
	"time"
)

type UserStore interface {
	Insert(fullName string, email string, phoneNumber string) (*entities.User, error)
	GetUserById(id int) (*entities.User, error)
	UpdateUser(id int, fullName string, email string, phoneNumber string) (*entities.User, error)
	DeleteUser(id int) (*entities.User, error)
}

type UserModel struct {
	Db *sql.DB
}

func NewUserStoreModel(db *sql.DB) *UserModel {
	return &UserModel{
		Db: db,
	}
}

func (store *UserModel) InsertUser(fullName string, email string, phoneNumber string) (*entities.User, error) {
	createdAt := time.Now()
	updatedAt := time.Now()
	user := entities.User{
		FullName:    fullName,
		Email:       email,
		PhoneNumber: phoneNumber,
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
	}
	_, err := store.Db.Exec("INSERT INTO User (full_name, email, phone_number, created_at, updated_at) VALUES(?, ?, ?, ?, ?)", fullName, email, phoneNumber, createdAt, updatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (store *UserModel) GetUserById(id int) (*[]entities.User, error) {
	var users []entities.User
	result, err := store.Db.Query("SELECT * FROM BankAccount.User WHERE id=?", id)
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

func (store *UserModel) UpdateUser(id int, fullName string, email string, phoneNumber string) (*entities.User, error) {
	updatedAt := time.Now()
	_, err := store.Db.Exec("UPDATE BankAccount.User SET full_name = ?, email = ?, phone_number = ?, updated_at =? WHERE id = ?", fullName, email, phoneNumber, updatedAt, id)
	if err != nil {
		return nil, err
	}
	return nil, err
}

func (store *UserModel) DeleteUser(id int) (*entities.User, error) {
	_, err := store.Db.Exec("DELETE FROM BankAccount.User WHERE id=?", id)
	if err != nil {
		return nil, err
	}
	return nil, err
}
