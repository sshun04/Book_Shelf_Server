package auth

import (
	"bookstorage_web/server/model"
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func TestSignUp(t *testing.T) {
	var user model.User
	demoUserJson := []byte(`{"user_name":"Alice","email_address":"alice@gmail.com","password":"passdesuyo"}`)

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
		fmt.Println("password: " + user.Password)
	}

	hashedPassWord := hashStringPassWord(user.Password)
	user.Password = hashedPassWord

	fmt.Println("hashed password: "+user.Password)

	jwtAccessToken := GetJwtAccessToken(user)

	fmt.Println("jwtAccessToken: "+jwtAccessToken)

}

func TestHashStringPassWord(t *testing.T) {
	demoPassword := "kskskksks"
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(demoPassword), 10)
	if err != nil {
		t.Error(err.Error())
	}
	fmt.Println("パスワード: ", demoPassword)
	fmt.Println("ハッシュ化されたパスワード",hashedPassword)
	fmt.Println("コンバート後のパスワード: ",hashedPassword)
}
