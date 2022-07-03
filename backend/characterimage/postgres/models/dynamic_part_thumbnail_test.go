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

func testDynamicPartThumbnails(t *testing.T) {
	t.Parallel()

	query := DynamicPartThumbnails()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testDynamicPartThumbnailsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DynamicPartThumbnail{}
	if err = randomize.Struct(seed, o, dynamicPartThumbnailDBTypes, true, dynamicPartThumbnailColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DynamicPartThumbnail struct: %s", err)
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

	count, err := DynamicPartThumbnails().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDynamicPartThumbnailsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DynamicPartThumbnail{}
	if err = randomize.Struct(seed, o, dynamicPartThumbnailDBTypes, true, dynamicPartThumbnailColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DynamicPartThumbnail struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := DynamicPartThumbnails().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := DynamicPartThumbnails().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDynamicPartThumbnailsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DynamicPartThumbnail{}
	if err = randomize.Struct(seed, o, dynamicPartThumbnailDBTypes, true, dynamicPartThumbnailColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DynamicPartThumbnail struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := DynamicPartThumbnailSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := DynamicPartThumbnails().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDynamicPartThumbnailsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DynamicPartThumbnail{}
	if err = randomize.Struct(seed, o, dynamicPartThumbnailDBTypes, true, dynamicPartThumbnailColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DynamicPartThumbnail struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := DynamicPartThumbnailExists(ctx, tx, o.DynamicPartID)
	if err != nil {
		t.Errorf("Unable to check if DynamicPartThumbnail exists: %s", err)
	}
	if !e {
		t.Errorf("Expected DynamicPartThumbnailExists to return true, but got false.")
	}
}

func testDynamicPartThumbnailsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DynamicPartThumbnail{}
	if err = randomize.Struct(seed, o, dynamicPartThumbnailDBTypes, true, dynamicPartThumbnailColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DynamicPartThumbnail struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	dynamicPartThumbnailFound, err := FindDynamicPartThumbnail(ctx, tx, o.DynamicPartID)
	if err != nil {
		t.Error(err)
	}

	if dynamicPartThumbnailFound == nil {
		t.Error("want a record, got nil")
	}
}

func testDynamicPartThumbnailsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DynamicPartThumbnail{}
	if err = randomize.Struct(seed, o, dynamicPartThumbnailDBTypes, true, dynamicPartThumbnailColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DynamicPartThumbnail struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = DynamicPartThumbnails().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testDynamicPartThumbnailsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DynamicPartThumbnail{}
	if err = randomize.Struct(seed, o, dynamicPartThumbnailDBTypes, true, dynamicPartThumbnailColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DynamicPartThumbnail struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := DynamicPartThumbnails().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testDynamicPartThumbnailsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dynamicPartThumbnailOne := &DynamicPartThumbnail{}
	dynamicPartThumbnailTwo := &DynamicPartThumbnail{}
	if err = randomize.Struct(seed, dynamicPartThumbnailOne, dynamicPartThumbnailDBTypes, false, dynamicPartThumbnailColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DynamicPartThumbnail struct: %s", err)
	}
	if err = randomize.Struct(seed, dynamicPartThumbnailTwo, dynamicPartThumbnailDBTypes, false, dynamicPartThumbnailColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DynamicPartThumbnail struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = dynamicPartThumbnailOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = dynamicPartThumbnailTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := DynamicPartThumbnails().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testDynamicPartThumbnailsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	dynamicPartThumbnailOne := &DynamicPartThumbnail{}
	dynamicPartThumbnailTwo := &DynamicPartThumbnail{}
	if err = randomize.Struct(seed, dynamicPartThumbnailOne, dynamicPartThumbnailDBTypes, false, dynamicPartThumbnailColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DynamicPartThumbnail struct: %s", err)
	}
	if err = randomize.Struct(seed, dynamicPartThumbnailTwo, dynamicPartThumbnailDBTypes, false, dynamicPartThumbnailColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DynamicPartThumbnail struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = dynamicPartThumbnailOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = dynamicPartThumbnailTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := DynamicPartThumbnails().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func dynamicPartThumbnailBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *DynamicPartThumbnail) error {
	*o = DynamicPartThumbnail{}
	return nil
}

func dynamicPartThumbnailAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *DynamicPartThumbnail) error {
	*o = DynamicPartThumbnail{}
	return nil
}

func dynamicPartThumbnailAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *DynamicPartThumbnail) error {
	*o = DynamicPartThumbnail{}
	return nil
}

func dynamicPartThumbnailBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *DynamicPartThumbnail) error {
	*o = DynamicPartThumbnail{}
	return nil
}

func dynamicPartThumbnailAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *DynamicPartThumbnail) error {
	*o = DynamicPartThumbnail{}
	return nil
}

func dynamicPartThumbnailBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *DynamicPartThumbnail) error {
	*o = DynamicPartThumbnail{}
	return nil
}

func dynamicPartThumbnailAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *DynamicPartThumbnail) error {
	*o = DynamicPartThumbnail{}
	return nil
}

func dynamicPartThumbnailBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *DynamicPartThumbnail) error {
	*o = DynamicPartThumbnail{}
	return nil
}

func dynamicPartThumbnailAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *DynamicPartThumbnail) error {
	*o = DynamicPartThumbnail{}
	return nil
}

func testDynamicPartThumbnailsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &DynamicPartThumbnail{}
	o := &DynamicPartThumbnail{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, dynamicPartThumbnailDBTypes, false); err != nil {
		t.Errorf("Unable to randomize DynamicPartThumbnail object: %s", err)
	}

	AddDynamicPartThumbnailHook(boil.BeforeInsertHook, dynamicPartThumbnailBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	dynamicPartThumbnailBeforeInsertHooks = []DynamicPartThumbnailHook{}

	AddDynamicPartThumbnailHook(boil.AfterInsertHook, dynamicPartThumbnailAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	dynamicPartThumbnailAfterInsertHooks = []DynamicPartThumbnailHook{}

	AddDynamicPartThumbnailHook(boil.AfterSelectHook, dynamicPartThumbnailAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	dynamicPartThumbnailAfterSelectHooks = []DynamicPartThumbnailHook{}

	AddDynamicPartThumbnailHook(boil.BeforeUpdateHook, dynamicPartThumbnailBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	dynamicPartThumbnailBeforeUpdateHooks = []DynamicPartThumbnailHook{}

	AddDynamicPartThumbnailHook(boil.AfterUpdateHook, dynamicPartThumbnailAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	dynamicPartThumbnailAfterUpdateHooks = []DynamicPartThumbnailHook{}

	AddDynamicPartThumbnailHook(boil.BeforeDeleteHook, dynamicPartThumbnailBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	dynamicPartThumbnailBeforeDeleteHooks = []DynamicPartThumbnailHook{}

	AddDynamicPartThumbnailHook(boil.AfterDeleteHook, dynamicPartThumbnailAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	dynamicPartThumbnailAfterDeleteHooks = []DynamicPartThumbnailHook{}

	AddDynamicPartThumbnailHook(boil.BeforeUpsertHook, dynamicPartThumbnailBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	dynamicPartThumbnailBeforeUpsertHooks = []DynamicPartThumbnailHook{}

	AddDynamicPartThumbnailHook(boil.AfterUpsertHook, dynamicPartThumbnailAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	dynamicPartThumbnailAfterUpsertHooks = []DynamicPartThumbnailHook{}
}

func testDynamicPartThumbnailsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DynamicPartThumbnail{}
	if err = randomize.Struct(seed, o, dynamicPartThumbnailDBTypes, true, dynamicPartThumbnailColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DynamicPartThumbnail struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := DynamicPartThumbnails().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testDynamicPartThumbnailsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DynamicPartThumbnail{}
	if err = randomize.Struct(seed, o, dynamicPartThumbnailDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DynamicPartThumbnail struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(dynamicPartThumbnailColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := DynamicPartThumbnails().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testDynamicPartThumbnailToOneDynamicPartUsingDynamicPart(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local DynamicPartThumbnail
	var foreign DynamicPart

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, dynamicPartThumbnailDBTypes, false, dynamicPartThumbnailColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DynamicPartThumbnail struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, dynamicPartDBTypes, false, dynamicPartColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DynamicPart struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.DynamicPartID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.DynamicPart().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := DynamicPartThumbnailSlice{&local}
	if err = local.L.LoadDynamicPart(ctx, tx, false, (*[]*DynamicPartThumbnail)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.DynamicPart == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.DynamicPart = nil
	if err = local.L.LoadDynamicPart(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.DynamicPart == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testDynamicPartThumbnailToOneSetOpDynamicPartUsingDynamicPart(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a DynamicPartThumbnail
	var b, c DynamicPart

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, dynamicPartThumbnailDBTypes, false, strmangle.SetComplement(dynamicPartThumbnailPrimaryKeyColumns, dynamicPartThumbnailColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, dynamicPartDBTypes, false, strmangle.SetComplement(dynamicPartPrimaryKeyColumns, dynamicPartColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, dynamicPartDBTypes, false, strmangle.SetComplement(dynamicPartPrimaryKeyColumns, dynamicPartColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*DynamicPart{&b, &c} {
		err = a.SetDynamicPart(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.DynamicPart != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.DynamicPartThumbnail != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.DynamicPartID != x.ID {
			t.Error("foreign key was wrong value", a.DynamicPartID)
		}

		if exists, err := DynamicPartThumbnailExists(ctx, tx, a.DynamicPartID); err != nil {
			t.Fatal(err)
		} else if !exists {
			t.Error("want 'a' to exist")
		}

	}
}

func testDynamicPartThumbnailsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DynamicPartThumbnail{}
	if err = randomize.Struct(seed, o, dynamicPartThumbnailDBTypes, true, dynamicPartThumbnailColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DynamicPartThumbnail struct: %s", err)
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

func testDynamicPartThumbnailsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DynamicPartThumbnail{}
	if err = randomize.Struct(seed, o, dynamicPartThumbnailDBTypes, true, dynamicPartThumbnailColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DynamicPartThumbnail struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := DynamicPartThumbnailSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testDynamicPartThumbnailsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DynamicPartThumbnail{}
	if err = randomize.Struct(seed, o, dynamicPartThumbnailDBTypes, true, dynamicPartThumbnailColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DynamicPartThumbnail struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := DynamicPartThumbnails().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	dynamicPartThumbnailDBTypes = map[string]string{`DynamicPartID`: `integer`, `ImageBytes`: `bytea`}
	_                           = bytes.MinRead
)

func testDynamicPartThumbnailsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(dynamicPartThumbnailPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(dynamicPartThumbnailAllColumns) == len(dynamicPartThumbnailPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &DynamicPartThumbnail{}
	if err = randomize.Struct(seed, o, dynamicPartThumbnailDBTypes, true, dynamicPartThumbnailColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DynamicPartThumbnail struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := DynamicPartThumbnails().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, dynamicPartThumbnailDBTypes, true, dynamicPartThumbnailPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DynamicPartThumbnail struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testDynamicPartThumbnailsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(dynamicPartThumbnailAllColumns) == len(dynamicPartThumbnailPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &DynamicPartThumbnail{}
	if err = randomize.Struct(seed, o, dynamicPartThumbnailDBTypes, true, dynamicPartThumbnailColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DynamicPartThumbnail struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := DynamicPartThumbnails().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, dynamicPartThumbnailDBTypes, true, dynamicPartThumbnailPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DynamicPartThumbnail struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(dynamicPartThumbnailAllColumns, dynamicPartThumbnailPrimaryKeyColumns) {
		fields = dynamicPartThumbnailAllColumns
	} else {
		fields = strmangle.SetComplement(
			dynamicPartThumbnailAllColumns,
			dynamicPartThumbnailPrimaryKeyColumns,
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

	slice := DynamicPartThumbnailSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testDynamicPartThumbnailsUpsert(t *testing.T) {
	t.Parallel()

	if len(dynamicPartThumbnailAllColumns) == len(dynamicPartThumbnailPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := DynamicPartThumbnail{}
	if err = randomize.Struct(seed, &o, dynamicPartThumbnailDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DynamicPartThumbnail struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert DynamicPartThumbnail: %s", err)
	}

	count, err := DynamicPartThumbnails().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, dynamicPartThumbnailDBTypes, false, dynamicPartThumbnailPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DynamicPartThumbnail struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert DynamicPartThumbnail: %s", err)
	}

	count, err = DynamicPartThumbnails().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
