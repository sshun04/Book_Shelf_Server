package config

import (
	"testing"
)

func TestGetConfig(t *testing.T) {
	config, err := GetConfig()
	if err != nil {
		t.Error(err.Error())
	}
	faceBookConfig := config.FaceBookAuth
	if len(faceBookConfig.ClientSecret) == 0 {
		t.Error("FacebookClientSecret is blank")
	}
	if len(faceBookConfig.ClientID) == 0 {
		t.Error("FacebookClientID is blank")
	}
	if len(faceBookConfig.CallBackUrl) == 0 {
		t.Error("FacebookCallbackUrl is blank")
	}
	githubConfig := config.GitHubAuth
	if len(githubConfig.ClientSecret) == 0 {
		t.Error("GithubClientSecret is blank")
	}
	if len(githubConfig.ClientID) == 0 {
		t.Error("GithubClientId is blank")
	}
	if len(githubConfig.CallBackUrl) == 0 {
		t.Error("GithubCallbackUrl is blank")
	}
	jwtConfig := config.JwtConfig
	if len(jwtConfig.SignInKey) == 0 {
		t.Error("JwtSignInKey is blank")
	}

}
