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
	b, err := characterimage.NewBodyType(
		req.Body.DisplayName,
		req.Body.PartMaps,
	)
	if err != nil {
		s.l.Printf("error creating a new body type object: %v", err)
		return "", err
	}

	id, err := s.db.CreateBodyType(ctx, b)
	if err != nil {
		s.l.Printf("error creating a body type: %v", err)
		return "", err
	}

	return id, nil
}
