package auth

import (
	"bookstorage_web/server/dao"
	"bookstorage_web/server/model"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)


func MustAuthenticated() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := ValidateAccessToken(c); err != nil {
			fmt.Println(err.Error())
			c.JSON(400, gin.H{"message": "Not authenticated"})
			c.Abort()
		}else {
			c.Next()
		}
	}
}

func SignUp(ctx *gin.Context) {
	var user model.User
	var error model.Error

	fmt.Println(ctx.Request.Body)

	err := json.NewDecoder(ctx.Request.Body).Decode(&user)
	if err != nil {
		fmt.Println("Error decoding request body")
	}

	if user.Name == "" {
		error.Message = "名前は必須です"
		ctx.Error(err).JSON()
		return
	}
	fmt.Println(user.Name)

	if user.EmailAddress == "" {
		error.Message = "Emailは必須です"
		ctx.Error(err).JSON()
		return
	}

	fmt.Println(user.EmailAddress)

	if user.Password == "" {
		error.Message = "パスワードは必須です。"
		ctx.Error(err).JSON()
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("パスワード: ", user.Password)
	fmt.Println("ハッシュ化されたパスワード", hash)

	user.Password = string(hash)
	fmt.Println("コンバート後のパスワード: ", user.Password)

	if savingerr := dao.Create(user, ""); savingerr != nil {
		fmt.Println(savingerr.Error())
		return
	}

	jwtAccessToken := GetJwtAccessToken(user)
	ctx.JSON(http.StatusOK, gin.H{"accessToken": jwtAccessToken})
}

func Login() {

}
