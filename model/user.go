package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name         string    `json:"user_name"`
	EmailAddress string    `json:"email_address"`
	Password     string    `json:"password"`
}
