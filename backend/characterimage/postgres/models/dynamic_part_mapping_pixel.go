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

// DynamicPartMappingPixel is an object representing the database table.
type DynamicPartMappingPixel struct {
	ColorStringID        int   `boil:"color_string_id" json:"color_string_id" toml:"color_string_id" yaml:"color_string_id"`
	DynamicPartMappingID int   `boil:"dynamic_part_mapping_id" json:"dynamic_part_mapping_id" toml:"dynamic_part_mapping_id" yaml:"dynamic_part_mapping_id"`
	X                    int16 `boil:"x" json:"x" toml:"x" yaml:"x"`
	Y                    int16 `boil:"y" json:"y" toml:"y" yaml:"y"`

	R *dynamicPartMappingPixelR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L dynamicPartMappingPixelL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var DynamicPartMappingPixelColumns = struct {
	ColorStringID        string
	DynamicPartMappingID string
	X                    string
	Y                    string
}{
	ColorStringID:        "color_string_id",
	DynamicPartMappingID: "dynamic_part_mapping_id",
	X:                    "x",
	Y:                    "y",
}

var DynamicPartMappingPixelTableColumns = struct {
	ColorStringID        string
	DynamicPartMappingID string
	X                    string
	Y                    string
}{
	ColorStringID:        "dynamic_part_mapping_pixel.color_string_id",
	DynamicPartMappingID: "dynamic_part_mapping_pixel.dynamic_part_mapping_id",
	X:                    "dynamic_part_mapping_pixel.x",
	Y:                    "dynamic_part_mapping_pixel.y",
}

// Generated where

var DynamicPartMappingPixelWhere = struct {
	ColorStringID        whereHelperint
	DynamicPartMappingID whereHelperint
	X                    whereHelperint16
	Y                    whereHelperint16
}{
	ColorStringID:        whereHelperint{field: "\"dynamic_part_mapping_pixel\".\"color_string_id\""},
	DynamicPartMappingID: whereHelperint{field: "\"dynamic_part_mapping_pixel\".\"dynamic_part_mapping_id\""},
	X:                    whereHelperint16{field: "\"dynamic_part_mapping_pixel\".\"x\""},
	Y:                    whereHelperint16{field: "\"dynamic_part_mapping_pixel\".\"y\""},
}

// DynamicPartMappingPixelRels is where relationship names are stored.
var DynamicPartMappingPixelRels = struct {
	ColorString        string
	DynamicPartMapping string
}{
	ColorString:        "ColorString",
	DynamicPartMapping: "DynamicPartMapping",
}

// dynamicPartMappingPixelR is where relationships are stored.
type dynamicPartMappingPixelR struct {
	ColorString        *ColorString        `boil:"ColorString" json:"ColorString" toml:"ColorString" yaml:"ColorString"`
	DynamicPartMapping *DynamicPartMapping `boil:"DynamicPartMapping" json:"DynamicPartMapping" toml:"DynamicPartMapping" yaml:"DynamicPartMapping"`
}

// NewStruct creates a new relationship struct
func (*dynamicPartMappingPixelR) NewStruct() *dynamicPartMappingPixelR {
	return &dynamicPartMappingPixelR{}
}

func (r *dynamicPartMappingPixelR) GetColorString() *ColorString {
	if r == nil {
		return nil
	}
	return r.ColorString
}

func (r *dynamicPartMappingPixelR) GetDynamicPartMapping() *DynamicPartMapping {
	if r == nil {
		return nil
	}
	return r.DynamicPartMapping
}

// dynamicPartMappingPixelL is where Load methods for each relationship are stored.
type dynamicPartMappingPixelL struct{}

var (
	dynamicPartMappingPixelAllColumns            = []string{"color_string_id", "dynamic_part_mapping_id", "x", "y"}
	dynamicPartMappingPixelColumnsWithoutDefault = []string{"color_string_id", "dynamic_part_mapping_id", "x", "y"}
	dynamicPartMappingPixelColumnsWithDefault    = []string{}
	dynamicPartMappingPixelPrimaryKeyColumns     = []string{"color_string_id", "dynamic_part_mapping_id"}
	dynamicPartMappingPixelGeneratedColumns      = []string{}
)

type (
	// DynamicPartMappingPixelSlice is an alias for a slice of pointers to DynamicPartMappingPixel.
	// This should almost always be used instead of []DynamicPartMappingPixel.
	DynamicPartMappingPixelSlice []*DynamicPartMappingPixel
	// DynamicPartMappingPixelHook is the signature for custom DynamicPartMappingPixel hook methods
	DynamicPartMappingPixelHook func(context.Context, boil.ContextExecutor, *DynamicPartMappingPixel) error

	dynamicPartMappingPixelQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	dynamicPartMappingPixelType                 = reflect.TypeOf(&DynamicPartMappingPixel{})
	dynamicPartMappingPixelMapping              = queries.MakeStructMapping(dynamicPartMappingPixelType)
	dynamicPartMappingPixelPrimaryKeyMapping, _ = queries.BindMapping(dynamicPartMappingPixelType, dynamicPartMappingPixelMapping, dynamicPartMappingPixelPrimaryKeyColumns)
	dynamicPartMappingPixelInsertCacheMut       sync.RWMutex
	dynamicPartMappingPixelInsertCache          = make(map[string]insertCache)
	dynamicPartMappingPixelUpdateCacheMut       sync.RWMutex
	dynamicPartMappingPixelUpdateCache          = make(map[string]updateCache)
	dynamicPartMappingPixelUpsertCacheMut       sync.RWMutex
	dynamicPartMappingPixelUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var dynamicPartMappingPixelAfterSelectHooks []DynamicPartMappingPixelHook

var dynamicPartMappingPixelBeforeInsertHooks []DynamicPartMappingPixelHook
var dynamicPartMappingPixelAfterInsertHooks []DynamicPartMappingPixelHook

var dynamicPartMappingPixelBeforeUpdateHooks []DynamicPartMappingPixelHook
var dynamicPartMappingPixelAfterUpdateHooks []DynamicPartMappingPixelHook

var dynamicPartMappingPixelBeforeDeleteHooks []DynamicPartMappingPixelHook
var dynamicPartMappingPixelAfterDeleteHooks []DynamicPartMappingPixelHook

var dynamicPartMappingPixelBeforeUpsertHooks []DynamicPartMappingPixelHook
var dynamicPartMappingPixelAfterUpsertHooks []DynamicPartMappingPixelHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *DynamicPartMappingPixel) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range dynamicPartMappingPixelAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *DynamicPartMappingPixel) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range dynamicPartMappingPixelBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *DynamicPartMappingPixel) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range dynamicPartMappingPixelAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *DynamicPartMappingPixel) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range dynamicPartMappingPixelBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *DynamicPartMappingPixel) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range dynamicPartMappingPixelAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *DynamicPartMappingPixel) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range dynamicPartMappingPixelBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *DynamicPartMappingPixel) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range dynamicPartMappingPixelAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *DynamicPartMappingPixel) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range dynamicPartMappingPixelBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *DynamicPartMappingPixel) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range dynamicPartMappingPixelAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddDynamicPartMappingPixelHook registers your hook function for all future operations.
func AddDynamicPartMappingPixelHook(hookPoint boil.HookPoint, dynamicPartMappingPixelHook DynamicPartMappingPixelHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		dynamicPartMappingPixelAfterSelectHooks = append(dynamicPartMappingPixelAfterSelectHooks, dynamicPartMappingPixelHook)
	case boil.BeforeInsertHook:
		dynamicPartMappingPixelBeforeInsertHooks = append(dynamicPartMappingPixelBeforeInsertHooks, dynamicPartMappingPixelHook)
	case boil.AfterInsertHook:
		dynamicPartMappingPixelAfterInsertHooks = append(dynamicPartMappingPixelAfterInsertHooks, dynamicPartMappingPixelHook)
	case boil.BeforeUpdateHook:
		dynamicPartMappingPixelBeforeUpdateHooks = append(dynamicPartMappingPixelBeforeUpdateHooks, dynamicPartMappingPixelHook)
	case boil.AfterUpdateHook:
		dynamicPartMappingPixelAfterUpdateHooks = append(dynamicPartMappingPixelAfterUpdateHooks, dynamicPartMappingPixelHook)
	case boil.BeforeDeleteHook:
		dynamicPartMappingPixelBeforeDeleteHooks = append(dynamicPartMappingPixelBeforeDeleteHooks, dynamicPartMappingPixelHook)
	case boil.AfterDeleteHook:
		dynamicPartMappingPixelAfterDeleteHooks = append(dynamicPartMappingPixelAfterDeleteHooks, dynamicPartMappingPixelHook)
	case boil.BeforeUpsertHook:
		dynamicPartMappingPixelBeforeUpsertHooks = append(dynamicPartMappingPixelBeforeUpsertHooks, dynamicPartMappingPixelHook)
	case boil.AfterUpsertHook:
		dynamicPartMappingPixelAfterUpsertHooks = append(dynamicPartMappingPixelAfterUpsertHooks, dynamicPartMappingPixelHook)
	}
}

