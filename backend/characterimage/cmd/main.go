package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	pb "github.com/lantspants/lootloadout/api/characterimage/v1"
	"github.com/lantspants/lootloadout/backend/characterimage/characterimage"
	grpcServer "github.com/lantspants/lootloadout/backend/characterimage/cmd/grpc"
	"github.com/lantspants/lootloadout/backend/characterimage/postgres"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	grpcEndpoint = flag.Int("grpc-server-endpoint", 9999, "gRPC server endpoint")
	// httpEndpoint = flag.String("http-server-endpoint", ":8888", "HTTP server endpoint")
)

func main() {
	ctx := context.Background()
	l := log.New(os.Stdout, "", 0)
	flag.Parse()

	db, err := postgres.NewImageDatabase(ctx, l)
	if err != nil {
		l.Fatalf("error initializing database: %v", err)
	}
	s := characterimage.NewCharacterImageService(l, db)
	// r := http.NewCharacterImageServer(l, s)
	r := grpcServer.NewCharacterImageServer(l, s)
	srv := grpc.NewServer()

	// e := echo.New()
	// api.RegisterHandlers(e, r)
	pb.RegisterImagesServer(srv, r)
	reflection.Register(srv)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *grpcEndpoint))
	if err != nil {
		l.Fatalf("Failed to listen: %v", err)
	}

	l.Printf("Initializing characterimage service at %s", lis.Addr().String())
	// l.Printf("Initializing characterimage service at %s", *httpEndpoint)
	// if err = e.Start(*httpEndpoint); err != nil {
	// 	l.Fatalf("failed to serve: %v", err)
	// }
	if err := srv.Serve(lis); err != nil {
		l.Fatalf("failed to serve: %v", err)
	}
}
