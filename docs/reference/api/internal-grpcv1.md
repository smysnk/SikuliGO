# API: `internal/grpcv1`

[Back to API Index](./)

<style>
  .api-type { color: #0f766e; font-weight: 700; }
  .api-func { color: #1d4ed8; font-weight: 700; }
  .api-method { color: #7c3aed; font-weight: 700; }
  .api-signature { font-family: ui-monospace, SFMono-Regular, Menlo, Consolas, "Liberation Mono", monospace; }
</style>

Legend: <span class="api-type">Type</span>, <span class="api-func">Function</span>, <span class="api-method">Method</span>

Package: `package grpcv1 // import "github.com/smysnk/sikuligo/internal/grpcv1"`

## Symbol Index

### Types

- <span class="api-type">[`MethodSnapshot`](#type-methodsnapshot)</span>
- <span class="api-type">[`MetricsRegistry`](#type-metricsregistry)</span>
- <span class="api-type">[`MetricsSnapshot`](#type-metricssnapshot)</span>
- <span class="api-type">[`MetricsSnapshotProvider`](#type-metricssnapshotprovider)</span>
- <span class="api-type">[`Server`](#type-server)</span>
- <span class="api-type">[`ServerOption`](#type-serveroption)</span>
- <span class="api-type">[`SessionTracker`](#type-sessiontracker)</span>
- <span class="api-type">[`StoreMetricsProvider`](#type-storemetricsprovider)</span>

### Functions

- <span class="api-func">[`NewAdminMux`](#func-newadminmux)</span>
- <span class="api-func">[`StreamInterceptors`](#func-streaminterceptors)</span>
- <span class="api-func">[`UnaryInterceptors`](#func-unaryinterceptors)</span>
- <span class="api-func">[`NewMetricsRegistry`](#func-newmetricsregistry)</span>
- <span class="api-func">[`NewServer`](#func-newserver)</span>
- <span class="api-func">[`WithCaptureScreen`](#func-withcapturescreen)</span>
- <span class="api-func">[`WithClickOnScreen`](#func-withclickonscreen)</span>
- <span class="api-func">[`WithFinderFactory`](#func-withfinderfactory)</span>
- <span class="api-func">[`WithFinderWithEngineFactory`](#func-withfinderwithenginefactory)</span>
- <span class="api-func">[`NewSessionTracker`](#func-newsessiontracker)</span>
- <span class="api-func">[`NewStoreMetricsProvider`](#func-newstoremetricsprovider)</span>

### Methods

- <span class="api-method">[`MetricsRegistry.FinishRequest`](#method-metricsregistry-finishrequest)</span>
- <span class="api-method">[`MetricsRegistry.Record`](#method-metricsregistry-record)</span>
- <span class="api-method">[`MetricsRegistry.RecordAuthFailure`](#method-metricsregistry-recordauthfailure)</span>
- <span class="api-method">[`MetricsRegistry.Snapshot`](#method-metricsregistry-snapshot)</span>
- <span class="api-method">[`MetricsRegistry.StartRequest`](#method-metricsregistry-startrequest)</span>
- <span class="api-method">[`Server.Click`](#method-server-click)</span>
- <span class="api-method">[`Server.ClickOnScreen`](#method-server-clickonscreen)</span>
- <span class="api-method">[`Server.CloseApp`](#method-server-closeapp)</span>
- <span class="api-method">[`Server.ExistsOnScreen`](#method-server-existsonscreen)</span>
- <span class="api-method">[`Server.Find`](#method-server-find)</span>
- <span class="api-method">[`Server.FindAll`](#method-server-findall)</span>
- <span class="api-method">[`Server.FindOnScreen`](#method-server-findonscreen)</span>
- <span class="api-method">[`Server.FindText`](#method-server-findtext)</span>
- <span class="api-method">[`Server.FocusApp`](#method-server-focusapp)</span>
- <span class="api-method">[`Server.Hotkey`](#method-server-hotkey)</span>
- <span class="api-method">[`Server.IsAppRunning`](#method-server-isapprunning)</span>
- <span class="api-method">[`Server.ListWindows`](#method-server-listwindows)</span>
- <span class="api-method">[`Server.MoveMouse`](#method-server-movemouse)</span>
- <span class="api-method">[`Server.ObserveAppear`](#method-server-observeappear)</span>
- <span class="api-method">[`Server.ObserveChange`](#method-server-observechange)</span>
- <span class="api-method">[`Server.ObserveVanish`](#method-server-observevanish)</span>
- <span class="api-method">[`Server.OpenApp`](#method-server-openapp)</span>
- <span class="api-method">[`Server.ReadText`](#method-server-readtext)</span>
- <span class="api-method">[`Server.TypeText`](#method-server-typetext)</span>
- <span class="api-method">[`Server.WaitOnScreen`](#method-server-waitonscreen)</span>
- <span class="api-method">[`SessionTracker.HandleConn`](#method-sessiontracker-handleconn)</span>
- <span class="api-method">[`SessionTracker.HandleRPC`](#method-sessiontracker-handlerpc)</span>
- <span class="api-method">[`SessionTracker.RecordInteraction`](#method-sessiontracker-recordinteraction)</span>
- <span class="api-method">[`SessionTracker.TagConn`](#method-sessiontracker-tagconn)</span>
- <span class="api-method">[`SessionTracker.TagRPC`](#method-sessiontracker-tagrpc)</span>
- <span class="api-method">[`StoreMetricsProvider.Snapshot`](#method-storemetricsprovider-snapshot)</span>

## Declarations

### Types

#### <a id="type-methodsnapshot"></a><span class="api-type">Type</span> `MethodSnapshot`

- Signature: <span class="api-signature">`type MethodSnapshot struct {`</span>

#### <a id="type-metricsregistry"></a><span class="api-type">Type</span> `MetricsRegistry`

- Signature: <span class="api-signature">`type MetricsRegistry struct {`</span>

#### <a id="type-metricssnapshot"></a><span class="api-type">Type</span> `MetricsSnapshot`

- Signature: <span class="api-signature">`type MetricsSnapshot struct {`</span>

#### <a id="type-metricssnapshotprovider"></a><span class="api-type">Type</span> `MetricsSnapshotProvider`

- Signature: <span class="api-signature">`type MetricsSnapshotProvider interface {`</span>

#### <a id="type-server"></a><span class="api-type">Type</span> `Server`

- Signature: <span class="api-signature">`type Server struct {`</span>

#### <a id="type-serveroption"></a><span class="api-type">Type</span> `ServerOption`

- Signature: <span class="api-signature">`type ServerOption func(*Server)`</span>
- Uses: [`Server`](#type-server)

#### <a id="type-sessiontracker"></a><span class="api-type">Type</span> `SessionTracker`

- Signature: <span class="api-signature">`type SessionTracker struct {`</span>

#### <a id="type-storemetricsprovider"></a><span class="api-type">Type</span> `StoreMetricsProvider`

- Signature: <span class="api-signature">`type StoreMetricsProvider struct {`</span>

### Functions

#### <a id="func-newadminmux"></a><span class="api-func">Function</span> `NewAdminMux`

- Signature: <span class="api-signature">`func NewAdminMux(provider MetricsSnapshotProvider, store *sessionstore.Store) *http.ServeMux`</span>
- Uses: [`MetricsSnapshotProvider`](#type-metricssnapshotprovider)

#### <a id="func-streaminterceptors"></a><span class="api-func">Function</span> `StreamInterceptors`

- Signature: <span class="api-signature">`func StreamInterceptors(authToken string, logger *log.Logger, metrics *MetricsRegistry, tracker *SessionTracker) []grpc.StreamServerInterceptor`</span>
- Uses: [`MetricsRegistry`](#type-metricsregistry), [`SessionTracker`](#type-sessiontracker)

#### <a id="func-unaryinterceptors"></a><span class="api-func">Function</span> `UnaryInterceptors`

- Signature: <span class="api-signature">`func UnaryInterceptors(authToken string, logger *log.Logger, metrics *MetricsRegistry, tracker *SessionTracker) []grpc.UnaryServerInterceptor`</span>
- Uses: [`MetricsRegistry`](#type-metricsregistry), [`SessionTracker`](#type-sessiontracker)

#### <a id="func-newmetricsregistry"></a><span class="api-func">Function</span> `NewMetricsRegistry`

- Signature: <span class="api-signature">`func NewMetricsRegistry() *MetricsRegistry`</span>
- Uses: [`MetricsRegistry`](#type-metricsregistry)

#### <a id="func-newserver"></a><span class="api-func">Function</span> `NewServer`

- Signature: <span class="api-signature">`func NewServer(opts ...ServerOption) *Server`</span>
- Uses: [`ServerOption`](#type-serveroption), [`Server`](#type-server)

#### <a id="func-withcapturescreen"></a><span class="api-func">Function</span> `WithCaptureScreen`

- Signature: <span class="api-signature">`func WithCaptureScreen(fn func(context.Context, string) (*sikuli.Image, error)) ServerOption`</span>
- Uses: [`ServerOption`](#type-serveroption)

#### <a id="func-withclickonscreen"></a><span class="api-func">Function</span> `WithClickOnScreen`

- Signature: <span class="api-signature">`func WithClickOnScreen(fn func(int, int, sikuli.InputOptions) error) ServerOption`</span>
- Uses: [`ServerOption`](#type-serveroption)

#### <a id="func-withfinderfactory"></a><span class="api-func">Function</span> `WithFinderFactory`

- Signature: <span class="api-signature">`func WithFinderFactory(fn func(*sikuli.Image) (*sikuli.Finder, error)) ServerOption`</span>
- Uses: [`ServerOption`](#type-serveroption)

#### <a id="func-withfinderwithenginefactory"></a><span class="api-func">Function</span> `WithFinderWithEngineFactory`

- Signature: <span class="api-signature">`func WithFinderWithEngineFactory(fn func(*sikuli.Image, cv.MatcherEngine) (*sikuli.Finder, error)) ServerOption`</span>
- Uses: [`ServerOption`](#type-serveroption)

#### <a id="func-newsessiontracker"></a><span class="api-func">Function</span> `NewSessionTracker`

- Signature: <span class="api-signature">`func NewSessionTracker(store *sessionstore.Store, apiSessionID uint, logger *log.Logger) *SessionTracker`</span>
- Uses: [`SessionTracker`](#type-sessiontracker)

#### <a id="func-newstoremetricsprovider"></a><span class="api-func">Function</span> `NewStoreMetricsProvider`

- Signature: <span class="api-signature">`func NewStoreMetricsProvider(store *sessionstore.Store) *StoreMetricsProvider`</span>
- Uses: [`StoreMetricsProvider`](#type-storemetricsprovider)

### Methods

#### <a id="method-metricsregistry-finishrequest"></a><span class="api-method">Method</span> `MetricsRegistry.FinishRequest`

- Signature: <span class="api-signature">`func (m *MetricsRegistry) FinishRequest()`</span>

#### <a id="method-metricsregistry-record"></a><span class="api-method">Method</span> `MetricsRegistry.Record`

- Signature: <span class="api-signature">`func (m *MetricsRegistry) Record(method string, code codes.Code, latency time.Duration, traceID string)`</span>

#### <a id="method-metricsregistry-recordauthfailure"></a><span class="api-method">Method</span> `MetricsRegistry.RecordAuthFailure`

- Signature: <span class="api-signature">`func (m *MetricsRegistry) RecordAuthFailure(method string)`</span>

#### <a id="method-metricsregistry-snapshot"></a><span class="api-method">Method</span> `MetricsRegistry.Snapshot`

- Signature: <span class="api-signature">`func (m *MetricsRegistry) Snapshot() MetricsSnapshot`</span>
- Uses: [`MetricsSnapshot`](#type-metricssnapshot)

#### <a id="method-metricsregistry-startrequest"></a><span class="api-method">Method</span> `MetricsRegistry.StartRequest`

- Signature: <span class="api-signature">`func (m *MetricsRegistry) StartRequest()`</span>

#### <a id="method-server-click"></a><span class="api-method">Method</span> `Server.Click`

- Signature: <span class="api-signature">`func (s *Server) Click(_ context.Context, req *pb.ClickRequest) (*pb.ActionResponse, error)`</span>

#### <a id="method-server-clickonscreen"></a><span class="api-method">Method</span> `Server.ClickOnScreen`

- Signature: <span class="api-signature">`func (s *Server) ClickOnScreen(ctx context.Context, req *pb.ClickOnScreenRequest) (*pb.FindResponse, error)`</span>

#### <a id="method-server-closeapp"></a><span class="api-method">Method</span> `Server.CloseApp`

- Signature: <span class="api-signature">`func (s *Server) CloseApp(_ context.Context, req *pb.AppActionRequest) (*pb.ActionResponse, error)`</span>

#### <a id="method-server-existsonscreen"></a><span class="api-method">Method</span> `Server.ExistsOnScreen`

- Signature: <span class="api-signature">`func (s *Server) ExistsOnScreen(ctx context.Context, req *pb.ExistsOnScreenRequest) (*pb.ExistsOnScreenResponse, error)`</span>

#### <a id="method-server-find"></a><span class="api-method">Method</span> `Server.Find`

- Signature: <span class="api-signature">`func (s *Server) Find(ctx context.Context, req *pb.FindRequest) (*pb.FindResponse, error)`</span>

#### <a id="method-server-findall"></a><span class="api-method">Method</span> `Server.FindAll`

- Signature: <span class="api-signature">`func (s *Server) FindAll(ctx context.Context, req *pb.FindRequest) (*pb.FindAllResponse, error)`</span>

#### <a id="method-server-findonscreen"></a><span class="api-method">Method</span> `Server.FindOnScreen`

- Signature: <span class="api-signature">`func (s *Server) FindOnScreen(ctx context.Context, req *pb.FindOnScreenRequest) (*pb.FindResponse, error)`</span>

#### <a id="method-server-findtext"></a><span class="api-method">Method</span> `Server.FindText`

- Signature: <span class="api-signature">`func (s *Server) FindText(_ context.Context, req *pb.FindTextRequest) (*pb.FindTextResponse, error)`</span>

#### <a id="method-server-focusapp"></a><span class="api-method">Method</span> `Server.FocusApp`

- Signature: <span class="api-signature">`func (s *Server) FocusApp(_ context.Context, req *pb.AppActionRequest) (*pb.ActionResponse, error)`</span>

#### <a id="method-server-hotkey"></a><span class="api-method">Method</span> `Server.Hotkey`

- Signature: <span class="api-signature">`func (s *Server) Hotkey(_ context.Context, req *pb.HotkeyRequest) (*pb.ActionResponse, error)`</span>

#### <a id="method-server-isapprunning"></a><span class="api-method">Method</span> `Server.IsAppRunning`

- Signature: <span class="api-signature">`func (s *Server) IsAppRunning(_ context.Context, req *pb.AppActionRequest) (*pb.IsAppRunningResponse, error)`</span>

#### <a id="method-server-listwindows"></a><span class="api-method">Method</span> `Server.ListWindows`

- Signature: <span class="api-signature">`func (s *Server) ListWindows(_ context.Context, req *pb.AppActionRequest) (*pb.ListWindowsResponse, error)`</span>

#### <a id="method-server-movemouse"></a><span class="api-method">Method</span> `Server.MoveMouse`

- Signature: <span class="api-signature">`func (s *Server) MoveMouse(_ context.Context, req *pb.MoveMouseRequest) (*pb.ActionResponse, error)`</span>

#### <a id="method-server-observeappear"></a><span class="api-method">Method</span> `Server.ObserveAppear`

- Signature: <span class="api-signature">`func (s *Server) ObserveAppear(_ context.Context, req *pb.ObserveRequest) (*pb.ObserveResponse, error)`</span>

#### <a id="method-server-observechange"></a><span class="api-method">Method</span> `Server.ObserveChange`

- Signature: <span class="api-signature">`func (s *Server) ObserveChange(_ context.Context, req *pb.ObserveChangeRequest) (*pb.ObserveResponse, error)`</span>

#### <a id="method-server-observevanish"></a><span class="api-method">Method</span> `Server.ObserveVanish`

- Signature: <span class="api-signature">`func (s *Server) ObserveVanish(_ context.Context, req *pb.ObserveRequest) (*pb.ObserveResponse, error)`</span>

#### <a id="method-server-openapp"></a><span class="api-method">Method</span> `Server.OpenApp`

- Signature: <span class="api-signature">`func (s *Server) OpenApp(_ context.Context, req *pb.AppActionRequest) (*pb.ActionResponse, error)`</span>

#### <a id="method-server-readtext"></a><span class="api-method">Method</span> `Server.ReadText`

- Signature: <span class="api-signature">`func (s *Server) ReadText(_ context.Context, req *pb.ReadTextRequest) (*pb.ReadTextResponse, error)`</span>

#### <a id="method-server-typetext"></a><span class="api-method">Method</span> `Server.TypeText`

- Signature: <span class="api-signature">`func (s *Server) TypeText(_ context.Context, req *pb.TypeTextRequest) (*pb.ActionResponse, error)`</span>

#### <a id="method-server-waitonscreen"></a><span class="api-method">Method</span> `Server.WaitOnScreen`

- Signature: <span class="api-signature">`func (s *Server) WaitOnScreen(ctx context.Context, req *pb.WaitOnScreenRequest) (*pb.FindResponse, error)`</span>

#### <a id="method-sessiontracker-handleconn"></a><span class="api-method">Method</span> `SessionTracker.HandleConn`

- Signature: <span class="api-signature">`func (s *SessionTracker) HandleConn(ctx context.Context, conn stats.ConnStats)`</span>

#### <a id="method-sessiontracker-handlerpc"></a><span class="api-method">Method</span> `SessionTracker.HandleRPC`

- Signature: <span class="api-signature">`func (s *SessionTracker) HandleRPC(context.Context, stats.RPCStats)`</span>

#### <a id="method-sessiontracker-recordinteraction"></a><span class="api-method">Method</span> `SessionTracker.RecordInteraction`

- Signature: <span class="api-signature">`func (s *SessionTracker) RecordInteraction(ctx context.Context, method string, code codes.Code, duration time.Duration, traceID string)`</span>

#### <a id="method-sessiontracker-tagconn"></a><span class="api-method">Method</span> `SessionTracker.TagConn`

- Signature: <span class="api-signature">`func (s *SessionTracker) TagConn(ctx context.Context, info *stats.ConnTagInfo) context.Context`</span>

#### <a id="method-sessiontracker-tagrpc"></a><span class="api-method">Method</span> `SessionTracker.TagRPC`

- Signature: <span class="api-signature">`func (s *SessionTracker) TagRPC(ctx context.Context, _ *stats.RPCTagInfo) context.Context`</span>

#### <a id="method-storemetricsprovider-snapshot"></a><span class="api-method">Method</span> `StoreMetricsProvider.Snapshot`

- Signature: <span class="api-signature">`func (p *StoreMetricsProvider) Snapshot() MetricsSnapshot`</span>
- Uses: [`MetricsSnapshot`](#type-metricssnapshot)

## Raw Package Doc

```text
package grpcv1 // import "github.com/smysnk/sikuligo/internal/grpcv1"


FUNCTIONS

func NewAdminMux(provider MetricsSnapshotProvider, store *sessionstore.Store) *http.ServeMux
func StreamInterceptors(authToken string, logger *log.Logger, metrics *MetricsRegistry, tracker *SessionTracker) []grpc.StreamServerInterceptor
func UnaryInterceptors(authToken string, logger *log.Logger, metrics *MetricsRegistry, tracker *SessionTracker) []grpc.UnaryServerInterceptor

TYPES

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

type MetricsRegistry struct {
	// Has unexported fields.
}

func NewMetricsRegistry() *MetricsRegistry

func (m *MetricsRegistry) FinishRequest()

func (m *MetricsRegistry) Record(method string, code codes.Code, latency time.Duration, traceID string)

func (m *MetricsRegistry) RecordAuthFailure(method string)

func (m *MetricsRegistry) Snapshot() MetricsSnapshot

func (m *MetricsRegistry) StartRequest()

type MetricsSnapshot struct {
	StartedAtRFC3339  string           `json:"started_at_rfc3339"`
	UptimeSeconds     int64            `json:"uptime_seconds"`
	Inflight          int64            `json:"inflight"`
	TotalRequests     uint64           `json:"total_requests"`
	TotalErrors       uint64           `json:"total_errors"`
	TotalAuthFailures uint64           `json:"total_auth_failures"`
	Methods           []MethodSnapshot `json:"methods"`
}

type MetricsSnapshotProvider interface {
	Snapshot() MetricsSnapshot
}

type Server struct {
	pb.UnimplementedSikuliServiceServer

	// Has unexported fields.
}

func NewServer(opts ...ServerOption) *Server

func (s *Server) Click(_ context.Context, req *pb.ClickRequest) (*pb.ActionResponse, error)

func (s *Server) ClickOnScreen(ctx context.Context, req *pb.ClickOnScreenRequest) (*pb.FindResponse, error)

func (s *Server) CloseApp(_ context.Context, req *pb.AppActionRequest) (*pb.ActionResponse, error)

func (s *Server) ExistsOnScreen(ctx context.Context, req *pb.ExistsOnScreenRequest) (*pb.ExistsOnScreenResponse, error)

func (s *Server) Find(ctx context.Context, req *pb.FindRequest) (*pb.FindResponse, error)

func (s *Server) FindAll(ctx context.Context, req *pb.FindRequest) (*pb.FindAllResponse, error)

func (s *Server) FindOnScreen(ctx context.Context, req *pb.FindOnScreenRequest) (*pb.FindResponse, error)

func (s *Server) FindText(_ context.Context, req *pb.FindTextRequest) (*pb.FindTextResponse, error)

func (s *Server) FocusApp(_ context.Context, req *pb.AppActionRequest) (*pb.ActionResponse, error)

func (s *Server) Hotkey(_ context.Context, req *pb.HotkeyRequest) (*pb.ActionResponse, error)

func (s *Server) IsAppRunning(_ context.Context, req *pb.AppActionRequest) (*pb.IsAppRunningResponse, error)

func (s *Server) ListWindows(_ context.Context, req *pb.AppActionRequest) (*pb.ListWindowsResponse, error)

func (s *Server) MoveMouse(_ context.Context, req *pb.MoveMouseRequest) (*pb.ActionResponse, error)

func (s *Server) ObserveAppear(_ context.Context, req *pb.ObserveRequest) (*pb.ObserveResponse, error)

func (s *Server) ObserveChange(_ context.Context, req *pb.ObserveChangeRequest) (*pb.ObserveResponse, error)

func (s *Server) ObserveVanish(_ context.Context, req *pb.ObserveRequest) (*pb.ObserveResponse, error)

func (s *Server) OpenApp(_ context.Context, req *pb.AppActionRequest) (*pb.ActionResponse, error)

func (s *Server) ReadText(_ context.Context, req *pb.ReadTextRequest) (*pb.ReadTextResponse, error)

func (s *Server) TypeText(_ context.Context, req *pb.TypeTextRequest) (*pb.ActionResponse, error)

func (s *Server) WaitOnScreen(ctx context.Context, req *pb.WaitOnScreenRequest) (*pb.FindResponse, error)

type ServerOption func(*Server)

func WithCaptureScreen(fn func(context.Context, string) (*sikuli.Image, error)) ServerOption

func WithClickOnScreen(fn func(int, int, sikuli.InputOptions) error) ServerOption

func WithFinderFactory(fn func(*sikuli.Image) (*sikuli.Finder, error)) ServerOption

func WithFinderWithEngineFactory(fn func(*sikuli.Image, cv.MatcherEngine) (*sikuli.Finder, error)) ServerOption

type SessionTracker struct {
	// Has unexported fields.
}

func NewSessionTracker(store *sessionstore.Store, apiSessionID uint, logger *log.Logger) *SessionTracker

func (s *SessionTracker) HandleConn(ctx context.Context, conn stats.ConnStats)

func (s *SessionTracker) HandleRPC(context.Context, stats.RPCStats)

func (s *SessionTracker) RecordInteraction(ctx context.Context, method string, code codes.Code, duration time.Duration, traceID string)

func (s *SessionTracker) TagConn(ctx context.Context, info *stats.ConnTagInfo) context.Context

func (s *SessionTracker) TagRPC(ctx context.Context, _ *stats.RPCTagInfo) context.Context

type StoreMetricsProvider struct {
	// Has unexported fields.
}

func NewStoreMetricsProvider(store *sessionstore.Store) *StoreMetricsProvider

func (p *StoreMetricsProvider) Snapshot() MetricsSnapshot

```
