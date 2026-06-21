package workspace_keys

import (
	"strconv"

	"github.com/gin-gonic/gin"

	httpresputils "github.com/i-sub135/redis-ui/source/common/glob_utils/http_resp_utils"
	"github.com/i-sub135/redis-ui/source/pkg/logger"
)

// ListKeys godoc
// @Summary List keys for a connection and DB
// @Router /api/v1/workspace/:id/keys [get]
func (h *Handler) ListKeys(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("id")

	db := 0
	if q := c.Query("db"); q != "" {
		if n, err := strconv.Atoi(q); err == nil {
			db = n
		}
	}

	conn, err := h.store.GetByID(id)
	if err != nil {
		errMsg := err.Error()
		httpresputils.HTTPRespNotFound(c, &errMsg)
		return
	}

	repo := injectRepository()
	keys, err := repo.ListKeys(ctx, conn.Addr, conn.Password, db)
	if err != nil {
		errMsg := err.Error()
		logger.Error().Err(err).Caller().Msg(errMsg)
		httpresputils.HTTPRespBadRequest(c, &errMsg)
		return
	}

	httpresputils.HTTPRespOK(c, keys, nil)
}
