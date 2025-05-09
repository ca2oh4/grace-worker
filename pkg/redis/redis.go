package redis

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

var Client *redis.Client

type SetupOption struct {
	Addr     string
	Username string
	Password string

	DB       int
	PoolSize int
}

// Setup 初始化全局 redis client
func Setup(option *SetupOption) error {
	const pingTimeout = time.Second * 5

	Client = redis.NewClient(&redis.Options{
		Addr:     option.Addr,
		Username: option.Username,
		Password: option.Password,
		DB:       option.DB,

		PoolSize: option.PoolSize,
	})
	ctx, cancel := context.WithTimeout(context.Background(), pingTimeout)
	defer cancel()

	return Client.Ping(ctx).Err()
}
