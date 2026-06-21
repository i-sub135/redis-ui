// Package httpresputils provides utility functions for standardized HTTP responses.
package httpresputils

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/i-sub135/redis-ui/source/config"
)

type response struct {
	Status     string    `json:"status"`
	Message    *string   `json:"message,omitempty"`
	Time       time.Time `json:"timestamp"`
	AppVersion string    `json:"app_version"`
	Data       any       `json:"data,omitempty"`
}

var (
	cfg = config.GetConfig()
)

func HTTPRespOK(c *gin.Context, data any, msg *string) {
	c.JSON(http.StatusOK, response{
		Status:     http.StatusText(http.StatusOK),
		Time:       time.Now(),
		AppVersion: cfg.App.Version,
		Data:       data,
		Message:    msg,
	})
}

func HTTPRespCreated(c *gin.Context, data any, msg *string) {
	c.JSON(http.StatusCreated, response{
		Status:     http.StatusText(http.StatusCreated),
		Time:       time.Now(),
		AppVersion: cfg.App.Version,
		Data:       data,
		Message:    msg,
	})
}

func HTTPRespAccepted(c *gin.Context, data any, msg *string) {
	c.JSON(http.StatusAccepted, response{
		Status:     http.StatusText(http.StatusAccepted),
		Time:       time.Now(),
		AppVersion: cfg.App.Version,
		Data:       data,
		Message:    msg,
	})
}

func HTTPRespNoContent(c *gin.Context) {
	c.JSON(http.StatusNoContent, response{
		Status:     http.StatusText(http.StatusNoContent),
		Time:       time.Now(),
		AppVersion: cfg.App.Version,
		Data:       nil,
		Message:    nil,
	})
}

func HTTPRespNotFound(c *gin.Context, msg *string) {
	c.JSON(http.StatusNotFound, response{
		Status:     http.StatusText(http.StatusNotFound),
		Message:    msg,
		AppVersion: cfg.App.Version,
		Time:       time.Now(),
	})
}

func HTTPRespBadRequest(c *gin.Context, msg *string) {
	c.JSON(http.StatusBadRequest, response{
		Status:     http.StatusText(http.StatusBadRequest),
		Message:    msg,
		AppVersion: cfg.App.Version,
		Time:       time.Now(),
	})
}

func HTTPRespBadGateway(c *gin.Context, msg *string) {
	c.JSON(http.StatusBadGateway, response{
		Status:     http.StatusText(http.StatusBadGateway),
		Message:    msg,
		AppVersion: cfg.App.Version,
		Time:       time.Now(),
	})
}
