package grpcv1

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/gorilla/websocket"
	"github.com/smysnk/sikuligo/internal/sessionstore"
	"google.golang.org/grpc/codes"
)

type methodStats struct {
	requests     uint64
	errors       uint64
	authFailures uint64
	totalLatency time.Duration
	maxLatency   time.Duration
	lastCode     string
	lastTraceID  string
	lastSeen     time.Time
}

type MethodSnapshot struct {
	Method           string  `json:"method"`
	Requests         uint64  `json:"requests"`
	Errors           uint64  `json:"errors"`
	AuthFailures     uint64  `json:"auth_failures"`
	AvgLatencyMS     float64 `json:"avg_latency_ms"`
	MaxLatencyMS     float64 `json:"max_latency_ms"`
	LastCode         string  `json:"last_code"`
	LastTraceID      string  `json:"last_trace_id"`
	LastSeenUnixMS   int64   `json:"last_seen_unix_ms"`
	LastSeenRFC3339  string  `json:"last_seen_rfc3339"`
	ErrorRatePercent float64 `json:"error_rate_percent"`
}

type MetricsSnapshot struct {
	StartedAtRFC3339  string           `json:"started_at_rfc3339"`
	UptimeSeconds     int64            `json:"uptime_seconds"`
	Inflight          int64            `json:"inflight"`
	TotalRequests     uint64           `json:"total_requests"`
	TotalErrors       uint64           `json:"total_errors"`
	TotalAuthFailures uint64           `json:"total_auth_failures"`
	Methods           []MethodSnapshot `json:"methods"`
}

type sessionSummary struct {
	ID          uint   `json:"id"`
	SessionUUID string `json:"session_uuid"`
	Date        string `json:"date"`
	Duration    string `json:"duration"`
}

type interactionSummary struct {
	ID       uint   `json:"id"`
	Method   string `json:"method"`
	TraceID  string `json:"trace_id"`
	GRPCCode string `json:"grpc_code"`
	Date     string `json:"date"`
	Duration string `json:"duration"`
}

type dashboardPushPayload struct {
	Snapshot    MetricsSnapshot  `json:"snapshot"`
	APISessions []sessionSummary `json:"api_sessions,omitempty"`
}

var dashboardUpgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type MetricsRegistry struct {
	startedAt time.Time

	inflight int64

	mu                sync.RWMutex
	totalRequests     uint64
	totalErrors       uint64
	totalAuthFailures uint64
	methods           map[string]*methodStats
}

func NewMetricsRegistry() *MetricsRegistry {
	return &MetricsRegistry{
		startedAt: time.Now().UTC(),
		methods:   make(map[string]*methodStats),
	}
}

func (m *MetricsRegistry) StartRequest() {
	atomic.AddInt64(&m.inflight, 1)
}

func (m *MetricsRegistry) FinishRequest() {
	atomic.AddInt64(&m.inflight, -1)
}

func (m *MetricsRegistry) Record(method string, code codes.Code, latency time.Duration, traceID string) {
	if m == nil {
		return
	}
	m.mu.Lock()
	defer m.mu.Unlock()

	m.totalRequests++
	if code != codes.OK {
		m.totalErrors++
	}

	stats := m.methods[method]
	if stats == nil {
		stats = &methodStats{}
		m.methods[method] = stats
	}
	stats.requests++
	if code != codes.OK {
		stats.errors++
	}
	stats.totalLatency += latency
	if latency > stats.maxLatency {
		stats.maxLatency = latency
	}
	stats.lastCode = code.String()
	stats.lastTraceID = traceID
	stats.lastSeen = time.Now().UTC()
}

func (m *MetricsRegistry) RecordAuthFailure(method string) {
	if m == nil {
		return
	}
	m.mu.Lock()
	defer m.mu.Unlock()

	m.totalAuthFailures++
	stats := m.methods[method]
	if stats == nil {
		stats = &methodStats{}
		m.methods[method] = stats
	}
	stats.authFailures++
	stats.lastCode = codes.Unauthenticated.String()
	stats.lastSeen = time.Now().UTC()
}

