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
	"github.com/volatiletech/strmangle"
)

// DynamicPartPixel is an object representing the database table.
type DynamicPartPixel struct {
	ID            int   `boil:"id" json:"id" toml:"id" yaml:"id"`
	ColorStringID int   `boil:"color_string_id" json:"color_string_id" toml:"color_string_id" yaml:"color_string_id"`
	DynamicPartID int   `boil:"dynamic_part_id" json:"dynamic_part_id" toml:"dynamic_part_id" yaml:"dynamic_part_id"`
	X             int16 `boil:"x" json:"x" toml:"x" yaml:"x"`
	Y             int16 `boil:"y" json:"y" toml:"y" yaml:"y"`

	R *dynamicPartPixelR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L dynamicPartPixelL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var DynamicPartPixelColumns = struct {
	ID            string
	ColorStringID string
	DynamicPartID string
	X             string
	Y             string
}{
	ID:            "id",
	ColorStringID: "color_string_id",
	DynamicPartID: "dynamic_part_id",
	X:             "x",
	Y:             "y",
}

var DynamicPartPixelTableColumns = struct {
	ID            string
	ColorStringID string
	DynamicPartID string
	X             string
	Y             string
}{
	ID:            "dynamic_part_pixel.id",
	ColorStringID: "dynamic_part_pixel.color_string_id",
	DynamicPartID: "dynamic_part_pixel.dynamic_part_id",
	X:             "dynamic_part_pixel.x",
	Y:             "dynamic_part_pixel.y",
}

// Generated where

var DynamicPartPixelWhere = struct {
	ID            whereHelperint
	ColorStringID whereHelperint
	DynamicPartID whereHelperint
	X             whereHelperint16
	Y             whereHelperint16
}{
	ID:            whereHelperint{field: "\"dynamic_part_pixel\".\"id\""},
	ColorStringID: whereHelperint{field: "\"dynamic_part_pixel\".\"color_string_id\""},
	DynamicPartID: whereHelperint{field: "\"dynamic_part_pixel\".\"dynamic_part_id\""},
	X:             whereHelperint16{field: "\"dynamic_part_pixel\".\"x\""},
	Y:             whereHelperint16{field: "\"dynamic_part_pixel\".\"y\""},
}

// DynamicPartPixelRels is where relationship names are stored.
var DynamicPartPixelRels = struct {
	ColorString string
	DynamicPart string
}{
	ColorString: "ColorString",
	DynamicPart: "DynamicPart",
}

// dynamicPartPixelR is where relationships are stored.
type dynamicPartPixelR struct {
	ColorString *ColorString `boil:"ColorString" json:"ColorString" toml:"ColorString" yaml:"ColorString"`
	DynamicPart *DynamicPart `boil:"DynamicPart" json:"DynamicPart" toml:"DynamicPart" yaml:"DynamicPart"`
}

// NewStruct creates a new relationship struct
func (*dynamicPartPixelR) NewStruct() *dynamicPartPixelR {
	return &dynamicPartPixelR{}
}

func (r *dynamicPartPixelR) GetColorString() *ColorString {
	if r == nil {
		return nil
	}
	return r.ColorString
}

func (r *dynamicPartPixelR) GetDynamicPart() *DynamicPart {
	if r == nil {
		return nil
	}
	return r.DynamicPart
}

// dynamicPartPixelL is where Load methods for each relationship are stored.
type dynamicPartPixelL struct{}

var (
	dynamicPartPixelAllColumns            = []string{"id", "color_string_id", "dynamic_part_id", "x", "y"}
	dynamicPartPixelColumnsWithoutDefault = []string{"color_string_id", "dynamic_part_id", "x", "y"}
	dynamicPartPixelColumnsWithDefault    = []string{"id"}
	dynamicPartPixelPrimaryKeyColumns     = []string{"id"}
	dynamicPartPixelGeneratedColumns      = []string{}
)

type (
	// DynamicPartPixelSlice is an alias for a slice of pointers to DynamicPartPixel.
	// This should almost always be used instead of []DynamicPartPixel.
	DynamicPartPixelSlice []*DynamicPartPixel
	// DynamicPartPixelHook is the signature for custom DynamicPartPixel hook methods
	DynamicPartPixelHook func(context.Context, boil.ContextExecutor, *DynamicPartPixel) error

	dynamicPartPixelQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	dynamicPartPixelType                 = reflect.TypeOf(&DynamicPartPixel{})
	dynamicPartPixelMapping              = queries.MakeStructMapping(dynamicPartPixelType)
	dynamicPartPixelPrimaryKeyMapping, _ = queries.BindMapping(dynamicPartPixelType, dynamicPartPixelMapping, dynamicPartPixelPrimaryKeyColumns)
	dynamicPartPixelInsertCacheMut       sync.RWMutex
	dynamicPartPixelInsertCache          = make(map[string]insertCache)
	dynamicPartPixelUpdateCacheMut       sync.RWMutex
	dynamicPartPixelUpdateCache          = make(map[string]updateCache)
	dynamicPartPixelUpsertCacheMut       sync.RWMutex
	dynamicPartPixelUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var dynamicPartPixelAfterSelectHooks []DynamicPartPixelHook

var dynamicPartPixelBeforeInsertHooks []DynamicPartPixelHook
var dynamicPartPixelAfterInsertHooks []DynamicPartPixelHook

var dynamicPartPixelBeforeUpdateHooks []DynamicPartPixelHook
var dynamicPartPixelAfterUpdateHooks []DynamicPartPixelHook

var dynamicPartPixelBeforeDeleteHooks []DynamicPartPixelHook
var dynamicPartPixelAfterDeleteHooks []DynamicPartPixelHook

var dynamicPartPixelBeforeUpsertHooks []DynamicPartPixelHook
var dynamicPartPixelAfterUpsertHooks []DynamicPartPixelHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *DynamicPartPixel) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range dynamicPartPixelAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *DynamicPartPixel) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range dynamicPartPixelBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *DynamicPartPixel) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range dynamicPartPixelAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *DynamicPartPixel) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range dynamicPartPixelBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *DynamicPartPixel) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range dynamicPartPixelAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *DynamicPartPixel) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range dynamicPartPixelBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *DynamicPartPixel) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range dynamicPartPixelAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *DynamicPartPixel) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range dynamicPartPixelBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *DynamicPartPixel) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range dynamicPartPixelAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddDynamicPartPixelHook registers your hook function for all future operations.
func AddDynamicPartPixelHook(hookPoint boil.HookPoint, dynamicPartPixelHook DynamicPartPixelHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		dynamicPartPixelAfterSelectHooks = append(dynamicPartPixelAfterSelectHooks, dynamicPartPixelHook)
	case boil.BeforeInsertHook:
		dynamicPartPixelBeforeInsertHooks = append(dynamicPartPixelBeforeInsertHooks, dynamicPartPixelHook)
	case boil.AfterInsertHook:
		dynamicPartPixelAfterInsertHooks = append(dynamicPartPixelAfterInsertHooks, dynamicPartPixelHook)
	case boil.BeforeUpdateHook:
		dynamicPartPixelBeforeUpdateHooks = append(dynamicPartPixelBeforeUpdateHooks, dynamicPartPixelHook)
	case boil.AfterUpdateHook:
		dynamicPartPixelAfterUpdateHooks = append(dynamicPartPixelAfterUpdateHooks, dynamicPartPixelHook)
	case boil.BeforeDeleteHook:
		dynamicPartPixelBeforeDeleteHooks = append(dynamicPartPixelBeforeDeleteHooks, dynamicPartPixelHook)
	case boil.AfterDeleteHook:
		dynamicPartPixelAfterDeleteHooks = append(dynamicPartPixelAfterDeleteHooks, dynamicPartPixelHook)
	case boil.BeforeUpsertHook:
		dynamicPartPixelBeforeUpsertHooks = append(dynamicPartPixelBeforeUpsertHooks, dynamicPartPixelHook)
	case boil.AfterUpsertHook:
		dynamicPartPixelAfterUpsertHooks = append(dynamicPartPixelAfterUpsertHooks, dynamicPartPixelHook)
	}
}

