package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	pb "github.com/lantspants/lootloadout/api/characterimage"
	"github.com/lantspants/lootloadout/backend/characterimage/characterimage"
	grpcServer "github.com/lantspants/lootloadout/backend/characterimage/cmd/grpc"
	"github.com/lantspants/lootloadout/backend/characterimage/postgres"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var grpcEndpoint = flag.Int("grpc-server-endpoint", 9999, "gRPC server endpoint")

func main() {
	l := log.New(os.Stdout, "", 0)
	flag.Parse()

	db := postgres.NewImageService(l)
	s := characterimage.NewCharacterImageService(l, db)
	r := grpcServer.NewCharacterImageServer(l, s)
	srv := grpc.NewServer()

	pb.RegisterImagesServer(srv, r)
	reflection.Register(srv)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *grpcEndpoint))
	if err != nil {
		l.Fatalf("Failed to listen: %v", err)
	}

	l.Printf("Initializing characterimage service at %s", lis.Addr().String())

	if err := srv.Serve(lis); err != nil {
		l.Fatalf("failed to serve: %v", err)
	}
}
