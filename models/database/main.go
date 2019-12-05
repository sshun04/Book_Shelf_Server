package database

import (
	"github.com/jinzhu/gorm"
)

const dbDeviceName string = "sqlite3"
const dbFileName  string = "test.sqlite3"

var db *gorm.DB
var transactionDB *gorm.DB

func DBInit(model interface{})  {
	db, err := gorm.Open(dbDeviceName, dbFileName)
	if err != nil {
		panic("failure open database:Init")
	}
	db.AutoMigrate(&model)
	defer db.Close()
}


func GormConnect() (*gorm.DB,error) {
	if transactionDB != nil {
		return transactionDB, nil
	}
	if db != nil {
		return db, nil
	}

	db, err := gorm.Open(dbDeviceName, dbFileName)

	if err != nil {
		return db, err
	}

	return db, nil
}

// Transaction トランザクション
func Transaction(db *gorm.DB) {
	transactionDB = db
}