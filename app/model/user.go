package model

type User struct {
	UserId      string `json:"user_id,omitempty"`
	FirstName   string `json:"first_name,omitempty"`
	LastName    string `json:"last_name,omitempty"`
	Email       string `json:"email,omitempty"`
	CountryCode string `json:"country_code,omitempty"`
}
