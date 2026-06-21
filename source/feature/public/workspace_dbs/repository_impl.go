package workspace_dbs

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"
)

func (r *repositoryImpl) ListDbs(ctx context.Context, addr, password string) ([]DbInfo, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:        addr,
		Password:    password,
		DB:          0,
		DialTimeout: 3 * time.Second,
		ReadTimeout: 3 * time.Second,
	})
	defer rdb.Close()

	info, err := rdb.Info(ctx, "keyspace").Result()
	if err != nil {
		return nil, fmt.Errorf("redis info: %w", err)
	}

	return parseKeyspace(info), nil
}

func parseKeyspace(info string) []DbInfo {
	var dbs []DbInfo
	for _, line := range strings.Split(info, "\n") {
		line = strings.TrimSpace(line)
		if !strings.HasPrefix(line, "db") {
			continue
		}
		// format: db0:keys=128,expires=2,avg_ttl=0
		parts := strings.SplitN(line, ":", 2)
		if len(parts) != 2 {
			continue
		}
		dbNum, err := strconv.Atoi(strings.TrimPrefix(parts[0], "db"))
		if err != nil {
			continue
		}
		keys := 0
		for _, kv := range strings.Split(parts[1], ",") {
			if strings.HasPrefix(kv, "keys=") {
				keys, _ = strconv.Atoi(strings.TrimPrefix(kv, "keys="))
				break
			}
		}
		dbs = append(dbs, DbInfo{DB: dbNum, Keys: keys})
	}
	return dbs
}
