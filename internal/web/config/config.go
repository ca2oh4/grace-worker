package config

import (
	"grace-worker/pkg/config"
)

var (
	Database config.Database
	Redis    config.Redis
	Server   config.Server
)

func Setup() error {
	type T struct {
		Database config.Database
		Redis    config.Redis
		Server   config.Server
	}
	var t T
	if err := config.Setup(&t); err != nil {
		return err
	}

	Database = t.Database
	Redis = t.Redis
	Server = t.Server

	return nil
}
