package facebook
//
//import (
//	"bookstorage_web/server/conf"
//	"golang.org/x/oauth2"
//)
//
//const (
//	authorizeEndpoint = "https://www.facebook.com/dialog/oauth"
//	tokenEndpoint     = "https://graph.facebook.com/oauth/access_token"
//)
//
//func GetConnect() *oauth2.Config {
//	authInfo := conf.GetConfig().Auth
//	config := &oauth2.Config{
//		ClientID:     authInfo.FaceBookClientID,
//		ClientSecret: authInfo.FaceBookClientSecret,
//		Endpoint: oauth2.Endpoint{
//			AuthURL:  authorizeEndpoint,
//			TokenURL: tokenEndpoint,
//		},
//		RedirectURL: authInfo.CallBackUrl + "apiのurl",
//		Scopes:      []string{"email"},
//	}
//
//	return config
//}