func (m *MetricsRegistry) Snapshot() MetricsSnapshot {
	if m == nil {
		return MetricsSnapshot{}
	}
	m.mu.RLock()
	defer m.mu.RUnlock()

	out := MetricsSnapshot{
		StartedAtRFC3339:  m.startedAt.Format(time.RFC3339),
		UptimeSeconds:     int64(time.Since(m.startedAt).Seconds()),
		Inflight:          atomic.LoadInt64(&m.inflight),
		TotalRequests:     m.totalRequests,
		TotalErrors:       m.totalErrors,
		TotalAuthFailures: m.totalAuthFailures,
		Methods:           make([]MethodSnapshot, 0, len(m.methods)),
	}
	for method, stats := range m.methods {
		avgLatencyMS := 0.0
		errorRatePercent := 0.0
		if stats.requests > 0 {
			avgLatencyMS = float64(stats.totalLatency.Microseconds()) / 1000 / float64(stats.requests)
			errorRatePercent = float64(stats.errors) / float64(stats.requests) * 100
		}
		out.Methods = append(out.Methods, MethodSnapshot{
			Method:           method,
			Requests:         stats.requests,
			Errors:           stats.errors,
			AuthFailures:     stats.authFailures,
			AvgLatencyMS:     avgLatencyMS,
			MaxLatencyMS:     float64(stats.maxLatency.Microseconds()) / 1000,
			LastCode:         stats.lastCode,
			LastTraceID:      stats.lastTraceID,
			LastSeenUnixMS:   stats.lastSeen.UnixMilli(),
			LastSeenRFC3339:  timestampOrEmpty(stats.lastSeen),
			ErrorRatePercent: errorRatePercent,
		})
	}
	sort.Slice(out.Methods, func(i, j int) bool {
		return out.Methods[i].Method < out.Methods[j].Method
	})
	return out
}

