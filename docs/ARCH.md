# ARCH.md — Architecture & Conventions

## Pattern Utama: Vertical Slice

**1 endpoint = 1 feature folder.** Tiap feature isolated — punya handler, logic, dan dependency sendiri.
Gak ada shared service layer yang jadi god-class.

Contoh benar:
```
feature/public/connections/list/    ← GET  /connections
feature/public/connections/create/  ← POST /connections
feature/public/connections/update/  ← PUT  /connections/:id
feature/public/connections/delete/  ← DELETE /connections/:id
```

```
source/feature/
  public/          ← fitur yang dipublish ke luar (accessible dari browser)
  private/         ← fitur/endpoint yang digunakan internal
  _example/        ← template scaffold
```

Tambah feature baru → jalanin `feature.sh`, mount di `service/route.go`. Gak perlu nyentuh file lain.

---

## Struktur Per Feature

```
source/feature/<scope>/<feature_name>/
  handler.go        ← struct Handler + NewHandler(rdb *redis.Client)
  handler_impl.go   ← method Impl(c *gin.Context) — logic endpoint
  repository.go     ← interface Repositories + repositoryImpl struct
  repository_impl.go← implementasi method repository
  body/
    request.go      ← request body struct
    response.go     ← response body struct
```

---

## Cara Bikin Feature Baru

```bash
./feature.sh public/list_keys
# → generate source/feature/public/list_keys/{handler.go, ...}
# → lalu daftarkan route di source/service/route.go
```

Setelah scaffold:
1. Edit `handler.go` — sesuaikan dependency kalau perlu
2. Implement logic di `handler_impl.go` dan `repository_impl.go`
3. Mount route di `source/service/route.go`

---

## Dependency Injection

- Redis client (`*redis.Client`) di-pass dari `main.go` ke `NewHandler()` ke `injectRepository()`
- Gak ada global state / singleton di luar `config` dan `logger`

---

## Config

- File: `config.yaml` (dev default)
- Override via env: `REDIS_ADDR`, `REDIS_PASSWORD`, `REDIS_DB`, `APP_PORT`, dll
- Semua field ada di `source/config/struct_cfg.go`
- Load via koanf: yaml dulu, env override setelahnya

---

## HTTP Response

Pakai helper di `source/common/glob_utils/http_resp_utils/`:

```go
httpresputils.HTTPRespOK(c, data, &msg)
httpresputils.HTTPRespBadRequest(c, &errMsg)
httpresputils.HTTPRespBadGateway(c, &errMsg)
```

---

## Logger

Zerolog via `source/pkg/logger`. Pakai:

```go
logger.Info().Str("key", val).Msg("message")
logger.Error().Err(err).Caller().Msg("message")
```

---

## Frontend

- HTMX untuk partial page swap (gak full reload)
- Tailwind via CDN (gak ada build step)
- Template HTML ada di `source/templates/[feature_name]/*.html`
- Tiap feature punya folder template sendiri sesuai nama feature-nya

### Mockup Workflow

Sebelum implement Go handler:
1. Bikin static HTML di `source/templates/` — data hardcode, tanpa HTMX dulu
2. Review di browser → approve visual + layout
3. Wire ke Go handler (HTMX partial swap)
4. Hapus hardcode data → pakai data dari backend

File mockup yang direncanakan:
```
source/templates/
  empty_state.html      ← Flow 1: belum ada connection
  connections.html      ← Flow 2: connection manager list
  workspace.html        ← Flow 3: split panel key browser
```

---

## Naming Conventions

| Hal             | Konvensi                          |
|-----------------|-----------------------------------|
| Feature folder  | `snake_case` (e.g. `list_keys`)   |
| Package name    | sama dengan folder name           |
| Handler struct  | `Handler`                         |
| Constructor     | `NewHandler(rdb *redis.Client)`   |
| Route mount     | di `source/service/route.go` only |
| Config field    | koanf tag `snake_case`            |
| Env key         | `UPPER_SNAKE_CASE`                |
