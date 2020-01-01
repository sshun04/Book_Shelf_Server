package config

import "github.com/BurntSushi/toml"

type Config struct {
	FaceBookAuth FaceBookConfig
	GitHubAuth GithubConfig
	JwtConfig JwtConfig
}

type FaceBookConfig struct {
	ClientID string  `toml:"facebookClientID"`
	ClientSecret string `toml:"facebookClientSecret"`
	CallBackUrl string `toml:"facebookCallbackUrl"`
}

type GithubConfig struct {
	ClientID string `toml:"githubClientID"`
	ClientSecret string `toml:"githubClientSecret"`
	CallBackUrl string `toml:"githubCallbackUrl"`
}

type JwtConfig struct {
	SignInKey string `toml:"signInKey"`
}

func GetConfig() (*Config,error) {
	var config Config
	_,err := toml.DecodeFile("/Users/shunsukeshoji/go/src/bookstorage_web/server/config/conf.toml",&config)
	if err != nil {
		print(err.Error())
		return  &config,err
	}
	return &config,nil
}