func NewAdminMux(provider MetricsSnapshotProvider, store *sessionstore.Store) *http.ServeMux {
	if provider == nil {
		provider = NewMetricsRegistry()
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		http.Redirect(w, r, "/dashboard", http.StatusTemporaryRedirect)
	})
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
		snap := provider.Snapshot()
		writeJSON(w, http.StatusOK, map[string]any{
			"status":           "ok",
			"started_at":       snap.StartedAtRFC3339,
			"uptime_seconds":   snap.UptimeSeconds,
			"inflight":         snap.Inflight,
			"total_requests":   snap.TotalRequests,
			"total_errors":     snap.TotalErrors,
			"auth_failures":    snap.TotalAuthFailures,
			"observed_methods": len(snap.Methods),
		})
	})
	mux.HandleFunc("/snapshot", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
		writeJSON(w, http.StatusOK, provider.Snapshot())
	})
	mux.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
		w.Header().Set("Content-Type", "text/plain; version=0.0.4")
		snap := provider.Snapshot()
		_, _ = fmt.Fprintln(w, "# HELP sikuligrpc_requests_total Total gRPC requests.")
		_, _ = fmt.Fprintln(w, "# TYPE sikuligrpc_requests_total counter")
		_, _ = fmt.Fprintf(w, "sikuligrpc_requests_total %d\n", snap.TotalRequests)
		_, _ = fmt.Fprintln(w, "# HELP sikuligrpc_errors_total Total gRPC requests with non-OK status.")
		_, _ = fmt.Fprintln(w, "# TYPE sikuligrpc_errors_total counter")
		_, _ = fmt.Fprintf(w, "sikuligrpc_errors_total %d\n", snap.TotalErrors)
		_, _ = fmt.Fprintln(w, "# HELP sikuligrpc_auth_failures_total Total gRPC authentication failures.")
		_, _ = fmt.Fprintln(w, "# TYPE sikuligrpc_auth_failures_total counter")
		_, _ = fmt.Fprintf(w, "sikuligrpc_auth_failures_total %d\n", snap.TotalAuthFailures)
		_, _ = fmt.Fprintln(w, "# HELP sikuligrpc_inflight Current in-flight gRPC requests.")
		_, _ = fmt.Fprintln(w, "# TYPE sikuligrpc_inflight gauge")
		_, _ = fmt.Fprintf(w, "sikuligrpc_inflight %d\n", snap.Inflight)

		_, _ = fmt.Fprintln(w, "# HELP sikuligrpc_method_requests_total Method request totals.")
		_, _ = fmt.Fprintln(w, "# TYPE sikuligrpc_method_requests_total counter")
		for _, m := range snap.Methods {
			_, _ = fmt.Fprintf(w, "sikuligrpc_method_requests_total{method=\"%s\"} %d\n", escapePromLabel(m.Method), m.Requests)
		}
		_, _ = fmt.Fprintln(w, "# HELP sikuligrpc_method_errors_total Method error totals.")
		_, _ = fmt.Fprintln(w, "# TYPE sikuligrpc_method_errors_total counter")
		for _, m := range snap.Methods {
			_, _ = fmt.Fprintf(w, "sikuligrpc_method_errors_total{method=\"%s\"} %d\n", escapePromLabel(m.Method), m.Errors)
		}
		_, _ = fmt.Fprintln(w, "# HELP sikuligrpc_method_avg_latency_ms Method average latency (ms).")
		_, _ = fmt.Fprintln(w, "# TYPE sikuligrpc_method_avg_latency_ms gauge")
		for _, m := range snap.Methods {
			_, _ = fmt.Fprintf(w, "sikuligrpc_method_avg_latency_ms{method=\"%s\"} %.3f\n", escapePromLabel(m.Method), m.AvgLatencyMS)
		}
	})
	mux.HandleFunc("/dashboard", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		snap := provider.Snapshot()
		_ = dashboardTemplate.Execute(w, snap)
	})
	mux.HandleFunc("/sessions", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		_ = sessionsTemplate.Execute(w, nil)
	})
	mux.HandleFunc("/sessions/api", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
		if store == nil {
			writeJSON(w, http.StatusOK, []sessionSummary{})
			return
		}
		rows, err := store.ListRecentAPISessions(r.Context(), 100)
		if err != nil {
			http.Error(w, fmt.Sprintf("list api sessions: %v", err), http.StatusInternalServerError)
			return
		}
		now := time.Now().UTC()
		out := make([]sessionSummary, 0, len(rows))
		for _, row := range rows {
			out = append(out, sessionSummary{
				ID:          row.ID,
				SessionUUID: row.SessionKey,
				Date:        row.StartedAt.UTC().Format(time.RFC3339),
				Duration:    durationDisplay(row.StartedAt, row.EndedAt, now),
			})
		}
		writeJSON(w, http.StatusOK, out)
	})
	mux.HandleFunc("/sessions/api/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
		if store == nil {
			writeJSON(w, http.StatusOK, []sessionSummary{})
			return
		}
		const prefix = "/sessions/api/"
		if !strings.HasPrefix(r.URL.Path, prefix) || !strings.HasSuffix(r.URL.Path, "/clients") {
			http.NotFound(w, r)
			return
		}
		idPart := strings.TrimSuffix(strings.TrimPrefix(r.URL.Path, prefix), "/clients")
		apiSessionID, ok := parseUintID(idPart)
		if !ok {
			http.Error(w, "invalid api session id", http.StatusBadRequest)
			return
		}
		rows, err := store.ListClientSessionsByAPI(r.Context(), apiSessionID)
		if err != nil {
			http.Error(w, fmt.Sprintf("list client sessions: %v", err), http.StatusInternalServerError)
			return
		}
		now := time.Now().UTC()
		out := make([]sessionSummary, 0, len(rows))
		for _, row := range rows {
			out = append(out, sessionSummary{
				ID:          row.ID,
				SessionUUID: row.SessionKey,
				Date:        row.StartedAt.UTC().Format(time.RFC3339),
				Duration:    durationDisplay(row.StartedAt, row.EndedAt, now),
			})
		}
		writeJSON(w, http.StatusOK, out)
	})
	mux.HandleFunc("/sessions/client/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
		if store == nil {
			writeJSON(w, http.StatusOK, []interactionSummary{})
			return
		}
		const prefix = "/sessions/client/"
		if !strings.HasPrefix(r.URL.Path, prefix) || !strings.HasSuffix(r.URL.Path, "/interactions") {
			http.NotFound(w, r)
			return
		}
		idPart := strings.TrimSuffix(strings.TrimPrefix(r.URL.Path, prefix), "/interactions")
		clientSessionID, ok := parseUintID(idPart)
		if !ok {
			http.Error(w, "invalid client session id", http.StatusBadRequest)
			return
		}
		rows, err := store.ListInteractionsByClient(r.Context(), clientSessionID, 500)
		if err != nil {
			http.Error(w, fmt.Sprintf("list interactions: %v", err), http.StatusInternalServerError)
			return
		}
		out := make([]interactionSummary, 0, len(rows))
		for _, row := range rows {
			out = append(out, interactionSummary{
				ID:       row.ID,
				Method:   row.Method,
				TraceID:  row.TraceID,
				GRPCCode: row.GRPCCode,
				Date:     row.StartedAt.UTC().Format(time.RFC3339),
				Duration: durationMSDisplay(row.DurationMS),
			})
		}
		writeJSON(w, http.StatusOK, out)
	})
	mux.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
		conn, err := dashboardUpgrader.Upgrade(w, r, nil)
		if err != nil {
			return
		}
		defer conn.Close()

		sendUpdate := func() error {
			msg := dashboardPushPayload{
				Snapshot: provider.Snapshot(),
			}
			if store != nil {
				rows, err := store.ListRecentAPISessions(r.Context(), 100)
				if err == nil {
					now := time.Now().UTC()
					msg.APISessions = make([]sessionSummary, 0, len(rows))
					for _, row := range rows {
						msg.APISessions = append(msg.APISessions, sessionSummary{
							ID:          row.ID,
							SessionUUID: row.SessionKey,
							Date:        row.StartedAt.UTC().Format(time.RFC3339),
							Duration:    durationDisplay(row.StartedAt, row.EndedAt, now),
						})
					}
				}
			}
			return conn.WriteJSON(msg)
		}

		_ = conn.SetWriteDeadline(time.Now().Add(5 * time.Second))
		if err := sendUpdate(); err != nil {
			return
		}
		ticker := time.NewTicker(1 * time.Second)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				_ = conn.SetWriteDeadline(time.Now().Add(5 * time.Second))
				if err := sendUpdate(); err != nil {
					return
				}
			case <-r.Context().Done():
				return
			}
		}
	})
	return mux
}

