package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/smysnk/sikuligo/internal/grpcv1"
	pb "github.com/smysnk/sikuligo/internal/grpcv1/pb"
	"github.com/smysnk/sikuligo/internal/sessionstore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	listenAddr := flag.String("listen", ":50051", "gRPC listen address")
	adminListenAddr := flag.String("admin-listen", ":8080", "admin HTTP listen address for health/metrics/dashboard; empty disables admin server")
	sqlitePath := flag.String("sqlite-path", "sikuligo.db", "sqlite datastore path for API sessions, client sessions, and interactions")
	authToken := flag.String("auth-token", os.Getenv("SIKULI_GRPC_AUTH_TOKEN"), "shared API token; accepted via metadata x-api-key or Authorization: Bearer <token>")
	enableReflection := flag.Bool("enable-reflection", true, "enable gRPC reflection")
	flag.Parse()

	logger := log.Default()
	metrics := grpcv1.NewMetricsRegistry()
	store, err := sessionstore.OpenSQLite(*sqlitePath)
	if err != nil {
		log.Fatalf("open sqlite store %s: %v", *sqlitePath, err)
	}
	defer func() {
		if err := store.Close(); err != nil {
			logger.Printf("sqlite close: %v", err)
		}
	}()
	apiSession, err := store.StartAPISession(context.Background(), sessionstore.APISessionStartInput{
		PID:             os.Getpid(),
		GRPCListenAddr:  *listenAddr,
		AdminListenAddr: *adminListenAddr,
	})
	if err != nil {
		log.Fatalf("create api session: %v", err)
	}
	tracker := grpcv1.NewSessionTracker(store, apiSession.ID, logger)
	defer func() {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := store.EndAPISession(ctx, apiSession.ID, time.Now().UTC()); err != nil {
			logger.Printf("api session close: %v", err)
		}
	}()

	lis, err := net.Listen("tcp", *listenAddr)
	if err != nil {
		log.Fatalf("listen %s: %v", *listenAddr, err)
	}

	srv := grpc.NewServer(
		grpc.StatsHandler(tracker),
		grpc.ChainUnaryInterceptor(grpcv1.UnaryInterceptors(*authToken, logger, metrics, tracker)...),
		grpc.ChainStreamInterceptor(grpcv1.StreamInterceptors(*authToken, logger, metrics, tracker)...),
	)
	pb.RegisterSikuliServiceServer(srv, grpcv1.NewServer())
	if *enableReflection {
		reflection.Register(srv)
	}

	grpcErrCh := make(chan error, 1)
	go func() {
		logger.Printf(
			"sikuligo listening grpc=%s auth=%t reflection=%t sqlite=%s api_session_id=%d",
			*listenAddr,
			*authToken != "",
			*enableReflection,
			*sqlitePath,
			apiSession.ID,
		)
		if err := srv.Serve(lis); err != nil {
			grpcErrCh <- fmt.Errorf("grpc serve: %w", err)
		}
	}()

	var adminSrv *http.Server
	adminErrCh := make(chan error, 1)
	if *adminListenAddr != "" {
		adminSrv = &http.Server{
			Addr:              *adminListenAddr,
			Handler:           grpcv1.NewAdminMux(metrics, store),
			ReadHeaderTimeout: 5 * time.Second,
		}
		go func() {
			logger.Printf("sikuligo admin listening http=%s endpoints=/healthz,/snapshot,/metrics,/dashboard", *adminListenAddr)
			if err := adminSrv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				adminErrCh <- fmt.Errorf("admin serve: %w", err)
			}
		}()
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	select {
	case sig := <-stop:
		logger.Printf("received signal: %s", sig)
	case err := <-grpcErrCh:
		logger.Printf("grpc server error: %v", err)
	case err := <-adminErrCh:
		logger.Printf("admin server error: %v", err)
	}

	if adminSrv != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := adminSrv.Shutdown(ctx); err != nil {
			logger.Printf("admin shutdown: %v", err)
		}
	}
	srv.GracefulStop()
}
