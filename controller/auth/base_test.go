package auth

import (
	"bookstorage_web/server/model"
	"encoding/json"
	"fmt"
	"testing"
)

func TestSignUp(t *testing.T) {
	var user model.User
	demoUserJson := []byte(`{"name":"Alice","email_address":"alice@gmail.com","password":"passdesuyo"}`)

	if err := json.Unmarshal(demoUserJson, &user); err != nil {
		t.Error(err.Error())
	}

	if user.Name == "" {
		t.Error("error user name is blank")
	} else {
		fmt.Println("name: " + user.Name)
	}

	if user.EmailAddress == "" {
		t.Error("error user email_address is blank")
	}else {
		fmt.Println("emailAddress: "+ user.EmailAddress)
	}

	if user.Password == "" {
		t.Error("error user password is blank")
	}else {
		fmt.Println("password" + user.Password)
	}
}

func TestHashStringPassWord(t *testing.T) {
	demoPassword := "kskskksks"
	hashedPassWord := HashStringPassWord(demoPassword)
	fmt.Println("パスワード: ", demoPassword)
	fmt.Println("ハッシュ化されたパスワード", hashedPassWord)
	fmt.Println("コンバート後のパスワード: ", hashedPassWord)
}
