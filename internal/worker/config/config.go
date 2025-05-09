package config

import (
	"grace-worker/pkg/config"
)

var (
	Database config.Database
	Redis    config.Redis
)

func Setup() error {
	type T struct {
		Database config.Database
		Redis    config.Redis
	}
	var t T
	if err := config.Setup(&t); err != nil {
		return err
	}

	Database = t.Database
	Redis = t.Redis

	return nil
}
