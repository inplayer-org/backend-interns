package entities

import "time"

type User struct {
	Id          int    `db:"id"`
	FullName    string `db:"full_name"`
	Email       string `db:"email"`
	PhoneNumber string `db:"phone_number"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

//type FullName struct {
//	FirstName string
//	LastName  string
//}
