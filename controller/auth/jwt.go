package auth

import (
	"bookstorage_web/server/config"
	"bookstorage_web/server/model"
	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"os"
	"time"
)

func GetJwtAccessToken(user model.User) string{
	token := jwt.New(jwt.SigningMethodHS256)

	// claimsのセット
	claims := token.Claims.(jwt.MapClaims)
	claims["admin"] = true
	claims["sub"] = user.ID
	claims["email"] = user.EmailAddress
	claims["name"] = user.Name
	claims["iat"] = time.Now()
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	// 電子署名
	conf, _ := config.GetConfig()
	tokenString, _ := token.SignedString([]byte(os.Getenv(conf.JwtConfig.SignInKey)))

	return tokenString
}

func ValidateAccessToken(c *gin.Context) error {
	return JwtMiddleware.CheckJWT(c.Writer,c.Request)
}

var JwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
		conf, _ := config.GetConfig()
		return []byte(os.Getenv(conf.JwtConfig.SignInKey)), nil
	},
	SigningMethod: jwt.SigningMethodHS256,
})