// One returns a single dynamicPartMappingPixel record from the query.
func (q dynamicPartMappingPixelQuery) One(ctx context.Context, exec boil.ContextExecutor) (*DynamicPartMappingPixel, error) {
	o := &DynamicPartMappingPixel{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for dynamic_part_mapping_pixel")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all DynamicPartMappingPixel records from the query.
func (q dynamicPartMappingPixelQuery) All(ctx context.Context, exec boil.ContextExecutor) (DynamicPartMappingPixelSlice, error) {
	var o []*DynamicPartMappingPixel

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to DynamicPartMappingPixel slice")
	}

	if len(dynamicPartMappingPixelAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all DynamicPartMappingPixel records in the query.
func (q dynamicPartMappingPixelQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count dynamic_part_mapping_pixel rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q dynamicPartMappingPixelQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if dynamic_part_mapping_pixel exists")
	}

	return count > 0, nil
}

// ColorString pointed to by the foreign key.
func (o *DynamicPartMappingPixel) ColorString(mods ...qm.QueryMod) colorStringQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.ColorStringID),
	}

	queryMods = append(queryMods, mods...)

	return ColorStrings(queryMods...)
}

// DynamicPartMapping pointed to by the foreign key.
func (o *DynamicPartMappingPixel) DynamicPartMapping(mods ...qm.QueryMod) dynamicPartMappingQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.DynamicPartMappingID),
	}

	queryMods = append(queryMods, mods...)

	return DynamicPartMappings(queryMods...)
}

