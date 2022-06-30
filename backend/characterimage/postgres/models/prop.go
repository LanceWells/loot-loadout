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

// Prop is an object representing the database table.
type Prop struct {
	ID          int          `boil:"id" json:"id" toml:"id" yaml:"id"`
	DisplayName string       `boil:"display_name" json:"display_name" toml:"display_name" yaml:"display_name"`
	PartType    NullPropType `boil:"part_type" json:"part_type,omitempty" toml:"part_type" yaml:"part_type,omitempty"`

	R *propR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L propL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var PropColumns = struct {
	ID          string
	DisplayName string
	PartType    string
}{
	ID:          "id",
	DisplayName: "display_name",
	PartType:    "part_type",
}

var PropTableColumns = struct {
	ID          string
	DisplayName string
	PartType    string
}{
	ID:          "prop.id",
	DisplayName: "prop.display_name",
	PartType:    "prop.part_type",
}

// Generated where

type whereHelperNullPropType struct{ field string }

func (w whereHelperNullPropType) EQ(x NullPropType) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelperNullPropType) NEQ(x NullPropType) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelperNullPropType) LT(x NullPropType) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelperNullPropType) LTE(x NullPropType) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelperNullPropType) GT(x NullPropType) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelperNullPropType) GTE(x NullPropType) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

func (w whereHelperNullPropType) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelperNullPropType) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }

var PropWhere = struct {
	ID          whereHelperint
	DisplayName whereHelperstring
	PartType    whereHelperNullPropType
}{
	ID:          whereHelperint{field: "\"prop\".\"id\""},
	DisplayName: whereHelperstring{field: "\"prop\".\"display_name\""},
	PartType:    whereHelperNullPropType{field: "\"prop\".\"part_type\""},
}

// PropRels is where relationship names are stored.
var PropRels = struct {
	PropImage string
}{
	PropImage: "PropImage",
}

// propR is where relationships are stored.
type propR struct {
	PropImage *PropImage `boil:"PropImage" json:"PropImage" toml:"PropImage" yaml:"PropImage"`
}

// NewStruct creates a new relationship struct
func (*propR) NewStruct() *propR {
	return &propR{}
}

func (r *propR) GetPropImage() *PropImage {
	if r == nil {
		return nil
	}
	return r.PropImage
}

// propL is where Load methods for each relationship are stored.
type propL struct{}

var (
	propAllColumns            = []string{"id", "display_name", "part_type"}
	propColumnsWithoutDefault = []string{"display_name"}
	propColumnsWithDefault    = []string{"id", "part_type"}
	propPrimaryKeyColumns     = []string{"id"}
	propGeneratedColumns      = []string{}
)

type (
	// PropSlice is an alias for a slice of pointers to Prop.
	// This should almost always be used instead of []Prop.
	PropSlice []*Prop
	// PropHook is the signature for custom Prop hook methods
	PropHook func(context.Context, boil.ContextExecutor, *Prop) error

	propQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	propType                 = reflect.TypeOf(&Prop{})
	propMapping              = queries.MakeStructMapping(propType)
	propPrimaryKeyMapping, _ = queries.BindMapping(propType, propMapping, propPrimaryKeyColumns)
	propInsertCacheMut       sync.RWMutex
	propInsertCache          = make(map[string]insertCache)
	propUpdateCacheMut       sync.RWMutex
	propUpdateCache          = make(map[string]updateCache)
	propUpsertCacheMut       sync.RWMutex
	propUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var propAfterSelectHooks []PropHook

var propBeforeInsertHooks []PropHook
var propAfterInsertHooks []PropHook

var propBeforeUpdateHooks []PropHook
var propAfterUpdateHooks []PropHook

var propBeforeDeleteHooks []PropHook
var propAfterDeleteHooks []PropHook

var propBeforeUpsertHooks []PropHook
var propAfterUpsertHooks []PropHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Prop) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range propAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Prop) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range propBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Prop) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range propAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Prop) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range propBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Prop) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range propAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Prop) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range propBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Prop) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range propAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Prop) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range propBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Prop) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range propAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddPropHook registers your hook function for all future operations.
func AddPropHook(hookPoint boil.HookPoint, propHook PropHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		propAfterSelectHooks = append(propAfterSelectHooks, propHook)
	case boil.BeforeInsertHook:
		propBeforeInsertHooks = append(propBeforeInsertHooks, propHook)
	case boil.AfterInsertHook:
		propAfterInsertHooks = append(propAfterInsertHooks, propHook)
	case boil.BeforeUpdateHook:
		propBeforeUpdateHooks = append(propBeforeUpdateHooks, propHook)
	case boil.AfterUpdateHook:
		propAfterUpdateHooks = append(propAfterUpdateHooks, propHook)
	case boil.BeforeDeleteHook:
		propBeforeDeleteHooks = append(propBeforeDeleteHooks, propHook)
	case boil.AfterDeleteHook:
		propAfterDeleteHooks = append(propAfterDeleteHooks, propHook)
	case boil.BeforeUpsertHook:
		propBeforeUpsertHooks = append(propBeforeUpsertHooks, propHook)
	case boil.AfterUpsertHook:
		propAfterUpsertHooks = append(propAfterUpsertHooks, propHook)
	}
}

