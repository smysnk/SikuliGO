package grpcv1

import (
	"context"
	"testing"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func TestAuthUnaryInterceptorRejectsMissingToken(t *testing.T) {
	metrics := NewMetricsRegistry()
	opts := newInterceptorOptions("secret-token", nil, metrics)
	interceptor := authUnaryInterceptor(opts)

	called := false
	_, err := interceptor(
		context.Background(),
		nil,
		&grpc.UnaryServerInfo{FullMethod: "/sikuli.v1.SikuliService/Find"},
		func(ctx context.Context, req any) (any, error) {
			called = true
			return "ok", nil
		},
	)
	if err == nil {
		t.Fatalf("expected unauthenticated error")
	}
	if code := status.Code(err); code != codes.Unauthenticated {
		t.Fatalf("expected unauthenticated code, got %s", code)
	}
	if called {
		t.Fatalf("handler should not be called when auth fails")
	}
	snap := metrics.Snapshot()
	if snap.TotalAuthFailures != 1 {
		t.Fatalf("expected auth failures=1, got %d", snap.TotalAuthFailures)
	}
}

func TestAuthUnaryInterceptorAcceptsBearerToken(t *testing.T) {
	opts := newInterceptorOptions("secret-token", nil, NewMetricsRegistry())
	interceptor := authUnaryInterceptor(opts)

	ctx := metadata.NewIncomingContext(context.Background(), metadata.Pairs("authorization", "Bearer secret-token"))

	called := false
	_, err := interceptor(
		ctx,
		nil,
		&grpc.UnaryServerInfo{FullMethod: "/sikuli.v1.SikuliService/Find"},
		func(ctx context.Context, req any) (any, error) {
			called = true
			return "ok", nil
		},
	)
	if err != nil {
		t.Fatalf("expected success, got error: %v", err)
	}
	if !called {
		t.Fatalf("handler should be called for valid token")
	}
}

func TestTracingUnaryInterceptorSetsTraceID(t *testing.T) {
	interceptor := tracingUnaryInterceptor(newInterceptorOptions("", nil, nil))

	var gotTraceID string
	_, err := interceptor(
		context.Background(),
		nil,
		&grpc.UnaryServerInfo{FullMethod: "/sikuli.v1.SikuliService/Find"},
		func(ctx context.Context, req any) (any, error) {
			gotTraceID = traceIDFromContext(ctx)
			return "ok", nil
		},
	)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotTraceID == "" {
		t.Fatalf("expected trace id to be present in context")
	}
}

func TestLoggingUnaryInterceptorRecordsMetrics(t *testing.T) {
	metrics := NewMetricsRegistry()
	interceptor := loggingUnaryInterceptor(newInterceptorOptions("", nil, metrics))

	_, err := interceptor(
		context.Background(),
		nil,
		&grpc.UnaryServerInfo{FullMethod: "/sikuli.v1.SikuliService/Find"},
		func(ctx context.Context, req any) (any, error) {
			time.Sleep(1 * time.Millisecond)
			return nil, status.Error(codes.NotFound, "missing")
		},
	)
	if err == nil {
		t.Fatalf("expected not found error")
	}
	if code := status.Code(err); code != codes.NotFound {
		t.Fatalf("expected not found code, got %s", code)
	}

	snap := metrics.Snapshot()
	if snap.TotalRequests != 1 {
		t.Fatalf("expected total requests=1, got %d", snap.TotalRequests)
	}
	if snap.TotalErrors != 1 {
		t.Fatalf("expected total errors=1, got %d", snap.TotalErrors)
	}
	if len(snap.Methods) != 1 {
		t.Fatalf("expected exactly one method metric, got %d", len(snap.Methods))
	}
	if snap.Methods[0].Method != "/sikuli.v1.SikuliService/Find" {
		t.Fatalf("unexpected method metric: %s", snap.Methods[0].Method)
	}
}
