package toml

import "github.com/BurntSushi/toml"

type Config struct {
	Auth AuthConfig
}

type AuthConfig struct {
	FaceBookClientID string  `toml:"facebookClientID"`
	FaceBookClientSecret string `toml:"facebookClientSecret"`
}

func GetConfig() *Config {
	var config Config
	_,err := toml.DecodeFile("/Users/shunsukeshoji/go/src/bookstorage_web/server/conf/conf.toml",&config)
	if err != nil {
		print(err.Error())
	}
	return &config
}

