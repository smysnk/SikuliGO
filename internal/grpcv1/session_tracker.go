package grpcv1

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"log"
	"time"

	"github.com/smysnk/sikuligo/internal/sessionstore"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/stats"
)

type contextKeyClientSessionID struct{}

func withClientSessionID(ctx context.Context, clientSessionID uint) context.Context {
	return context.WithValue(ctx, contextKeyClientSessionID{}, clientSessionID)
}

func clientSessionIDFromContext(ctx context.Context) uint {
	if ctx == nil {
		return 0
	}
	id, _ := ctx.Value(contextKeyClientSessionID{}).(uint)
	return id
}

type SessionTracker struct {
	store        *sessionstore.Store
	apiSessionID uint
	logger       *log.Logger
}

func NewSessionTracker(store *sessionstore.Store, apiSessionID uint, logger *log.Logger) *SessionTracker {
	if store == nil || apiSessionID == 0 {
		return nil
	}
	return &SessionTracker{
		store:        store,
		apiSessionID: apiSessionID,
		logger:       logger,
	}
}

func (s *SessionTracker) TagConn(ctx context.Context, info *stats.ConnTagInfo) context.Context {
	if s == nil || s.store == nil {
		return ctx
	}
	connID := newConnectionUUID()
	remote := ""
	local := ""
	if info != nil {
		if info.RemoteAddr != nil {
			remote = info.RemoteAddr.String()
		}
		if info.LocalAddr != nil {
			local = info.LocalAddr.String()
		}
	}
	row, err := s.store.StartClientSession(ctx, sessionstore.ClientSessionStartInput{
		APISessionID: s.apiSessionID,
		ConnectionID: connID,
		RemoteAddr:   remote,
		LocalAddr:    local,
	})
	if err != nil {
		if s.logger != nil {
			s.logger.Printf("session tracker failed to create client session: %v", err)
		}
		return ctx
	}
	return withClientSessionID(ctx, row.ID)
}

func newConnectionUUID() string {
	var b [16]byte
	if _, err := rand.Read(b[:]); err != nil {
		seed := uint64(time.Now().UnixNano())
		for i := 0; i < 8; i++ {
			shift := uint(56 - (i * 8))
			b[i] = byte(seed >> shift)
			b[i+8] = byte(seed >> shift)
		}
	}
	// RFC 4122 version 4 UUID bits.
	b[6] = (b[6] & 0x0f) | 0x40
	b[8] = (b[8] & 0x3f) | 0x80
	buf := make([]byte, 36)
	hex.Encode(buf[0:8], b[0:4])
	buf[8] = '-'
	hex.Encode(buf[9:13], b[4:6])
	buf[13] = '-'
	hex.Encode(buf[14:18], b[6:8])
	buf[18] = '-'
	hex.Encode(buf[19:23], b[8:10])
	buf[23] = '-'
	hex.Encode(buf[24:36], b[10:16])
	return string(buf)
}

func (s *SessionTracker) HandleConn(ctx context.Context, conn stats.ConnStats) {
	if s == nil || s.store == nil {
		return
	}
	if _, ok := conn.(*stats.ConnEnd); !ok {
		return
	}
	clientSessionID := clientSessionIDFromContext(ctx)
	if clientSessionID == 0 {
		return
	}
	if err := s.store.EndClientSession(context.Background(), clientSessionID, time.Now().UTC()); err != nil && s.logger != nil {
		s.logger.Printf("session tracker failed to close client session=%d: %v", clientSessionID, err)
	}
}

func (s *SessionTracker) TagRPC(ctx context.Context, _ *stats.RPCTagInfo) context.Context {
	return ctx
}

func (s *SessionTracker) HandleRPC(context.Context, stats.RPCStats) {}

func (s *SessionTracker) RecordInteraction(ctx context.Context, method string, code codes.Code, duration time.Duration, traceID string) {
	if s == nil || s.store == nil {
		return
	}
	clientSessionID := clientSessionIDFromContext(ctx)
	if clientSessionID == 0 {
		return
	}
	completedAt := time.Now().UTC()
	startedAt := completedAt
	if duration > 0 {
		startedAt = completedAt.Add(-duration)
	}
	if err := s.store.RecordInteraction(context.Background(), sessionstore.InteractionInput{
		APISessionID:    s.apiSessionID,
		ClientSessionID: clientSessionID,
		Method:          method,
		TraceID:         traceID,
		GRPCCode:        code.String(),
		DurationMS:      duration.Milliseconds(),
		StartedAt:       startedAt,
		CompletedAt:     completedAt,
	}); err != nil && s.logger != nil {
		s.logger.Printf("session tracker failed to record interaction method=%s client_session=%d: %v", method, clientSessionID, err)
	}
}