// One returns a single dynamicPartPixel record from the query.
func (q dynamicPartPixelQuery) One(ctx context.Context, exec boil.ContextExecutor) (*DynamicPartPixel, error) {
	o := &DynamicPartPixel{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for dynamic_part_pixel")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all DynamicPartPixel records from the query.
func (q dynamicPartPixelQuery) All(ctx context.Context, exec boil.ContextExecutor) (DynamicPartPixelSlice, error) {
	var o []*DynamicPartPixel

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to DynamicPartPixel slice")
	}

	if len(dynamicPartPixelAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all DynamicPartPixel records in the query.
func (q dynamicPartPixelQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count dynamic_part_pixel rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q dynamicPartPixelQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if dynamic_part_pixel exists")
	}

	return count > 0, nil
}

// ColorString pointed to by the foreign key.
func (o *DynamicPartPixel) ColorString(mods ...qm.QueryMod) colorStringQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.ColorStringID),
	}

	queryMods = append(queryMods, mods...)

	return ColorStrings(queryMods...)
}

// DynamicPart pointed to by the foreign key.
func (o *DynamicPartPixel) DynamicPart(mods ...qm.QueryMod) dynamicPartQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.DynamicPartID),
	}

	queryMods = append(queryMods, mods...)

	return DynamicParts(queryMods...)
}

