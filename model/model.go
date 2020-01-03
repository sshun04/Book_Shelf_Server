package model

import (
	"time"
)
type Error struct {
	Message string `json:"message"`
	error
}

type BaseModel struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
}

func (_ Error)Error() string {
	return ""
}
