package sessionstore

import (
	"context"
	"testing"
	"time"
)

func TestStoreAPISessionClientSessionAndInteraction(t *testing.T) {
	t.Parallel()

	store, err := OpenSQLite(":memory:")
	if err != nil {
		t.Fatalf("open sqlite: %v", err)
	}
	defer func() {
		_ = store.Close()
	}()

	ctx := context.Background()
	api, err := store.StartAPISession(ctx, APISessionStartInput{
		PID:             1234,
		GRPCListenAddr:  "127.0.0.1:50051",
		AdminListenAddr: "127.0.0.1:8080",
	})
	if err != nil {
		t.Fatalf("start api session: %v", err)
	}
	if api.ID == 0 {
		t.Fatalf("expected api session id")
	}
	if api.SessionKey == "" {
		t.Fatalf("expected api session key")
	}

	client, err := store.StartClientSession(ctx, ClientSessionStartInput{
		APISessionID: api.ID,
		ConnectionID: "conn-1",
		RemoteAddr:   "10.0.0.10:50000",
		LocalAddr:    "127.0.0.1:50051",
	})
	if err != nil {
		t.Fatalf("start client session: %v", err)
	}
	if client.ID == 0 {
		t.Fatalf("expected client session id")
	}

	now := time.Now().UTC()
	if err := store.RecordInteraction(ctx, InteractionInput{
		APISessionID:    api.ID,
		ClientSessionID: client.ID,
		Method:          "/sikuli.v1.SikuliService/FindOnScreen",
		TraceID:         "trace-abc",
		GRPCCode:        "OK",
		DurationMS:      12,
		StartedAt:       now.Add(-12 * time.Millisecond),
		CompletedAt:     now,
	}); err != nil {
		t.Fatalf("record interaction: %v", err)
	}

	if err := store.EndClientSession(ctx, client.ID, now); err != nil {
		t.Fatalf("end client session: %v", err)
	}
	if err := store.EndAPISession(ctx, api.ID, now); err != nil {
		t.Fatalf("end api session: %v", err)
	}

	apiCount, err := store.CountAPISessions(ctx)
	if err != nil {
		t.Fatalf("count api sessions: %v", err)
	}
	if apiCount != 1 {
		t.Fatalf("expected 1 api session, got %d", apiCount)
	}
	clientCount, err := store.CountClientSessions(ctx)
	if err != nil {
		t.Fatalf("count client sessions: %v", err)
	}
	if clientCount != 1 {
		t.Fatalf("expected 1 client session, got %d", clientCount)
	}
	interactionCount, err := store.CountInteractions(ctx)
	if err != nil {
		t.Fatalf("count interactions: %v", err)
	}
	if interactionCount != 1 {
		t.Fatalf("expected 1 interaction, got %d", interactionCount)
	}
}
