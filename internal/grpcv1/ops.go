package grpcv1

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"sort"
	"strings"
	"sync"
	"sync/atomic"
	"time"

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

func NewAdminMux(metrics *MetricsRegistry) *http.ServeMux {
	if metrics == nil {
		metrics = NewMetricsRegistry()
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
		snap := metrics.Snapshot()
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
		writeJSON(w, http.StatusOK, metrics.Snapshot())
	})
	mux.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
		w.Header().Set("Content-Type", "text/plain; version=0.0.4")
		snap := metrics.Snapshot()
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
		snap := metrics.Snapshot()
		_ = dashboardTemplate.Execute(w, snap)
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
    body { font-family: ui-sans-serif, system-ui, -apple-system, Segoe UI, Roboto, Arial, sans-serif; margin: 20px; color: #111827; }
    h1 { margin-bottom: 0.25rem; }
    .muted { color: #6b7280; margin-top: 0; }
    .cards { display: grid; grid-template-columns: repeat(4, minmax(160px, 1fr)); gap: 12px; margin: 16px 0; }
    .card { border: 1px solid #e5e7eb; border-radius: 8px; padding: 12px; background: #f9fafb; }
    .k { font-size: 12px; color: #6b7280; }
    .v { font-size: 24px; font-weight: 700; }
    table { border-collapse: collapse; width: 100%; margin-top: 12px; }
    th, td { border: 1px solid #e5e7eb; padding: 8px; text-align: left; font-size: 13px; }
    th { background: #f3f4f6; }
    code { background: #f3f4f6; padding: 1px 4px; border-radius: 4px; }
  </style>
</head>
<body>
  <h1>Sikuli gRPC Dashboard</h1>
  <p class="muted">Started {{ .StartedAtRFC3339 }} · Uptime {{ .UptimeSeconds }}s</p>

  <div class="cards">
    <div class="card"><div class="k">In-flight</div><div class="v">{{ .Inflight }}</div></div>
    <div class="card"><div class="k">Total Requests</div><div class="v">{{ .TotalRequests }}</div></div>
    <div class="card"><div class="k">Total Errors</div><div class="v">{{ .TotalErrors }}</div></div>
    <div class="card"><div class="k">Auth Failures</div><div class="v">{{ .TotalAuthFailures }}</div></div>
  </div>

  <p>Raw endpoints: <code>/healthz</code>, <code>/snapshot</code>, <code>/metrics</code></p>

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
    <tbody>
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
</body>
</html>`))
