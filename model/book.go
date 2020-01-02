package model

import "github.com/jinzhu/gorm"

type BookCommon struct {
	gorm.Model
	ISBN      int    `json:"isbn"`
	Title     string `json:"title"`
	Author    string `json:"author"`
	Publisher string `json:"publisher"`
}

type BookPersonal struct {
	gorm.Model
	OwnerId uint   `json:"owner_id"`
	ISBN    int    `json:"isbn"`
	State   string `json:"state"`
}
