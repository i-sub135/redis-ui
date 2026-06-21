package connections

import (
	"github.com/gin-gonic/gin"

	connectionlist "github.com/i-sub135/redis-ui/source/pkg/connection-list"
)

type Handler struct {
	repo Repositories
}

func NewHandler(store *connectionlist.Store) *Handler {
	return &Handler{repo: injectRepository(store)}
}

func (h *Handler) RegisterRoutes(rg *gin.RouterGroup) {
	rg.GET("", h.List)
	rg.POST("", h.Create)
	rg.PUT("/:id", h.Update)
	rg.DELETE("/:id", h.Delete)
}
