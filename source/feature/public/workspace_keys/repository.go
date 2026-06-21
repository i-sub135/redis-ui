package workspace_keys

import "context"

type KeyInfo struct {
	Key  string `json:"key"`
	Type string `json:"type"`
	TTL  int64  `json:"ttl"`
}

type Repositories interface {
	ListKeys(ctx context.Context, addr, password string, db int) ([]KeyInfo, error)
}

type repositoryImpl struct{}

func injectRepository() Repositories {
	return &repositoryImpl{}
}
