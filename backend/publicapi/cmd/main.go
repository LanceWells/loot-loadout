package main

import (
	"context"
	"flag"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	roomSocket "github.com/lantspants/lootloadout/backend/publicapi/cmd/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"

	papb "github.com/lantspants/lootloadout/api/publicapi"
	rspb "github.com/lantspants/lootloadout/api/roomsocket"
)

var (
	grpcEndpoint = flag.String("grpc-server-endpoint", "localhost:9090", "gRPC server endpoint")
	httpEndpoint = flag.String("http-server-endpoint", "localhost:8080", "HTTP server endpoint")
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	l := log.New(os.Stdout, "", 0)

	srv := grpc.NewServer()

	// Rooms Service
	r := getRoomSocketServer(ctx, l)
	papb.RegisterRoomsServer(srv, r)

	// Registering those gRPC endpoints makes many tools work 👌
	reflection.Register(srv)

	// Run the gRPC server in a routine. This lets us talk to it from the multiplexer.
	go func() {
		var err error

		lis, err := net.Listen("tcp", *grpcEndpoint)
		if err != nil {
			l.Fatal(err)
		}

		l.Printf("Serving gRPC traffic on: %v", lis.Addr())
		err = srv.Serve(lis)
		if err != nil {
			l.Fatal(err)
		}
	}()

	// From here-on down is HTTP handling, and forwarding to the gRPC endpoints via the multiplexer.
	// This HTTP server makes the assumption that we've already gotten the gRPC server going.
	mux := runtime.NewServeMux()
	err := papb.RegisterRoomsHandlerFromEndpoint(
		ctx,
		mux,
		*grpcEndpoint,
		[]grpc.DialOption{
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		},
	)
	if err != nil {
		l.Fatal(err)
	}

	l.Printf("Serving HTTP traffic on: %v", *httpEndpoint)
	err = http.ListenAndServe(*httpEndpoint, mux)
	if err != nil {
		l.Fatal(err)
	}
}

func getRoomSocketServer(ctx context.Context, l *log.Logger) *roomSocket.RoomSocketServer {
	address, ok := os.LookupEnv("ROOMSOCKET_SERVICE")
	if !ok {
		address = "room-socket:80"
	}

	conn, err := grpc.DialContext(
		ctx,
		address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		l.Fatal(err)
	}

	c := rspb.NewRoomSocketClient(conn)
	r := roomSocket.NewRoomSocketServer(l, c)

	return r
}
