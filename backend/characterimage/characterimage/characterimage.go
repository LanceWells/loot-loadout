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

	id, err := s.db.AddProp(ctx, req.Prop)
	if err != nil {
		s.l.Println("error adding a prop:", err)
		return nil, err
	}

	return &id, nil
}

func (s CharacterImageService) ListProps(
	ctx context.Context,
	req *api.ListPropsRequest,
) (map[string]*api.Prop, error) {
	_, span := otel.Tracer("CharacterImageService").Start(ctx, "ListProps")
	defer span.End()

	props, err := s.db.ListProps(ctx, nil)
	if err != nil {
		s.l.Println("error getting props:", err)
		return nil, err
	}

	return props, nil
}

func (s CharacterImageService) AddBody(
	ctx context.Context,
	req *api.AddBodyRequest,
) (*characterimage.ID, error) {
	_, span := otel.Tracer("CharacterImageService").Start(ctx, "AddBody")
	defer span.End()

	id, err := s.db.AddBody(ctx, req.Body)
	if err != nil {
		s.l.Println("error adding a prop:", err)
		return nil, err
	}

	return &id, nil
}

func (s CharacterImageService) ListBodies(
	ctx context.Context,
	req *api.ListBodiesRequest,
) (map[string]*api.Body, error) {
	_, span := otel.Tracer("CharacterImageService").Start(ctx, "ListBodies")
	defer span.End()

	bodies, err := s.db.ListBodies(ctx, nil)
	if err != nil {
		s.l.Println("error getting props:", err)
		return nil, err
	}

	return bodies, nil
}

func (s CharacterImageService) AddDynamicMapping(
	ctx context.Context,
	req *api.AddDynamicMappingRequest,
) (characterimage.ID, error) {
	_, span := otel.Tracer("CharacterImageService").Start(ctx, "AddDynamicMapping")
	defer span.End()

	id, err := s.db.AddDynamicMapping(ctx, req.Mapping, req.BodyID)
	if err != nil {
		s.l.Println("error adding dynamic mapping:", err)
	}

	return id, nil
}
