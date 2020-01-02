package auth

import (
	"bookstorage_web/server/model"
	"fmt"
	"testing"
)

func TestGetJwtAccessToken(t *testing.T) {
	user := model.User{
		Name:"sh",
		Password:"sssssssss",
		EmailAddress:"ssss@gmail.com",
	}
	accessToken := GetJwtAccessToken(user)

	fmt.Println(accessToken)
}

func TestValidateAccessToken(t *testing.T) {

}
