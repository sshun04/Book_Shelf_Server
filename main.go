package main

import (
	"bookstorage_web/server/controller/auth"
	"bookstorage_web/server/controller/auth/facebook"
	"bookstorage_web/server/controller/auth/github"
	"bookstorage_web/server/controller/books"
	"bookstorage_web/server/dao"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	if err := dao.DBInit(); err != nil{
		fmt.Println(err.Error())
	}

	authorized := router.Group("/user")
	authorized.Use(auth.MustAuthenticated())
	{
		authorized.GET("/books",books.GetAll)
	}
	router.POST("/signup", auth.SignUp)
	router.POST("/facebook", facebook.Login)
	router.POST("/github", github.Login)
	router.POST("/auth/signup", auth.SignUp)
	router.POST("/auth/facebook", facebook.Login)
	router.GET("/auth/facebook/callback", facebook.CallBack)
	router.POST("/auth/github", github.Login)
	router.GET("/auth/github/callback", github.CallBack)

	router.Run(":8080")
}
