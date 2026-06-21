// Package logger provides structured logging functionality.
package logger

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"

	"github.com/i-sub135/redis-ui/source/config"
	"github.com/i-sub135/redis-ui/source/service/constant"
)

var Log zerolog.Logger

// Init initializes global logger.
// levelStr example: "debug", "info"
// prettyConsole: when true, use human-friendly console writer
// callerSkip: frames to skip so caller points to original caller (use 2 if wrapping)
func Init(prettyConsole bool) {
	// parse level
	lvl, err := zerolog.ParseLevel(config.GetConfig().Log.Level)
	if err != nil {
		lvl = zerolog.InfoLevel
	}

	zerolog.CallerMarshalFunc = func(pc uintptr, file string, line int) string {
		return fmt.Sprintf("%s:%d", file, line) // Full path instead of filepath.Base(file)
	}
	// zerolog.CallerSkipFrameCount = 0
	zerolog.TimeFieldFormat = time.RFC3339

	if prettyConsole {
		out := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
		Log = zerolog.New(out).Level(lvl).With().Timestamp().Str("app", config.GetConfig().App.Name).Str("app_version", config.GetConfig().App.Version).Logger()
	} else {
		Log = zerolog.New(os.Stdout).Level(lvl).With().Timestamp().Str("app", config.GetConfig().App.Name).Str("app_version", config.GetConfig().App.Version).Logger()
	}
}

// Debug returns a debug level event.
func Debug() *zerolog.Event { return Log.Debug() }

// Info returns an info level event.
func Info() *zerolog.Event { return Log.Info() }

// Warn returns a warn level event.
func Warn() *zerolog.Event { return Log.Warn() }

// Error returns an error level event.
func Error() *zerolog.Event { return Log.Error() }

// GinZLogger returns middleware that logs request after handler runs.
func GinZLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		dur := time.Since(start)
		status := c.Writer.Status()
		method := c.Request.Method
		path := c.FullPath()
		if path == "" {
			path = c.Request.URL.Path
		}
		ip := c.ClientIP()

		var ev *zerolog.Event
		switch {
		case status >= 500:
			ev = Log.Error()
		case status >= 400:
			ev = Log.Warn()
		case status >= 300:
			ev = Log.Debug()
		default:
			ev = Log.Info()
		}

		// Get request body and query parameters
		body, _ := c.GetRawData()
		query := c.Request.URL.Query()
		email, _ := c.Get(constant.ContextUserEmail)

		ev = ev.Str("method", method).
			Str("path", path).
			Str("client_ip", ip).
			Str(constant.RequestIDKey, c.GetString(constant.RequestIDKey)).
			Interface("user", email).
			Int("status", status).
			Interface("query", query).
			Dur("latency", dur)

		if len(body) == 0 {
			ev.Str("body", "")
		} else if json.Valid(body) {
			ev.RawJSON("body", body)
		} else {
			ev.Str("body", string(body))
		}

		ev.Msg("http request")
	}
}
