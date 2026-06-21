package main

import (
	"context"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/i-sub135/redis-ui/source/common/glob_utils/helpers"
	"github.com/i-sub135/redis-ui/source/config"
	"github.com/i-sub135/redis-ui/source/feature/public/healtcheck"
	"github.com/i-sub135/redis-ui/source/pkg/logger"
	"github.com/i-sub135/redis-ui/source/service"
	"github.com/i-sub135/redis-ui/source/service/middleware"
)

func loadTemplates(root string) *template.Template {
	t := template.New("")
	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() || !strings.HasSuffix(path, ".html") {
			return nil
		}
		name := strings.TrimPrefix(path, root+"/")
		b, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		template.Must(t.New(name).Parse(string(b)))
		return nil
	})
	return t
}

func main() {
	if err := config.LoadConfig("config.yaml"); err != nil {
		panic(err)
	}
	cfg := config.GetConfig()

	logger.Init(cfg.Log.PrettyConsole)

	if cfg.App.Mode == gin.DebugMode {
		fmt.Printf("Running in mode\n%s\n", helpers.JSONDebuger(cfg))
		fmt.Println("===========================================")
	}

	gin.SetMode(cfg.App.Mode)
	r := gin.New()
	r.Use(middleware.RequestIDMiddleware())
	r.Use(logger.GinZLogger())
	r.Use(gin.Recovery())

	r.SetHTMLTemplate(loadTemplates("source/templates"))
	r.Static("/static", "web/static")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "connections/index.html", nil)
	})
	r.GET("/workspace/:id", func(c *gin.Context) {
		c.HTML(http.StatusOK, "workspace/index.html", gin.H{"id": c.Param("id")})
	})

	healthcheck := healtcheck.NewHandler()
	r.GET("/health", healthcheck.HealtCheck)

	// Mounting routers
	routeAPI := r.Group("/api/v1")
	service.NewRouters().MountRouters(routeAPI)

	svc := &http.Server{
		Addr:           fmt.Sprintf(":%v", cfg.App.Port),
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		IdleTimeout:    120 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	logger.Info().Str("mode", cfg.App.Mode).Msgf("listening on port %v", cfg.App.Port)

	go func() {
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
		<-sigChan

		logger.Info().Msg("shutdown signal received, gracefully shutting down...")

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := svc.Shutdown(ctx); err != nil {
			logger.Error().Err(err).Msg("server forced to shutdown")
		}
	}()

	if err := svc.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Error().Err(err).Msg("server error")
	}

	logger.Info().Msg("server stopped")
}
