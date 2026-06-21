package workspace_key_value

import "context"

type KeyValue struct {
	Key   string `json:"key"`
	Type  string `json:"type"`
	TTL   int64  `json:"ttl"`
	Value any    `json:"value"`
}

type Repositories interface {
	GetValue(ctx context.Context, addr, password string, db int, key string) (*KeyValue, error)
}

type repositoryImpl struct{}

func injectRepository() Repositories {
	return &repositoryImpl{}
}
