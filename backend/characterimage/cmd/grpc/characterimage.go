package grpc

import (
	"context"
	"log"

	pb "github.com/lantspants/lootloadout/api/characterimage"
	service "github.com/lantspants/lootloadout/backend/characterimage/characterimage"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
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
	_, span := otel.Tracer("CharacterImageServer").Start(ctx, "CreateBodyType")
	defer span.End()

	err := req.Validate()
	if err != nil {
		r.l.Printf("error validating: %v", err)
		return nil, err
	}

	id, err := r.s.CreateBodyType(ctx, req)
	if err != nil {
		r.l.Printf("error creating a body type: %v", err)
		return nil, err
	}

	span.SetAttributes(
		attribute.String("BodyTypeID", id),
	)

	return &pb.CreateBodyTypeResponse{
		Id: id,
	}, nil
}

func (r CharacterImageServer) GetBodyType(
	ctx context.Context,
	req *pb.GetBodyTypeRequest,
) (*pb.GetBodyTypeResponse, error) {
	_, span := otel.Tracer("CharacterImageServer").Start(ctx, "GetBodyType")
	defer span.End()

	err := req.Validate()
	if err != nil {
		r.l.Printf("error validating: %v", err)
		return nil, err
	}

	body, err := r.s.GetBodyType(ctx, req)
	if err != nil {
		r.l.Printf("error getting body type: %v", err)
		return nil, err
	}

	return &pb.GetBodyTypeResponse{
		Body: body,
	}, nil
}
