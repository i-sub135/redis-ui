// Package config provides configuration management for the application.
package config

import (
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

func readVersionFile() string {
	content, err := os.ReadFile("version")
	if err != nil {
		return "unknown"
	}
	return strings.TrimSpace(string(content))
}

var k = koanf.New(".")
var cfg Config

var envKeyMap = map[string]string{
	"APP_NAME":           "app.name",
	"APP_MODE":           "app.mode",
	"APP_PORT":           "app.port",
	"LOG_LEVEL":          "log.level",
	"LOG_PRETTY":         "log.pretty_console",
	"LOG_PRETTY_CONSOLE": "log.pretty_console",
	"REDIS_ADDR":         "redis.addr",
	"REDIS_PASSWORD":     "redis.password",
	"REDIS_DB":           "redis.db",
}

func loadMappedEnv() error {
	for envKey, configPath := range envKeyMap {
		value, exists := os.LookupEnv(envKey)
		if !exists {
			continue
		}
		if err := setConfigValue(configPath, value); err != nil {
			return err
		}
	}
	return nil
}

func setConfigValue(path, value string) error {
	switch path {
	case "app.port", "redis.db":
		n, err := strconv.Atoi(value)
		if err != nil {
			return err
		}
		k.Set(path, n)
	case "log.pretty_console":
		boolVal, err := strconv.ParseBool(value)
		if err != nil {
			return err
		}
		k.Set(path, boolVal)
	default:
		k.Set(path, value)
	}
	return nil
}

// LoadConfig loads config from a YAML file (optional) and env overrides.
// Call once at bootstrap.
func LoadConfig(path string) error {
	if path != "" {
		if err := k.Load(file.Provider(path), yaml.Parser()); err != nil {
			log.Printf("warning: config file not loaded: %v", err)
		}
	}

	if err := loadMappedEnv(); err != nil {
		return err
	}

	if k.String("app.name") == "" {
		k.Set("app.name", "redis-ui")
	}
	if k.String("app.mode") == "" {
		k.Set("app.mode", "local")
	}
	if k.Int("app.port") == 0 {
		k.Set("app.port", 8080)
	}
	if k.String("app.version") == "" {
		k.Set("app.version", readVersionFile())
	}
	if k.String("log.level") == "" {
		k.Set("log.level", "debug")
	}
	if k.String("redis.addr") == "" {
		k.Set("redis.addr", "localhost:6379")
	}

	if err := k.Unmarshal("", &cfg); err != nil {
		return err
	}
	return nil
}

func GetConfig() *Config   { return &cfg }
func Koanf() *koanf.Koanf { return k }

func ResetConfig() {
	k = koanf.New(".")
	cfg = Config{}
}
