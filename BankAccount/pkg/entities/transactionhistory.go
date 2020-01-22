package entities

import (
	"time"
)

type TransactionHistory struct {
	Id        int       `db: "id"`
	UserId    int       `db: "user_id"`
	AccountId int       `db:"account_id"`
	Amount    float64   `db: "amount"`
	Action    string    `db: "action"`
	CreatedAt time.Time `db: "created_at"`
}
