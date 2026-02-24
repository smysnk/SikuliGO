package grpcv1

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/gorilla/websocket"
	"github.com/smysnk/sikuligo/internal/sessionstore"
	"google.golang.org/grpc/codes"
)

func TestAdminMuxHealthSnapshotAndMetrics(t *testing.T) {
	metrics := NewMetricsRegistry()
	metrics.StartRequest()
	metrics.Record("/sikuli.v1.SikuliService/Find", codes.OK, 12*time.Millisecond, "trace-1")
	metrics.Record("/sikuli.v1.SikuliService/Find", codes.NotFound, 20*time.Millisecond, "trace-2")
	metrics.RecordAuthFailure("/sikuli.v1.SikuliService/Find")
	metrics.FinishRequest()

	mux := NewAdminMux(metrics, nil)

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

	mux := NewAdminMux(metrics, nil)
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

func TestAdminMuxSessionsViewerEndpoints(t *testing.T) {
	metrics := NewMetricsRegistry()
	store, err := sessionstore.OpenSQLite(":memory:")
	if err != nil {
		t.Fatalf("open sqlite: %v", err)
	}
	defer func() {
		_ = store.Close()
	}()

	api, err := store.StartAPISession(context.Background(), sessionstore.APISessionStartInput{
		PID:            111,
		GRPCListenAddr: "127.0.0.1:50051",
	})
	if err != nil {
		t.Fatalf("start api session: %v", err)
	}
	client, err := store.StartClientSession(context.Background(), sessionstore.ClientSessionStartInput{
		APISessionID: api.ID,
		ConnectionID: "conn-1",
		RemoteAddr:   "127.0.0.1:60000",
		LocalAddr:    "127.0.0.1:50051",
	})
	if err != nil {
		t.Fatalf("start client session: %v", err)
	}
	if err := store.RecordInteraction(context.Background(), sessionstore.InteractionInput{
		APISessionID:    api.ID,
		ClientSessionID: client.ID,
		Method:          "/sikuli.v1.SikuliService/FindOnScreen",
		TraceID:         "trace-1",
		GRPCCode:        "OK",
		DurationMS:      23,
		StartedAt:       time.Now().UTC().Add(-23 * time.Millisecond),
		CompletedAt:     time.Now().UTC(),
	}); err != nil {
		t.Fatalf("record interaction: %v", err)
	}

	mux := NewAdminMux(metrics, store)

	viewReq := httptest.NewRequest(http.MethodGet, "/sessions", nil)
	viewRes := httptest.NewRecorder()
	mux.ServeHTTP(viewRes, viewReq)
	if viewRes.Code != http.StatusOK {
		t.Fatalf("sessions view status mismatch: %d", viewRes.Code)
	}
	if !strings.Contains(viewRes.Body.String(), "Session Viewer") {
		t.Fatalf("sessions view title missing")
	}

	apiReq := httptest.NewRequest(http.MethodGet, "/sessions/api", nil)
	apiRes := httptest.NewRecorder()
	mux.ServeHTTP(apiRes, apiReq)
	if apiRes.Code != http.StatusOK {
		t.Fatalf("sessions api list status mismatch: %d", apiRes.Code)
	}
	var apiRows []sessionSummary
	if err := json.Unmarshal(apiRes.Body.Bytes(), &apiRows); err != nil {
		t.Fatalf("sessions api parse failed: %v", err)
	}
	if len(apiRows) != 1 {
		t.Fatalf("expected 1 api session row, got %d", len(apiRows))
	}

	clientReq := httptest.NewRequest(http.MethodGet, "/sessions/api/"+strconv.FormatUint(uint64(api.ID), 10)+"/clients", nil)
	clientRes := httptest.NewRecorder()
	mux.ServeHTTP(clientRes, clientReq)
	if clientRes.Code != http.StatusOK {
		t.Fatalf("sessions client list status mismatch: %d", clientRes.Code)
	}
	var clientRows []sessionSummary
	if err := json.Unmarshal(clientRes.Body.Bytes(), &clientRows); err != nil {
		t.Fatalf("sessions client parse failed: %v", err)
	}
	if len(clientRows) != 1 {
		t.Fatalf("expected 1 client session row, got %d", len(clientRows))
	}

	interactionReq := httptest.NewRequest(http.MethodGet, "/sessions/client/"+strconv.FormatUint(uint64(client.ID), 10)+"/interactions", nil)
	interactionRes := httptest.NewRecorder()
	mux.ServeHTTP(interactionRes, interactionReq)
	if interactionRes.Code != http.StatusOK {
		t.Fatalf("sessions interaction list status mismatch: %d", interactionRes.Code)
	}
	var interactions []interactionSummary
	if err := json.Unmarshal(interactionRes.Body.Bytes(), &interactions); err != nil {
		t.Fatalf("sessions interaction parse failed: %v", err)
	}
	if len(interactions) != 1 {
		t.Fatalf("expected 1 interaction row, got %d", len(interactions))
	}
	if interactions[0].Method != "/sikuli.v1.SikuliService/FindOnScreen" {
		t.Fatalf("unexpected interaction method: %s", interactions[0].Method)
	}
}

func TestAdminMuxWebSocketPush(t *testing.T) {
	metrics := NewMetricsRegistry()
	metrics.Record("/sikuli.v1.SikuliService/Find", codes.OK, 5*time.Millisecond, "trace-ws")

	server := httptest.NewServer(NewAdminMux(metrics, nil))
	defer server.Close()

	wsURL := "ws" + strings.TrimPrefix(server.URL, "http") + "/ws"
	conn, _, err := websocket.DefaultDialer.Dial(wsURL, nil)
	if err != nil {
		t.Fatalf("websocket dial failed: %v", err)
	}
	defer conn.Close()

	var payload dashboardPushPayload
	if err := conn.ReadJSON(&payload); err != nil {
		t.Fatalf("websocket read json failed: %v", err)
	}
	if payload.Snapshot.TotalRequests < 1 {
		t.Fatalf("expected snapshot requests >= 1, got %d", payload.Snapshot.TotalRequests)
	}
}
