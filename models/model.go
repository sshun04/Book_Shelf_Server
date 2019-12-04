package models

/**
データベースに登録されるモデルの共通ロジック
をまとめたファイル
**/

import (
	"bookstorage_web/server/models/database"
	"github.com/jinzhu/gorm"
	"strings"
	"time"
)

// ErrRecordeNotFound レコードなし
const ErrRecordeNotFound = "record not found"

// ErrFileTypeUnMatch レコードなし
const ErrFileTypeUnMatch = "file type unmatch"

const fileName string = "fileName"

type BaseModel struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
}

// getBindAndPlaceHolder バインドとプレースホルダの結果を取得する
func getBindAndPlaceHolder(where string, bindList []map[string]interface{}) (string, []interface{}) {
	bind := []interface{}{}
	var holder string

	for _, list := range bindList {
		for key, value := range list {

			switch value := value.(type) {
			// 配列対応
			case []int:
				holder = " ("

				i := 0
				for _, data := range value {
					if i > 0 {
						holder += ", "
					}
					holder += "?"
					bind = append(bind, data)
					i++
				}

				holder += ") "
			default:
				holder = "?"
				bind = append(bind, value)
			}

			where = strings.Replace(where, ":"+key, holder, 1)
		}
	}

	return where, bind
}

func getDbOption(where string, bindList []map[string]interface{}, option map[string]interface{}) (*gorm.DB, error) {
	db, err := database.GormConnect()
	if err != nil {
		return db, err
	}

	if where != "" {
		w, bind := getBindAndPlaceHolder(where, bindList)
		db = db.Where(w, bind...)
	}

	if order, ok := option["order"].(string); ok {
		db = db.Order(order)
	}

	if limit, ok := option["limit"].(int); ok {
		db = db.Limit(limit)
	}

	if offset, ok := option["offset"].(int); ok {
		db = db.Offset(offset)
	}

	if sel, ok := option["select"].(string); ok {
		db = db.Select(sel)
	}

	return db, nil
}

func checkError(err error) error {
	if err == nil {
		return nil
	}

	if err.Error() == ErrRecordeNotFound {
		return nil
	}

	if err.Error() == ErrFileTypeUnMatch {
		return nil
	}

	return err
}

// データベースへの登録
func Create(dbModel interface{}, tableName string) error {
	db, err := database.GormConnect()
	if err != nil {
		return err
	}
	return db.Table(tableName).Create(dbModel).Error
}

// Where条件から取得
// GetWhere 条件から取得する
func GetWhere(dbModel interface{}, where string, bindList []map[string]interface{}, option map[string]interface{}) (*gorm.DB, error) {
	db, err := getDbOption(where, bindList, option)
	if err != nil {
		return db, err
	}

	err = db.First(dbModel).Error
	if err = checkError(err); err != nil {
		return db, nil
	}

	return db, err
}
