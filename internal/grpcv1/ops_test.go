package grpcv1

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"google.golang.org/grpc/codes"
)

func TestAdminMuxHealthSnapshotAndMetrics(t *testing.T) {
	metrics := NewMetricsRegistry()
	metrics.StartRequest()
	metrics.Record("/sikuli.v1.SikuliService/Find", codes.OK, 12*time.Millisecond, "trace-1")
	metrics.Record("/sikuli.v1.SikuliService/Find", codes.NotFound, 20*time.Millisecond, "trace-2")
	metrics.RecordAuthFailure("/sikuli.v1.SikuliService/Find")
	metrics.FinishRequest()

	mux := NewAdminMux(metrics)

	healthReq := httptest.NewRequest(http.MethodGet, "/healthz", nil)
	healthRes := httptest.NewRecorder()
	mux.ServeHTTP(healthRes, healthReq)
	if healthRes.Code != http.StatusOK {
		t.Fatalf("healthz status mismatch: %d", healthRes.Code)
	}
	var health map[string]any
	if err := json.Unmarshal(healthRes.Body.Bytes(), &health); err != nil {
		t.Fatalf("healthz json parse failed: %v", err)
	}
	if health["status"] != "ok" {
		t.Fatalf("healthz status mismatch: %#v", health["status"])
	}

	snapshotReq := httptest.NewRequest(http.MethodGet, "/snapshot", nil)
	snapshotRes := httptest.NewRecorder()
	mux.ServeHTTP(snapshotRes, snapshotReq)
	if snapshotRes.Code != http.StatusOK {
		t.Fatalf("snapshot status mismatch: %d", snapshotRes.Code)
	}
	var snapshot MetricsSnapshot
	if err := json.Unmarshal(snapshotRes.Body.Bytes(), &snapshot); err != nil {
		t.Fatalf("snapshot json parse failed: %v", err)
	}
	if snapshot.TotalRequests != 2 {
		t.Fatalf("snapshot total requests mismatch: %d", snapshot.TotalRequests)
	}
	if snapshot.TotalErrors != 1 {
		t.Fatalf("snapshot total errors mismatch: %d", snapshot.TotalErrors)
	}
	if snapshot.TotalAuthFailures != 1 {
		t.Fatalf("snapshot auth failures mismatch: %d", snapshot.TotalAuthFailures)
	}

	metricsReq := httptest.NewRequest(http.MethodGet, "/metrics", nil)
	metricsRes := httptest.NewRecorder()
	mux.ServeHTTP(metricsRes, metricsReq)
	if metricsRes.Code != http.StatusOK {
		t.Fatalf("metrics status mismatch: %d", metricsRes.Code)
	}
	body := metricsRes.Body.String()
	if !strings.Contains(body, "sikuligrpc_requests_total 2") {
		t.Fatalf("metrics missing total request counter")
	}
	if !strings.Contains(body, "sikuligrpc_auth_failures_total 1") {
		t.Fatalf("metrics missing auth failure counter")
	}
	if !strings.Contains(body, "method=\"/sikuli.v1.SikuliService/Find\"") {
		t.Fatalf("metrics missing method label")
	}
}

func TestAdminMuxDashboard(t *testing.T) {
	metrics := NewMetricsRegistry()
	metrics.Record("/sikuli.v1.SikuliService/Find", codes.OK, 5*time.Millisecond, "trace-1")

	mux := NewAdminMux(metrics)
	req := httptest.NewRequest(http.MethodGet, "/dashboard", nil)
	res := httptest.NewRecorder()
	mux.ServeHTTP(res, req)
	if res.Code != http.StatusOK {
		t.Fatalf("dashboard status mismatch: %d", res.Code)
	}
	body := res.Body.String()
	if !strings.Contains(body, "Sikuli gRPC Dashboard") {
		t.Fatalf("dashboard title missing")
	}
	if !strings.Contains(body, "/sikuli.v1.SikuliService/Find") {
		t.Fatalf("dashboard missing method row")
	}
}
