package model

type User struct {
	UserId      string `uri:"userid"`
	FirstName   string
	SecondName  string
	Email       string
	CountryCode string
}
