package models

type User struct {
	ID        uint64
	Username  string
	Password  string
	Avatar    string
	Sex       int
	Phone     string
	Email     string
	Address   string
	Hobby     string
	Deleted   int
	CreatedAt int
	UpdatedAt int
}
