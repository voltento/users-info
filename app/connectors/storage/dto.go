package storage

type User struct {
	UserId      int `sql:",pk"`
	FirstName   string
	LastName    string
	Email       string
	CountryCode string
}
