package grpc

import (
	"context"
	"log"

	pb "github.com/lantspants/lootloadout/api/roomsocket"
	"github.com/lantspants/lootloadout/backend/roomsocket"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
)

const RoomIDHeader = "room_id"

// RoomSocketServer is a server used to handle incoming GRPC requests that are then forwarded to the
// underlying service.
type RoomSocketServer struct {
	pb.UnimplementedRoomSocketServer
	l *log.Logger
	h roomsocket.RoomService
}

// NewRoomSocketServer creates a new instance of a RoomSocketServer.
func NewRoomSocketServer(l *log.Logger, h roomsocket.RoomService) RoomSocketServer {
	return RoomSocketServer{
		l: l,
		h: h,
	}
}

// ConnectToChat is a server-streaming endpoint for interacting with the room.
func (r RoomSocketServer) ConnectToRoom(req *pb.ConnectToRoomRequest, stream pb.RoomSocket_ConnectToRoomServer) error {
	ctx := stream.Context()
	_, span := otel.Tracer("RoomSocketServer").Start(ctx, "Chat")
	defer span.End()

	err := req.Validate()
	if err != nil {
		r.l.Printf("error validating: %v", err)
		return err
	}

	id := roomsocket.NewRoomIdentifier(req.RoomId)
	if !id.Validate() {
		err := roomsocket.ErrRoomIDHeaderInvalid
		r.l.Printf("error in room identifier: %v", err)
		return err
	}

	span.SetAttributes(
		attribute.String("roomid", id.String()),
	)

	err = r.h.Subscribe(ctx, id, stream)
	if err != nil {
		r.l.Printf("error in subscription: %v", err)
		return err
	}

	return nil
}

// CreateRoom creates a room with a new ID.
func (r RoomSocketServer) CreateRoom(
	ctx context.Context,
	req *pb.CreateRoomRequest,
) (*pb.CreateRoomResponse, error) {
	_, span := otel.Tracer("RoomSocketServer").Start(ctx, "CreateRoom")
	defer span.End()

	err := req.Validate()
	if err != nil {
		r.l.Printf("error validating: %v", err)
		return nil, err
	}

	resp, err := r.h.Create(ctx, req)
	if err != nil {
		r.l.Printf("error creating a room: %v", err)
		return nil, err
	}

	span.SetAttributes(
		attribute.String("roomid", resp.RoomId),
	)

	return resp, nil
}

// DeleteRoom deletes a room with the specified ID.
func (r RoomSocketServer) DeleteRoom(
	ctx context.Context,
	req *pb.DeleteRoomRequest,
) (*pb.DeleteRoomResponse, error) {
	_, span := otel.Tracer("RoomSocketServer").Start(ctx, "DeleteRoom")
	defer span.End()

	err := req.Validate()
	if err != nil {
		r.l.Printf("error validating: %v", err)
		return nil, err
	}

	id := roomsocket.NewRoomIdentifier(req.RoomId)
	if !id.Validate() {
		err := roomsocket.ErrRoomIDHeaderInvalid
		r.l.Printf("error in metadata: %v", err)
		return nil, err
	}

	span.SetAttributes(
		attribute.String("roomid", id.String()),
	)

	resp, err := r.h.Delete(ctx, id, req)
	if err != nil {
		r.l.Printf("error deleting a room: %v", err)
		return nil, err
	}

	return resp, nil
}

func (r RoomSocketServer) SendMessage(
	ctx context.Context,
	command *pb.ChatCommand,
) (*pb.ChatCommandResponse, error) {
	_, span := otel.Tracer("RoomSocketServer").Start(ctx, "SendMessage")
	defer span.End()

	err := command.Validate()
	if err != nil {
		r.l.Printf("error validating: %v", err)
		return nil, err
	}

	id := roomsocket.NewRoomIdentifier(command.RoomId)
	if !id.Validate() {
		err := roomsocket.ErrRoomIDHeaderInvalid
		r.l.Printf("error in metadata: %v", err)
		return nil, err
	}

	span.SetAttributes(
		attribute.String("roomid", id.String()),
	)

	resp, err := r.h.SendMessage(ctx, id, command)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
