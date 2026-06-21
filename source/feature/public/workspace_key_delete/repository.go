package workspace_key_delete

import "context"

type Repositories interface {
	DeleteKey(ctx context.Context, addr, password string, db int, key string) error
}

type repositoryImpl struct{}

func injectRepository() Repositories {
	return &repositoryImpl{}
}
