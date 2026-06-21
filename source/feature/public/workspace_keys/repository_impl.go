package workspace_keys

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

func (r *repositoryImpl) ListKeys(ctx context.Context, addr, password string, db int) ([]KeyInfo, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:        addr,
		Password:    password,
		DB:          db,
		DialTimeout: 3 * time.Second,
		ReadTimeout: 3 * time.Second,
	})
	defer rdb.Close()

	var keys []string
	var cursor uint64
	for {
		batch, next, err := rdb.Scan(ctx, cursor, "*", 100).Result()
		if err != nil {
			return nil, fmt.Errorf("scan: %w", err)
		}
		keys = append(keys, batch...)
		cursor = next
		if cursor == 0 {
			break
		}
	}

	if len(keys) == 0 {
		return []KeyInfo{}, nil
	}

	pipe := rdb.Pipeline()
	typeCmds := make([]*redis.StatusCmd, len(keys))
	ttlCmds := make([]*redis.DurationCmd, len(keys))
	for i, k := range keys {
		typeCmds[i] = pipe.Type(ctx, k)
		ttlCmds[i] = pipe.TTL(ctx, k)
	}
	if _, err := pipe.Exec(ctx); err != nil && err != redis.Nil {
		return nil, fmt.Errorf("type pipeline: %w", err)
	}

	result := make([]KeyInfo, len(keys))
	for i, k := range keys {
		t, _ := typeCmds[i].Result()
		ttlDur, _ := ttlCmds[i].Result()
		result[i] = KeyInfo{Key: k, Type: t, TTL: int64(ttlDur / time.Second)}
	}
	return result, nil
}
