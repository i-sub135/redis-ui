package healtcheck

import (
	"github.com/gin-gonic/gin"

	httpresputils "github.com/i-sub135/redis-ui/source/common/glob_utils/http_resp_utils"
)

// HealtCheck godoc
// @Summary App alive check
// @Router /api/v1/health [get]
func (h *Handler) HealtCheck(c *gin.Context) {
	msg := "ok"
	httpresputils.HTTPRespOK(c, nil, &msg)
}