// LoadColorString allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (dynamicPartPixelL) LoadColorString(ctx context.Context, e boil.ContextExecutor, singular bool, maybeDynamicPartPixel interface{}, mods queries.Applicator) error {
	var slice []*DynamicPartPixel
	var object *DynamicPartPixel

	if singular {
		object = maybeDynamicPartPixel.(*DynamicPartPixel)
	} else {
		slice = *maybeDynamicPartPixel.(*[]*DynamicPartPixel)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &dynamicPartPixelR{}
		}
		args = append(args, object.ColorStringID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &dynamicPartPixelR{}
			}

			for _, a := range args {
				if a == obj.ColorStringID {
					continue Outer
				}
			}

			args = append(args, obj.ColorStringID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`color_string`),
		qm.WhereIn(`color_string.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load ColorString")
	}

	var resultSlice []*ColorString
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice ColorString")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for color_string")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for color_string")
	}

	if len(dynamicPartPixelAfterSelectHooks) != 0 {
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
		object.R.ColorString = foreign
		if foreign.R == nil {
			foreign.R = &colorStringR{}
		}
		foreign.R.DynamicPartPixels = append(foreign.R.DynamicPartPixels, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.ColorStringID == foreign.ID {
				local.R.ColorString = foreign
				if foreign.R == nil {
					foreign.R = &colorStringR{}
				}
				foreign.R.DynamicPartPixels = append(foreign.R.DynamicPartPixels, local)
				break
			}
		}
	}

	return nil
}

// LoadDynamicPart allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (dynamicPartPixelL) LoadDynamicPart(ctx context.Context, e boil.ContextExecutor, singular bool, maybeDynamicPartPixel interface{}, mods queries.Applicator) error {
	var slice []*DynamicPartPixel
	var object *DynamicPartPixel

	if singular {
		object = maybeDynamicPartPixel.(*DynamicPartPixel)
	} else {
		slice = *maybeDynamicPartPixel.(*[]*DynamicPartPixel)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &dynamicPartPixelR{}
		}
		args = append(args, object.DynamicPartID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &dynamicPartPixelR{}
			}

			for _, a := range args {
				if a == obj.DynamicPartID {
					continue Outer
				}
			}

			args = append(args, obj.DynamicPartID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`dynamic_part`),
		qm.WhereIn(`dynamic_part.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load DynamicPart")
	}

	var resultSlice []*DynamicPart
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice DynamicPart")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for dynamic_part")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for dynamic_part")
	}

	if len(dynamicPartPixelAfterSelectHooks) != 0 {
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
		object.R.DynamicPart = foreign
		if foreign.R == nil {
			foreign.R = &dynamicPartR{}
		}
		foreign.R.DynamicPartPixels = append(foreign.R.DynamicPartPixels, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.DynamicPartID == foreign.ID {
				local.R.DynamicPart = foreign
				if foreign.R == nil {
					foreign.R = &dynamicPartR{}
				}
				foreign.R.DynamicPartPixels = append(foreign.R.DynamicPartPixels, local)
				break
			}
		}
	}

	return nil
}

// SetColorString of the dynamicPartPixel to the related item.
// Sets o.R.ColorString to related.
// Adds o to related.R.DynamicPartPixels.
func (o *DynamicPartPixel) SetColorString(ctx context.Context, exec boil.ContextExecutor, insert bool, related *ColorString) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"dynamic_part_pixel\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"color_string_id"}),
		strmangle.WhereClause("\"", "\"", 2, dynamicPartPixelPrimaryKeyColumns),
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

	o.ColorStringID = related.ID
	if o.R == nil {
		o.R = &dynamicPartPixelR{
			ColorString: related,
		}
	} else {
		o.R.ColorString = related
	}

	if related.R == nil {
		related.R = &colorStringR{
			DynamicPartPixels: DynamicPartPixelSlice{o},
		}
	} else {
		related.R.DynamicPartPixels = append(related.R.DynamicPartPixels, o)
	}

	return nil
}

// SetDynamicPart of the dynamicPartPixel to the related item.
// Sets o.R.DynamicPart to related.
// Adds o to related.R.DynamicPartPixels.
func (o *DynamicPartPixel) SetDynamicPart(ctx context.Context, exec boil.ContextExecutor, insert bool, related *DynamicPart) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"dynamic_part_pixel\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"dynamic_part_id"}),
		strmangle.WhereClause("\"", "\"", 2, dynamicPartPixelPrimaryKeyColumns),
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

	o.DynamicPartID = related.ID
	if o.R == nil {
		o.R = &dynamicPartPixelR{
			DynamicPart: related,
		}
	} else {
		o.R.DynamicPart = related
	}

	if related.R == nil {
		related.R = &dynamicPartR{
			DynamicPartPixels: DynamicPartPixelSlice{o},
		}
	} else {
		related.R.DynamicPartPixels = append(related.R.DynamicPartPixels, o)
	}

	return nil
}

