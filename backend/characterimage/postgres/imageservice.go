package postgres

import (
	"context"
	"log"

	"github.com/lantspants/lootloadout/backend/characterimage"
)

type CharacterImageDB struct {
	l *log.Logger
}

func NewImageService(
	l *log.Logger,
) *CharacterImageDB {
	return &CharacterImageDB{
		l: l,
	}
}

func (r CharacterImageDB) CreateBodyType(
	context.Context,
	*characterimage.BodyType,
) (characterimage.ID, error) {
	return "", nil
}

func (r *CharacterImageDB) GetBodyType(
	context.Context,
	characterimage.ID,
) (*characterimage.BodyType, error) {
	return nil, nil
}

func (r *CharacterImageDB) ListBodyTypes(
	context.Context,
	*characterimage.BodyTypesFilter,
) ([]characterimage.BodyType, error) {
	return nil, nil
}

func (r *CharacterImageDB) CreateStaticPart(
	context.Context,
	*characterimage.StaticPart,
) (characterimage.ID, error) {
	return "", nil
}

func (r *CharacterImageDB) GetStaticPart(
	context.Context,
	characterimage.ID,
) (*characterimage.StaticPart, error) {
	return nil, nil
}

func (r *CharacterImageDB) ListStaticParts(
	context.Context,
	*characterimage.StaticPartsFilter,
) ([]characterimage.StaticPart, error) {
	return nil, nil
}

func (r *CharacterImageDB) CreateDynamicPart(
	context.Context,
	*characterimage.DynamicPart,
) (characterimage.ID, error) {
	return "", nil
}

func (r *CharacterImageDB) GetDynamicPart(
	context.Context,
	characterimage.ID,
) (*characterimage.DynamicPart, error) {
	return nil, nil
}

func (r *CharacterImageDB) ListDynamicParts(
	context.Context,
	*characterimage.DynamicPartsFilter,
) ([]characterimage.DynamicPart, error) {
	return nil, nil
}

func (r *CharacterImageDB) CreateProp(
	context.Context,
	*characterimage.Prop,
) (characterimage.ID, error) {
	return "", nil
}

func (r *CharacterImageDB) GetProp(
	context.Context,
	characterimage.ID,
) (*characterimage.Prop, error) {
	return nil, nil
}

func (r *CharacterImageDB) ListProps(
	context.Context,
	*characterimage.PropsFilter,
) ([]characterimage.Prop, error) {
	return nil, nil
}

func (r *CharacterImageDB) CreateAnimation(
	context.Context,
	*characterimage.Animation,
) (characterimage.ID, error) {
	return "", nil
}

func (r *CharacterImageDB) GetAnimation(
	context.Context,
	characterimage.ID,
) (*characterimage.Animation, error) {
	return nil, nil
}

func (r *CharacterImageDB) ListAnimations(
	context.Context,
	*characterimage.AnimationsFilter,
) ([]characterimage.BodyType, error) {
	return nil, nil
}
