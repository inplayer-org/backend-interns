package entities


type User struct {
	Id          string    `db:"id"`
	FullName    *FullName `db:"full_name"`
	Email       string    `db:"email"`
	PhoneNumber string    `db:"phone_number"`
	Created     string    `db:"created_at"`
	Updated     string    `db:"updated_at"`

}

type FullName struct {
	FirstName string
	LastName  string
}
