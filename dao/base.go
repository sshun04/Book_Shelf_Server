package dao

import "github.com/jinzhu/gorm"

/**
データベースの共通ロジックをまとめたファイル
**/

const dbDeviceName string = "sqlite3"
const dbFileName string = "test"

func GormConnect() (*gorm.DB,error)  {
	return gorm.Open(dbDeviceName,dbFileName)
}

func DBInit(tableName string, model interface{})  {
	db,err := GormConnect()
	if err != nil {

	}

	if !db.HasTable(model) {

	}


	defer db.Close()
}

// データベースへの登録
func Create(dbModel interface{}, tableName string) error {
	db, err := GormConnect()
	if err != nil {
		return err
	}
	defer db.Close()
	return db.Table(tableName).Create(&dbModel).Error
}

// Where条件から取得
// GetWhere 条件から取得する
