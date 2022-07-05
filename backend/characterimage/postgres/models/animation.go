// Code generated by SQLBoiler 4.11.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/v4/types"
	"github.com/volatiletech/strmangle"
)

// Animation is an object representing the database table.
type Animation struct {
	ID          int               `boil:"id" json:"id" toml:"id" yaml:"id"`
	BodyTypeID  int               `boil:"body_type_id" json:"body_type_id" toml:"body_type_id" yaml:"body_type_id"`
	DisplayName string            `boil:"display_name" json:"display_name" toml:"display_name" yaml:"display_name"`
	PartType    types.StringArray `boil:"part_type" json:"part_type" toml:"part_type" yaml:"part_type"`

	R *animationR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L animationL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var AnimationColumns = struct {
	ID          string
	BodyTypeID  string
	DisplayName string
	PartType    string
}{
	ID:          "id",
	BodyTypeID:  "body_type_id",
	DisplayName: "display_name",
	PartType:    "part_type",
}

var AnimationTableColumns = struct {
	ID          string
	BodyTypeID  string
	DisplayName string
	PartType    string
}{
	ID:          "animation.id",
	BodyTypeID:  "animation.body_type_id",
	DisplayName: "animation.display_name",
	PartType:    "animation.part_type",
}

// Generated where

type whereHelperint struct{ field string }

func (w whereHelperint) EQ(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperint) NEQ(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperint) LT(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperint) LTE(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperint) GT(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperint) GTE(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperint) IN(slice []int) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperint) NIN(slice []int) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelperstring struct{ field string }

func (w whereHelperstring) EQ(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperstring) NEQ(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperstring) LT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperstring) LTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperstring) GT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperstring) GTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperstring) IN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperstring) NIN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelpertypes_StringArray struct{ field string }

func (w whereHelpertypes_StringArray) EQ(x types.StringArray) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.EQ, x)
}
func (w whereHelpertypes_StringArray) NEQ(x types.StringArray) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelpertypes_StringArray) LT(x types.StringArray) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpertypes_StringArray) LTE(x types.StringArray) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpertypes_StringArray) GT(x types.StringArray) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpertypes_StringArray) GTE(x types.StringArray) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

var AnimationWhere = struct {
	ID          whereHelperint
	BodyTypeID  whereHelperint
	DisplayName whereHelperstring
	PartType    whereHelpertypes_StringArray
}{
	ID:          whereHelperint{field: "\"animation\".\"id\""},
	BodyTypeID:  whereHelperint{field: "\"animation\".\"body_type_id\""},
	DisplayName: whereHelperstring{field: "\"animation\".\"display_name\""},
	PartType:    whereHelpertypes_StringArray{field: "\"animation\".\"part_type\""},
}

// AnimationRels is where relationship names are stored.
var AnimationRels = struct {
	BodyType        string
	AnimationFrames string
}{
	BodyType:        "BodyType",
	AnimationFrames: "AnimationFrames",
}

// animationR is where relationships are stored.
type animationR struct {
	BodyType        *BodyType           `boil:"BodyType" json:"BodyType" toml:"BodyType" yaml:"BodyType"`
	AnimationFrames AnimationFrameSlice `boil:"AnimationFrames" json:"AnimationFrames" toml:"AnimationFrames" yaml:"AnimationFrames"`
}

// NewStruct creates a new relationship struct
func (*animationR) NewStruct() *animationR {
	return &animationR{}
}

func (r *animationR) GetBodyType() *BodyType {
	if r == nil {
		return nil
	}
	return r.BodyType
}

func (r *animationR) GetAnimationFrames() AnimationFrameSlice {
	if r == nil {
		return nil
	}
	return r.AnimationFrames
}

// animationL is where Load methods for each relationship are stored.
type animationL struct{}

var (
	animationAllColumns            = []string{"id", "body_type_id", "display_name", "part_type"}
	animationColumnsWithoutDefault = []string{"body_type_id", "display_name", "part_type"}
	animationColumnsWithDefault    = []string{"id"}
	animationPrimaryKeyColumns     = []string{"id"}
	animationGeneratedColumns      = []string{}
)

