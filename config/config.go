package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Host     string
	Port     string
	Password string
	Db       int
}

func Load() (*Config, error) {
	v := viper.New()
	v.SetConfigName("grc")
	v.AddConfigPath(".")

	err := v.ReadInConfig()
	if err != nil {
		return nil, err
	}

	c := &Config{}
	err = v.Unmarshal(c)
	return c, err
}
