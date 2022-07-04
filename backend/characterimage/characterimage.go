package characterimage

import (
	"context"

	// We use some pb types here. This is intentional. It's intended to preserve the contract between
	// the API and what we are actually able to store.
	api "github.com/lantspants/lootloadout/api/characterimage/v1"
)

// ID is an alias for an identifier that refers to a specific resour.ce
type ID = string

type GeneratedFrame struct {
	Frame  *api.Frame
	Images map[api.DynamicPartType][]byte
}

type GeneratedAnimation struct {
	Frames  []GeneratedFrame
	Statics []*api.Static
	Prop    *api.Prop
}

// BodyTypesFilter is used to provide a limited list of the given item.
type BodyTypesFilter struct {
	BodyTypeIdentifiers []string
}

// StaticPartsFilter is used to provide a limited list of the given item.
type StaticPartsFilter struct {
	StaticPartTypes []api.StaticPartType
}

// DynamicPartsFilter is used to provide a limited list of the given item.
type DynamicPartsFilter struct {
	DynamicPartTypes []api.DynamicPartType
}

// PropsFilter is used to provide a limited list of the given item.
type PropsFilter struct {
	PropTypes []api.PropType
}

// AnimationsFilter is used to provide a limited list of the given item.
type AnimationsFilter struct {
	BodyTypeIdentifiers []string
}

type CharacterImageDatabase interface {
	AddBody(context.Context, *api.Body) (ID, error)
	ListBodies(context.Context, *BodyTypesFilter) (map[string]*api.Body, error)

	AddDynamicMapping(ctx context.Context, m *api.DynamicMapping, bodyID ID) (ID, error)

	AddStatic(ctx context.Context, s *api.Static, bodyID ID) (ID, error)
	ListStatics(context.Context, *StaticPartsFilter) (map[string]*api.Static, error)

	AddDynamic(ctx context.Context, d *api.Dynamic, bodyID ID, image []byte) (ID, error)
	ListDynamics(context.Context, *DynamicPartsFilter) (map[string]*api.Dynamic, map[string][]byte, error)

	AddAnimation(ctx context.Context, a *api.Animation, bodyID ID) (ID, error)
	ListAnimations(context.Context, *AnimationsFilter) (map[string]*api.Animation, error)

	AddFrame(ctx context.Context, f *api.Frame, animationID ID, image []byte) (ID, error)

	AddProp(context.Context, *api.Prop) (ID, error)
	ListProps(context.Context, *PropsFilter) (map[string]*api.Prop, error)

	GenerateAnimation(
		ctx context.Context,
		bodyID string,
		animationID string,
		dynamicPartIDs map[api.DynamicPartType]string,
		staticPartIDs map[api.StaticPartType]string,
		propID string,
	) (*GeneratedAnimation, error)
}
