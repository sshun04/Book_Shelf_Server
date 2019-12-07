package main

import (
	"bookstorage_web/server/controller/auth"
	"bookstorage_web/server/controller/auth/facebook"
	"bookstorage_web/server/controller/auth/github"
	"github.com/gin-gonic/gin"
)



func main() {
	router := gin.Default()
	router.POST("/auth/signup", auth.SignUp)
	router.POST("/auth/facebook",facebook.Login)
	router.GET("/auth/facebook/callback",facebook.CallBack)
	router.POST("/auth/github",github.Login)
	router.GET("/auth/github/callback",github.CallBack)

	router.Run(":8080")
}
