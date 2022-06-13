package grpc

import (
	"context"
	"io"
	"log"

	papb "github.com/lantspants/lootloadout/api/publicapi"
	rspb "github.com/lantspants/lootloadout/api/roomsocket"
)

type RoomSocketServer struct {
	papb.UnimplementedRoomsServer
	l      *log.Logger
	client rspb.RoomSocketClient
}

func NewRoomSocketServer(l *log.Logger, client rspb.RoomSocketClient) *RoomSocketServer {
	return &RoomSocketServer{
		l:      l,
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

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("Error when receiving a message: %v", err)
			break
		}

		srv.Send(msg)
	}

	return nil
}