type (
	// AnimationSlice is an alias for a slice of pointers to Animation.
	// This should almost always be used instead of []Animation.
	AnimationSlice []*Animation
	// AnimationHook is the signature for custom Animation hook methods
	AnimationHook func(context.Context, boil.ContextExecutor, *Animation) error

	animationQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	animationType                 = reflect.TypeOf(&Animation{})
	animationMapping              = queries.MakeStructMapping(animationType)
	animationPrimaryKeyMapping, _ = queries.BindMapping(animationType, animationMapping, animationPrimaryKeyColumns)
	animationInsertCacheMut       sync.RWMutex
	animationInsertCache          = make(map[string]insertCache)
	animationUpdateCacheMut       sync.RWMutex
	animationUpdateCache          = make(map[string]updateCache)
	animationUpsertCacheMut       sync.RWMutex
	animationUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var animationAfterSelectHooks []AnimationHook

var animationBeforeInsertHooks []AnimationHook
var animationAfterInsertHooks []AnimationHook

var animationBeforeUpdateHooks []AnimationHook
var animationAfterUpdateHooks []AnimationHook

var animationBeforeDeleteHooks []AnimationHook
var animationAfterDeleteHooks []AnimationHook

var animationBeforeUpsertHooks []AnimationHook
var animationAfterUpsertHooks []AnimationHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Animation) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range animationAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Animation) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range animationBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Animation) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range animationAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Animation) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range animationBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Animation) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range animationAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Animation) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range animationBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Animation) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range animationAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Animation) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range animationBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Animation) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range animationAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddAnimationHook registers your hook function for all future operations.
func AddAnimationHook(hookPoint boil.HookPoint, animationHook AnimationHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		animationAfterSelectHooks = append(animationAfterSelectHooks, animationHook)
	case boil.BeforeInsertHook:
		animationBeforeInsertHooks = append(animationBeforeInsertHooks, animationHook)
	case boil.AfterInsertHook:
		animationAfterInsertHooks = append(animationAfterInsertHooks, animationHook)
	case boil.BeforeUpdateHook:
		animationBeforeUpdateHooks = append(animationBeforeUpdateHooks, animationHook)
	case boil.AfterUpdateHook:
		animationAfterUpdateHooks = append(animationAfterUpdateHooks, animationHook)
	case boil.BeforeDeleteHook:
		animationBeforeDeleteHooks = append(animationBeforeDeleteHooks, animationHook)
	case boil.AfterDeleteHook:
		animationAfterDeleteHooks = append(animationAfterDeleteHooks, animationHook)
	case boil.BeforeUpsertHook:
		animationBeforeUpsertHooks = append(animationBeforeUpsertHooks, animationHook)
	case boil.AfterUpsertHook:
		animationAfterUpsertHooks = append(animationAfterUpsertHooks, animationHook)
	}
}

// One returns a single animation record from the query.
func (q animationQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Animation, error) {
	o := &Animation{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for animation")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Animation records from the query.
func (q animationQuery) All(ctx context.Context, exec boil.ContextExecutor) (AnimationSlice, error) {
	var o []*Animation

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Animation slice")
	}

	if len(animationAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Animation records in the query.
func (q animationQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count animation rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q animationQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if animation exists")
	}

	return count > 0, nil
}

// BodyType pointed to by the foreign key.
func (o *Animation) BodyType(mods ...qm.QueryMod) bodyTypeQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.BodyTypeID),
	}

	queryMods = append(queryMods, mods...)

	return BodyTypes(queryMods...)
}

// AnimationFrames retrieves all the animation_frame's AnimationFrames with an executor.
func (o *Animation) AnimationFrames(mods ...qm.QueryMod) animationFrameQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"animation_frame\".\"animation_id\"=?", o.ID),
	)

	return AnimationFrames(queryMods...)
}

// LoadBodyType allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (animationL) LoadBodyType(ctx context.Context, e boil.ContextExecutor, singular bool, maybeAnimation interface{}, mods queries.Applicator) error {
	var slice []*Animation
	var object *Animation

	if singular {
		object = maybeAnimation.(*Animation)
	} else {
		slice = *maybeAnimation.(*[]*Animation)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &animationR{}
		}
		args = append(args, object.BodyTypeID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &animationR{}
			}

			for _, a := range args {
				if a == obj.BodyTypeID {
					continue Outer
				}
			}

			args = append(args, obj.BodyTypeID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`body_type`),
		qm.WhereIn(`body_type.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load BodyType")
	}

	var resultSlice []*BodyType
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice BodyType")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for body_type")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for body_type")
	}

	if len(animationAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.BodyType = foreign
		if foreign.R == nil {
			foreign.R = &bodyTypeR{}
		}
		foreign.R.Animations = append(foreign.R.Animations, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.BodyTypeID == foreign.ID {
				local.R.BodyType = foreign
				if foreign.R == nil {
					foreign.R = &bodyTypeR{}
				}
				foreign.R.Animations = append(foreign.R.Animations, local)
				break
			}
		}
	}

	return nil
}

// LoadAnimationFrames allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (animationL) LoadAnimationFrames(ctx context.Context, e boil.ContextExecutor, singular bool, maybeAnimation interface{}, mods queries.Applicator) error {
	var slice []*Animation
	var object *Animation

	if singular {
		object = maybeAnimation.(*Animation)
	} else {
		slice = *maybeAnimation.(*[]*Animation)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &animationR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &animationR{}
			}

			for _, a := range args {
				if a == obj.ID {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`animation_frame`),
		qm.WhereIn(`animation_frame.animation_id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load animation_frame")
	}

	var resultSlice []*AnimationFrame
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice animation_frame")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on animation_frame")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for animation_frame")
	}

	if len(animationFrameAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.AnimationFrames = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &animationFrameR{}
			}
			foreign.R.Animation = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.AnimationID {
				local.R.AnimationFrames = append(local.R.AnimationFrames, foreign)
				if foreign.R == nil {
					foreign.R = &animationFrameR{}
				}
				foreign.R.Animation = local
				break
			}
		}
	}

	return nil
}

