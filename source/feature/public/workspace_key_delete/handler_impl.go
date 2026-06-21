package workspace_key_delete

import (
	"strconv"

	"github.com/gin-gonic/gin"

	httpresputils "github.com/i-sub135/redis-ui/source/common/glob_utils/http_resp_utils"
	"github.com/i-sub135/redis-ui/source/pkg/logger"
)

func (h *Handler) DeleteKey(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("id")
	key := c.Query("key")
	if key == "" {
		errMsg := "key is required"
		httpresputils.HTTPRespBadRequest(c, &errMsg)
		return
	}

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
	if err := repo.DeleteKey(ctx, conn.Addr, conn.Password, db, key); err != nil {
		errMsg := err.Error()
		logger.Error().Err(err).Caller().Msg(errMsg)
		httpresputils.HTTPRespBadRequest(c, &errMsg)
		return
	}

	c.Status(204)
}
