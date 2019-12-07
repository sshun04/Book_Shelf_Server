package facebook

import (
	"bookstorage_web/server/config"
	"bookstorage_web/server/controller/auth"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)


var server auth.Server

var EndPoint = oauth2.Endpoint{
	AuthURL:  "https://www.facebook.com/dialog/oauth",
	TokenURL: "https://graph.facebook.com/oauth/access_token",
}

func Login(ctx *gin.Context) {
	conf, err := config.GetConfig()
	if err != nil {
		fmt.Println(err)
		return
	}
	authConf := conf.FaceBookAuth
	server = *auth.NewServer(authConf.ClientID, authConf.ClientSecret, authConf.CallBackUrl, EndPoint)
	server.Authorize(ctx)
}

func CallBack(ctx *gin.Context)  {
	server.Callback(ctx)
}
