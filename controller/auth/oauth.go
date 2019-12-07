package auth

import (
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"log"
	"math/rand"
	"net/http"
	"time"
)

const stateCookieName = "oauthState"

type Server struct {
	clientID     string
	clientSecret string
	callbackUrl  string
	endPoint     oauth2.Endpoint
}

func NewServer(clientId string, clientSecret string, callbackUrl string, endpoint oauth2.Endpoint) *Server {
	return &Server{
		clientID:     clientId,
		clientSecret: clientSecret,
		callbackUrl:  callbackUrl,
		endPoint:     endpoint,
	}
}

func (s *Server) GetConnect() *oauth2.Config {
	config := &oauth2.Config{
		ClientID:     s.clientID,
		ClientSecret: s.clientSecret,
		Endpoint: s.endPoint,
		RedirectURL: s.callbackUrl,
		Scopes:      []string{"email"},
	}

	return config
}

func (s *Server) Authorize(ctx *gin.Context) {
	state, err := generateState()
	if err != nil {
		s.writeError(ctx.Writer, http.StatusInternalServerError, err)
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
	http.SetCookie(ctx.Writer, cookie)
	//ctx.JSON(http.StatusOK,state)

	ctx.Redirect(http.StatusFound, url)
	ctx.JSON(http.StatusOK,url)
}

func (s *Server) Callback(ctx *gin.Context) {
	log.Printf("Callback: state=%v, code=%v", ctx.Request.FormValue("state"), ctx.Request.FormValue("code"))

	if e := ctx.Request.FormValue("error"); e != "" {
		//TODO　エラーの内容に応じたエラー画面の表示
		s.writeError(ctx.Writer, http.StatusBadRequest, fmt.Errorf("error returned in authorization: %v", e))
		return
	}

	if err := validateState(ctx.Request); err != nil {
		s.writeError(ctx.Writer, http.StatusBadRequest, err)
		return
	}

	code := ctx.Request.FormValue("code")
	if code == "" {
		s.writeError(ctx.Writer, http.StatusBadRequest, fmt.Errorf("code is required"))
		return
	}
	context, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	token := s.exchange(context, code)

	ctx.Writer.WriteHeader(http.StatusOK)

	// TODO　アクセストークンをレスポンス
	ctx.JSON(http.StatusOK,token.AccessToken)
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

func (s *Server) exchange(ctx context.Context, authCode string) *oauth2.Token {
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

func (s *Server) writeError(w http.ResponseWriter, code int, err error) {
	w.WriteHeader(code)
	fmt.Fprint(w, err.Error())
}
