package grpc

import (
	"context"
	"log"

	pb "github.com/lantspants/lootloadout/api/roomsocket"
	"github.com/lantspants/lootloadout/backend/roomsocket"
	"github.com/segmentio/ksuid"
	"go.opentelemetry.io/otel"
)

// RoomService is a type used to handle room-based requests. These include commands and chat
// messages from characters.
type RoomService struct {
	l     *log.Logger
	rooms map[roomsocket.RoomIdentifier]*Room
}

// NewRoomService returns a new RoomHandler.
func NewRoomService(l *log.Logger) *RoomService {
	return &RoomService{
		l:     l,
		rooms: make(map[roomsocket.RoomIdentifier]*Room),
	}
}

// Subscribe adds the provided stream to the desired room. That stream is now considered
// "subscribed" to the feed associated with that room and will recieve messages generated within
// that room.
func (h RoomService) Subscribe(
	ctx context.Context,
	id *roomsocket.RoomIdentifier,
	stream pb.RoomSocket_ConnectToRoomServer,
) error {
	_, span := otel.Tracer("RoomService").Start(ctx, "Subscribe")
	defer span.End()

	room, ok := h.rooms[*id]
	if !ok {
		err := roomsocket.ErrRoomDoesNotExist
		h.l.Printf("err when subscribing to room: %v", err)
		return err
	}

	err := room.Subscribe(ctx, stream)
	if err != nil {
		return err
	}

	return nil
}

// Create creates a new room for a user to subscribe to.
func (h RoomService) Create(
	ctx context.Context,
	req *pb.CreateRoomRequest,
) (*pb.CreateRoomResponse, error) {
	ksuid := ksuid.New()
	id := roomsocket.NewRoomIdentifierFromKSUID(ksuid)
	ok := id.Validate()
	if !ok {
		err := roomsocket.ErrUnexpected
		h.l.Printf("error when creating ID: %v", err)
		return nil, err
	}

	h.rooms[*id] = NewRoom(h.l)

	return &pb.CreateRoomResponse{
		RoomId: id.String(),
	}, nil
}

// Delete deletes a specified room.
func (h RoomService) Delete(
	ctx context.Context,
	id *roomsocket.RoomIdentifier,
	req *pb.DeleteRoomRequest,
) (*pb.DeleteRoomResponse, error) {
	room, ok := h.rooms[*id]
	if !ok {
		err := roomsocket.ErrRoomDoesNotExist
		h.l.Printf("err when subscribing to room: %v", err)
		return nil, err
	}

	h.l.Printf("closing room %v", id)
	room.Close(ctx)

	return &pb.DeleteRoomResponse{}, nil
}

func (h RoomService) SendMessage(
	ctx context.Context,
	id *roomsocket.RoomIdentifier,
	command *pb.ChatCommand,
) (*pb.ChatCommandResponse, error) {
	_, span := otel.Tracer("RoomService").Start(ctx, "SendMessage")
	defer span.End()

	room, ok := h.rooms[*id]
	if !ok {
		err := roomsocket.ErrRoomDoesNotExist
		h.l.Printf("err when sending a message to room %v", err)
		return nil, err
	}

	err := room.PublishMessage(ctx, command)
	if err != nil {
		return nil, err
	}

	return &pb.ChatCommandResponse{}, nil
}
