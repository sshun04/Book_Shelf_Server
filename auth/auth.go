package auth

import (
	"bookstorage_web/server/conf"
	"encoding/base64"
	"fmt"
	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

const stateCookieName = "oauthState"

type server struct {
	clientID          string
	clientSecret      string
	authorizeEndPoint string
	tokenEndPoint     string
}

func NewServer(clientId string, clientSecret string, authorizeEndPoint string, tokenEndPoint string) *server {
	return &server{
		clientID:          clientId,
		clientSecret:      clientSecret,
		authorizeEndPoint: authorizeEndPoint,
		tokenEndPoint:     tokenEndPoint,
	}
}

func (s *server) GetConnect() *oauth2.Config {
	authInfo := conf.GetConfig().Auth
	config := &oauth2.Config{
		ClientID:     s.clientID,
		ClientSecret: s.clientSecret,
		Endpoint: oauth2.Endpoint{
			AuthURL:  s.authorizeEndPoint,
			TokenURL: s.tokenEndPoint,
		},
		RedirectURL: authInfo.CallBackUrl,
		Scopes:      []string{"email"},
	}

	return config
}

func (s *server) authorize(w http.ResponseWriter, r *http.Request) {
	state, err := generateState()
	if err != nil {
		s.writeError(w, http.StatusInternalServerError, err)
		return
	}
	authConfig := s.GetConnect()
	url := authConfig.AuthCodeURL(state)

	cookie := &http.Cookie{
		Name:     stateCookieName,
		Value:    state,
		Path:     "/",
		Expires:  time.Now().Add(10 * time.Minute),
		HttpOnly: true,
	}
	http.SetCookie(w, cookie)

	http.Redirect(w, r, url, http.StatusFound)
}

func (s *server) callback(w http.ResponseWriter, r *http.Request) {
	log.Printf("callback: state=%v, code=%v", r.FormValue("state"), r.FormValue("code"))

	if e := r.FormValue("error"); e != "" {
		//TODO　エラーの内容に応じたエラー画面の表示
		s.writeError(w, http.StatusBadRequest, fmt.Errorf("error returned in authorization: %v", e))
		return
	}

	if err := validateState(r); err != nil {
		s.writeError(w, http.StatusBadRequest, err)
		return
	}

	code := r.URL.String()
	if code == "" {
		s.writeError(w, http.StatusBadRequest, fmt.Errorf("code is required"))
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	token := s.exchange(ctx, code)

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "accessToken = %v", token)

	// TODO　アクセストークンをレスポンス
	w.Write([]byte(token.AccessToken))
}

func validateState(r *http.Request) error {
	state := r.FormValue("state")
	oauthState, err := r.Cookie(stateCookieName)
	if err != nil {
		return fmt.Errorf("failed to get cookie %v", stateCookieName)
	}
	if state != oauthState.Value {
		return fmt.Errorf("state doesn't match")
	}
	return nil
}

func (s *server) exchange(ctx context.Context, authCode string) *oauth2.Token {
	authConfig := s.GetConnect()
	token, err := authConfig.Exchange(context.Background(), authCode)
	if err != nil {
		log.Fatal(err)
	}

	return token
}

func generateState() (string, error) {
	b := make([]byte, 64)
	if _, err := rand.Read(b); err != nil {
		return "", fmt.Errorf("failed to generate state: %v", err)
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

func (s *server) writeError(w http.ResponseWriter, code int, err error) {
	w.WriteHeader(code)
	fmt.Fprint(w, err.Error())
}



//　JWTのお遊び用コード
var GetTokenHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// headerのセット
	token := jwt.New(jwt.SigningMethodHS256)

	// claimsのセット
	claims := token.Claims.(jwt.MapClaims)
	claims["admin"] = true
	claims["sub"] = "54546557354"
	claims["name"] = "shun"
	claims["iat"] = time.Now()
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	// 電子署名
	tokenString, _ := token.SignedString([]byte(os.Getenv("SIGNINGKEY")))

	// JWTを返却
	w.Write([]byte(tokenString))
})

var JwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SIGNINGKEY")), nil
	},
	SigningMethod: jwt.SigningMethodHS256,
})
