package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	pb "github.com/lantspants/lootloadout/api/roomsocket"
	grpcServer "github.com/lantspants/lootloadout/backend/roomsocket/cmd/grpc"
	grpcService "github.com/lantspants/lootloadout/backend/roomsocket/grpc"
	googleGRPC "google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"
)

var (
	port = flag.Int("port", 8080, "The server port")
)

var kaep = keepalive.EnforcementPolicy{
	PermitWithoutStream: true,
	MinTime:             time.Second * 60,
}

func main() {
	l := log.New(os.Stdout, "", 0)
	flag.Parse()

	s := grpcService.NewRoomService(l)
	r := grpcServer.NewRoomSocketServer(l, s)
	srv := googleGRPC.NewServer(
		googleGRPC.KeepaliveEnforcementPolicy(kaep),
	)

	pb.RegisterRoomSocketServer(srv, r)
	reflection.Register(srv)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		l.Fatalf("failed to listen: %v", err)
	}

	l.Printf("Initializing roomsocket service at %s", lis.Addr().String())

	if err := srv.Serve(lis); err != nil {
		l.Fatalf("failed to serve: %v", err)
	}
}
