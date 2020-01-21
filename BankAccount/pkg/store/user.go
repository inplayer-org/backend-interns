package store

import (
	"bankacc/pkg/config"
	"bankacc/pkg/entities"
	"database/sql"
)

type UserStore interface {
	Insert(full_name string, email string, phone_number string, created_at string, updated_at string) (*entities.User, error)
}
type UserModel struct {
	Db *sql.DB
}

func NewUserStoreModel(db *sql.DB) *UserModel {
	return &UserModel{
		Db: db,
	}
}
func (store *UserModel) Insert(full_name string, email string, phone_number string, created_at string, updated_at string) (*entities.User, error) {
	db, err := config.GetMySQLDB()
	if err != nil {
		return nil, err
	} else {
		UModel := UserModel{
			Db: db,
		}
		user := entities.User{
			FullName:    full_name,
			Email:       email,
			PhoneNumber: phone_number,
			Created:     created_at,
			Updated:     updated_at,
		}
		_, err := UModel.Db.Exec("INSERT INTO User (full_name, email, phone_number, created_at, updated_at) VALUES(?, ?, ?, ?, ?)",full_name, email, phone_number, created_at, updated_at)
		if err != nil {
			return nil, err
		}
		return &user, nil
	}

}
