package internal

import "github.com/Netflix/go-env"

type Config struct {
	DatabaseURL    string `env:"DATABASE_URL"`
	Port           string `env:"PORT"`
	DatabaseDriver string `env:"DATABASE_DRIVER"`
}

func LoadConfig() (*Config, error) {
	var config Config
	_, err := env.UnmarshalFromEnviron(&config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
