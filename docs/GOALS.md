# GOALS.md — Task Running & Acceptance Criteria

Status: `[ ]` todo · `[~]` in progress · `[x]` done

---

## G1 — Project Bootstrap

- [x] `go mod init github.com/i-sub135/redis-ui`
- [x] `git init`
- [x] Copy & adapt struktur dari `v2.gorest-bluepint`
- [x] Drop pkg irrelevant (db, jwt, minio)
- [x] `source/pkg/redis/` — go-redis client init
- [x] Config stripped → hanya App + Log + Redis
- [x] `go build ./...` clean

**AC:** `go build ./...` jalan tanpa error, struktur folder sesuai pattern gorest.

---

## G2 — Docs

- [x] `docs/SCOPE.md`
- [x] `docs/README.md`
- [x] `docs/GOALS.md`
- [x] `docs/checkpoint.md`
- [x] `docs/ARCH.md`

**AC:** Semua docs ada, isi sesuai scope yang disepakati.

---

## G2.5 — Static HTML Mockup ✅

- [x] `source/templates/empty_state.html` — Flow 1: no connection state
- [x] `source/templates/connections.html` — Flow 2: connection manager list
- [x] `source/templates/workspace.html` — Flow 3: split panel key browser
- [x] `source/templates/workspace_detail_mobile.html` — mobile detail state
- [x] Review di browser, approved Iyan 2026-06-20

**AC:** ✅ Semua flow approved. Desktop split panel, mobile fullscreen per state, DB selector onchange auto-refresh (no manual refresh button).

---

## G3 — Multi Connection ✅ (fondasi)

- [x] JSON store untuk menyimpan connection list (`source/pkg/connection-list/`)
- [x] CRUD API: GET / POST / PUT /:id / DELETE /:id via `source/feature/public/connections/`
- [x] Store di-init di `service/route.go`, dimount di `/api/v1/connections`
- [x] Healthcheck → pure app-alive (no Redis ping)
- [ ] UI: form add/edit/delete connection (wire ke API) ← next
- [ ] UI: switch connection aktif → re-render key list ← next

**AC saat ini:** API CRUD jalan, `go run .` bersih, data persist di `./data/connections.json`.

---

## G4 — DB Selector ✅

- [x] Backend: `GET /api/v1/workspace/:id/dbs` — INFO keyspace, return `[{db, keys}]`
- [x] UI: dropdown populated real dari API, onDbChange reload

**AC:** ✅ DB 0 (241 keys) vs DB 7 (11 keys) tampil berbeda. Tested live dengan localhost:6379.

---

## G5 — Browse Keys ✅

- [x] `GET /api/v1/workspace/:id/keys?db=0` — SCAN + pipeline TYPE
- [x] Key list render dinamis, type badge color-coded
- [x] Search/filter client-side
- [x] Klik key → trigger detail load (G6)

**AC:** ✅ 241 keys live, search filter jalan, klik key load detail.

---

## G6 — View Value ✅

- [x] `GET /api/v1/workspace/:id/key-value?key=...&db=0`
- [x] Detect type, fetch per type: string/hash/list/set/zset/stream
- [x] Render detail panel dinamis per type
- [x] XSS protection via escHtml()

**AC:** ✅ Semua type (termasuk stream) bisa ditampilkan. Tested dengan `ai_jobs` stream dari langchain Redis.

---

## G7 — Edit Value *(keep)*

- [ ] Edit value per type
- [ ] Konfirmasi sebelum save

---

## G8 — Delete Key *(keep)*

- [ ] Delete key dengan konfirmasi
- [ ] Key list refresh setelah delete

---

## G9 — TTL View & Set *(keep)*

- [ ] Tampilkan TTL sisa (detik / no expiry)
- [ ] Form set TTL baru

---
