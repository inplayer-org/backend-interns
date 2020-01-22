package entities

type Account struct {
	Id        int     `db:"id"`
	UserId    int  	  `db:"user_id"`
	Balance   float64 `db:"balance"`
	Currency  string  `db:"currency"`
	Status    bool    `db:"status"`
	CreatedAt string  `db:"created_at"`
	UpdatedAt string  `db:"updated_at"`
}
