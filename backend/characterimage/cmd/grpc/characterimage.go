package grpc

import (
	"context"
	"log"

	api "github.com/lantspants/lootloadout/api/characterimage/v1"
	service "github.com/lantspants/lootloadout/backend/characterimage/characterimage"
	"go.opentelemetry.io/otel"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CharacterImageServer struct {
	api.UnimplementedImagesServer
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

func (r CharacterImageServer) AddProp(
	ctx context.Context,
	req *api.AddPropRequest,
) (*api.AddPropResponse, error) {
	_, span := otel.Tracer("CharacterImageServer").Start(ctx, "AddProp")
	defer span.End()

	err := req.Validate()
	if err != nil {
		r.l.Println("error validating:", err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	id, err := r.s.AddProp(ctx, req)
	if err != nil {
		r.l.Println("error adding a prop:", err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &api.AddPropResponse{
		Id: *id,
	}, nil
}

func (r CharacterImageServer) ListProps(
	ctx context.Context,
	req *api.ListPropsRequest,
) (*api.ListPropsResponse, error) {
	_, span := otel.Tracer("CharacterImageServer").Start(ctx, "ListProps")
	defer span.End()

	err := req.Validate()
	if err != nil {
		r.l.Println("error validating:", err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	props, err := r.s.ListProps(ctx, req)
	if err != nil {
		r.l.Println("error listing props:", err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &api.ListPropsResponse{
		Props: props,
	}, nil
}

func (r CharacterImageServer) AddBody(
	ctx context.Context,
	req *api.AddBodyRequest,
) (*api.AddBodyResponse, error) {
	_, span := otel.Tracer("CharacterImageServer").Start(ctx, "AddBody")
	defer span.End()

	err := req.Validate()
	if err != nil {
		r.l.Println("error validating:", err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	id, err := r.s.AddBody(ctx, req)
	if err != nil {
		r.l.Println("error adding a body:", err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &api.AddBodyResponse{
		Id: *id,
	}, nil
}

func (r CharacterImageServer) ListBodies(
	ctx context.Context,
	req *api.ListBodiesRequest,
) (*api.ListBodiesResponse, error) {
	_, span := otel.Tracer("CharacterImageServer").Start(ctx, "ListBodies")
	defer span.End()

	err := req.Validate()
	if err != nil {
		r.l.Println("error validating:", err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	bodies, err := r.s.ListBodies(ctx, req)
	if err != nil {
		r.l.Println("error listing bodies:", err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &api.ListBodiesResponse{
		Bodies: bodies,
	}, nil
}

func (r CharacterImageServer) AddDynamicMapping(
	ctx context.Context,
	req *api.AddDynamicMappingRequest,
) (*api.AddDynamicMappingResponse, error) {
	_, span := otel.Tracer("CharacterImageServer").Start(ctx, "AddDynamicMapping")
	defer span.End()

	err := req.Validate()
	if err != nil {
		r.l.Println("error validating:", err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	id, err := r.s.AddDynamicMapping(ctx, req)
	if err != nil {
		r.l.Println("error adding dynamic mapping:", err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &api.AddDynamicMappingResponse{
		Id: id,
	}, nil
}

func (r CharacterImageServer) AddDynamic(
	ctx context.Context,
	req *api.AddDynamicRequest,
) (*api.AddDynamicResponse, error) {
	_, span := otel.Tracer("CharacterImageServer").Start(ctx, "AddDynamic")
	defer span.End()

	err := req.Validate()
	if err != nil {
		r.l.Println("error validating:", err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	id, err := r.s.AddDynamic(ctx, req)
	if err != nil {
		r.l.Println("error adding dynamic part:", err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &api.AddDynamicResponse{
		Id: id,
	}, nil
}

func (r CharacterImageServer) ListDynamics(
	ctx context.Context,
	req *api.ListDynamicsRequest,
) (*api.ListDynamicsResponse, error) {
	_, span := otel.Tracer("CharacterImageServer").Start(ctx, "ListDynamics")
	defer span.End()

	err := req.Validate()
	if err != nil {
		r.l.Println("error validating:", err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	dynamics, thumbnails, err := r.s.ListDynamics(ctx, req)
	if err != nil {
		r.l.Println("error listing dynamics:", err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &api.ListDynamicsResponse{
		Dynamics:   dynamics,
		Thumbnails: thumbnails,
	}, nil
}

func (r CharacterImageServer) AddStatic(
	ctx context.Context,
	req *api.AddStaticRequest,
) (*api.AddStaticResponse, error) {
	_, span := otel.Tracer("CharacterImageServer").Start(ctx, "AddStatic")
	defer span.End()

	err := req.Validate()
	if err != nil {
		r.l.Println("error validating:", err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	id, err := r.s.AddStatic(ctx, req)
	if err != nil {
		r.l.Println("error adding static part:", err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &api.AddStaticResponse{
		Id: id,
	}, nil
}

func (r CharacterImageServer) ListStatics(
	ctx context.Context,
	req *api.ListStaticRequest,
) (*api.ListStaticsResponse, error) {
	_, span := otel.Tracer("CharacterImageServer").Start(ctx, "ListStatics")
	defer span.End()

	err := req.Validate()
	if err != nil {
		r.l.Println("error validating:", err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	statics, err := r.s.ListStatics(ctx, req)
	if err != nil {
		r.l.Println("error listing bodies:", err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &api.ListStaticsResponse{
		Statics: statics,
	}, nil
}

func (r CharacterImageServer) AddAnimation(
	ctx context.Context,
	req *api.AddAnimationRequest,
) (*api.AddAnimationResponse, error) {
	_, span := otel.Tracer("CharacterImageServer").Start(ctx, "AddAnimation")
	defer span.End()

	err := req.Validate()
	if err != nil {
		r.l.Println("error validating:", err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	id, err := r.s.AddAnimation(ctx, req)
	if err != nil {
		r.l.Println("error adding animation:", err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &api.AddAnimationResponse{
		Id: id,
	}, nil
}

func (r CharacterImageServer) ListAnimations(
	ctx context.Context,
	req *api.ListAnimationsRequest,
) (*api.ListAnimationsResponse, error) {
	_, span := otel.Tracer("CharacterImageServer").Start(ctx, "ListAnimations")
	defer span.End()

	err := req.Validate()
	if err != nil {
		r.l.Println("error validating:", err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	animations, err := r.s.ListAnimations(ctx, req)
	if err != nil {
		r.l.Println("error listing animations:", err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &api.ListAnimationsResponse{
		Animations: animations,
	}, nil
}

func (r CharacterImageServer) AddFrame(
	ctx context.Context,
	req *api.AddFrameRequest,
) (*api.AddFrameResponse, error) {
	_, span := otel.Tracer("CharacterImageServer").Start(ctx, "AddFrame")
	defer span.End()

	err := req.Validate()
	if err != nil {
		r.l.Println("error validating:", err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	id, err := r.s.AddFrame(ctx, req)
	if err != nil {
		r.l.Println("error adding frame:", err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &api.AddFrameResponse{
		Id: id,
	}, nil
}

func (r CharacterImageServer) GenerateAnimation(
	ctx context.Context,
	req *api.GenerateAnimationRequest,
) (*api.GenerateAnimationResponse, error) {
	_, span := otel.Tracer("CharacterImageServer").Start(ctx, "GenerateAnimation")
	defer span.End()

	err := req.Validate()
	if err != nil {
		r.l.Println("error validating:", err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	a, err := r.s.GenerateAnimation(ctx, req)
	if err != nil {
		r.l.Println("error generating animation:", err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	frames := []*api.GenerateAnimationResponse_FrameImage{}
	for _, f := range a.Frames {
		dynamicImages := make(map[string][]byte, len(f.Images))
		for ik, iv := range f.Images {
			dynamicImages[ik.Enum().String()] = iv
		}

		frames = append(frames, &api.GenerateAnimationResponse_FrameImage{
			Frame: f.Frame,
			Image: dynamicImages,
		})
	}

	return &api.GenerateAnimationResponse{
		Statics: a.Statics,
		Prop:    a.Prop,
		Frames:  frames,
	}, nil
}
