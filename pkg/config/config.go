package config

import (
	"github.com/spf13/viper"
)

func Setup[T any](cfg *T) error {
	v := viper.New()
	v.SetConfigFile(".env")
	if err := v.ReadInConfig(); err != nil {
		return err
	}
	return v.Unmarshal(cfg)
}
