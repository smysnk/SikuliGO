package sessionstore

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
)

type APISession struct {
	ID              uint `gorm:"primaryKey"`
	SessionKey      string
	PID             int
	GRPCListenAddr  string
	AdminListenAddr string
	StartedAt       time.Time
	EndedAt         *time.Time
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
	db *gorm.DB
}

type APISessionStartInput struct {
	PID             int
	GRPCListenAddr  string
	AdminListenAddr string
}

type ClientSessionStartInput struct {
	APISessionID uint
	ConnectionID string
	RemoteAddr   string
	LocalAddr    string
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

func OpenSQLite(path string) (*Store, error) {
	path = strings.TrimSpace(path)
	if path == "" {
		return nil, fmt.Errorf("sqlite path is empty")
	}
	if path != ":memory:" {
		if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
			return nil, fmt.Errorf("create sqlite dir: %w", err)
		}
	}

	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{
		Logger: gormlogger.Default.LogMode(gormlogger.Silent),
	})
	if err != nil {
		return nil, fmt.Errorf("open sqlite: %w", err)
	}

	if err := db.AutoMigrate(&APISession{}, &ClientSession{}, &Interaction{}); err != nil {
		return nil, fmt.Errorf("auto migrate session models: %w", err)
	}
	if err := db.Exec(`CREATE UNIQUE INDEX IF NOT EXISTS idx_api_sessions_session_key ON api_sessions(session_key);`).Error; err != nil {
		return nil, fmt.Errorf("create api session index: %w", err)
	}
	if err := db.Exec(`CREATE UNIQUE INDEX IF NOT EXISTS idx_client_sessions_session_key ON client_sessions(session_key);`).Error; err != nil {
		return nil, fmt.Errorf("create client session index: %w", err)
	}
	if err := db.Exec(`CREATE UNIQUE INDEX IF NOT EXISTS idx_client_sessions_connection_id ON client_sessions(connection_id);`).Error; err != nil {
		return nil, fmt.Errorf("create client connection index: %w", err)
	}

	return &Store{db: db}, nil
}

