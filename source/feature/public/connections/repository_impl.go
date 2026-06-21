package connections

import connectionlist "github.com/i-sub135/redis-ui/source/pkg/connection-list"

func (r *repositoryImpl) List() ([]connectionlist.Connection, error) {
	return r.store.Load()
}

func (r *repositoryImpl) Add(conn connectionlist.Connection) (connectionlist.Connection, error) {
	return r.store.Add(conn)
}

func (r *repositoryImpl) Update(id string, conn connectionlist.Connection) error {
	return r.store.Update(id, conn)
}

func (r *repositoryImpl) Delete(id string) error {
	return r.store.Delete(id)
}
