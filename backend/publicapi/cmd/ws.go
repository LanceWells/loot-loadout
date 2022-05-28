package main

import (
	"context"
	"log"
	"regexp"

	"github.com/lantspants/lootloadout/publicapi"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
)

type HubIdentifier struct {
	id string
}

func (i *HubIdentifier) Validate() bool {
	ok, err := regexp.MatchString("\\w{12}", i.id)
	if err != nil || !ok {
		return false
	}

	return true
}

type WSHandler struct {
	hubs map[string]*publicapi.Hub
	log  *log.Logger
}

func NewWSHandler(logger *log.Logger) *WSHandler {
	return &WSHandler{
		hubs: make(map[string]*publicapi.Hub),
		log:  logger,
	}
}

func (h *WSHandler) GetHub(ctx context.Context, id *HubIdentifier) (*publicapi.Hub, error) {
	_, span := otel.Tracer("WSHandler").Start(ctx, "GetHub")
	defer span.End()

	span.SetAttributes(
		attribute.String("id.id", id.id),
	)

	ok := id.Validate()
	if !ok {
		h.log.Println(publicapi.ErrHubNameInvalid)
		return nil, publicapi.ErrHubNameInvalid
	}

	hub, ok := h.hubs[id.id]
	if !ok {
		h.log.Println(publicapi.ErrHubNotFound)
		return nil, publicapi.ErrHubNotFound
	}

	return hub, nil
}

func (h *WSHandler) CreateHub(ctx context.Context, id *HubIdentifier) (*publicapi.Hub, error) {
	_, span := otel.Tracer("WSHandler").Start(ctx, "CreateHub")
	defer span.End()

	span.SetAttributes(
		attribute.String("id.id", id.id),
	)

	ok := id.Validate()
	if !ok {
		h.log.Println(publicapi.ErrHubNameInvalid)
		return nil, publicapi.ErrHubNameInvalid
	}

	_, ok = h.hubs[id.id]
	if !ok {
		h.log.Println(publicapi.ErrHubAlreadyExists)
		return nil, publicapi.ErrHubAlreadyExists
	}

	hub := publicapi.NewHub(h.log)
	h.hubs[id.id] = hub

	return hub, nil
}

func (h *WSHandler) DeleteHub(ctx context.Context, id HubIdentifier) error {
	_, span := otel.Tracer("WSHandler").Start(ctx, "DeleteHub")
	defer span.End()

	span.SetAttributes(
		attribute.String("id.id", id.id),
	)

	ok := id.Validate()
	if !ok {
		h.log.Println(publicapi.ErrHubNameInvalid)
		return publicapi.ErrHubNameInvalid
	}

	_, ok = h.hubs[id.id]
	if !ok {
		h.log.Println(publicapi.ErrHubNotFound)
		return publicapi.ErrHubNotFound
	}

	h.hubs[id.id] = nil

	return nil
}