// DynamicPartPixels retrieves all the records using an executor.
func DynamicPartPixels(mods ...qm.QueryMod) dynamicPartPixelQuery {
	mods = append(mods, qm.From("\"dynamic_part_pixel\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"dynamic_part_pixel\".*"})
	}

	return dynamicPartPixelQuery{q}
}

// FindDynamicPartPixel retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindDynamicPartPixel(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*DynamicPartPixel, error) {
	dynamicPartPixelObj := &DynamicPartPixel{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"dynamic_part_pixel\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, dynamicPartPixelObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from dynamic_part_pixel")
	}

	if err = dynamicPartPixelObj.doAfterSelectHooks(ctx, exec); err != nil {
		return dynamicPartPixelObj, err
	}

	return dynamicPartPixelObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *DynamicPartPixel) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no dynamic_part_pixel provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(dynamicPartPixelColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	dynamicPartPixelInsertCacheMut.RLock()
	cache, cached := dynamicPartPixelInsertCache[key]
	dynamicPartPixelInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			dynamicPartPixelAllColumns,
			dynamicPartPixelColumnsWithDefault,
			dynamicPartPixelColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(dynamicPartPixelType, dynamicPartPixelMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(dynamicPartPixelType, dynamicPartPixelMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"dynamic_part_pixel\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"dynamic_part_pixel\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into dynamic_part_pixel")
	}

	if !cached {
		dynamicPartPixelInsertCacheMut.Lock()
		dynamicPartPixelInsertCache[key] = cache
		dynamicPartPixelInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the DynamicPartPixel.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *DynamicPartPixel) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	dynamicPartPixelUpdateCacheMut.RLock()
	cache, cached := dynamicPartPixelUpdateCache[key]
	dynamicPartPixelUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			dynamicPartPixelAllColumns,
			dynamicPartPixelPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update dynamic_part_pixel, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"dynamic_part_pixel\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, dynamicPartPixelPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(dynamicPartPixelType, dynamicPartPixelMapping, append(wl, dynamicPartPixelPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update dynamic_part_pixel row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for dynamic_part_pixel")
	}

	if !cached {
		dynamicPartPixelUpdateCacheMut.Lock()
		dynamicPartPixelUpdateCache[key] = cache
		dynamicPartPixelUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q dynamicPartPixelQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for dynamic_part_pixel")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for dynamic_part_pixel")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o DynamicPartPixelSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), dynamicPartPixelPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"dynamic_part_pixel\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, dynamicPartPixelPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in dynamicPartPixel slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all dynamicPartPixel")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *DynamicPartPixel) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no dynamic_part_pixel provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(dynamicPartPixelColumnsWithDefault, o)

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

	dynamicPartPixelUpsertCacheMut.RLock()
	cache, cached := dynamicPartPixelUpsertCache[key]
	dynamicPartPixelUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			dynamicPartPixelAllColumns,
			dynamicPartPixelColumnsWithDefault,
			dynamicPartPixelColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			dynamicPartPixelAllColumns,
			dynamicPartPixelPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert dynamic_part_pixel, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(dynamicPartPixelPrimaryKeyColumns))
			copy(conflict, dynamicPartPixelPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"dynamic_part_pixel\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(dynamicPartPixelType, dynamicPartPixelMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(dynamicPartPixelType, dynamicPartPixelMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert dynamic_part_pixel")
	}

	if !cached {
		dynamicPartPixelUpsertCacheMut.Lock()
		dynamicPartPixelUpsertCache[key] = cache
		dynamicPartPixelUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single DynamicPartPixel record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *DynamicPartPixel) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no DynamicPartPixel provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), dynamicPartPixelPrimaryKeyMapping)
	sql := "DELETE FROM \"dynamic_part_pixel\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from dynamic_part_pixel")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for dynamic_part_pixel")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q dynamicPartPixelQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no dynamicPartPixelQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from dynamic_part_pixel")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for dynamic_part_pixel")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o DynamicPartPixelSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(dynamicPartPixelBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), dynamicPartPixelPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"dynamic_part_pixel\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, dynamicPartPixelPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from dynamicPartPixel slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for dynamic_part_pixel")
	}

	if len(dynamicPartPixelAfterDeleteHooks) != 0 {
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
func (o *DynamicPartPixel) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindDynamicPartPixel(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *DynamicPartPixelSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := DynamicPartPixelSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), dynamicPartPixelPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"dynamic_part_pixel\".* FROM \"dynamic_part_pixel\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, dynamicPartPixelPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in DynamicPartPixelSlice")
	}

	*o = slice

	return nil
}

// DynamicPartPixelExists checks if the DynamicPartPixel row exists.
func DynamicPartPixelExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"dynamic_part_pixel\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if dynamic_part_pixel exists")
	}

	return exists, nil
}
