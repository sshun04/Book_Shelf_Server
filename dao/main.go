package dao

import (
	"bookstorage_web/server/model"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

/**
データベースの共通ロジックをまとめたファイル
**/

const dbDeviceName string = "sqlite3"
const dbFileName string = "debug"

func GormConnect() (*gorm.DB, error) {
	return gorm.Open(dbDeviceName, dbFileName)
}

func DBInit() error {
	db, err := GormConnect()
	if err != nil {
		fmt.Println(err.Error())
		return err
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
	return nil
}

// データベースへの登録
func Create(dbModel interface{}) error {
	db, err := GormConnect()
	if err != nil {
		return err
	}
	defer db.Close()
	db.Create(dbModel)
	return nil
}

func GetUser()  {

}

func GetBooksById(ownerId uint)  {

}

// Where条件から取得
// GetWhere 条件から取得する
