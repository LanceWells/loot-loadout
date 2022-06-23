package grpc

import (
	"context"
	"log"

	pb "github.com/lantspants/lootloadout/api/characterimage"
	service "github.com/lantspants/lootloadout/backend/characterimage/characterimage"
)

type CharacterImageServer struct {
	pb.UnimplementedImagesServer
	l *log.Logger
	s service.CharacterImageService
}

func NewCharacterImageServer(
	l *log.Logger,
	s service.CharacterImageService,
) CharacterImageServer {
	return CharacterImageServer{
		l: l,
		s: s,
	}
}

func (r CharacterImageServer) CreateBodyType(
	ctx context.Context,
	req *pb.CreateBodyTypeRequest,
) (*pb.CreateBodyTypeResponse, error) {
	return nil, nil
}
