# API Electron App (macOS)

Desktop shell for the sikuli-go dashboard/session viewer with API process control.

## Run

```bash
yarn workspace @sikuligo/api-electron dev
```

Start the Milestone 1 editor flow from the repo root:

```bash
yarn dev:ide
```

Environment overrides:

- `SIKULI_GO_BINARY_PATH` (default: `../../../sikuli-go` relative to this package)
- `SIKULI_GO_API_LISTEN` (default: `127.0.0.1:50051`)
- `SIKULI_GO_ADMIN_LISTEN` (default: `127.0.0.1:8080`)
- `SIKULI_GO_API_AUTO_START` (default: `1`; set `0` to disable auto-start)
- `SIKULI_GO_DASHBOARD_URL` (default: `http://127.0.0.1:8080/dashboard`)
- `SIKULI_GO_SESSION_VIEW_URL` (default: `http://127.0.0.1:8080/sessions`)
- `SIKULI_GO_EDITOR_URL` (optional Next.js editor URL; used by `yarn dev:ide`)
- `SIKULI_GO_INITIAL_VIEW` (`editor` or `dashboard`)
