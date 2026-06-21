package example

import (
	"github.com/gin-gonic/gin"

	httpresputils "github.com/i-sub135/redis-ui/source/common/glob_utils/http_resp_utils"
	"github.com/i-sub135/redis-ui/source/pkg/logger"
)

// Impl godoc
// @Summary Example endpoint
// @Router /api/v1/example [post]
func (h *Handler) Impl(c *gin.Context) {
	ctx := c.Request.Context()

	// TODO: ambil param / query / body sesuai kebutuhan
	// key := c.Param("key")

	result, err := h.repo.Execute(ctx)
	if err != nil {
		errMsg := err.Error()
		logger.Error().Err(err).Caller().Msg(errMsg)
		httpresputils.HTTPRespBadRequest(c, &errMsg)
		return
	}

	httpresputils.HTTPRespOK(c, result, nil)
}
