package connections

import connectionlist "github.com/i-sub135/redis-ui/source/pkg/connection-list"

type Repositories interface {
	List() ([]connectionlist.Connection, error)
	Add(conn connectionlist.Connection) (connectionlist.Connection, error)
	Update(id string, conn connectionlist.Connection) error
	Delete(id string) error
}

type repositoryImpl struct {
	store *connectionlist.Store
}

func injectRepository(store *connectionlist.Store) Repositories {
	return &repositoryImpl{store: store}
}
