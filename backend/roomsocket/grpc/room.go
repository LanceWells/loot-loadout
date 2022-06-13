package grpc

import (
	"context"
	"log"
	"sync"

	pb "github.com/lantspants/lootloadout/api/roomsocket"
	"go.opentelemetry.io/otel"
)

type Room struct {
	l       *log.Logger
	clients map[pb.RoomSocket_ConnectToRoomServer]bool
	mu      sync.Mutex
}

func NewRoom(l *log.Logger) *Room {
	return &Room{
		l:       l,
		clients: make(map[pb.RoomSocket_ConnectToRoomServer]bool),
		mu:      sync.Mutex{},
	}
}

func (r *Room) Subscribe(ctx context.Context, stream pb.RoomSocket_ConnectToRoomServer) error {
	_, span := otel.Tracer("Room").Start(ctx, "Subscribe")
	defer span.End()

	r.mu.Lock()
	r.clients[stream] = true
	r.mu.Unlock()

	for range ctx.Done() {
		return ctx.Err()
	}

	return nil
}

func (r *Room) PublishMessage(ctx context.Context, message *pb.ChatCommand) error {
	_, span := otel.Tracer("Room").Start(ctx, "PublishMessage")
	defer span.End()

	for c := range r.clients {
		err := c.Send(message)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *Room) Close(ctx context.Context) error {
	_, span := otel.Tracer("Room").Start(ctx, "Close")
	defer span.End()

	r.mu.Lock()
	defer r.mu.Unlock()

	for c := range r.clients {
		_, cancel := context.WithCancel(c.Context())
		cancel()
	}

	return nil
}
