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
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// StaticPart is an object representing the database table.
type StaticPart struct {
	ID          int                `boil:"id" json:"id" toml:"id" yaml:"id"`
	BodyTypeID  null.Int           `boil:"body_type_id" json:"body_type_id,omitempty" toml:"body_type_id" yaml:"body_type_id,omitempty"`
	DisplayName string             `boil:"display_name" json:"display_name" toml:"display_name" yaml:"display_name"`
	PartType    NullStaticPartType `boil:"part_type" json:"part_type,omitempty" toml:"part_type" yaml:"part_type,omitempty"`

	R *staticPartR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L staticPartL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var StaticPartColumns = struct {
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

var StaticPartTableColumns = struct {
	ID          string
	BodyTypeID  string
	DisplayName string
	PartType    string
}{
	ID:          "static_part.id",
	BodyTypeID:  "static_part.body_type_id",
	DisplayName: "static_part.display_name",
	PartType:    "static_part.part_type",
}

// Generated where

type whereHelperNullStaticPartType struct{ field string }

func (w whereHelperNullStaticPartType) EQ(x NullStaticPartType) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelperNullStaticPartType) NEQ(x NullStaticPartType) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelperNullStaticPartType) LT(x NullStaticPartType) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelperNullStaticPartType) LTE(x NullStaticPartType) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelperNullStaticPartType) GT(x NullStaticPartType) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelperNullStaticPartType) GTE(x NullStaticPartType) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

func (w whereHelperNullStaticPartType) IsNull() qm.QueryMod { return qmhelper.WhereIsNull(w.field) }
func (w whereHelperNullStaticPartType) IsNotNull() qm.QueryMod {
	return qmhelper.WhereIsNotNull(w.field)
}

var StaticPartWhere = struct {
	ID          whereHelperint
	BodyTypeID  whereHelpernull_Int
	DisplayName whereHelperstring
	PartType    whereHelperNullStaticPartType
}{
	ID:          whereHelperint{field: "\"static_part\".\"id\""},
	BodyTypeID:  whereHelpernull_Int{field: "\"static_part\".\"body_type_id\""},
	DisplayName: whereHelperstring{field: "\"static_part\".\"display_name\""},
	PartType:    whereHelperNullStaticPartType{field: "\"static_part\".\"part_type\""},
}

// StaticPartRels is where relationship names are stored.
var StaticPartRels = struct {
	BodyType        string
	StaticPartImage string
}{
	BodyType:        "BodyType",
	StaticPartImage: "StaticPartImage",
}

// staticPartR is where relationships are stored.
type staticPartR struct {
	BodyType        *BodyType        `boil:"BodyType" json:"BodyType" toml:"BodyType" yaml:"BodyType"`
	StaticPartImage *StaticPartImage `boil:"StaticPartImage" json:"StaticPartImage" toml:"StaticPartImage" yaml:"StaticPartImage"`
}

// NewStruct creates a new relationship struct
func (*staticPartR) NewStruct() *staticPartR {
	return &staticPartR{}
}

func (r *staticPartR) GetBodyType() *BodyType {
	if r == nil {
		return nil
	}
	return r.BodyType
}

func (r *staticPartR) GetStaticPartImage() *StaticPartImage {
	if r == nil {
		return nil
	}
	return r.StaticPartImage
}

// staticPartL is where Load methods for each relationship are stored.
type staticPartL struct{}

var (
	staticPartAllColumns            = []string{"id", "body_type_id", "display_name", "part_type"}
	staticPartColumnsWithoutDefault = []string{"display_name"}
	staticPartColumnsWithDefault    = []string{"id", "body_type_id", "part_type"}
	staticPartPrimaryKeyColumns     = []string{"id"}
	staticPartGeneratedColumns      = []string{}
)

