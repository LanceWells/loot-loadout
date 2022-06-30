package characterimage

import (
	"context"
	"log"

	api "github.com/lantspants/lootloadout/api/characterimage/v1"
	"github.com/lantspants/lootloadout/backend/characterimage"
	"go.opentelemetry.io/otel"
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

func (s CharacterImageService) AddProp(
	ctx context.Context,
	req *api.AddPropRequest,
) (*characterimage.ID, error) {
	_, span := otel.Tracer("CharacterImageService").Start(ctx, "AddProp")
	defer span.End()

	id, err := s.db.CreateProp(ctx, req.Prop)
	if err != nil {
		s.l.Printf("error adding a prop: %v", err)
		return nil, err
	}

	return &id, nil
}

func (s CharacterImageService) ListProps(
	ctx context.Context,
	req *api.ListPropsRequest,
) ([]api.Prop, error) {
	_, span := otel.Tracer("CharacterImageService").Start(ctx, "ListProps")
	defer span.End()

	props, err := s.db.ListProps(ctx, nil)
	if err != nil {
		s.l.Printf("error getting props: %v", err)
		return nil, err
	}

	return props, nil
}

// func (s CharacterImageService) CreateBodyType(
// 	ctx context.Context,
// 	req *api.CreateBodyTypeRequest,
// ) (characterimage.ID, error) {
// 	id, err := s.db.CreateBodyType(ctx, req.Body)
// 	if err != nil {
// 		s.l.Printf("error creating a body type: %v", err)
// 		return "", err
// 	}

// 	return id, nil
// }

// func (s CharacterImageService) GetBodyType(
// 	ctx context.Context,
// 	req *api.GetBodyTypeRequest,
// ) (*api.BodyType, error) {
// 	body, err := s.db.GetBodyType(ctx, req.Id)
// 	if err != nil {
// 		s.l.Printf("error getting a body type: %v", err)
// 		return nil, err
// 	}

// 	return body, nil
// }

// func (s CharacterImageService) CreateDynamicPart(
// 	ctx context.Context,
// 	*pb.DynamicPart,
// ) (ID, error) {
// 	s.db.Create
// }

// func (s CharacterImageService) GetDynamicPart(
// 	ctx context.Context,
// 	id ID,
// ) (*pb.DynamicPart, error) {

// }
