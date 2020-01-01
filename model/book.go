package model

import "github.com/jinzhu/gorm"


type BookCommon struct {
	gorm.Model
	ISBN      int    `json:"isbn"`
	Title     string `json:"title"`
	Author    string `json:"author"`
	Publisher string  `json:"publisher"`
}

type BookPersonal struct {
	gorm.Model
	State string `json:"state"`
}
