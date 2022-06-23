package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	pb "github.com/lantspants/lootloadout/api/roomsocket"
	grpcServer "github.com/lantspants/lootloadout/backend/roomsocket/cmd/grpc"
	grpcService "github.com/lantspants/lootloadout/backend/roomsocket/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	grpcEndpoint = flag.Int("port", 8888, "The server port")
)

func main() {
	l := log.New(os.Stdout, "", 0)
	flag.Parse()

	s := grpcService.NewRoomService(l)
	r := grpcServer.NewRoomSocketServer(l, s)
	srv := grpc.NewServer()

	pb.RegisterRoomSocketServer(srv, r)
	reflection.Register(srv)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *grpcEndpoint))
	if err != nil {
		l.Fatalf("Failed to listen: %v", err)
	}

	l.Printf("Initializing roomsocket service at %s", lis.Addr().String())

	if err := srv.Serve(lis); err != nil {
		l.Fatalf("failed to serve: %v", err)
	}
}
