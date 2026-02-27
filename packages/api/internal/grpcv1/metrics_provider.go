package grpcv1

import (
	"context"
	"sort"
	"time"

	"github.com/smysnk/sikuligo/internal/sessionstore"
)

type MetricsSnapshotProvider interface {
	Snapshot() MetricsSnapshot
}

type StoreMetricsProvider struct {
	store *sessionstore.Store
}

func NewStoreMetricsProvider(store *sessionstore.Store) *StoreMetricsProvider {
	if store == nil {
		return nil
	}
	return &StoreMetricsProvider{store: store}
}

func (p *StoreMetricsProvider) Snapshot() MetricsSnapshot {
	if p == nil || p.store == nil {
		return MetricsSnapshot{}
	}
	ctx := context.Background()
	api, ok, err := p.store.LatestAPISession(ctx)
	if err != nil || !ok {
		return MetricsSnapshot{}
	}
	methods, err := p.store.MethodMetricsByAPISession(ctx, api.ID)
	if err != nil {
		return MetricsSnapshot{
			StartedAtRFC3339: api.StartedAt.UTC().Format(time.RFC3339),
			UptimeSeconds:    int64(time.Since(api.StartedAt).Seconds()),
		}
	}

	out := MetricsSnapshot{
		StartedAtRFC3339: api.StartedAt.UTC().Format(time.RFC3339),
		UptimeSeconds:    int64(time.Since(api.StartedAt).Seconds()),
		Inflight:         0,
		Methods:          make([]MethodSnapshot, 0, len(methods)),
	}
	for _, m := range methods {
		out.TotalRequests += m.Requests
		out.TotalErrors += m.Errors
		out.TotalAuthFailures += m.AuthFailures
		errorRate := 0.0
		if m.Requests > 0 {
			errorRate = float64(m.Errors) / float64(m.Requests) * 100
		}
		out.Methods = append(out.Methods, MethodSnapshot{
			Method:           m.Method,
			Requests:         m.Requests,
			Errors:           m.Errors,
			AuthFailures:     m.AuthFailures,
			AvgLatencyMS:     m.AvgDuration,
			MaxLatencyMS:     float64(m.MaxDuration),
			LastCode:         m.LastCode,
			LastTraceID:      m.LastTraceID,
			LastSeenUnixMS:   m.LastSeen.UnixMilli(),
			LastSeenRFC3339:  timestampOrEmpty(m.LastSeen),
			ErrorRatePercent: errorRate,
		})
	}
	sort.Slice(out.Methods, func(i, j int) bool {
		return out.Methods[i].Method < out.Methods[j].Method
	})
	return out
}