type (
	// StaticPartSlice is an alias for a slice of pointers to StaticPart.
	// This should almost always be used instead of []StaticPart.
	StaticPartSlice []*StaticPart
	// StaticPartHook is the signature for custom StaticPart hook methods
	StaticPartHook func(context.Context, boil.ContextExecutor, *StaticPart) error

	staticPartQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	staticPartType                 = reflect.TypeOf(&StaticPart{})
	staticPartMapping              = queries.MakeStructMapping(staticPartType)
	staticPartPrimaryKeyMapping, _ = queries.BindMapping(staticPartType, staticPartMapping, staticPartPrimaryKeyColumns)
	staticPartInsertCacheMut       sync.RWMutex
	staticPartInsertCache          = make(map[string]insertCache)
	staticPartUpdateCacheMut       sync.RWMutex
	staticPartUpdateCache          = make(map[string]updateCache)
	staticPartUpsertCacheMut       sync.RWMutex
	staticPartUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var staticPartAfterSelectHooks []StaticPartHook

var staticPartBeforeInsertHooks []StaticPartHook
var staticPartAfterInsertHooks []StaticPartHook

var staticPartBeforeUpdateHooks []StaticPartHook
var staticPartAfterUpdateHooks []StaticPartHook

var staticPartBeforeDeleteHooks []StaticPartHook
var staticPartAfterDeleteHooks []StaticPartHook

var staticPartBeforeUpsertHooks []StaticPartHook
var staticPartAfterUpsertHooks []StaticPartHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *StaticPart) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range staticPartAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *StaticPart) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range staticPartBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *StaticPart) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range staticPartAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *StaticPart) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range staticPartBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *StaticPart) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range staticPartAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *StaticPart) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range staticPartBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *StaticPart) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range staticPartAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *StaticPart) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range staticPartBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *StaticPart) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range staticPartAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddStaticPartHook registers your hook function for all future operations.
func AddStaticPartHook(hookPoint boil.HookPoint, staticPartHook StaticPartHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		staticPartAfterSelectHooks = append(staticPartAfterSelectHooks, staticPartHook)
	case boil.BeforeInsertHook:
		staticPartBeforeInsertHooks = append(staticPartBeforeInsertHooks, staticPartHook)
	case boil.AfterInsertHook:
		staticPartAfterInsertHooks = append(staticPartAfterInsertHooks, staticPartHook)
	case boil.BeforeUpdateHook:
		staticPartBeforeUpdateHooks = append(staticPartBeforeUpdateHooks, staticPartHook)
	case boil.AfterUpdateHook:
		staticPartAfterUpdateHooks = append(staticPartAfterUpdateHooks, staticPartHook)
	case boil.BeforeDeleteHook:
		staticPartBeforeDeleteHooks = append(staticPartBeforeDeleteHooks, staticPartHook)
	case boil.AfterDeleteHook:
		staticPartAfterDeleteHooks = append(staticPartAfterDeleteHooks, staticPartHook)
	case boil.BeforeUpsertHook:
		staticPartBeforeUpsertHooks = append(staticPartBeforeUpsertHooks, staticPartHook)
	case boil.AfterUpsertHook:
		staticPartAfterUpsertHooks = append(staticPartAfterUpsertHooks, staticPartHook)
	}
}

// One returns a single staticPart record from the query.
func (q staticPartQuery) One(ctx context.Context, exec boil.ContextExecutor) (*StaticPart, error) {
	o := &StaticPart{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for static_part")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all StaticPart records from the query.
func (q staticPartQuery) All(ctx context.Context, exec boil.ContextExecutor) (StaticPartSlice, error) {
	var o []*StaticPart

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to StaticPart slice")
	}

	if len(staticPartAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all StaticPart records in the query.
func (q staticPartQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count static_part rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q staticPartQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if static_part exists")
	}

	return count > 0, nil
}

// BodyType pointed to by the foreign key.
func (o *StaticPart) BodyType(mods ...qm.QueryMod) bodyTypeQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.BodyTypeID),
	}

	queryMods = append(queryMods, mods...)

	return BodyTypes(queryMods...)
}

// StaticPartImage pointed to by the foreign key.
func (o *StaticPart) StaticPartImage(mods ...qm.QueryMod) staticPartImageQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"static_part_id\" = ?", o.ID),
	}

	queryMods = append(queryMods, mods...)

	return StaticPartImages(queryMods...)
}

