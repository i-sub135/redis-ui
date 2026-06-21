package example

// HOW TO USE:
// Run: ./feature.sh public/feature_name OR private/feature_name
// Then implement logic in handler_impl.go and repository_impl.go
// Register route in source/service/route.go

import (
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type Handler struct {
	rdb  *redis.Client
	repo Repositories
}

func NewHandler(rdb *redis.Client) gin.HandlerFunc {
	repo := injectRepository(rdb)
	handler := Handler{
		rdb:  rdb,
		repo: repo,
	}
	return handler.Impl
}
