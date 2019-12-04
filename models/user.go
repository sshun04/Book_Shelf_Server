package models

type User struct {
	BaseModel    `model:"true"`
	Name         string    `json:"user_name"`
	EmailAddress string    `json:"email_address"`
	Password     string    `json:"password"`
}
