package grpcv1

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"log"
	"strings"
	"sync/atomic"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"
)

const traceHeader = "x-trace-id"

type interceptorOptions struct {
	authToken     string
	logger        *log.Logger
	metrics       *MetricsRegistry
	publicMethods map[string]struct{}
}

type contextKeyTraceID struct{}

type wrappedServerStream struct {
	grpc.ServerStream
	ctx context.Context
}

func (w *wrappedServerStream) Context() context.Context {
	return w.ctx
}

var traceCounter uint64

func newInterceptorOptions(authToken string, logger *log.Logger, metrics *MetricsRegistry) interceptorOptions {
	return interceptorOptions{
		authToken: strings.TrimSpace(authToken),
		logger:    logger,
		metrics:   metrics,
	}
}

func (o *interceptorOptions) allowMethod(method string) bool {
	if len(o.publicMethods) == 0 {
		return false
	}
	_, ok := o.publicMethods[method]
	return ok
}

func UnaryInterceptors(authToken string, logger *log.Logger, metrics *MetricsRegistry) []grpc.UnaryServerInterceptor {
	opts := newInterceptorOptions(authToken, logger, metrics)
	return []grpc.UnaryServerInterceptor{
		tracingUnaryInterceptor(opts),
		loggingUnaryInterceptor(opts),
		authUnaryInterceptor(opts),
	}
}

func StreamInterceptors(authToken string, logger *log.Logger, metrics *MetricsRegistry) []grpc.StreamServerInterceptor {
	opts := newInterceptorOptions(authToken, logger, metrics)
	return []grpc.StreamServerInterceptor{
		tracingStreamInterceptor(opts),
		loggingStreamInterceptor(opts),
		authStreamInterceptor(opts),
	}
}

func tracingUnaryInterceptor(_ interceptorOptions) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		traceID := incomingTraceID(ctx)
		if traceID == "" {
			traceID = generateTraceID()
		}
		ctx = context.WithValue(ctx, contextKeyTraceID{}, traceID)
		_ = grpc.SetHeader(ctx, metadata.Pairs(traceHeader, traceID))
		_ = grpc.SetTrailer(ctx, metadata.Pairs(traceHeader, traceID))
		return handler(ctx, req)
	}
}

func tracingStreamInterceptor(_ interceptorOptions) grpc.StreamServerInterceptor {
	return func(srv any, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		traceID := incomingTraceID(ss.Context())
		if traceID == "" {
			traceID = generateTraceID()
		}
		ctx := context.WithValue(ss.Context(), contextKeyTraceID{}, traceID)
		wrapped := &wrappedServerStream{ServerStream: ss, ctx: ctx}
		_ = ss.SetHeader(metadata.Pairs(traceHeader, traceID))
		ss.SetTrailer(metadata.Pairs(traceHeader, traceID))
		return handler(srv, wrapped)
	}
}

func loggingUnaryInterceptor(opts interceptorOptions) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		start := time.Now()
		if opts.metrics != nil {
			opts.metrics.StartRequest()
			defer opts.metrics.FinishRequest()
		}
		resp, err := handler(ctx, req)
		code := status.Code(err)
		duration := time.Since(start)
		traceID := traceIDFromContext(ctx)
		if opts.metrics != nil {
			opts.metrics.Record(info.FullMethod, code, duration, traceID)
		}
		if opts.logger != nil {
			opts.logger.Printf("grpc unary method=%s code=%s duration=%s trace_id=%s peer=%s", info.FullMethod, code.String(), duration, traceID, peerAddress(ctx))
		}
		return resp, err
	}
}

func loggingStreamInterceptor(opts interceptorOptions) grpc.StreamServerInterceptor {
	return func(srv any, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		start := time.Now()
		if opts.metrics != nil {
			opts.metrics.StartRequest()
			defer opts.metrics.FinishRequest()
		}
		err := handler(srv, ss)
		code := status.Code(err)
		duration := time.Since(start)
		traceID := traceIDFromContext(ss.Context())
		if opts.metrics != nil {
			opts.metrics.Record(info.FullMethod, code, duration, traceID)
		}
		if opts.logger != nil {
			opts.logger.Printf("grpc stream method=%s code=%s duration=%s trace_id=%s peer=%s", info.FullMethod, code.String(), duration, traceID, peerAddress(ss.Context()))
		}
		return err
	}
}

func authUnaryInterceptor(opts interceptorOptions) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		if opts.authToken == "" || opts.allowMethod(info.FullMethod) {
			return handler(ctx, req)
		}
		if !isAuthorized(ctx, opts.authToken) {
			if opts.metrics != nil {
				opts.metrics.RecordAuthFailure(info.FullMethod)
			}
			return nil, status.Error(codes.Unauthenticated, "missing or invalid auth token")
		}
		return handler(ctx, req)
	}
}

func authStreamInterceptor(opts interceptorOptions) grpc.StreamServerInterceptor {
	return func(srv any, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		if opts.authToken == "" || opts.allowMethod(info.FullMethod) {
			return handler(srv, ss)
		}
		if !isAuthorized(ss.Context(), opts.authToken) {
			if opts.metrics != nil {
				opts.metrics.RecordAuthFailure(info.FullMethod)
			}
			return status.Error(codes.Unauthenticated, "missing or invalid auth token")
		}
		return handler(srv, ss)
	}
}

func isAuthorized(ctx context.Context, token string) bool {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return false
	}
	token = strings.TrimSpace(token)
	if token == "" {
		return true
	}
	for _, v := range md.Get("x-api-key") {
		if strings.TrimSpace(v) == token {
			return true
		}
	}
	for _, v := range md.Get("authorization") {
		v = strings.TrimSpace(v)
		if strings.HasPrefix(strings.ToLower(v), "bearer ") && strings.TrimSpace(v[7:]) == token {
			return true
		}
	}
	return false
}

func incomingTraceID(ctx context.Context) string {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ""
	}
	for _, v := range md.Get(traceHeader) {
		v = strings.TrimSpace(v)
		if v != "" {
			return v
		}
	}
	return ""
}

func traceIDFromContext(ctx context.Context) string {
	if v, ok := ctx.Value(contextKeyTraceID{}).(string); ok {
		return v
	}
	return ""
}

func generateTraceID() string {
	buf := make([]byte, 16)
	if _, err := rand.Read(buf); err == nil {
		return hex.EncodeToString(buf)
	}
	counter := atomic.AddUint64(&traceCounter, 1)
	now := uint64(time.Now().UnixNano())
	fallback := make([]byte, 16)
	for i := 0; i < 8; i++ {
		fallback[i] = byte(now >> (i * 8))
		fallback[8+i] = byte(counter >> (i * 8))
	}
	return hex.EncodeToString(fallback)
}

func peerAddress(ctx context.Context) string {
	p, ok := peer.FromContext(ctx)
	if !ok || p == nil || p.Addr == nil {
		return "unknown"
	}
	return p.Addr.String()
}
