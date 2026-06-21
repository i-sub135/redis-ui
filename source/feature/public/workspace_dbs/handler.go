package workspace_dbs

import (
	"github.com/gin-gonic/gin"

	connectionlist "github.com/i-sub135/redis-ui/source/pkg/connection-list"
)

type Handler struct {
	store *connectionlist.Store
}

func NewHandler(store *connectionlist.Store) gin.HandlerFunc {
	h := &Handler{store: store}
	return h.ListDbs
}
