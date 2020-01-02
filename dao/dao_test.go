package dao

import (
	"bookstorage_web/server/model"
	"fmt"
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

	user := &model.User{
		Name:         "S",
		EmailAddress: "ss@gmail.com",
		Password:     "ss",}
	if err := db.Create(user).Error; err != nil {
		t.Error("Error create User :", err.Error())
	}

	user2 := &model.User{
		Name:         "k",
		EmailAddress: "kk@gmail.com",
		Password:     "kk",
	}
	if err := Create(user2); err != nil {
		t.Error(err.Error())
	}

	bookCommon := &model.BookCommon{
		ISBN:      11111,
		Title:     "学校のジュリエット",
		Author:    "ssss",
		Publisher: "sc.inc",
	}
	if err := db.Create(bookCommon).Error; err != nil {
		t.Error("Error create BookCommon :", err.Error())
	}

	bookPersonal := &model.BookPersonal{
		ISBN:    11111,
		OwnerId: 1,
		State:   "未読",
	}
	if err := db.Create(bookPersonal).Error; err != nil {
		t.Error("Error create BookPersonal :", err.Error())
	}

	if err := db.Close(); err != nil {
		t.Error(err.Error())
	}
}

func TestGetUser(t *testing.T) {
	db ,_:= GormConnect()
	var users []model.User
	table := db.Order("created_at desc").Find(&users)
	fmt.Print(table)
}