func writeJSON(w http.ResponseWriter, statusCode int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	_ = enc.Encode(payload)
}

func parseUintID(raw string) (uint, bool) {
	raw = strings.TrimSpace(strings.Trim(raw, "/"))
	if raw == "" {
		return 0, false
	}
	id64, err := strconv.ParseUint(raw, 10, 64)
	if err != nil {
		return 0, false
	}
	return uint(id64), true
}

func durationDisplay(start time.Time, end *time.Time, now time.Time) string {
	if start.IsZero() {
		return ""
	}
	finish := now
	if end != nil && !end.IsZero() {
		finish = end.UTC()
	}
	if finish.Before(start) {
		finish = start
	}
	d := finish.Sub(start)
	if d < 0 {
		d = 0
	}
	return d.Round(time.Millisecond).String()
}

func durationMSDisplay(ms int64) string {
	if ms < 0 {
		ms = 0
	}
	return (time.Duration(ms) * time.Millisecond).Round(time.Millisecond).String()
}

func timestampOrEmpty(t time.Time) string {
	if t.IsZero() {
		return ""
	}
	return t.Format(time.RFC3339)
}

func escapePromLabel(v string) string {
	v = strings.ReplaceAll(v, `\`, `\\`)
	v = strings.ReplaceAll(v, `"`, `\"`)
	return v
}

var dashboardTemplate = template.Must(template.New("dashboard").Parse(`<!doctype html>
<html lang="en">
<head>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <title>Sikuli gRPC Dashboard</title>
  <style>
    body { font-family: ui-sans-serif, system-ui, -apple-system, Segoe UI, Roboto, Arial, sans-serif; margin: 0; color: #111827; background: #f8fafc; }
    .wrap { padding: 16px 20px; }
    .menubar { display: flex; justify-content: space-between; align-items: center; border-bottom: 1px solid #d1d5db; background: #ffffff; padding: 12px 20px; }
    .title { margin: 0; font-size: 20px; }
    .btn-group { display: inline-flex; border: 1px solid #cbd5e1; border-radius: 8px; overflow: hidden; background: #ffffff; }
    .btn-group a { padding: 8px 14px; font-size: 13px; text-decoration: none; color: #1f2937; border-right: 1px solid #cbd5e1; background: #ffffff; }
    .btn-group a:last-child { border-right: none; }
    .btn-group a.active { background: #111827; color: #f9fafb; }
    .muted { color: #6b7280; margin-top: 0; margin-bottom: 12px; }
    .cards { display: grid; grid-template-columns: repeat(4, minmax(160px, 1fr)); gap: 12px; margin: 16px 0; }
    .card { border: 1px solid #e5e7eb; border-radius: 8px; padding: 12px; background: #f9fafb; }
    .k { font-size: 12px; color: #6b7280; }
    .v { font-size: 24px; font-weight: 700; }
    table { border-collapse: collapse; width: 100%; margin-top: 12px; }
    th, td { border: 1px solid #e5e7eb; padding: 8px; text-align: left; font-size: 13px; }
    th { background: #f3f4f6; }
    code { background: #f3f4f6; padding: 1px 4px; border-radius: 4px; }
    @media (max-width: 900px) {
      .cards { grid-template-columns: repeat(2, minmax(160px, 1fr)); }
    }
  </style>
</head>
<body>
  <div class="menubar">
    <h1 class="title">Sikuli Dashboard</h1>
    <div class="btn-group">
      <a href="/dashboard" class="active">live</a>
      <a href="/sessions">session viewer</a>
    </div>
  </div>
  <div class="wrap">
    <p id="statusLine" class="muted">Started {{ .StartedAtRFC3339 }} · Uptime {{ .UptimeSeconds }}s</p>

    <div class="cards">
      <div class="card"><div class="k">In-flight</div><div id="inflight" class="v">{{ .Inflight }}</div></div>
      <div class="card"><div class="k">Total Requests</div><div id="totalRequests" class="v">{{ .TotalRequests }}</div></div>
      <div class="card"><div class="k">Total Errors</div><div id="totalErrors" class="v">{{ .TotalErrors }}</div></div>
      <div class="card"><div class="k">Auth Failures</div><div id="totalAuthFailures" class="v">{{ .TotalAuthFailures }}</div></div>
    </div>

    <p>Raw endpoints: <code>/healthz</code>, <code>/snapshot</code>, <code>/metrics</code>, <code>/sessions</code>, <code>/ws</code></p>

    <table>
      <thead>
        <tr>
          <th>Method</th>
          <th>Requests</th>
          <th>Errors</th>
          <th>Auth Failures</th>
          <th>Avg Latency (ms)</th>
          <th>Max Latency (ms)</th>
          <th>Error Rate (%)</th>
          <th>Last Code</th>
          <th>Last Seen</th>
        </tr>
      </thead>
      <tbody id="methodsBody">
      {{ range .Methods }}
        <tr>
          <td><code>{{ .Method }}</code></td>
          <td>{{ .Requests }}</td>
          <td>{{ .Errors }}</td>
          <td>{{ .AuthFailures }}</td>
          <td>{{ printf "%.3f" .AvgLatencyMS }}</td>
          <td>{{ printf "%.3f" .MaxLatencyMS }}</td>
          <td>{{ printf "%.2f" .ErrorRatePercent }}</td>
          <td>{{ .LastCode }}</td>
          <td>{{ .LastSeenRFC3339 }}</td>
        </tr>
      {{ end }}
      </tbody>
    </table>
  </div>
  <script>
    const statusLineEl = document.getElementById("statusLine");
    const inflightEl = document.getElementById("inflight");
    const totalRequestsEl = document.getElementById("totalRequests");
    const totalErrorsEl = document.getElementById("totalErrors");
    const totalAuthFailuresEl = document.getElementById("totalAuthFailures");
    const methodsBodyEl = document.getElementById("methodsBody");

    function esc(v) {
      return String(v ?? "")
        .replaceAll("&", "&amp;")
        .replaceAll("<", "&lt;")
        .replaceAll(">", "&gt;")
        .replaceAll("\"", "&quot;")
        .replaceAll("'", "&#39;");
    }

    function renderSnapshot(snap) {
      if (!snap) {
        return;
      }
      statusLineEl.textContent = "Started " + (snap.started_at_rfc3339 || "") + " · Uptime " + (snap.uptime_seconds || 0) + "s";
      inflightEl.textContent = String(snap.inflight ?? 0);
      totalRequestsEl.textContent = String(snap.total_requests ?? 0);
      totalErrorsEl.textContent = String(snap.total_errors ?? 0);
      totalAuthFailuresEl.textContent = String(snap.total_auth_failures ?? 0);
      const methods = Array.isArray(snap.methods) ? snap.methods : [];
      if (methods.length === 0) {
        methodsBodyEl.innerHTML = '<tr><td colspan="9">No method metrics yet.</td></tr>';
        return;
      }
      methodsBodyEl.innerHTML = methods.map((m) => {
        return '<tr>' +
          '<td><code>' + esc(m.method) + '</code></td>' +
          '<td>' + esc(m.requests) + '</td>' +
          '<td>' + esc(m.errors) + '</td>' +
          '<td>' + esc(m.auth_failures) + '</td>' +
          '<td>' + Number(m.avg_latency_ms || 0).toFixed(3) + '</td>' +
          '<td>' + Number(m.max_latency_ms || 0).toFixed(3) + '</td>' +
          '<td>' + Number(m.error_rate_percent || 0).toFixed(2) + '</td>' +
          '<td>' + esc(m.last_code || "") + '</td>' +
          '<td>' + esc(m.last_seen_rfc3339 || "") + '</td>' +
          '</tr>';
      }).join('');
    }

    function connectWS() {
      const scheme = window.location.protocol === "https:" ? "wss" : "ws";
      const ws = new WebSocket(scheme + "://" + window.location.host + "/ws");
      ws.onmessage = (event) => {
        try {
          const payload = JSON.parse(event.data);
          renderSnapshot(payload.snapshot);
        } catch (e) {
          // ignore malformed payloads
        }
      };
      ws.onclose = () => {
        setTimeout(connectWS, 1000);
      };
    }

    connectWS();
  </script>
</body>
</html>`))

var sessionsTemplate = template.Must(template.New("sessions").Parse(`<!doctype html>
<html lang="en">
<head>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <title>Sikuli Session Viewer</title>
  <style>
    body { font-family: ui-sans-serif, system-ui, -apple-system, Segoe UI, Roboto, Arial, sans-serif; margin: 0; color: #111827; background: #f8fafc; }
    .menubar { display: flex; justify-content: space-between; align-items: center; border-bottom: 1px solid #d1d5db; background: #ffffff; padding: 12px 20px; }
    .title { margin: 0; font-size: 20px; }
    .btn-group { display: inline-flex; border: 1px solid #cbd5e1; border-radius: 8px; overflow: hidden; background: #ffffff; }
    .btn-group a { padding: 8px 14px; font-size: 13px; text-decoration: none; color: #1f2937; border-right: 1px solid #cbd5e1; background: #ffffff; }
    .btn-group a:last-child { border-right: none; }
    .btn-group a.active { background: #111827; color: #f9fafb; }
    .wrap { padding: 16px 20px; }
    .muted { color: #6b7280; margin-top: 0; margin-bottom: 12px; }
    .grid { display: grid; grid-template-columns: 1.8fr 1.2fr 1.4fr; gap: 12px; height: calc(100vh - 140px); }
    .col { border: 1px solid #e5e7eb; border-radius: 8px; background: #f9fafb; min-height: 0; display: flex; flex-direction: column; height: 100%; }
    .col h2 { font-size: 14px; margin: 0; padding: 10px 12px; border-bottom: 1px solid #e5e7eb; background: #f3f4f6; }
    .rows { overflow: auto; flex: 1; }
    .row { padding: 10px 12px; border-bottom: 1px solid #e5e7eb; cursor: pointer; }
    .row:last-child { border-bottom: none; }
    .row:hover { background: #eef2ff; }
    .row.selected { background: #dbeafe; }
    .k { font-size: 12px; color: #6b7280; }
    .v { font-size: 13px; font-family: ui-monospace, SFMono-Regular, Menlo, monospace; word-break: break-all; }
    .empty { padding: 14px; color: #6b7280; font-size: 13px; }
    @media (max-width: 980px) {
      .grid { grid-template-columns: 1fr; }
      .col { min-height: 260px; }
    }
  </style>
</head>
<body>
  <div class="menubar">
    <h1 class="title">Sikuli Dashboard</h1>
    <div class="btn-group">
      <a href="/dashboard">live</a>
      <a href="/sessions" class="active">session viewer</a>
    </div>
  </div>
  <div class="wrap">
    <p class="muted">API sessions contain client sessions, and client sessions contain interactions.</p>
    <div class="grid">
      <section class="col">
        <h2>API Sessions</h2>
        <div id="apiRows" class="rows"><div class="empty">Loading API sessions...</div></div>
      </section>
      <section class="col">
        <h2>Client Sessions</h2>
        <div id="clientRows" class="rows"><div class="empty">Select an API session.</div></div>
      </section>
      <section class="col">
        <h2>Client Interactions</h2>
        <div id="interactionRows" class="rows"><div class="empty">Select a client session.</div></div>
      </section>
    </div>
  </div>

  <script>
    const apiRowsEl = document.getElementById("apiRows");
    const clientRowsEl = document.getElementById("clientRows");
    const interactionRowsEl = document.getElementById("interactionRows");
    let selectedApiID = null;
    let selectedClientID = null;
    let refreshing = false;

    function rowTemplate(title, date, duration, extra) {
      return '<div class="k">session uuid</div><div class="v">' + title + '</div>' +
        '<div class="k">date</div><div>' + date + '</div>' +
        '<div class="k">duration</div><div>' + duration + '</div>' +
        (extra || "");
    }

    function esc(v) {
      return String(v ?? "")
        .replaceAll("&", "&amp;")
        .replaceAll("<", "&lt;")
        .replaceAll(">", "&gt;")
        .replaceAll("\"", "&quot;")
        .replaceAll("'", "&#39;");
    }

    function interactionTemplate(item) {
      const trace = item.trace_id || "-";
      return '<div class="k">method</div><div class="v">' + esc(item.method) + '</div>' +
        '<div class="k">trace</div><div>' + esc(trace) + '</div>' +
        '<div class="k">code</div><div>' + esc(item.grpc_code) + '</div>' +
        '<div class="k">date</div><div>' + esc(item.date) + '</div>' +
        '<div class="k">duration</div><div>' + esc(item.duration) + '</div>';
    }

    async function loadAPISessions() {
      const res = await fetch('/sessions/api');
      const rows = await res.json();
      applyAPISessions(rows);
    }

    function applyAPISessions(rows) {
      if (!Array.isArray(rows) || rows.length === 0) {
        apiRowsEl.innerHTML = '<div class="empty">No API sessions yet.</div>';
        clientRowsEl.innerHTML = '<div class="empty">Select an API session.</div>';
        interactionRowsEl.innerHTML = '<div class="empty">Select a client session.</div>';
        selectedApiID = null;
        selectedClientID = null;
        return;
      }
      if (!rows.some((r) => r.id === selectedApiID)) {
        selectedApiID = rows[0].id;
        selectedClientID = null;
      }
      apiRowsEl.innerHTML = '';
      rows.forEach((row) => {
        const el = document.createElement('div');
        el.className = 'row' + (selectedApiID === row.id ? ' selected' : '');
        el.innerHTML = rowTemplate(esc(row.session_uuid), esc(row.date), esc(row.duration), '');
        el.addEventListener('click', () => selectAPI(row.id));
        apiRowsEl.appendChild(el);
      });
    }

    async function selectAPI(apiID) {
      selectedApiID = apiID;
      selectedClientID = null;
      await refreshSelected();
    }

    async function loadClientSessions(apiID, autoSelect) {
      const res = await fetch('/sessions/api/' + apiID + '/clients');
      const rows = await res.json();
      if (!Array.isArray(rows) || rows.length === 0) {
        clientRowsEl.innerHTML = '<div class="empty">No client sessions for this API session.</div>';
        return;
      }
      clientRowsEl.innerHTML = '';
      rows.forEach((row) => {
        const el = document.createElement('div');
        el.className = 'row' + (selectedClientID === row.id ? ' selected' : '');
        el.innerHTML = rowTemplate(esc(row.session_uuid), esc(row.date), esc(row.duration), '');
        el.addEventListener('click', () => selectClient(row.id));
        clientRowsEl.appendChild(el);
      });
      if (autoSelect) {
        if (!rows.some((r) => r.id === selectedClientID)) {
          selectedClientID = rows[0].id;
        }
        await loadInteractions(selectedClientID);
      }
    }

    async function selectClient(clientID) {
      selectedClientID = clientID;
      await loadClientSessions(selectedApiID, false);
      await loadInteractions(clientID);
    }

    async function loadInteractions(clientID) {
      const res = await fetch('/sessions/client/' + clientID + '/interactions');
      const rows = await res.json();
      if (!Array.isArray(rows) || rows.length === 0) {
        interactionRowsEl.innerHTML = '<div class="empty">No interactions for this client session.</div>';
        return;
      }
      interactionRowsEl.innerHTML = '';
      rows.forEach((row) => {
        const el = document.createElement('div');
        el.className = 'row';
        el.innerHTML = interactionTemplate(row);
        interactionRowsEl.appendChild(el);
      });
    }

    loadAPISessions().catch((err) => {
      apiRowsEl.innerHTML = '<div class="empty">Failed to load sessions: ' + String(err) + '</div>';
    });

    async function refreshSelected() {
      if (refreshing || selectedApiID === null) {
        return;
      }
      refreshing = true;
      try {
        await loadClientSessions(selectedApiID, true);
      } finally {
        refreshing = false;
      }
    }

    function connectWS() {
      const scheme = window.location.protocol === "https:" ? "wss" : "ws";
      const ws = new WebSocket(scheme + "://" + window.location.host + "/ws");
      ws.onmessage = (event) => {
        try {
          const payload = JSON.parse(event.data);
          if (Array.isArray(payload.api_sessions)) {
            applyAPISessions(payload.api_sessions);
            refreshSelected();
          }
        } catch (e) {
          // ignore malformed payload
        }
      };
      ws.onclose = () => {
        setTimeout(connectWS, 1000);
      };
    }

    connectWS();
  </script>
</body>
</html>`))
