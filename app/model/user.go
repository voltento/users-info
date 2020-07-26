package model

type User struct {
	UserId      string `uri:"userid" binding:"required"`
	FirstName   string
	SecondName  string
	Email       string
	CountryCode string
}
