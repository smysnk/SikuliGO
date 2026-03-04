# API: `internal/sessionstore`

[Back to API Index](./)

<style>
  .api-type { color: #0f766e; font-weight: 700; }
  .api-func { color: #1d4ed8; font-weight: 700; }
  .api-method { color: #7c3aed; font-weight: 700; }
  .api-signature { font-family: ui-monospace, SFMono-Regular, Menlo, Consolas, "Liberation Mono", monospace; }
</style>

Legend: <span class="api-type">Type</span>, <span class="api-func">Function</span>, <span class="api-method">Method</span>

Package: `package sessionstore // import "github.com/smysnk/sikuligo/internal/sessionstore"`

## Symbol Index

### Types

- <span class="api-type">[`APISession`](#type-apisession)</span>
- <span class="api-type">[`APISessionStartInput`](#type-apisessionstartinput)</span>
- <span class="api-type">[`ClientSession`](#type-clientsession)</span>
- <span class="api-type">[`ClientSessionStartInput`](#type-clientsessionstartinput)</span>
- <span class="api-type">[`Interaction`](#type-interaction)</span>
- <span class="api-type">[`InteractionInput`](#type-interactioninput)</span>
- <span class="api-type">[`MethodMetrics`](#type-methodmetrics)</span>
- <span class="api-type">[`Store`](#type-store)</span>

### Functions

- <span class="api-func">[`OpenSQLite`](#func-opensqlite)</span>

### Methods

- <span class="api-method">[`Store.Close`](#method-store-close)</span>
- <span class="api-method">[`Store.CountAPISessions`](#method-store-countapisessions)</span>
- <span class="api-method">[`Store.CountClientSessions`](#method-store-countclientsessions)</span>
- <span class="api-method">[`Store.CountInteractions`](#method-store-countinteractions)</span>
- <span class="api-method">[`Store.EndAPISession`](#method-store-endapisession)</span>
- <span class="api-method">[`Store.EndClientSession`](#method-store-endclientsession)</span>
- <span class="api-method">[`Store.LatestAPISession`](#method-store-latestapisession)</span>
- <span class="api-method">[`Store.ListClientSessionsByAPI`](#method-store-listclientsessionsbyapi)</span>
- <span class="api-method">[`Store.ListInteractionsByClient`](#method-store-listinteractionsbyclient)</span>
- <span class="api-method">[`Store.ListRecentAPISessions`](#method-store-listrecentapisessions)</span>
- <span class="api-method">[`Store.MethodMetricsByAPISession`](#method-store-methodmetricsbyapisession)</span>
- <span class="api-method">[`Store.RecordInteraction`](#method-store-recordinteraction)</span>
- <span class="api-method">[`Store.StartAPISession`](#method-store-startapisession)</span>
- <span class="api-method">[`Store.StartClientSession`](#method-store-startclientsession)</span>

## Declarations

### Types

#### <a id="type-apisession"></a><span class="api-type">Type</span> `APISession`

- Signature: <span class="api-signature">`type APISession struct {`</span>

#### <a id="type-apisessionstartinput"></a><span class="api-type">Type</span> `APISessionStartInput`

- Signature: <span class="api-signature">`type APISessionStartInput struct {`</span>

#### <a id="type-clientsession"></a><span class="api-type">Type</span> `ClientSession`

- Signature: <span class="api-signature">`type ClientSession struct {`</span>

#### <a id="type-clientsessionstartinput"></a><span class="api-type">Type</span> `ClientSessionStartInput`

- Signature: <span class="api-signature">`type ClientSessionStartInput struct {`</span>

#### <a id="type-interaction"></a><span class="api-type">Type</span> `Interaction`

- Signature: <span class="api-signature">`type Interaction struct {`</span>

#### <a id="type-interactioninput"></a><span class="api-type">Type</span> `InteractionInput`

- Signature: <span class="api-signature">`type InteractionInput struct {`</span>

#### <a id="type-methodmetrics"></a><span class="api-type">Type</span> `MethodMetrics`

- Signature: <span class="api-signature">`type MethodMetrics struct {`</span>

#### <a id="type-store"></a><span class="api-type">Type</span> `Store`

- Signature: <span class="api-signature">`type Store struct {`</span>

### Functions

#### <a id="func-opensqlite"></a><span class="api-func">Function</span> `OpenSQLite`

- Signature: <span class="api-signature">`func OpenSQLite(path string) (*Store, error)`</span>
- Uses: [`Store`](#type-store)

### Methods

#### <a id="method-store-close"></a><span class="api-method">Method</span> `Store.Close`

- Signature: <span class="api-signature">`func (s *Store) Close() error`</span>

#### <a id="method-store-countapisessions"></a><span class="api-method">Method</span> `Store.CountAPISessions`

- Signature: <span class="api-signature">`func (s *Store) CountAPISessions(ctx context.Context) (int64, error)`</span>

#### <a id="method-store-countclientsessions"></a><span class="api-method">Method</span> `Store.CountClientSessions`

- Signature: <span class="api-signature">`func (s *Store) CountClientSessions(ctx context.Context) (int64, error)`</span>

#### <a id="method-store-countinteractions"></a><span class="api-method">Method</span> `Store.CountInteractions`

- Signature: <span class="api-signature">`func (s *Store) CountInteractions(ctx context.Context) (int64, error)`</span>

#### <a id="method-store-endapisession"></a><span class="api-method">Method</span> `Store.EndAPISession`

- Signature: <span class="api-signature">`func (s *Store) EndAPISession(ctx context.Context, apiSessionID uint, endedAt time.Time) error`</span>

#### <a id="method-store-endclientsession"></a><span class="api-method">Method</span> `Store.EndClientSession`

- Signature: <span class="api-signature">`func (s *Store) EndClientSession(ctx context.Context, clientSessionID uint, endedAt time.Time) error`</span>

#### <a id="method-store-latestapisession"></a><span class="api-method">Method</span> `Store.LatestAPISession`

- Signature: <span class="api-signature">`func (s *Store) LatestAPISession(ctx context.Context) (APISession, bool, error)`</span>
- Uses: [`APISession`](#type-apisession)

#### <a id="method-store-listclientsessionsbyapi"></a><span class="api-method">Method</span> `Store.ListClientSessionsByAPI`

- Signature: <span class="api-signature">`func (s *Store) ListClientSessionsByAPI(ctx context.Context, apiSessionID uint) ([]ClientSession, error)`</span>
- Uses: [`ClientSession`](#type-clientsession)

#### <a id="method-store-listinteractionsbyclient"></a><span class="api-method">Method</span> `Store.ListInteractionsByClient`

- Signature: <span class="api-signature">`func (s *Store) ListInteractionsByClient(ctx context.Context, clientSessionID uint, limit int) ([]Interaction, error)`</span>
- Uses: [`Interaction`](#type-interaction)

#### <a id="method-store-listrecentapisessions"></a><span class="api-method">Method</span> `Store.ListRecentAPISessions`

- Signature: <span class="api-signature">`func (s *Store) ListRecentAPISessions(ctx context.Context, limit int) ([]APISession, error)`</span>
- Uses: [`APISession`](#type-apisession)

#### <a id="method-store-methodmetricsbyapisession"></a><span class="api-method">Method</span> `Store.MethodMetricsByAPISession`

- Signature: <span class="api-signature">`func (s *Store) MethodMetricsByAPISession(ctx context.Context, apiSessionID uint) ([]MethodMetrics, error)`</span>
- Uses: [`MethodMetrics`](#type-methodmetrics)

#### <a id="method-store-recordinteraction"></a><span class="api-method">Method</span> `Store.RecordInteraction`

- Signature: <span class="api-signature">`func (s *Store) RecordInteraction(ctx context.Context, in InteractionInput) error`</span>
- Uses: [`InteractionInput`](#type-interactioninput)

#### <a id="method-store-startapisession"></a><span class="api-method">Method</span> `Store.StartAPISession`

- Signature: <span class="api-signature">`func (s *Store) StartAPISession(ctx context.Context, in APISessionStartInput) (APISession, error)`</span>
- Uses: [`APISessionStartInput`](#type-apisessionstartinput), [`APISession`](#type-apisession)

#### <a id="method-store-startclientsession"></a><span class="api-method">Method</span> `Store.StartClientSession`

- Signature: <span class="api-signature">`func (s *Store) StartClientSession(ctx context.Context, in ClientSessionStartInput) (ClientSession, error)`</span>
- Uses: [`ClientSessionStartInput`](#type-clientsessionstartinput), [`ClientSession`](#type-clientsession)

## Raw Package Doc

```text
package sessionstore // import "github.com/smysnk/sikuligo/internal/sessionstore"


TYPES

type APISession struct {
	ID              uint `gorm:"primaryKey"`
	SessionKey      string
	PID             int
	GRPCListenAddr  string
	AdminListenAddr string
	StartedAt       time.Time
	EndedAt         *time.Time
}

type APISessionStartInput struct {
	PID             int
	GRPCListenAddr  string
	AdminListenAddr string
}

type ClientSession struct {
	ID           uint `gorm:"primaryKey"`
	APISessionID uint
	SessionKey   string
	ConnectionID string
	RemoteAddr   string
	LocalAddr    string
	StartedAt    time.Time
	EndedAt      *time.Time
	LastSeenAt   time.Time
}

type ClientSessionStartInput struct {
	APISessionID uint
	ConnectionID string
	RemoteAddr   string
	LocalAddr    string
}

type Interaction struct {
	ID              uint `gorm:"primaryKey"`
	APISessionID    uint
	ClientSessionID uint
	Method          string
	TraceID         string
	GRPCCode        string
	DurationMS      int64
	StartedAt       time.Time
	CompletedAt     time.Time
}

type InteractionInput struct {
	APISessionID    uint
	ClientSessionID uint
	Method          string
	TraceID         string
	GRPCCode        string
	DurationMS      int64
	StartedAt       time.Time
	CompletedAt     time.Time
}

type MethodMetrics struct {
	Method       string
	Requests     uint64
	Errors       uint64
	AuthFailures uint64
	AvgDuration  float64
	MaxDuration  int64
	LastCode     string
	LastTraceID  string
	LastSeen     time.Time
}

type Store struct {
	// Has unexported fields.
}

func OpenSQLite(path string) (*Store, error)

func (s *Store) Close() error

func (s *Store) CountAPISessions(ctx context.Context) (int64, error)

func (s *Store) CountClientSessions(ctx context.Context) (int64, error)

func (s *Store) CountInteractions(ctx context.Context) (int64, error)

func (s *Store) EndAPISession(ctx context.Context, apiSessionID uint, endedAt time.Time) error

func (s *Store) EndClientSession(ctx context.Context, clientSessionID uint, endedAt time.Time) error

func (s *Store) LatestAPISession(ctx context.Context) (APISession, bool, error)

func (s *Store) ListClientSessionsByAPI(ctx context.Context, apiSessionID uint) ([]ClientSession, error)

func (s *Store) ListInteractionsByClient(ctx context.Context, clientSessionID uint, limit int) ([]Interaction, error)

func (s *Store) ListRecentAPISessions(ctx context.Context, limit int) ([]APISession, error)

func (s *Store) MethodMetricsByAPISession(ctx context.Context, apiSessionID uint) ([]MethodMetrics, error)

func (s *Store) RecordInteraction(ctx context.Context, in InteractionInput) error

func (s *Store) StartAPISession(ctx context.Context, in APISessionStartInput) (APISession, error)

func (s *Store) StartClientSession(ctx context.Context, in ClientSessionStartInput) (ClientSession, error)

```
