package main

import (
	"bookstorage_web/server/auth"
	"bookstorage_web/server/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
	_ "github.com/mattn/go-sqlite3"
)

var dbDeviceName string = "sqlite3"
var dbFileName string = "users.sqlite3"

func main() {
	r := mux.NewRouter()
	// localhost:8080/publicでpublicハンドラーを実行
	r.Handle("/public", public)
	r.Handle("/private", auth.JwtMiddleware.Handler(private))
	r.Handle("/auth", auth.GetTokenHandler)

	//サーバー起動
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal("ListenAndServe:", nil)
	}
}


func dbInitUserTable() {
	db, err := gorm.Open(dbDeviceName, dbFileName)
	if err != nil {
		panic("failure open database:Init")
	}
	db.Table("active_users").AutoMigrate(&models.User{})
	defer db.Close()
}


// example of authentication using jwtmiddleware

var private = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	post := &models.Post{
		Title: "VGolangとGoogle Cloud Vision APIで画像から文字認識するCLIを速攻でつくる",
		Tag:   "Go",
		URL:   "https://qiita.com/po3rin/items/bf439424e38757c1e69b",
	}
	json.NewEncoder(w).Encode(post)
})

var public = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	post := &models.Post{
		Title: "VueCLIからVue.js入門①【VueCLIで出てくるファイルを概要図で理解】",
		Tag:   "Vue.js",
		URL:   "https://qiita.com/po3rin/items/3968f825f3c86f9c4e21",
	}
	json.NewEncoder(w).Encode(post)
})