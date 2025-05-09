package config

import (
	"grace-worker/pkg/redis"
)

type Redis struct {
	Addr string
	DB   int

	Username string
	Password string

	PoolSize int
}

func (t *Redis) ToRedisOptions() *redis.SetupOption {
	return &redis.SetupOption{
		Addr:     t.Addr,
		DB:       t.DB,
		Username: t.Username,
		Password: t.Password,
		PoolSize: t.PoolSize,
	}
}
