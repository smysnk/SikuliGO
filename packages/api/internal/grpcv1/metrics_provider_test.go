package grpcv1

import (
	"context"
	"testing"
	"time"

	"github.com/smysnk/sikuligo/internal/sessionstore"
)

func TestStoreMetricsProviderSnapshot(t *testing.T) {
	t.Parallel()

	store, err := sessionstore.OpenSQLite(":memory:")
	if err != nil {
		t.Fatalf("open sqlite: %v", err)
	}
	defer func() {
		_ = store.Close()
	}()

	api, err := store.StartAPISession(context.Background(), sessionstore.APISessionStartInput{
		PID:            11,
		GRPCListenAddr: "127.0.0.1:50051",
	})
	if err != nil {
		t.Fatalf("start api session: %v", err)
	}
	client, err := store.StartClientSession(context.Background(), sessionstore.ClientSessionStartInput{
		APISessionID: api.ID,
		ConnectionID: "conn-1",
	})
	if err != nil {
		t.Fatalf("start client session: %v", err)
	}
	now := time.Now().UTC()
	if err := store.RecordInteraction(context.Background(), sessionstore.InteractionInput{
		APISessionID:    api.ID,
		ClientSessionID: client.ID,
		Method:          "/sikuli.v1.SikuliService/FindOnScreen",
		TraceID:         "trace-1",
		GRPCCode:        "OK",
		DurationMS:      10,
		StartedAt:       now.Add(-10 * time.Millisecond),
		CompletedAt:     now,
	}); err != nil {
		t.Fatalf("record interaction 1: %v", err)
	}
	if err := store.RecordInteraction(context.Background(), sessionstore.InteractionInput{
		APISessionID:    api.ID,
		ClientSessionID: client.ID,
		Method:          "/sikuli.v1.SikuliService/FindOnScreen",
		TraceID:         "trace-2",
		GRPCCode:        "NotFound",
		DurationMS:      30,
		StartedAt:       now.Add(-30 * time.Millisecond),
		CompletedAt:     now,
	}); err != nil {
		t.Fatalf("record interaction 2: %v", err)
	}
	totalInteractions, err := store.CountInteractions(context.Background())
	if err != nil {
		t.Fatalf("count interactions: %v", err)
	}
	if totalInteractions != 2 {
		t.Fatalf("expected 2 interactions in store, got %d", totalInteractions)
	}
	latestAPI, ok, err := store.LatestAPISession(context.Background())
	if err != nil {
		t.Fatalf("latest api session: %v", err)
	}
	if !ok || latestAPI.ID == 0 {
		t.Fatalf("expected latest api session")
	}
	methodRows, err := store.MethodMetricsByAPISession(context.Background(), latestAPI.ID)
	if err != nil {
		t.Fatalf("method metrics: %v", err)
	}
	if len(methodRows) != 1 {
		t.Fatalf("expected 1 method metric row, got %d", len(methodRows))
	}

	provider := NewStoreMetricsProvider(store)
	snap := provider.Snapshot()
	if snap.TotalRequests != 2 {
		t.Fatalf("expected total requests 2, got %d", snap.TotalRequests)
	}
	if snap.TotalErrors != 1 {
		t.Fatalf("expected total errors 1, got %d", snap.TotalErrors)
	}
	if len(snap.Methods) != 1 {
		t.Fatalf("expected 1 method row, got %d", len(snap.Methods))
	}
	if snap.Methods[0].Method != "/sikuli.v1.SikuliService/FindOnScreen" {
		t.Fatalf("unexpected method: %s", snap.Methods[0].Method)
	}
}
