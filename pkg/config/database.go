package config

import (
	"grace-worker/pkg/database"
)

type Database struct {
	Host string
	Port int
	DB   string

	Username string
	Password string

	PoolSize int
}

func (t *Database) ToDatabaseOption() *database.SetupOption {
	return &database.SetupOption{
		DatabaseName: t.DB,
		Host:         t.Host,
		Password:     t.Password,
		PoolSize:     t.PoolSize,
		Port:         t.Port,
		Username:     t.Username,
	}
}
