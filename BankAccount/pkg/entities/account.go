package entities

import "time"

type Account struct {
	Id        int     `db:"id"`
	UserId    int  	  `db:"user_id"`
	Balance   float64 `db:"balance"`
	Currency  string  `db:"currency"`
	Status    bool    `db:"status"`
	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt time.Time  `db:"updated_at"`
}
