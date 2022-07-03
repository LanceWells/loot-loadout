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
		r.l.Printf("error validating: %v", err)
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
		r.l.Printf("error validating: %v", err)
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
		r.l.Printf("error validating: %v", err)
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
		r.l.Printf("error validating: %v", err)
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

func (r CharacterImageServer) AddStatic(
	ctx context.Context,
	req *api.AddStaticRequest,
) (*api.AddStaticResponse, error) {
	_, span := otel.Tracer("CharacterImageServer").Start(ctx, "AddStatic")
	defer span.End()

	err := req.Validate()
	if err != nil {
		r.l.Printf("error validating: %v", err)
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
		r.l.Printf("error validating: %v", err)
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
