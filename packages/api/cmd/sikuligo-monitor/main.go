package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/smysnk/sikuligo/internal/cv"
	"github.com/smysnk/sikuligo/internal/grpcv1"
	"github.com/smysnk/sikuligo/internal/sessionstore"
)

func main() {
	if !cv.OpenCVEnabled() {
		log.Fatal("sikuligo-monitor requires OpenCV-enabled builds; rebuild with -tags \"gosseract opencv gocv_specific_modules gocv_features2d gocv_calib3d\"")
	}

	listenAddr := flag.String("listen", ":8080", "HTTP listen address for monitor dashboard/session viewer")
	sqlitePath := flag.String("sqlite-path", "sikuligo.db", "sqlite datastore path shared with sikuligo")
	flag.Parse()

	logger := log.Default()
	store, err := sessionstore.OpenSQLite(*sqlitePath)
	if err != nil {
		log.Fatalf("open sqlite store %s: %v", *sqlitePath, err)
	}
	defer func() {
		if err := store.Close(); err != nil {
			logger.Printf("sqlite close: %v", err)
		}
	}()

	provider := grpcv1.NewStoreMetricsProvider(store)
	srv := &http.Server{
		Addr:              *listenAddr,
		Handler:           grpcv1.NewAdminMux(provider, store),
		ReadHeaderTimeout: 5 * time.Second,
	}

	errCh := make(chan error, 1)
	go func() {
		logger.Printf("sikuligo-monitor listening http=%s sqlite=%s endpoints=/dashboard,/sessions,/ws", *listenAddr, *sqlitePath)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			errCh <- fmt.Errorf("monitor serve: %w", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	select {
	case sig := <-stop:
		logger.Printf("received signal: %s", sig)
	case err := <-errCh:
		logger.Printf("monitor server error: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Printf("monitor shutdown: %v", err)
	}
}
