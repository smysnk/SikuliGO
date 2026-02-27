# API Electron App (macOS)

Desktop shell for SikuliGO dashboard/session viewer with API process control.

## Run

```bash
yarn workspace @sikuligo/api-electron dev
```

Environment overrides:

- `SIKULIGO_BINARY_PATH` (default: `../../../sikuligo` relative to this package)
- `SIKULIGO_API_LISTEN` (default: `127.0.0.1:50051`)
- `SIKULIGO_ADMIN_LISTEN` (default: `127.0.0.1:8080`)
- `SIKULIGO_API_AUTO_START` (default: `1`; set `0` to disable auto-start)
- `SIKULIGO_DASHBOARD_URL` (default: `http://127.0.0.1:8080/dashboard`)
- `SIKULIGO_SESSION_VIEW_URL` (default: `${SIKULIGO_DASHBOARD_URL}?view=session-viewer`)
