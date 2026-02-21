package main

import (
	"flag"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/sikulix/portgo/internal/grpcv1"
	pb "github.com/sikulix/portgo/internal/grpcv1/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	listenAddr := flag.String("listen", ":50051", "gRPC listen address")
	flag.Parse()

	lis, err := net.Listen("tcp", *listenAddr)
	if err != nil {
		log.Fatalf("listen %s: %v", *listenAddr, err)
	}

	srv := grpc.NewServer()
	pb.RegisterSikuliServiceServer(srv, grpcv1.NewServer())
	reflection.Register(srv)

	go func() {
		log.Printf("sikuligrpc listening on %s", *listenAddr)
		if err := srv.Serve(lis); err != nil {
			log.Fatalf("serve: %v", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop
	srv.GracefulStop()
}
