package grpcv1

import (
	"context"
	"log"
	"regexp"
	"testing"
	"time"

	"github.com/smysnk/sikuligo/internal/sessionstore"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/stats"
)

func TestSessionTrackerCreatesClientSessionAndInteraction(t *testing.T) {
	t.Parallel()

	store, err := sessionstore.OpenSQLite(":memory:")
	if err != nil {
		t.Fatalf("open sqlite: %v", err)
	}
	defer func() {
		_ = store.Close()
	}()

	apiSession, err := store.StartAPISession(context.Background(), sessionstore.APISessionStartInput{
		PID:            1,
		GRPCListenAddr: "127.0.0.1:50051",
	})
	if err != nil {
		t.Fatalf("start api session: %v", err)
	}

	tracker := NewSessionTracker(store, apiSession.ID, log.Default())
	if tracker == nil {
		t.Fatalf("expected tracker")
	}

	ctx := tracker.TagConn(context.Background(), &stats.ConnTagInfo{})
	clientSessionID := clientSessionIDFromContext(ctx)
	if clientSessionID == 0 {
		t.Fatalf("expected client session id in context")
	}

	tracker.RecordInteraction(ctx, "/sikuli.v1.SikuliService/FindOnScreen", codes.OK, 15*time.Millisecond, "trace-1")
	tracker.HandleConn(ctx, &stats.ConnEnd{})

	clientCount, err := store.CountClientSessions(context.Background())
	if err != nil {
		t.Fatalf("count client sessions: %v", err)
	}
	if clientCount != 1 {
		t.Fatalf("expected 1 client session, got %d", clientCount)
	}
	interactionCount, err := store.CountInteractions(context.Background())
	if err != nil {
		t.Fatalf("count interactions: %v", err)
	}
	if interactionCount != 1 {
		t.Fatalf("expected 1 interaction, got %d", interactionCount)
	}
}

func TestSessionTrackerConnectionIDsDoNotCollideAcrossAPISessions(t *testing.T) {
	t.Parallel()

	store, err := sessionstore.OpenSQLite(":memory:")
	if err != nil {
		t.Fatalf("open sqlite: %v", err)
	}
	defer func() {
		_ = store.Close()
	}()

	ctx := context.Background()
	apiOne, err := store.StartAPISession(ctx, sessionstore.APISessionStartInput{
		PID:            1,
		GRPCListenAddr: "127.0.0.1:50051",
	})
	if err != nil {
		t.Fatalf("start api session 1: %v", err)
	}
	apiTwo, err := store.StartAPISession(ctx, sessionstore.APISessionStartInput{
		PID:            2,
		GRPCListenAddr: "127.0.0.1:50052",
	})
	if err != nil {
		t.Fatalf("start api session 2: %v", err)
	}

	trackerOne := NewSessionTracker(store, apiOne.ID, log.Default())
	trackerTwo := NewSessionTracker(store, apiTwo.ID, log.Default())

	ctxOne := trackerOne.TagConn(context.Background(), &stats.ConnTagInfo{})
	ctxTwo := trackerTwo.TagConn(context.Background(), &stats.ConnTagInfo{})
	if clientSessionIDFromContext(ctxOne) == 0 || clientSessionIDFromContext(ctxTwo) == 0 {
		t.Fatalf("expected both trackers to create client sessions")
	}

	clientCount, err := store.CountClientSessions(ctx)
	if err != nil {
		t.Fatalf("count client sessions: %v", err)
	}
	if clientCount != 2 {
		t.Fatalf("expected 2 client sessions, got %d", clientCount)
	}
}

func TestNewConnectionUUIDFormat(t *testing.T) {
	t.Parallel()
	uuidRe := regexp.MustCompile(`^[0-9a-f]{8}-[0-9a-f]{4}-4[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$`)

	id := newConnectionUUID()
	if !uuidRe.MatchString(id) {
		t.Fatalf("unexpected UUID format: %s", id)
	}
}
