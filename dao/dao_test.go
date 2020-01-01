package dao

import (
	"bookstorage_web/server/model"
	"testing"
)

func TestDBInit(t *testing.T) {
	db, err := GormConnect()
	if err != nil {
		t.Error(err.Error())
		return
	}
	db.AutoMigrate(&model.User{}, &model.BookCommon{}, &model.BookPersonal{})
	if !db.HasTable(&model.User{}) {
		db.CreateTable(&model.User{})
	}
	if !db.HasTable(&model.BookCommon{}) {
		db.CreateTable(&model.BookCommon{})
	}
	if !db.HasTable(&model.BookPersonal{}) {
		db.CreateTable(&model.BookPersonal{})
	}
	defer db.Close()
}

func TestGormConnect(t *testing.T) {

}

func TestCreate(t *testing.T) {
	if err := DBInit(); err != nil {
		t.Error(err.Error())
	}
	db, error := GormConnect()
	if error != nil {
		t.Error(error.Error())
		return
	}
	if err := db.Create(&model.User{
		Name:         "S",
		EmailAddress: "ss@gmail.com",
		Password:     "ss",}).Error; err != nil {
		t.Error("Error create User :", err.Error())
	}
	if err := db.Create(&model.BookCommon{
		ISBN:      11111,
		Title:     "学校のジュリエット",
		Author:    "ssss",
		Publisher: "sc.inc",
	}).Error; err != nil {
		t.Error("Error create BookCommon :", err.Error())
	}
	if err := db.Create(&model.BookPersonal{State: "未読"}).Error; err != nil {
		t.Error("Error create BookPersonal :", err.Error())
	}
	if err := db.Close(); err != nil {
		t.Error(err.Error())
	}
}
