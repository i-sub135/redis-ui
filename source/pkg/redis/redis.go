package redisclient

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"

	"github.com/i-sub135/redis-ui/source/config"
)

func Init() (*redis.Client, error) {
	cfg := config.GetConfig()

	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Addr,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	})

	if err := rdb.Ping(context.Background()).Err(); err != nil {
		return nil, fmt.Errorf("redis connect failed: %w", err)
	}

	return rdb, nil
}
