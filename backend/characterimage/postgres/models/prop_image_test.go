// Code generated by SQLBoiler 4.11.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/randomize"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testPropImages(t *testing.T) {
	t.Parallel()

	query := PropImages()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testPropImagesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PropImage{}
	if err = randomize.Struct(seed, o, propImageDBTypes, true, propImageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PropImage struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := PropImages().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPropImagesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PropImage{}
	if err = randomize.Struct(seed, o, propImageDBTypes, true, propImageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PropImage struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := PropImages().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := PropImages().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPropImagesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PropImage{}
	if err = randomize.Struct(seed, o, propImageDBTypes, true, propImageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PropImage struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := PropImageSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := PropImages().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPropImagesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PropImage{}
	if err = randomize.Struct(seed, o, propImageDBTypes, true, propImageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PropImage struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := PropImageExists(ctx, tx, o.PropID)
	if err != nil {
		t.Errorf("Unable to check if PropImage exists: %s", err)
	}
	if !e {
		t.Errorf("Expected PropImageExists to return true, but got false.")
	}
}

func testPropImagesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PropImage{}
	if err = randomize.Struct(seed, o, propImageDBTypes, true, propImageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PropImage struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	propImageFound, err := FindPropImage(ctx, tx, o.PropID)
	if err != nil {
		t.Error(err)
	}

	if propImageFound == nil {
		t.Error("want a record, got nil")
	}
}

func testPropImagesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PropImage{}
	if err = randomize.Struct(seed, o, propImageDBTypes, true, propImageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PropImage struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = PropImages().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testPropImagesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PropImage{}
	if err = randomize.Struct(seed, o, propImageDBTypes, true, propImageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PropImage struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := PropImages().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testPropImagesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	propImageOne := &PropImage{}
	propImageTwo := &PropImage{}
	if err = randomize.Struct(seed, propImageOne, propImageDBTypes, false, propImageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PropImage struct: %s", err)
	}
	if err = randomize.Struct(seed, propImageTwo, propImageDBTypes, false, propImageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PropImage struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = propImageOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = propImageTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := PropImages().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testPropImagesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	propImageOne := &PropImage{}
	propImageTwo := &PropImage{}
	if err = randomize.Struct(seed, propImageOne, propImageDBTypes, false, propImageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PropImage struct: %s", err)
	}
	if err = randomize.Struct(seed, propImageTwo, propImageDBTypes, false, propImageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PropImage struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = propImageOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = propImageTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := PropImages().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func propImageBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *PropImage) error {
	*o = PropImage{}
	return nil
}

func propImageAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *PropImage) error {
	*o = PropImage{}
	return nil
}

func propImageAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *PropImage) error {
	*o = PropImage{}
	return nil
}

func propImageBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *PropImage) error {
	*o = PropImage{}
	return nil
}

func propImageAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *PropImage) error {
	*o = PropImage{}
	return nil
}

func propImageBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *PropImage) error {
	*o = PropImage{}
	return nil
}

func propImageAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *PropImage) error {
	*o = PropImage{}
	return nil
}

func propImageBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *PropImage) error {
	*o = PropImage{}
	return nil
}

func propImageAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *PropImage) error {
	*o = PropImage{}
	return nil
}

func testPropImagesHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &PropImage{}
	o := &PropImage{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, propImageDBTypes, false); err != nil {
		t.Errorf("Unable to randomize PropImage object: %s", err)
	}

	AddPropImageHook(boil.BeforeInsertHook, propImageBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	propImageBeforeInsertHooks = []PropImageHook{}

	AddPropImageHook(boil.AfterInsertHook, propImageAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	propImageAfterInsertHooks = []PropImageHook{}

	AddPropImageHook(boil.AfterSelectHook, propImageAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	propImageAfterSelectHooks = []PropImageHook{}

	AddPropImageHook(boil.BeforeUpdateHook, propImageBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	propImageBeforeUpdateHooks = []PropImageHook{}

	AddPropImageHook(boil.AfterUpdateHook, propImageAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	propImageAfterUpdateHooks = []PropImageHook{}

	AddPropImageHook(boil.BeforeDeleteHook, propImageBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	propImageBeforeDeleteHooks = []PropImageHook{}

	AddPropImageHook(boil.AfterDeleteHook, propImageAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	propImageAfterDeleteHooks = []PropImageHook{}

	AddPropImageHook(boil.BeforeUpsertHook, propImageBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	propImageBeforeUpsertHooks = []PropImageHook{}

	AddPropImageHook(boil.AfterUpsertHook, propImageAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	propImageAfterUpsertHooks = []PropImageHook{}
}

func testPropImagesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PropImage{}
	if err = randomize.Struct(seed, o, propImageDBTypes, true, propImageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PropImage struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := PropImages().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testPropImagesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PropImage{}
	if err = randomize.Struct(seed, o, propImageDBTypes, true); err != nil {
		t.Errorf("Unable to randomize PropImage struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(propImageColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := PropImages().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testPropImageToOnePropUsingProp(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local PropImage
	var foreign Prop

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, propImageDBTypes, false, propImageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PropImage struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, propDBTypes, false, propColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Prop struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.PropID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Prop().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := PropImageSlice{&local}
	if err = local.L.LoadProp(ctx, tx, false, (*[]*PropImage)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Prop == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Prop = nil
	if err = local.L.LoadProp(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Prop == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testPropImageToOneSetOpPropUsingProp(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a PropImage
	var b, c Prop

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, propImageDBTypes, false, strmangle.SetComplement(propImagePrimaryKeyColumns, propImageColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, propDBTypes, false, strmangle.SetComplement(propPrimaryKeyColumns, propColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, propDBTypes, false, strmangle.SetComplement(propPrimaryKeyColumns, propColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Prop{&b, &c} {
		err = a.SetProp(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Prop != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.PropImage != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.PropID != x.ID {
			t.Error("foreign key was wrong value", a.PropID)
		}

		if exists, err := PropImageExists(ctx, tx, a.PropID); err != nil {
			t.Fatal(err)
		} else if !exists {
			t.Error("want 'a' to exist")
		}

	}
}

func testPropImagesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PropImage{}
	if err = randomize.Struct(seed, o, propImageDBTypes, true, propImageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PropImage struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testPropImagesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PropImage{}
	if err = randomize.Struct(seed, o, propImageDBTypes, true, propImageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PropImage struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := PropImageSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testPropImagesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PropImage{}
	if err = randomize.Struct(seed, o, propImageDBTypes, true, propImageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PropImage struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := PropImages().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	propImageDBTypes = map[string]string{`PropID`: `integer`, `X`: `smallint`, `Y`: `smallint`, `ImageBytes`: `bytea`}
	_                = bytes.MinRead
)

func testPropImagesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(propImagePrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(propImageAllColumns) == len(propImagePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &PropImage{}
	if err = randomize.Struct(seed, o, propImageDBTypes, true, propImageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PropImage struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := PropImages().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, propImageDBTypes, true, propImagePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize PropImage struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testPropImagesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(propImageAllColumns) == len(propImagePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &PropImage{}
	if err = randomize.Struct(seed, o, propImageDBTypes, true, propImageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PropImage struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := PropImages().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, propImageDBTypes, true, propImagePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize PropImage struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(propImageAllColumns, propImagePrimaryKeyColumns) {
		fields = propImageAllColumns
	} else {
		fields = strmangle.SetComplement(
			propImageAllColumns,
			propImagePrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := PropImageSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testPropImagesUpsert(t *testing.T) {
	t.Parallel()

	if len(propImageAllColumns) == len(propImagePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := PropImage{}
	if err = randomize.Struct(seed, &o, propImageDBTypes, true); err != nil {
		t.Errorf("Unable to randomize PropImage struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert PropImage: %s", err)
	}

	count, err := PropImages().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, propImageDBTypes, false, propImagePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize PropImage struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert PropImage: %s", err)
	}

	count, err = PropImages().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}