package grpc

import (
	"context"
	"io"
	"log"
	"time"

	papb "github.com/lantspants/lootloadout/api/publicapi"
	rspb "github.com/lantspants/lootloadout/api/roomsocket"
)

type RoomSocketServerConfig struct {
	sendKeepAliveToClients bool
}

func (config *RoomSocketServerConfig) WithSendKeepAliveToClients() *RoomSocketServerConfig {
	config.sendKeepAliveToClients = true
	return config
}

type RoomSocketServer struct {
	papb.UnimplementedRoomsServer
	l      *log.Logger
	config *RoomSocketServerConfig
	client rspb.RoomSocketClient
}

func NewRoomSocketServer(
	l *log.Logger,
	client rspb.RoomSocketClient,
	config *RoomSocketServerConfig,
) *RoomSocketServer {
	return &RoomSocketServer{
		l:      l,
		config: config,
		client: client,
	}
}

func (s RoomSocketServer) CreateRoom(
	ctx context.Context,
	req *rspb.CreateRoomRequest,
) (*rspb.CreateRoomResponse, error) {
	res, err := s.client.CreateRoom(ctx, req)
	if err != nil {
		s.l.Printf("Error when creating a room: %v", err)
		return nil, err
	}

	return res, nil
}

func (s RoomSocketServer) DeleteRoom(
	ctx context.Context,
	req *rspb.DeleteRoomRequest,
) (*rspb.DeleteRoomResponse, error) {
	res, err := s.client.DeleteRoom(ctx, req)
	if err != nil {
		s.l.Printf("Error when deleting a room: %v", err)
		return nil, err
	}

	return res, nil
}

func (s RoomSocketServer) SendMessage(
	ctx context.Context,
	req *rspb.ChatCommand,
) (*rspb.ChatCommandResponse, error) {
	res, err := s.client.SendMessage(ctx, req)
	if err != nil {
		s.l.Printf("Error when sending a message: %v", err)
		return nil, err
	}

	return res, nil
}

func (s RoomSocketServer) ConnectToRoom(
	req *rspb.ConnectToRoomRequest,
	srv papb.Rooms_ConnectToRoomServer,
) error {
	ctx := srv.Context()
	stream, err := s.client.ConnectToRoom(ctx, req)
	if err != nil {
		s.l.Printf("Error when connecting to the room: %v", err)
	}

	// This feels so dirty, but after spending enough time trying to configure the ingress controller
	// on a free service (and finding out that that's not viable) I'm using this as a configurable
	// workaround.
	//
	// The problem is that nginx will kill streaming gRPC connections that haven't seen traffic within
	// a set deadline (default 60s). So far the only viable, free, option I've found is Okteto, which
	// uses nginx as its default ingress controller. The problem with Okteto is that it's a little
	// extra in its ingress management in its cloud service, specifically it overwrites the server
	// config annotation on its default ingress controller that we could otherwise use to up the
	// timeout to something like 6 hours.
	//
	// [ingress-nginx](https://kubernetes.github.io/ingress-nginx/)
	//
	// As a result, the default timeouts cannot be overwritten on the cloud service using the default
	// ingress. We can't use helm charts to set up traefik, because adding helm charts is a paid
	// option. And while I'm not certain if self-configuring traefik will work (I should clarify that
	// traefik *strongly* suggests that you use their helm chart, and there aren't any modern guides
	// otherwise that I could find), I'm somewhat confident that Okteto cloud wouldn't like us trying
	// to change its ingress controller (it seems to auto-inject nginx, but that needs a fact check).
	//
	// So in this case, in ~20 lines of code we can send a 'keep-alive' ping instead to all of our
	// clients and it avoids those problems; gross but effective.
	if s.config.sendKeepAliveToClients {
		t := time.NewTicker(time.Second * 40)
		tickerDone := make(chan bool)
		defer t.Stop()

		go func() {
			for {
				select {
				case <-t.C:
					srv.Send(&papb.ChatResponse{
						ChatResponseOptions: &papb.ChatResponse_KeepAlive{},
					})
				case <-tickerDone:
					t.Stop()
					return
				}
			}
		}()
	}

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Printf("Error when receiving a message: %v", err)
			return err
		}

		srv.Send(&papb.ChatResponse{
			ChatResponseOptions: &papb.ChatResponse_Command{
				Command: msg,
			},
		})
	}
}