// SetBodyType of the animation to the related item.
// Sets o.R.BodyType to related.
// Adds o to related.R.Animations.
func (o *Animation) SetBodyType(ctx context.Context, exec boil.ContextExecutor, insert bool, related *BodyType) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"animation\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"body_type_id"}),
		strmangle.WhereClause("\"", "\"", 2, animationPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.BodyTypeID = related.ID
	if o.R == nil {
		o.R = &animationR{
			BodyType: related,
		}
	} else {
		o.R.BodyType = related
	}

	if related.R == nil {
		related.R = &bodyTypeR{
			Animations: AnimationSlice{o},
		}
	} else {
		related.R.Animations = append(related.R.Animations, o)
	}

	return nil
}

// AddAnimationFrames adds the given related objects to the existing relationships
// of the animation, optionally inserting them as new records.
// Appends related to o.R.AnimationFrames.
// Sets related.R.Animation appropriately.
func (o *Animation) AddAnimationFrames(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*AnimationFrame) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.AnimationID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"animation_frame\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"animation_id"}),
				strmangle.WhereClause("\"", "\"", 2, animationFramePrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.AnimationID = o.ID
		}
	}

	if o.R == nil {
		o.R = &animationR{
			AnimationFrames: related,
		}
	} else {
		o.R.AnimationFrames = append(o.R.AnimationFrames, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &animationFrameR{
				Animation: o,
			}
		} else {
			rel.R.Animation = o
		}
	}
	return nil
}

// Animations retrieves all the records using an executor.
func Animations(mods ...qm.QueryMod) animationQuery {
	mods = append(mods, qm.From("\"animation\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"animation\".*"})
	}

	return animationQuery{q}
}

// FindAnimation retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindAnimation(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*Animation, error) {
	animationObj := &Animation{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"animation\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, animationObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from animation")
	}

	if err = animationObj.doAfterSelectHooks(ctx, exec); err != nil {
		return animationObj, err
	}

	return animationObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Animation) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no animation provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(animationColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	animationInsertCacheMut.RLock()
	cache, cached := animationInsertCache[key]
	animationInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			animationAllColumns,
			animationColumnsWithDefault,
			animationColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(animationType, animationMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(animationType, animationMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"animation\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"animation\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into animation")
	}

	if !cached {
		animationInsertCacheMut.Lock()
		animationInsertCache[key] = cache
		animationInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Animation.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Animation) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	animationUpdateCacheMut.RLock()
	cache, cached := animationUpdateCache[key]
	animationUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			animationAllColumns,
			animationPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update animation, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"animation\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, animationPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(animationType, animationMapping, append(wl, animationPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update animation row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for animation")
	}

	if !cached {
		animationUpdateCacheMut.Lock()
		animationUpdateCache[key] = cache
		animationUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q animationQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for animation")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for animation")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o AnimationSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), animationPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"animation\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, animationPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in animation slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all animation")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Animation) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no animation provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(animationColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	animationUpsertCacheMut.RLock()
	cache, cached := animationUpsertCache[key]
	animationUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			animationAllColumns,
			animationColumnsWithDefault,
			animationColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			animationAllColumns,
			animationPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert animation, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(animationPrimaryKeyColumns))
			copy(conflict, animationPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"animation\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(animationType, animationMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(animationType, animationMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if errors.Is(err, sql.ErrNoRows) {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert animation")
	}

	if !cached {
		animationUpsertCacheMut.Lock()
		animationUpsertCache[key] = cache
		animationUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Animation record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Animation) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Animation provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), animationPrimaryKeyMapping)
	sql := "DELETE FROM \"animation\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from animation")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for animation")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q animationQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no animationQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from animation")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for animation")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o AnimationSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(animationBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), animationPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"animation\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, animationPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from animation slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for animation")
	}

	if len(animationAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Animation) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindAnimation(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *AnimationSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := AnimationSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), animationPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"animation\".* FROM \"animation\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, animationPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in AnimationSlice")
	}

	*o = slice

	return nil
}

// AnimationExists checks if the Animation row exists.
func AnimationExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"animation\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if animation exists")
	}

	return exists, nil
}