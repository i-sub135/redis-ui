package workspace_key_delete

import (
	"context"

	"github.com/redis/go-redis/v9"
)

func (r *repositoryImpl) DeleteKey(ctx context.Context, addr, password string, db int, key string) error {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
	defer rdb.Close()

	return rdb.Del(ctx, key).Err()
}
