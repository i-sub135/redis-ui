package workspace_key_value

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

func (r *repositoryImpl) GetValue(ctx context.Context, addr, password string, db int, key string) (*KeyValue, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:        addr,
		Password:    password,
		DB:          db,
		DialTimeout: 3 * time.Second,
		ReadTimeout: 3 * time.Second,
	})
	defer rdb.Close()

	keyType, err := rdb.Type(ctx, key).Result()
	if err != nil {
		return nil, fmt.Errorf("type: %w", err)
	}

	ttl, err := rdb.TTL(ctx, key).Result()
	if err != nil {
		return nil, fmt.Errorf("ttl: %w", err)
	}
	ttlSec := int64(ttl.Seconds())

	var value any
	switch keyType {
	case "string":
		value, err = rdb.Get(ctx, key).Result()
	case "hash":
		value, err = rdb.HGetAll(ctx, key).Result()
	case "list":
		value, err = rdb.LRange(ctx, key, 0, -1).Result()
	case "set":
		value, err = rdb.SMembers(ctx, key).Result()
	case "zset":
		value, err = rdb.ZRangeWithScores(ctx, key, 0, -1).Result()
	case "stream":
		value, err = rdb.XRange(ctx, key, "-", "+").Result()
	default:
		value = nil
	}
	if err != nil && err != redis.Nil {
		return nil, fmt.Errorf("get value (%s): %w", keyType, err)
	}

	return &KeyValue{Key: key, Type: keyType, TTL: ttlSec, Value: value}, nil
}
