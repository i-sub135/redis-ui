package workspace_dbs

import "context"

type DbInfo struct {
	DB   int `json:"db"`
	Keys int `json:"keys"`
}

type Repositories interface {
	ListDbs(ctx context.Context, addr, password string) ([]DbInfo, error)
}

type repositoryImpl struct{}

func injectRepository() Repositories {
	return &repositoryImpl{}
}
