package dao

import (
	"bookstorage_web/server/model"
	"fmt"
	"golang.org/x/crypto/bcrypt"
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
	db, err := GormConnect()
	if err != nil {
		t.Error(err.Error())
		return
	}

	user := &model.User{
		Name:         "S",
		EmailAddress: "ss@gmail.com",
		Password:     "ss",}
	if err := db.Create(&user).Error; err != nil {
		t.Error("Error create User :", err.Error())
	}

	user2 := &model.User{
		Name:         "k",
		EmailAddress: "kk@gmail.com",
		Password:     "kk",
	}
	if err := Create(&user2,"users"); err != nil {
		t.Error(err.Error())
	}

	user3 := &model.User{
		Name:"c",
		EmailAddress:"cc@gmail.com",
		Password:"cc",
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user3.Password), 10)
	if err != nil {
		t.Error(err.Error())
	}
	user3.Password = string(hashedPassword)

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
	defer db.Close()
}

func TestGetUser(t *testing.T) {
	db, _ := GormConnect()
	var users []model.User
	table := db.Order("created_at desc").Find(&users)
	fmt.Print(table)
	defer db.Close()
}

func TestSearchUserByEmail(t *testing.T) {
	if user, err := SearchUserByEmail("kk@gmail.com"); err != nil {
		t.Error(err.Error())
	} else if user.Name == "" && user.EmailAddress == "" && user.Password == "" {
		t.Error("user is blank")
	} else {
		fmt.Println(user.Name)
	}
}

func TestSearchUser(t *testing.T) {
	user1 := model.User{EmailAddress: "kk@gmail.com", Password: "kk"}
	if SearchUser(user1) {
		fmt.Println("got an appropriate user")
	} else {
		t.Error("Error Could not get an appropriate user ")
	}

	user2 := model.User{EmailAddress: "kk@gmail.com", Password: "aa"}
	if  SearchUser(user2) {
		t.Error("expect get an error but didn't ")
	} else {
		fmt.Println("got appropriate result")
	}

	user3 := model.User{EmailAddress:"aaaaaa@gmail.com",Password:"kk"}
	if  SearchUser(user3) {
		t.Error("expect get an error but didn't ")
	} else {
		fmt.Println("got appropriate result")
	}

	user4 := model.User{EmailAddress:"cc@gmail.com",Password:"cc"}
	if  SearchUser(user4) {
		t.Error("expect get an error but didn't ")
	} else {
		fmt.Println("got appropriate result")
	}
}
