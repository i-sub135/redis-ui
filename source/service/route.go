package service

import (
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"

	"github.com/i-sub135/redis-ui/source/feature/public/connections"
	workspace_dbs "github.com/i-sub135/redis-ui/source/feature/public/workspace_dbs"
	workspace_key_value "github.com/i-sub135/redis-ui/source/feature/public/workspace_key_value"
	workspace_keys "github.com/i-sub135/redis-ui/source/feature/public/workspace_keys"
	connectionlist "github.com/i-sub135/redis-ui/source/pkg/connection-list"
)

type Routers struct {
	rdb   *redis.Client
	store *connectionlist.Store
}

func NewRouters(rdb *redis.Client) *Routers {
	return &Routers{
		rdb:   rdb,
		store: connectionlist.NewStore("./data/connections.json"),
	}
}

func (r *Routers) MountRouters(routeGroup *gin.RouterGroup) {
	connections.NewHandler(r.store).RegisterRoutes(routeGroup.Group("/connections"))
	routeGroup.GET("/workspace/:id/dbs", workspace_dbs.NewHandler(r.store))
	routeGroup.GET("/workspace/:id/keys", workspace_keys.NewHandler(r.store))
	routeGroup.GET("/workspace/:id/key-value", workspace_key_value.NewHandler(r.store))
}
