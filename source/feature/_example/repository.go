package example

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type Repositories interface {
	Execute(ctx context.Context) (any, error)
}

type repositoryImpl struct {
	rdb *redis.Client
}

func injectRepository(rdb *redis.Client) Repositories {
	return &repositoryImpl{rdb: rdb}
}
