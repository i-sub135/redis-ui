package workspace_dbs

import (
	"github.com/gin-gonic/gin"

	httpresputils "github.com/i-sub135/redis-ui/source/common/glob_utils/http_resp_utils"
	"github.com/i-sub135/redis-ui/source/pkg/logger"
)

// ListDbs godoc
// @Summary List Redis DBs with key counts for a connection
// @Router /api/v1/workspace/:id/dbs [get]
func (h *Handler) ListDbs(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("id")

	conn, err := h.store.GetByID(id)
	if err != nil {
		errMsg := err.Error()
		httpresputils.HTTPRespNotFound(c, &errMsg)
		return
	}

	repo := injectRepository()
	dbs, err := repo.ListDbs(ctx, conn.Addr, conn.Password)
	if err != nil {
		errMsg := err.Error()
		logger.Error().Err(err).Caller().Msg(errMsg)
		httpresputils.HTTPRespBadRequest(c, &errMsg)
		return
	}

	if dbs == nil {
		dbs = []DbInfo{}
	}
	httpresputils.HTTPRespOK(c, dbs, nil)
}