// LoadBodyType allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (staticPartL) LoadBodyType(ctx context.Context, e boil.ContextExecutor, singular bool, maybeStaticPart interface{}, mods queries.Applicator) error {
	var slice []*StaticPart
	var object *StaticPart

	if singular {
		object = maybeStaticPart.(*StaticPart)
	} else {
		slice = *maybeStaticPart.(*[]*StaticPart)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &staticPartR{}
		}
		if !queries.IsNil(object.BodyTypeID) {
			args = append(args, object.BodyTypeID)
		}

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &staticPartR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.BodyTypeID) {
					continue Outer
				}
			}

			if !queries.IsNil(obj.BodyTypeID) {
				args = append(args, obj.BodyTypeID)
			}

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

	if len(staticPartAfterSelectHooks) != 0 {
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
		foreign.R.StaticParts = append(foreign.R.StaticParts, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if queries.Equal(local.BodyTypeID, foreign.ID) {
				local.R.BodyType = foreign
				if foreign.R == nil {
					foreign.R = &bodyTypeR{}
				}
				foreign.R.StaticParts = append(foreign.R.StaticParts, local)
				break
			}
		}
	}

	return nil
}

// LoadStaticPartImage allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-1 relationship.
func (staticPartL) LoadStaticPartImage(ctx context.Context, e boil.ContextExecutor, singular bool, maybeStaticPart interface{}, mods queries.Applicator) error {
	var slice []*StaticPart
	var object *StaticPart

	if singular {
		object = maybeStaticPart.(*StaticPart)
	} else {
		slice = *maybeStaticPart.(*[]*StaticPart)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &staticPartR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &staticPartR{}
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
		qm.From(`static_part_image`),
		qm.WhereIn(`static_part_image.static_part_id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load StaticPartImage")
	}

	var resultSlice []*StaticPartImage
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice StaticPartImage")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for static_part_image")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for static_part_image")
	}

	if len(staticPartAfterSelectHooks) != 0 {
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
		object.R.StaticPartImage = foreign
		if foreign.R == nil {
			foreign.R = &staticPartImageR{}
		}
		foreign.R.StaticPart = object
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.ID == foreign.StaticPartID {
				local.R.StaticPartImage = foreign
				if foreign.R == nil {
					foreign.R = &staticPartImageR{}
				}
				foreign.R.StaticPart = local
				break
			}
		}
	}

	return nil
}

// SetBodyType of the staticPart to the related item.
// Sets o.R.BodyType to related.
// Adds o to related.R.StaticParts.
func (o *StaticPart) SetBodyType(ctx context.Context, exec boil.ContextExecutor, insert bool, related *BodyType) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"static_part\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"body_type_id"}),
		strmangle.WhereClause("\"", "\"", 2, staticPartPrimaryKeyColumns),
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

	queries.Assign(&o.BodyTypeID, related.ID)
	if o.R == nil {
		o.R = &staticPartR{
			BodyType: related,
		}
	} else {
		o.R.BodyType = related
	}

	if related.R == nil {
		related.R = &bodyTypeR{
			StaticParts: StaticPartSlice{o},
		}
	} else {
		related.R.StaticParts = append(related.R.StaticParts, o)
	}

	return nil
}

// RemoveBodyType relationship.
// Sets o.R.BodyType to nil.
// Removes o from all passed in related items' relationships struct.
func (o *StaticPart) RemoveBodyType(ctx context.Context, exec boil.ContextExecutor, related *BodyType) error {
	var err error

	queries.SetScanner(&o.BodyTypeID, nil)
	if _, err = o.Update(ctx, exec, boil.Whitelist("body_type_id")); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	if o.R != nil {
		o.R.BodyType = nil
	}
	if related == nil || related.R == nil {
		return nil
	}

	for i, ri := range related.R.StaticParts {
		if queries.Equal(o.BodyTypeID, ri.BodyTypeID) {
			continue
		}

		ln := len(related.R.StaticParts)
		if ln > 1 && i < ln-1 {
			related.R.StaticParts[i] = related.R.StaticParts[ln-1]
		}
		related.R.StaticParts = related.R.StaticParts[:ln-1]
		break
	}
	return nil
}

// SetStaticPartImage of the staticPart to the related item.
// Sets o.R.StaticPartImage to related.
// Adds o to related.R.StaticPart.
func (o *StaticPart) SetStaticPartImage(ctx context.Context, exec boil.ContextExecutor, insert bool, related *StaticPartImage) error {
	var err error

	if insert {
		related.StaticPartID = o.ID

		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"static_part_image\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"static_part_id"}),
			strmangle.WhereClause("\"", "\"", 2, staticPartImagePrimaryKeyColumns),
		)
		values := []interface{}{o.ID, related.StaticPartID}

		if boil.IsDebug(ctx) {
			writer := boil.DebugWriterFrom(ctx)
			fmt.Fprintln(writer, updateQuery)
			fmt.Fprintln(writer, values)
		}
		if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.StaticPartID = o.ID
	}

	if o.R == nil {
		o.R = &staticPartR{
			StaticPartImage: related,
		}
	} else {
		o.R.StaticPartImage = related
	}

	if related.R == nil {
		related.R = &staticPartImageR{
			StaticPart: o,
		}
	} else {
		related.R.StaticPart = o
	}
	return nil
}

// StaticParts retrieves all the records using an executor.
func StaticParts(mods ...qm.QueryMod) staticPartQuery {
	mods = append(mods, qm.From("\"static_part\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"static_part\".*"})
	}

	return staticPartQuery{q}
}

// FindStaticPart retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindStaticPart(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*StaticPart, error) {
	staticPartObj := &StaticPart{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"static_part\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, staticPartObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from static_part")
	}

	if err = staticPartObj.doAfterSelectHooks(ctx, exec); err != nil {
		return staticPartObj, err
	}

	return staticPartObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *StaticPart) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no static_part provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(staticPartColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	staticPartInsertCacheMut.RLock()
	cache, cached := staticPartInsertCache[key]
	staticPartInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			staticPartAllColumns,
			staticPartColumnsWithDefault,
			staticPartColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(staticPartType, staticPartMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(staticPartType, staticPartMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"static_part\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"static_part\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into static_part")
	}

	if !cached {
		staticPartInsertCacheMut.Lock()
		staticPartInsertCache[key] = cache
		staticPartInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the StaticPart.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *StaticPart) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	staticPartUpdateCacheMut.RLock()
	cache, cached := staticPartUpdateCache[key]
	staticPartUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			staticPartAllColumns,
			staticPartPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update static_part, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"static_part\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, staticPartPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(staticPartType, staticPartMapping, append(wl, staticPartPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update static_part row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for static_part")
	}

	if !cached {
		staticPartUpdateCacheMut.Lock()
		staticPartUpdateCache[key] = cache
		staticPartUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q staticPartQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for static_part")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for static_part")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o StaticPartSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), staticPartPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"static_part\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, staticPartPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in staticPart slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all staticPart")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *StaticPart) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no static_part provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(staticPartColumnsWithDefault, o)

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

	staticPartUpsertCacheMut.RLock()
	cache, cached := staticPartUpsertCache[key]
	staticPartUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			staticPartAllColumns,
			staticPartColumnsWithDefault,
			staticPartColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			staticPartAllColumns,
			staticPartPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert static_part, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(staticPartPrimaryKeyColumns))
			copy(conflict, staticPartPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"static_part\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(staticPartType, staticPartMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(staticPartType, staticPartMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert static_part")
	}

	if !cached {
		staticPartUpsertCacheMut.Lock()
		staticPartUpsertCache[key] = cache
		staticPartUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single StaticPart record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *StaticPart) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no StaticPart provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), staticPartPrimaryKeyMapping)
	sql := "DELETE FROM \"static_part\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from static_part")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for static_part")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q staticPartQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no staticPartQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from static_part")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for static_part")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o StaticPartSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(staticPartBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), staticPartPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"static_part\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, staticPartPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from staticPart slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for static_part")
	}

	if len(staticPartAfterDeleteHooks) != 0 {
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
func (o *StaticPart) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindStaticPart(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *StaticPartSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := StaticPartSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), staticPartPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"static_part\".* FROM \"static_part\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, staticPartPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in StaticPartSlice")
	}

	*o = slice

	return nil
}

// StaticPartExists checks if the StaticPart row exists.
func StaticPartExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"static_part\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if static_part exists")
	}

	return exists, nil
}
