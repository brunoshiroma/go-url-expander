package internal

import (
	"log"

	env "github.com/Netflix/go-env"
)

type EnvConfig struct {
	LogLevel       int    `env:"LOG_LEVEL,default=1"`
	Port           int    `env:"PORT,default=60080"`
	BindHost       string `env:"BIND_HOST,default=127.0.0.1"`
	RequestTimeout int    `env:"REQUEST_TIMEOUT,default=5"`
	MaxRedirects   int    `env:"MAX_REDIRECTS,default=5"`
}

var Config EnvConfig

func init() {
	_, err := env.UnmarshalFromEnviron(&Config)
	if err != nil {
		log.Fatal(err)
	}
}
