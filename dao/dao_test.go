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
	db.AutoMigrate(&model.User{})
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
	e := db.Create(&model.User{
		Name:         "Shu",
		EmailAddress: "sksks@gmail.com",
		Password:     "skskskksks",}).Error

	if e != nil {
		t.Error(e.Error())
	}

	if err := db.Close(); err != nil {
		t.Error(err.Error())
	}
}