func (s *Store) Close() error {
	if s == nil || s.db == nil {
		return nil
	}
	sqlDB, err := s.db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

func (s *Store) StartAPISession(ctx context.Context, in APISessionStartInput) (APISession, error) {
	if s == nil || s.db == nil {
		return APISession{}, fmt.Errorf("store is nil")
	}
	now := time.Now().UTC()
	row := APISession{
		SessionKey:      newSessionKey("api"),
		PID:             in.PID,
		GRPCListenAddr:  in.GRPCListenAddr,
		AdminListenAddr: in.AdminListenAddr,
		StartedAt:       now,
	}
	if err := s.db.WithContext(ctx).Create(&row).Error; err != nil {
		return APISession{}, err
	}
	return row, nil
}

func (s *Store) EndAPISession(ctx context.Context, apiSessionID uint, endedAt time.Time) error {
	if s == nil || s.db == nil || apiSessionID == 0 {
		return nil
	}
	return s.db.WithContext(ctx).
		Model(&APISession{}).
		Where("id = ?", apiSessionID).
		Update("ended_at", endedAt.UTC()).
		Error
}

func (s *Store) StartClientSession(ctx context.Context, in ClientSessionStartInput) (ClientSession, error) {
	if s == nil || s.db == nil {
		return ClientSession{}, fmt.Errorf("store is nil")
	}
	now := time.Now().UTC()
	row := ClientSession{
		APISessionID: in.APISessionID,
		SessionKey:   newSessionKey("client"),
		ConnectionID: in.ConnectionID,
		RemoteAddr:   in.RemoteAddr,
		LocalAddr:    in.LocalAddr,
		StartedAt:    now,
		LastSeenAt:   now,
	}
	if err := s.db.WithContext(ctx).Create(&row).Error; err != nil {
		return ClientSession{}, err
	}
	return row, nil
}

func (s *Store) EndClientSession(ctx context.Context, clientSessionID uint, endedAt time.Time) error {
	if s == nil || s.db == nil || clientSessionID == 0 {
		return nil
	}
	return s.db.WithContext(ctx).
		Model(&ClientSession{}).
		Where("id = ?", clientSessionID).
		Updates(map[string]any{
			"ended_at":     endedAt.UTC(),
			"last_seen_at": endedAt.UTC(),
		}).
		Error
}

func (s *Store) RecordInteraction(ctx context.Context, in InteractionInput) error {
	if s == nil || s.db == nil {
		return fmt.Errorf("store is nil")
	}
	if in.ClientSessionID == 0 {
		return nil
	}
	if in.CompletedAt.IsZero() {
		in.CompletedAt = time.Now().UTC()
	}
	if in.StartedAt.IsZero() {
		in.StartedAt = in.CompletedAt
	}

	row := Interaction{
		APISessionID:    in.APISessionID,
		ClientSessionID: in.ClientSessionID,
		Method:          in.Method,
		TraceID:         in.TraceID,
		GRPCCode:        in.GRPCCode,
		DurationMS:      in.DurationMS,
		StartedAt:       in.StartedAt.UTC(),
		CompletedAt:     in.CompletedAt.UTC(),
	}
	tx := s.db.WithContext(ctx).Begin()
	if tx.Error != nil {
		return tx.Error
	}
	if err := tx.Create(&row).Error; err != nil {
		_ = tx.Rollback().Error
		return err
	}
	if err := tx.Model(&ClientSession{}).
		Where("id = ?", in.ClientSessionID).
		Update("last_seen_at", row.CompletedAt).
		Error; err != nil {
		_ = tx.Rollback().Error
		return err
	}
	return tx.Commit().Error
}

func (s *Store) ListRecentAPISessions(ctx context.Context, limit int) ([]APISession, error) {
	if s == nil || s.db == nil {
		return nil, fmt.Errorf("store is nil")
	}
	if limit <= 0 {
		limit = 50
	}
	var rows []APISession
	if err := s.db.WithContext(ctx).
		Order("started_at DESC").
		Limit(limit).
		Find(&rows).
		Error; err != nil {
		return nil, err
	}
	return rows, nil
}

func (s *Store) ListClientSessionsByAPI(ctx context.Context, apiSessionID uint) ([]ClientSession, error) {
	if s == nil || s.db == nil {
		return nil, fmt.Errorf("store is nil")
	}
	var rows []ClientSession
	if err := s.db.WithContext(ctx).
		Where("api_session_id = ?", apiSessionID).
		Order("started_at DESC").
		Find(&rows).
		Error; err != nil {
		return nil, err
	}
	return rows, nil
}

func (s *Store) ListInteractionsByClient(ctx context.Context, clientSessionID uint, limit int) ([]Interaction, error) {
	if s == nil || s.db == nil {
		return nil, fmt.Errorf("store is nil")
	}
	if limit <= 0 {
		limit = 200
	}
	var rows []Interaction
	if err := s.db.WithContext(ctx).
		Where("client_session_id = ?", clientSessionID).
		Order("started_at DESC").
		Limit(limit).
		Find(&rows).
		Error; err != nil {
		return nil, err
	}
	return rows, nil
}

func (s *Store) LatestAPISession(ctx context.Context) (APISession, bool, error) {
	if s == nil || s.db == nil {
		return APISession{}, false, fmt.Errorf("store is nil")
	}
	var row APISession
	err := s.db.WithContext(ctx).
		Where("ended_at IS NULL").
		Order("started_at DESC").
		Limit(1).
		Find(&row).
		Error
	if err != nil {
		return APISession{}, false, err
	}
	if row.ID != 0 {
		return row, true, nil
	}
	if err := s.db.WithContext(ctx).
		Order("started_at DESC").
		Limit(1).
		Find(&row).
		Error; err != nil {
		return APISession{}, false, err
	}
	if row.ID == 0 {
		return APISession{}, false, nil
	}
	return row, true, nil
}

func (s *Store) MethodMetricsByAPISession(ctx context.Context, apiSessionID uint) ([]MethodMetrics, error) {
	if s == nil || s.db == nil {
		return nil, fmt.Errorf("store is nil")
	}
	type row struct {
		Method       string
		Requests     uint64
		Errors       uint64
		AuthFailures uint64
		AvgDuration  float64
		MaxDuration  int64
	}
	var grouped []row
	grpcCodeCol := s.db.NamingStrategy.ColumnName("", "GRPCCode")
	durationCol := s.db.NamingStrategy.ColumnName("", "DurationMS")
	apiSessionCol := s.db.NamingStrategy.ColumnName("", "APISessionID")
	selectExpr := fmt.Sprintf(
		`method,
		 COUNT(*) AS requests,
		 SUM(CASE WHEN %s != 'OK' THEN 1 ELSE 0 END) AS errors,
		 SUM(CASE WHEN %s = 'Unauthenticated' THEN 1 ELSE 0 END) AS auth_failures,
		 AVG(%s) AS avg_duration,
		 MAX(%s) AS max_duration`,
		grpcCodeCol,
		grpcCodeCol,
		durationCol,
		durationCol,
	)
	if err := s.db.WithContext(ctx).
		Model(&Interaction{}).
		Select(selectExpr).
		Where(apiSessionCol+" = ?", apiSessionID).
		Group("method").
		Scan(&grouped).
		Error; err != nil {
		return nil, err
	}
	out := make([]MethodMetrics, 0, len(grouped))
	for _, g := range grouped {
		last := Interaction{}
		if err := s.db.WithContext(ctx).
			Where(apiSessionCol+" = ? AND method = ?", apiSessionID, g.Method).
			Order("started_at DESC").
			Limit(1).
			Find(&last).
			Error; err != nil {
			return nil, err
		}
		out = append(out, MethodMetrics{
			Method:       g.Method,
			Requests:     g.Requests,
			Errors:       g.Errors,
			AuthFailures: g.AuthFailures,
			AvgDuration:  g.AvgDuration,
			MaxDuration:  g.MaxDuration,
			LastCode:     last.GRPCCode,
			LastTraceID:  last.TraceID,
			LastSeen:     last.StartedAt.UTC(),
		})
	}
	return out, nil
}

func (s *Store) CountAPISessions(ctx context.Context) (int64, error) {
	if s == nil || s.db == nil {
		return 0, fmt.Errorf("store is nil")
	}
	var count int64
	if err := s.db.WithContext(ctx).Model(&APISession{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (s *Store) CountClientSessions(ctx context.Context) (int64, error) {
	if s == nil || s.db == nil {
		return 0, fmt.Errorf("store is nil")
	}
	var count int64
	if err := s.db.WithContext(ctx).Model(&ClientSession{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (s *Store) CountInteractions(ctx context.Context) (int64, error) {
	if s == nil || s.db == nil {
		return 0, fmt.Errorf("store is nil")
	}
	var count int64
	if err := s.db.WithContext(ctx).Model(&Interaction{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func newSessionKey(prefix string) string {
	buf := make([]byte, 12)
	if _, err := rand.Read(buf); err != nil {
		now := time.Now().UnixNano()
		return fmt.Sprintf("%s-%x", prefix, now)
	}
	return fmt.Sprintf("%s-%s", prefix, hex.EncodeToString(buf))
}
