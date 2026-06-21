# Redis UI — Project Scope

## Overview

Web-based Redis management dashboard — konsep seperti DataGrip/DB tooling tapi untuk Redis.
Dibangun dengan Go + HTMX, single binary, composable.

User harus input/select connection Redis sendiri. Tidak ada default connection yang langsung ke-connect otomatis.
Tujuan: bisa ngurus Redis instance(s) dari browser tanpa dependency external FE.

---

## Tech Stack

| Layer     | Tech                        |
|-----------|-----------------------------|
| Backend   | Go + Gin                    |
| Frontend  | HTMX + Tailwind CDN         |
| Client    | go-redis/v9                 |
| Config    | koanf (yaml + env override) |
| Logger    | zerolog                     |
| Deploy    | Docker Compose (single svc) |

Arsitektur: vertical slice — 1 endpoint = 1 feature folder, isolated.

Responsive: desktop first. Mobile gracefully degraded — layout tidak rusak, panel stack vertikal di layar kecil. Tidak ada optimasi khusus mobile.

Design approach: HTML mockup static dulu di `source/templates/`, review di browser, baru wire ke backend.

---

## API

- `GET /health` — cek app server up (bukan Redis). Dipakai untuk Docker Compose healthcheck.
- Status Redis per-connection ditampilkan sebagai status indicator di UI, bukan lewat health API.
- Endpoint lain murni server-side rendered via HTMX partial.

---

## Feature Plan

### v1 — MVP

| # | Feature            | Deskripsi                                                                 |
|---|--------------------|---------------------------------------------------------------------------|
| 1 | Multi Connection   | Kelola multiple Redis connections. Select/switch koneksi aktif dari UI.  |
| 2 | DB Selector        | Pilih DB number (0–15) dari koneksi yang aktif.                          |
| 3 | Browse Keys        | List semua key di DB aktif. Supports search/filter by key name.          |
| 4 | View Value         | Lihat value key. Tampil per type: String, Hash, List, Set, ZSet (tabs/submenu). |

### Keep (implemented later, folder scaffolded now)

| # | Feature      | Deskripsi                                                    |
|---|--------------|--------------------------------------------------------------|
| 5 | Edit Value   | Edit value key sesuai type-nya.                              |
| 6 | Delete Key   | Hapus key dari DB aktif.                                     |
| 7 | TTL View     | Lihat TTL sisa (dalam detik) untuk key tertentu.             |
| 8 | TTL Set      | Set atau update TTL untuk key tertentu.                      |

---

## Connection Model

- **Tidak ada default connection.** User wajib input atau select connection Redis secara explicit.
- Connection list disimpan sebagai JSON di `source/pkg/connection-list/` — bukan di config/yaml.
- Connection list bisa di-CRUD (add, edit, delete) dari UI.
- Tiap connection punya: `name`, `addr`, `password`, `db` (default).
- Konsep: mirip DataGrip — connection management explicit, bisa connect ke multiple Redis instance.

---

## UI Flow & Layout

Tiga state utama:

### Flow 1 — Empty State
User belum ada connection. Tampil halaman kosong dengan satu tombol `[+ Add Connection]`.

```
┌──────────────────────────────────────┐
│           Redis UI                   │
│                                      │
│     Tidak ada koneksi aktif.         │
│     [+ Add Connection]               │
│                                      │
└──────────────────────────────────────┘
```

### Flow 2 — Connection Manager
List connection yang tersimpan (JSON), tiap item punya status indicator (● connected / ○ not connected),
tombol `[▶ Connect]`, dan action edit/delete. Plus tombol `[+ Add Connection]` di bawah.

```
┌──────────────────────────────────────┐
│  Connections                [+ Add]  │
│                                      │
│  ● prod       localhost:6379  [▶]   │
│  ○ staging    10.0.0.1:6379   [▶]   │
│  ○ local      127.0.0.1:6380  [▶]   │
│                                      │
└──────────────────────────────────────┘
```

### Flow 3 — Main Workspace
Split panel: kiri = key list + search + DB selector, kanan = detail key + tabs per type.
Header kecil kiri atas: nama connection aktif + tombol `[← back]` ke Connection Manager.

```
┌────────────────────────────────────────────────────────┐
│  [← prod]                          DB: [0 ▼]           │
├──────────────────┬─────────────────────────────────────┤
│  [search keys...]│  key: user:123                       │
│                  │  type: hash  |  ttl: 3600s           │
│  user:123    ●   │                                      │
│  session:abc     │  [String][Hash][List][Set][ZSet]     │
│  cache:xyz       │  field1: value1                      │
│  ...             │  field2: value2                      │
└──────────────────┴─────────────────────────────────────┘
```

Mobile: panel kiri jadi list fullscreen, klik key → panel kanan push-in (stack vertikal).

---

## Mockup Approach

- Static HTML mockup dulu di `source/templates/` — review di browser sebelum wire ke backend
- Tiap flow = 1 file HTML terpisah (`empty_state.html`, `connections.html`, `workspace.html`)
- Styling: Tailwind CDN, tanpa build step
- Interaksi mockup: bisa hardcode data, belum perlu HTMX

---

## Out of Scope (v1)

- Pub/Sub monitor
- Memory/stats dashboard
- Import/export