// One returns a single prop record from the query.
func (q propQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Prop, error) {
	o := &Prop{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for prop")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Prop records from the query.
func (q propQuery) All(ctx context.Context, exec boil.ContextExecutor) (PropSlice, error) {
	var o []*Prop

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Prop slice")
	}

	if len(propAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Prop records in the query.
func (q propQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count prop rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q propQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if prop exists")
	}

	return count > 0, nil
}

// PropImage pointed to by the foreign key.
func (o *Prop) PropImage(mods ...qm.QueryMod) propImageQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"prop_id\" = ?", o.ID),
	}

	queryMods = append(queryMods, mods...)

	return PropImages(queryMods...)
}

// LoadPropImage allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-1 relationship.
func (propL) LoadPropImage(ctx context.Context, e boil.ContextExecutor, singular bool, maybeProp interface{}, mods queries.Applicator) error {
	var slice []*Prop
	var object *Prop

	if singular {
		object = maybeProp.(*Prop)
	} else {
		slice = *maybeProp.(*[]*Prop)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &propR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &propR{}
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
		qm.From(`prop_image`),
		qm.WhereIn(`prop_image.prop_id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load PropImage")
	}

	var resultSlice []*PropImage
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice PropImage")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for prop_image")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for prop_image")
	}

	if len(propAfterSelectHooks) != 0 {
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
		object.R.PropImage = foreign
		if foreign.R == nil {
			foreign.R = &propImageR{}
		}
		foreign.R.Prop = object
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.ID == foreign.PropID {
				local.R.PropImage = foreign
				if foreign.R == nil {
					foreign.R = &propImageR{}
				}
				foreign.R.Prop = local
				break
			}
		}
	}

	return nil
}

// SetPropImage of the prop to the related item.
// Sets o.R.PropImage to related.
// Adds o to related.R.Prop.
func (o *Prop) SetPropImage(ctx context.Context, exec boil.ContextExecutor, insert bool, related *PropImage) error {
	var err error

	if insert {
		related.PropID = o.ID

		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"prop_image\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"prop_id"}),
			strmangle.WhereClause("\"", "\"", 2, propImagePrimaryKeyColumns),
		)
		values := []interface{}{o.ID, related.PropID}

		if boil.IsDebug(ctx) {
			writer := boil.DebugWriterFrom(ctx)
			fmt.Fprintln(writer, updateQuery)
			fmt.Fprintln(writer, values)
		}
		if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.PropID = o.ID
	}

	if o.R == nil {
		o.R = &propR{
			PropImage: related,
		}
	} else {
		o.R.PropImage = related
	}

	if related.R == nil {
		related.R = &propImageR{
			Prop: o,
		}
	} else {
		related.R.Prop = o
	}
	return nil
}

// Props retrieves all the records using an executor.
func Props(mods ...qm.QueryMod) propQuery {
	mods = append(mods, qm.From("\"prop\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"prop\".*"})
	}

	return propQuery{q}
}

// FindProp retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindProp(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*Prop, error) {
	propObj := &Prop{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"prop\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, propObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from prop")
	}

	if err = propObj.doAfterSelectHooks(ctx, exec); err != nil {
		return propObj, err
	}

	return propObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Prop) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no prop provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(propColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	propInsertCacheMut.RLock()
	cache, cached := propInsertCache[key]
	propInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			propAllColumns,
			propColumnsWithDefault,
			propColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(propType, propMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(propType, propMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"prop\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"prop\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into prop")
	}

	if !cached {
		propInsertCacheMut.Lock()
		propInsertCache[key] = cache
		propInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Prop.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Prop) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	propUpdateCacheMut.RLock()
	cache, cached := propUpdateCache[key]
	propUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			propAllColumns,
			propPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update prop, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"prop\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, propPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(propType, propMapping, append(wl, propPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update prop row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for prop")
	}

	if !cached {
		propUpdateCacheMut.Lock()
		propUpdateCache[key] = cache
		propUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q propQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for prop")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for prop")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o PropSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), propPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"prop\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, propPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in prop slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all prop")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Prop) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no prop provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(propColumnsWithDefault, o)

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

	propUpsertCacheMut.RLock()
	cache, cached := propUpsertCache[key]
	propUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			propAllColumns,
			propColumnsWithDefault,
			propColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			propAllColumns,
			propPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert prop, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(propPrimaryKeyColumns))
			copy(conflict, propPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"prop\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(propType, propMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(propType, propMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert prop")
	}

	if !cached {
		propUpsertCacheMut.Lock()
		propUpsertCache[key] = cache
		propUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Prop record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Prop) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Prop provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), propPrimaryKeyMapping)
	sql := "DELETE FROM \"prop\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from prop")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for prop")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q propQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no propQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from prop")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for prop")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o PropSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(propBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), propPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"prop\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, propPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from prop slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for prop")
	}

	if len(propAfterDeleteHooks) != 0 {
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
func (o *Prop) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindProp(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *PropSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := PropSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), propPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"prop\".* FROM \"prop\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, propPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in PropSlice")
	}

	*o = slice

	return nil
}

// PropExists checks if the Prop row exists.
func PropExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"prop\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if prop exists")
	}

	return exists, nil
}
