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
		s.l.Println("error adding prop:", err)
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
		s.l.Println("error listing props:", err)
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
		s.l.Println("error adding body:", err)
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
		s.l.Println("error listing bodies:", err)
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

func (s CharacterImageService) AddDynamic(
	ctx context.Context,
	req *api.AddDynamicRequest,
) (characterimage.ID, error) {
	_, span := otel.Tracer("CharacterImageService").Start(ctx, "AddDynamic")
	defer span.End()

	id, err := s.db.AddDynamic(ctx, req.Dynamic, req.BodyID, req.Image)
	if err != nil {
		s.l.Println("error adding dynamic:", err)
	}

	return id, nil
}

func (s CharacterImageService) ListDynamics(
	ctx context.Context,
	req *api.ListDynamicsRequest,
) (map[string]*api.Dynamic, map[string][]byte, error) {
	_, span := otel.Tracer("CharacterImageService").Start(ctx, "ListDynamics")
	defer span.End()

	dynamics, thumbnails, err := s.db.ListDynamics(ctx, nil)
	if err != nil {
		s.l.Println("error listing dynamics:", err)
	}

	return dynamics, thumbnails, nil
}

func (s CharacterImageService) AddStatic(
	ctx context.Context,
	req *api.AddStaticRequest,
) (characterimage.ID, error) {
	_, span := otel.Tracer("CharacterImageService").Start(ctx, "AddStatic")
	defer span.End()

	id, err := s.db.AddStatic(ctx, req.Static, req.BodyID)
	if err != nil {
		s.l.Println("error adding static:", err)
	}

	return id, nil
}

func (s CharacterImageService) ListStatics(
	ctx context.Context,
	req *api.ListStaticRequest,
) (map[characterimage.ID]*api.Static, error) {
	_, span := otel.Tracer("CharacterImageService").Start(ctx, "ListStatics")
	defer span.End()

	statics, err := s.db.ListStatics(ctx, nil)
	if err != nil {
		s.l.Println("error listing statics:", err)
	}

	return statics, nil
}

func (s CharacterImageService) AddAnimation(
	ctx context.Context,
	req *api.AddAnimationRequest,
) (characterimage.ID, error) {
	_, span := otel.Tracer("CharacterImageService").Start(ctx, "AddAnimation")
	defer span.End()

	id, err := s.db.AddAnimation(ctx, req.Animation, req.BodyID)
	if err != nil {
		s.l.Println("error adding animation:", err)
	}

	return id, nil
}

func (s CharacterImageService) ListAnimations(
	ctx context.Context,
	req *api.ListAnimationsRequest,
) (map[characterimage.ID]*api.Animation, error) {
	_, span := otel.Tracer("CharacterImageService").Start(ctx, "ListAnimations")
	defer span.End()

	animations, err := s.db.ListAnimations(ctx, nil)
	if err != nil {
		s.l.Println("error listing animations:", err)
	}

	return animations, nil
}

func (s CharacterImageService) AddFrame(
	ctx context.Context,
	req *api.AddFrameRequest,
) (characterimage.ID, error) {
	_, span := otel.Tracer("CharacterImageService").Start(ctx, "AddFrame")
	defer span.End()

	id, err := s.db.AddFrame(ctx, req.Frame, req.AnimationID, req.Image)
	if err != nil {
		s.l.Println("error adding frame:", err)
	}

	return id, nil
}

func (s CharacterImageService) GenerateAnimation(
	ctx context.Context,
	req *api.GenerateAnimationRequest,
) (*characterimage.GeneratedAnimation, error) {
	_, span := otel.Tracer("CharacterImageService").Start(ctx, "GenerateAnimation")
	defer span.End()

	return nil, nil
}
