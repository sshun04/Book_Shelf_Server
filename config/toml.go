package config

import "github.com/BurntSushi/toml"

type Config struct {
	FaceBookAuth FaceBookConfig
	GitHubAuth GithubConfig
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

func GetConfig() (*Config,error) {
	var config Config
	_,err := toml.DecodeFile("server/config/conf.toml",&config)
	if err != nil {
		print(err.Error())
		return  &config,err
	}
	return &config,nil
}

