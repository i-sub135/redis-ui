package example

import "context"

// Implementasi method tambahan yang ada di interface Repositories
// tapi bukan bagian dari shared repo (internal logic feature ini).
//
// Kalau semua method sudah terpenuhi dari embedded shared repo,
// file ini bisa dikosongkan (tapi tetap ada biar konsisten).

func (r *repositoryImpl) Execute(ctx context.Context) (any, error) {
	// TODO: implementasi logic
	return nil, nil
}