# checkpoint.md — Progress Summary

Last updated: 2026-06-20

---

## Done

### Bootstrap (G1 ✅)
- Folder `~/Documents/home-project/redis-ui` dibuat
- `go mod init github.com/i-sub135/redis-ui` + `git init`
- Struktur di-copy dari `v2.gorest-bluepint`, di-adapt: drop db/jwt/minio, tambah `pkg/redis/`, strip config
- `go build ./...` clean

### Docs (G2 ✅)
- `docs/SCOPE.md`, `README.md`, `GOALS.md`, `ARCH.md`, `checkpoint.md`

### Mockup (G2.5 ✅)
- `source/templates/connections/index.html` — Flow 2: connection manager list
- `source/templates/connections/empty.html` — Flow 1: no connection state
- `source/templates/workspace/index.html` — Flow 3: split panel key browser
- `source/templates/workspace/detail_mobile.html` — mobile detail state
- All approved by Iyan 2026-06-20

### G3 — Connection Store + API ✅
- `source/pkg/connection-list/connection.go` — JSON file-backed store, thread-safe (sync.RWMutex), `Load/Add/Update/Delete/GetByID`
- `source/feature/public/connections/` — CRUD handler (bundled, approved exception):
  - `GET /api/v1/connections` → List
  - `POST /api/v1/connections` → Create (201)
  - `PUT /api/v1/connections/:id` → Update (404 if not found)
  - `DELETE /api/v1/connections/:id` → Delete (204 / 404)
- `source/service/route.go` — Store init `./data/connections.json`, routes dimount
- `source/feature/public/healtcheck/` — pure app-alive, no Redis ping
- `go run .` clean, semua routes terdaftar

### Template Serving (Pre-G4) ✅
- `main.go` — `loadTemplates()` walk `source/templates/` dengan relative path sebagai nama template
- `r.SetHTMLTemplate()`, `r.Static("/static", "web/static")`
- `GET /` → `connections/index.html`
- `GET /workspace/:id` → `workspace/index.html`

### G4 — DB Selector ✅
- Feature: `source/feature/public/workspace_dbs/`
- `GET /api/v1/workspace/:id/dbs` — stateless connect ke Redis, parse `INFO keyspace`, return `[{db, keys}]`
- `workspace/index.html` — fetch on load, populate dropdown real, `onDbChange` reload

### G5 — Browse Keys ✅
- Feature: `source/feature/public/workspace_keys/`
- `GET /api/v1/workspace/:id/keys?db=0` — SCAN all keys + pipeline TYPE, return `[{key, type}]`
- `workspace/index.html` — key list rendered dari API, type badge color-coded, search filter client-side, `onDbChange` reload keys

### G6 — View Value ✅
- Feature: `source/feature/public/workspace_key_value/`
- `GET /api/v1/workspace/:id/key-value?key=...&db=0` — detect type, fetch value per type
  - string → GET, hash → HGETALL, list → LRANGE, set → SMEMBERS, zset → ZRANGE WITHSCORES, stream → XRANGE
- `workspace/index.html` — klik key → fetch → render detail panel dinamis per type, `escHtml()` XSS protection

### Real Test Connection ✅
- Connection `localhost:6379` terdaftar via API, ID: `ed990297-37f1-4e49-9042-485c3d91710b`
- DB 0: 241 keys (stream), DB 7: 11 keys — data live
- `data/connections.json` persist di project root

---

## Active Routes

```
GET  /
GET  /workspace/:id
GET  /health
GET  /static/*filepath

GET    /api/v1/connections
POST   /api/v1/connections
PUT    /api/v1/connections/:id
DELETE /api/v1/connections/:id

GET  /api/v1/workspace/:id/dbs
GET  /api/v1/workspace/:id/keys
GET  /api/v1/workspace/:id/key-value
```

---

## Next

- **G7:** Edit Value — `PUT /api/v1/workspace/:id/key-value?key=...&db=0`, type-specific update
- **G8:** Delete Key — `DELETE /api/v1/workspace/:id/key-value?key=...&db=0`
- **G9:** TTL View & Set — tampil TTL dari key-value response (sudah ada), add endpoint `PATCH /api/v1/workspace/:id/key-ttl`

---

## Catatan Penting

- **Arsitektur:** `pkg` = zero HTTP, feature handler consume pkg + expose route. 1 endpoint = 1 feature folder (connections/ = exception approved)
- `feature.sh` scaffold: template `_example` pakai `example`/`Example` literal. Body folder kadang generate `{{FEATURE_CAMEL}}` placeholder — hapus manual kalau gak dipakai
- Template naming: `source/templates/(feature_name)/*.html` — relative path jadi nama template di Gin
- Connection store stateless per-request: fresh dial ke Redis target, close setelah done
- `data/connections.json` di-gitignore atau perlu di-add — belum dicek
- Project alias: `"redis ui"` → `~/Documents/home-project/redis-ui`, topic Telegram 1324
