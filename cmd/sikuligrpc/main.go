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

	"github.com/sikulix/portgo/internal/grpcv1"
	pb "github.com/sikulix/portgo/internal/grpcv1/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	listenAddr := flag.String("listen", ":50051", "gRPC listen address")
	adminListenAddr := flag.String("admin-listen", ":8080", "admin HTTP listen address for health/metrics/dashboard; empty disables admin server")
	authToken := flag.String("auth-token", os.Getenv("SIKULI_GRPC_AUTH_TOKEN"), "shared API token; accepted via metadata x-api-key or Authorization: Bearer <token>")
	enableReflection := flag.Bool("enable-reflection", true, "enable gRPC reflection")
	flag.Parse()

	logger := log.Default()
	metrics := grpcv1.NewMetricsRegistry()

	lis, err := net.Listen("tcp", *listenAddr)
	if err != nil {
		log.Fatalf("listen %s: %v", *listenAddr, err)
	}

	srv := grpc.NewServer(
		grpc.ChainUnaryInterceptor(grpcv1.UnaryInterceptors(*authToken, logger, metrics)...),
		grpc.ChainStreamInterceptor(grpcv1.StreamInterceptors(*authToken, logger, metrics)...),
	)
	pb.RegisterSikuliServiceServer(srv, grpcv1.NewServer())
	if *enableReflection {
		reflection.Register(srv)
	}

	grpcErrCh := make(chan error, 1)
	go func() {
		logger.Printf("sikuligrpc listening grpc=%s auth=%t reflection=%t", *listenAddr, *authToken != "", *enableReflection)
		if err := srv.Serve(lis); err != nil {
			grpcErrCh <- fmt.Errorf("grpc serve: %w", err)
		}
	}()

	var adminSrv *http.Server
	adminErrCh := make(chan error, 1)
	if *adminListenAddr != "" {
		adminSrv = &http.Server{
			Addr:              *adminListenAddr,
			Handler:           grpcv1.NewAdminMux(metrics),
			ReadHeaderTimeout: 5 * time.Second,
		}
		go func() {
			logger.Printf("sikuligrpc admin listening http=%s endpoints=/healthz,/snapshot,/metrics,/dashboard", *adminListenAddr)
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
