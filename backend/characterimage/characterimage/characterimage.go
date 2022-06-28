package characterimage

import (
	"context"
	"log"

	pb "github.com/lantspants/lootloadout/api/characterimage"
	"github.com/lantspants/lootloadout/backend/characterimage"
)

type CharacterImageService struct {
	l  *log.Logger
	db characterimage.CharacterImageDatabase
}

func NewCharacterImageService(
	l *log.Logger,
	db characterimage.CharacterImageDatabase,
) CharacterImageService {
	return CharacterImageService{
		l:  l,
		db: db,
	}
}

func (s CharacterImageService) CreateBodyType(
	ctx context.Context,
	req *pb.CreateBodyTypeRequest,
) (characterimage.ID, error) {
	id, err := s.db.CreateBodyType(ctx, req.Body)
	if err != nil {
		s.l.Printf("error creating a body type: %v", err)
		return "", err
	}

	return id, nil
}

func (s CharacterImageService) GetBodyType(
	ctx context.Context,
	req *pb.GetBodyTypeRequest,
) (*pb.BodyType, error) {
	body, err := s.db.GetBodyType(ctx, req.Id)
	if err != nil {
		s.l.Printf("error getting a body type: %v", err)
		return nil, err
	}

	return body, nil
}
