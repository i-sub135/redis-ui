package connections

import (
	"errors"

	"github.com/gin-gonic/gin"

	httpresputils "github.com/i-sub135/redis-ui/source/common/glob_utils/http_resp_utils"
	connectionlist "github.com/i-sub135/redis-ui/source/pkg/connection-list"
	"github.com/i-sub135/redis-ui/source/pkg/logger"

	"github.com/i-sub135/redis-ui/source/feature/public/connections/body"
)

// List godoc
// @Summary List all connections
// @Router /api/v1/connections [get]
func (h *Handler) List(c *gin.Context) {
	conns, err := h.repo.List()
	if err != nil {
		errMsg := err.Error()
		logger.Error().Err(err).Caller().Msg(errMsg)
		httpresputils.HTTPRespBadRequest(c, &errMsg)
		return
	}
	httpresputils.HTTPRespOK(c, conns, nil)
}

// Create godoc
// @Summary Add a new connection
// @Router /api/v1/connections [post]
func (h *Handler) Create(c *gin.Context) {
	var req body.CreateConnectionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		errMsg := err.Error()
		httpresputils.HTTPRespBadRequest(c, &errMsg)
		return
	}

	conn, err := h.repo.Add(connectionlist.Connection{
		Name:     req.Name,
		Addr:     req.Addr,
		Password: req.Password,
		DB:       req.DB,
	})
	if err != nil {
		errMsg := err.Error()
		logger.Error().Err(err).Caller().Msg(errMsg)
		httpresputils.HTTPRespBadRequest(c, &errMsg)
		return
	}
	httpresputils.HTTPRespCreated(c, conn, nil)
}

// Update godoc
// @Summary Update a connection
// @Router /api/v1/connections/:id [put]
func (h *Handler) Update(c *gin.Context) {
	id := c.Param("id")

	var req body.UpdateConnectionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		errMsg := err.Error()
		httpresputils.HTTPRespBadRequest(c, &errMsg)
		return
	}

	err := h.repo.Update(id, connectionlist.Connection{
		Name:     req.Name,
		Addr:     req.Addr,
		Password: req.Password,
		DB:       req.DB,
	})
	if err != nil {
		errMsg := err.Error()
		if errors.Is(err, connectionlist.ErrNotFound) {
			httpresputils.HTTPRespNotFound(c, &errMsg)
			return
		}
		logger.Error().Err(err).Caller().Msg(errMsg)
		httpresputils.HTTPRespBadRequest(c, &errMsg)
		return
	}
	httpresputils.HTTPRespOK(c, nil, nil)
}

// Delete godoc
// @Summary Delete a connection
// @Router /api/v1/connections/:id [delete]
func (h *Handler) Delete(c *gin.Context) {
	id := c.Param("id")

	err := h.repo.Delete(id)
	if err != nil {
		errMsg := err.Error()
		if errors.Is(err, connectionlist.ErrNotFound) {
			httpresputils.HTTPRespNotFound(c, &errMsg)
			return
		}
		logger.Error().Err(err).Caller().Msg(errMsg)
		httpresputils.HTTPRespBadRequest(c, &errMsg)
		return
	}
	httpresputils.HTTPRespNoContent(c)
}
