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
		r.l.Printf("error validating: %v", err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	id, err := r.s.AddProp(ctx, req)
	if err != nil {
		r.l.Printf("error adding a prop: %v", err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &api.AddPropResponse{
		Id: *id,
	}, nil
}

// func (r CharacterImageServer) CreateBodyType(
// 	ctx context.Context,
// 	req *pb.CreateBodyTypeRequest,
// ) (*pb.CreateBodyTypeResponse, error) {
// 	_, span := otel.Tracer("CharacterImageServer").Start(ctx, "CreateBodyType")
// 	defer span.End()

// 	err := req.Validate()
// 	if err != nil {
// 		r.l.Printf("error validating: %v", err)
// 		return nil, err
// 	}

// 	id, err := r.s.CreateBodyType(ctx, req)
// 	if err != nil {
// 		r.l.Printf("error creating a body type: %v", err)
// 		return nil, err
// 	}

// 	span.SetAttributes(
// 		attribute.String("BodyTypeID", id),
// 	)

// 	return &pb.CreateBodyTypeResponse{
// 		Id: id,
// 	}, nil
// }

// func (r CharacterImageServer) GetBodyType(
// 	ctx context.Context,
// 	req *pb.GetBodyTypeRequest,
// ) (*pb.GetBodyTypeResponse, error) {
// 	_, span := otel.Tracer("CharacterImageServer").Start(ctx, "GetBodyType")
// 	defer span.End()

// 	err := req.Validate()
// 	if err != nil {
// 		r.l.Printf("error validating: %v", err)
// 		return nil, status.Errorf(codes.InvalidArgument, err.Error())
// 	}

// 	body, err := r.s.GetBodyType(ctx, req)
// 	if err != nil {
// 		r.l.Printf("error getting body type: %v", err)
// 		return nil, status.Errorf(codes.Internal, err.Error())
// 	}

// 	return &pb.GetBodyTypeResponse{
// 		Body: body,
// 	}, nil
// }

// func (r CharacterImageServer) CreateDynamicPartMapping(
// 	ctx context.Context,
// 	req *pb.CreateDynamicPartMappingRequest,
// ) (*pb.CreateDynamicPartMappingResponse, error) {
// 	_, span := otel.Tracer("CharacterImageServer").Start(ctx, "CreateDynamicPartMapping")
// 	defer span.End()

// 	err := req.Validate()
// 	if err != nil {
// 		r.l.Printf("error validating: %v", err)
// 		return nil, status.Errorf(codes.InvalidArgument, err.Error())
// 	}

// 	return nil, status.Errorf(codes.Unimplemented, "method CreateDynamicPartMapping not implemented")
// }
