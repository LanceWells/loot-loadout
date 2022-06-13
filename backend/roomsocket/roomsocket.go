package roomsocket

import (
	"context"

	pb "github.com/lantspants/lootloadout/api/roomsocket"
	"github.com/segmentio/ksuid"
)

// RoomService is an interface used to handle room-based requests. These include commands and chat
// messages from characters.
type RoomService interface {
	// Subscribe adds the provided stream to the desired room. That stream is now considered
	// "subscribed" to the feed associated with that room and will recieve messages generated within
	// that room.
	Subscribe(ctx context.Context, id *RoomIdentifier, stream pb.RoomSocket_ConnectToRoomServer) error

	// Create creates a new room for a user to subscribe to.
	Create(ctx context.Context, req *pb.CreateRoomRequest) (*pb.CreateRoomResponse, error)

	// Delete deletes a specified room.
	Delete(ctx context.Context, id *RoomIdentifier, req *pb.DeleteRoomRequest) (*pb.DeleteRoomResponse, error)

	// SendMessage sends a message to the specified room.
	SendMessage(ctx context.Context, id *RoomIdentifier, command *pb.ChatCommand) (*pb.ChatCommandResponse, error)
}

type RoomIdentifier struct {
	roomID string
}

func NewRoomIdentifier(id string) *RoomIdentifier {
	return &RoomIdentifier{
		roomID: id,
	}
}

func NewRoomIdentifierFromKSUID(id ksuid.KSUID) *RoomIdentifier {
	return &RoomIdentifier{
		roomID: id.String(),
	}
}

func (id RoomIdentifier) Validate() bool {
	return len(id.roomID) != 0
}

func (id RoomIdentifier) String() string {
	return id.roomID
}