// LoadColorString allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (dynamicPartMappingPixelL) LoadColorString(ctx context.Context, e boil.ContextExecutor, singular bool, maybeDynamicPartMappingPixel interface{}, mods queries.Applicator) error {
	var slice []*DynamicPartMappingPixel
	var object *DynamicPartMappingPixel

	if singular {
		object = maybeDynamicPartMappingPixel.(*DynamicPartMappingPixel)
	} else {
		slice = *maybeDynamicPartMappingPixel.(*[]*DynamicPartMappingPixel)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &dynamicPartMappingPixelR{}
		}
		args = append(args, object.ColorStringID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &dynamicPartMappingPixelR{}
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

	if len(dynamicPartMappingPixelAfterSelectHooks) != 0 {
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
		foreign.R.DynamicPartMappingPixels = append(foreign.R.DynamicPartMappingPixels, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.ColorStringID == foreign.ID {
				local.R.ColorString = foreign
				if foreign.R == nil {
					foreign.R = &colorStringR{}
				}
				foreign.R.DynamicPartMappingPixels = append(foreign.R.DynamicPartMappingPixels, local)
				break
			}
		}
	}

	return nil
}

// LoadDynamicPartMapping allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (dynamicPartMappingPixelL) LoadDynamicPartMapping(ctx context.Context, e boil.ContextExecutor, singular bool, maybeDynamicPartMappingPixel interface{}, mods queries.Applicator) error {
	var slice []*DynamicPartMappingPixel
	var object *DynamicPartMappingPixel

	if singular {
		object = maybeDynamicPartMappingPixel.(*DynamicPartMappingPixel)
	} else {
		slice = *maybeDynamicPartMappingPixel.(*[]*DynamicPartMappingPixel)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &dynamicPartMappingPixelR{}
		}
		args = append(args, object.DynamicPartMappingID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &dynamicPartMappingPixelR{}
			}

			for _, a := range args {
				if a == obj.DynamicPartMappingID {
					continue Outer
				}
			}

			args = append(args, obj.DynamicPartMappingID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`dynamic_part_mapping`),
		qm.WhereIn(`dynamic_part_mapping.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load DynamicPartMapping")
	}

	var resultSlice []*DynamicPartMapping
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice DynamicPartMapping")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for dynamic_part_mapping")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for dynamic_part_mapping")
	}

	if len(dynamicPartMappingPixelAfterSelectHooks) != 0 {
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
		object.R.DynamicPartMapping = foreign
		if foreign.R == nil {
			foreign.R = &dynamicPartMappingR{}
		}
		foreign.R.DynamicPartMappingPixels = append(foreign.R.DynamicPartMappingPixels, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.DynamicPartMappingID == foreign.ID {
				local.R.DynamicPartMapping = foreign
				if foreign.R == nil {
					foreign.R = &dynamicPartMappingR{}
				}
				foreign.R.DynamicPartMappingPixels = append(foreign.R.DynamicPartMappingPixels, local)
				break
			}
		}
	}

	return nil
}

// SetColorString of the dynamicPartMappingPixel to the related item.
// Sets o.R.ColorString to related.
// Adds o to related.R.DynamicPartMappingPixels.
func (o *DynamicPartMappingPixel) SetColorString(ctx context.Context, exec boil.ContextExecutor, insert bool, related *ColorString) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"dynamic_part_mapping_pixel\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"color_string_id"}),
		strmangle.WhereClause("\"", "\"", 2, dynamicPartMappingPixelPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ColorStringID, o.DynamicPartMappingID}

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
		o.R = &dynamicPartMappingPixelR{
			ColorString: related,
		}
	} else {
		o.R.ColorString = related
	}

	if related.R == nil {
		related.R = &colorStringR{
			DynamicPartMappingPixels: DynamicPartMappingPixelSlice{o},
		}
	} else {
		related.R.DynamicPartMappingPixels = append(related.R.DynamicPartMappingPixels, o)
	}

	return nil
}

// SetDynamicPartMapping of the dynamicPartMappingPixel to the related item.
// Sets o.R.DynamicPartMapping to related.
// Adds o to related.R.DynamicPartMappingPixels.
func (o *DynamicPartMappingPixel) SetDynamicPartMapping(ctx context.Context, exec boil.ContextExecutor, insert bool, related *DynamicPartMapping) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"dynamic_part_mapping_pixel\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"dynamic_part_mapping_id"}),
		strmangle.WhereClause("\"", "\"", 2, dynamicPartMappingPixelPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ColorStringID, o.DynamicPartMappingID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.DynamicPartMappingID = related.ID
	if o.R == nil {
		o.R = &dynamicPartMappingPixelR{
			DynamicPartMapping: related,
		}
	} else {
		o.R.DynamicPartMapping = related
	}

	if related.R == nil {
		related.R = &dynamicPartMappingR{
			DynamicPartMappingPixels: DynamicPartMappingPixelSlice{o},
		}
	} else {
		related.R.DynamicPartMappingPixels = append(related.R.DynamicPartMappingPixels, o)
	}

	return nil
}

// DynamicPartMappingPixels retrieves all the records using an executor.
func DynamicPartMappingPixels(mods ...qm.QueryMod) dynamicPartMappingPixelQuery {
	mods = append(mods, qm.From("\"dynamic_part_mapping_pixel\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"dynamic_part_mapping_pixel\".*"})
	}

	return dynamicPartMappingPixelQuery{q}
}

// FindDynamicPartMappingPixel retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindDynamicPartMappingPixel(ctx context.Context, exec boil.ContextExecutor, colorStringID int, dynamicPartMappingID int, selectCols ...string) (*DynamicPartMappingPixel, error) {
	dynamicPartMappingPixelObj := &DynamicPartMappingPixel{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"dynamic_part_mapping_pixel\" where \"color_string_id\"=$1 AND \"dynamic_part_mapping_id\"=$2", sel,
	)

	q := queries.Raw(query, colorStringID, dynamicPartMappingID)

	err := q.Bind(ctx, exec, dynamicPartMappingPixelObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from dynamic_part_mapping_pixel")
	}

	if err = dynamicPartMappingPixelObj.doAfterSelectHooks(ctx, exec); err != nil {
		return dynamicPartMappingPixelObj, err
	}

	return dynamicPartMappingPixelObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *DynamicPartMappingPixel) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no dynamic_part_mapping_pixel provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(dynamicPartMappingPixelColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	dynamicPartMappingPixelInsertCacheMut.RLock()
	cache, cached := dynamicPartMappingPixelInsertCache[key]
	dynamicPartMappingPixelInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			dynamicPartMappingPixelAllColumns,
			dynamicPartMappingPixelColumnsWithDefault,
			dynamicPartMappingPixelColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(dynamicPartMappingPixelType, dynamicPartMappingPixelMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(dynamicPartMappingPixelType, dynamicPartMappingPixelMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"dynamic_part_mapping_pixel\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"dynamic_part_mapping_pixel\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into dynamic_part_mapping_pixel")
	}

	if !cached {
		dynamicPartMappingPixelInsertCacheMut.Lock()
		dynamicPartMappingPixelInsertCache[key] = cache
		dynamicPartMappingPixelInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the DynamicPartMappingPixel.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *DynamicPartMappingPixel) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	dynamicPartMappingPixelUpdateCacheMut.RLock()
	cache, cached := dynamicPartMappingPixelUpdateCache[key]
	dynamicPartMappingPixelUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			dynamicPartMappingPixelAllColumns,
			dynamicPartMappingPixelPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update dynamic_part_mapping_pixel, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"dynamic_part_mapping_pixel\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, dynamicPartMappingPixelPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(dynamicPartMappingPixelType, dynamicPartMappingPixelMapping, append(wl, dynamicPartMappingPixelPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update dynamic_part_mapping_pixel row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for dynamic_part_mapping_pixel")
	}

	if !cached {
		dynamicPartMappingPixelUpdateCacheMut.Lock()
		dynamicPartMappingPixelUpdateCache[key] = cache
		dynamicPartMappingPixelUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q dynamicPartMappingPixelQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for dynamic_part_mapping_pixel")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for dynamic_part_mapping_pixel")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o DynamicPartMappingPixelSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), dynamicPartMappingPixelPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"dynamic_part_mapping_pixel\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, dynamicPartMappingPixelPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in dynamicPartMappingPixel slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all dynamicPartMappingPixel")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *DynamicPartMappingPixel) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no dynamic_part_mapping_pixel provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(dynamicPartMappingPixelColumnsWithDefault, o)

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

	dynamicPartMappingPixelUpsertCacheMut.RLock()
	cache, cached := dynamicPartMappingPixelUpsertCache[key]
	dynamicPartMappingPixelUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			dynamicPartMappingPixelAllColumns,
			dynamicPartMappingPixelColumnsWithDefault,
			dynamicPartMappingPixelColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			dynamicPartMappingPixelAllColumns,
			dynamicPartMappingPixelPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert dynamic_part_mapping_pixel, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(dynamicPartMappingPixelPrimaryKeyColumns))
			copy(conflict, dynamicPartMappingPixelPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"dynamic_part_mapping_pixel\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(dynamicPartMappingPixelType, dynamicPartMappingPixelMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(dynamicPartMappingPixelType, dynamicPartMappingPixelMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert dynamic_part_mapping_pixel")
	}

	if !cached {
		dynamicPartMappingPixelUpsertCacheMut.Lock()
		dynamicPartMappingPixelUpsertCache[key] = cache
		dynamicPartMappingPixelUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single DynamicPartMappingPixel record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *DynamicPartMappingPixel) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no DynamicPartMappingPixel provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), dynamicPartMappingPixelPrimaryKeyMapping)
	sql := "DELETE FROM \"dynamic_part_mapping_pixel\" WHERE \"color_string_id\"=$1 AND \"dynamic_part_mapping_id\"=$2"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from dynamic_part_mapping_pixel")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for dynamic_part_mapping_pixel")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q dynamicPartMappingPixelQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no dynamicPartMappingPixelQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from dynamic_part_mapping_pixel")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for dynamic_part_mapping_pixel")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o DynamicPartMappingPixelSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(dynamicPartMappingPixelBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), dynamicPartMappingPixelPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"dynamic_part_mapping_pixel\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, dynamicPartMappingPixelPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from dynamicPartMappingPixel slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for dynamic_part_mapping_pixel")
	}

	if len(dynamicPartMappingPixelAfterDeleteHooks) != 0 {
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
func (o *DynamicPartMappingPixel) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindDynamicPartMappingPixel(ctx, exec, o.ColorStringID, o.DynamicPartMappingID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *DynamicPartMappingPixelSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := DynamicPartMappingPixelSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), dynamicPartMappingPixelPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"dynamic_part_mapping_pixel\".* FROM \"dynamic_part_mapping_pixel\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, dynamicPartMappingPixelPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in DynamicPartMappingPixelSlice")
	}

	*o = slice

	return nil
}

// DynamicPartMappingPixelExists checks if the DynamicPartMappingPixel row exists.
func DynamicPartMappingPixelExists(ctx context.Context, exec boil.ContextExecutor, colorStringID int, dynamicPartMappingID int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"dynamic_part_mapping_pixel\" where \"color_string_id\"=$1 AND \"dynamic_part_mapping_id\"=$2 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, colorStringID, dynamicPartMappingID)
	}
	row := exec.QueryRowContext(ctx, sql, colorStringID, dynamicPartMappingID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if dynamic_part_mapping_pixel exists")
	}

	return exists, nil
}