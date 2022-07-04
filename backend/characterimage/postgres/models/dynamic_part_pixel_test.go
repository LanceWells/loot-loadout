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

func testDynamicPartPixels(t *testing.T) {
	t.Parallel()

	query := DynamicPartPixels()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testDynamicPartPixelsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DynamicPartPixel{}
	if err = randomize.Struct(seed, o, dynamicPartPixelDBTypes, true, dynamicPartPixelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DynamicPartPixel struct: %s", err)
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

	count, err := DynamicPartPixels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDynamicPartPixelsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DynamicPartPixel{}
	if err = randomize.Struct(seed, o, dynamicPartPixelDBTypes, true, dynamicPartPixelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DynamicPartPixel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := DynamicPartPixels().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := DynamicPartPixels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDynamicPartPixelsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DynamicPartPixel{}
	if err = randomize.Struct(seed, o, dynamicPartPixelDBTypes, true, dynamicPartPixelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DynamicPartPixel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := DynamicPartPixelSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := DynamicPartPixels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDynamicPartPixelsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DynamicPartPixel{}
	if err = randomize.Struct(seed, o, dynamicPartPixelDBTypes, true, dynamicPartPixelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DynamicPartPixel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := DynamicPartPixelExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if DynamicPartPixel exists: %s", err)
	}
	if !e {
		t.Errorf("Expected DynamicPartPixelExists to return true, but got false.")
	}
}

func testDynamicPartPixelsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DynamicPartPixel{}
	if err = randomize.Struct(seed, o, dynamicPartPixelDBTypes, true, dynamicPartPixelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DynamicPartPixel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	dynamicPartPixelFound, err := FindDynamicPartPixel(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if dynamicPartPixelFound == nil {
		t.Error("want a record, got nil")
	}
}

func testDynamicPartPixelsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DynamicPartPixel{}
	if err = randomize.Struct(seed, o, dynamicPartPixelDBTypes, true, dynamicPartPixelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DynamicPartPixel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = DynamicPartPixels().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testDynamicPartPixelsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DynamicPartPixel{}
	if err = randomize.Struct(seed, o, dynamicPartPixelDBTypes, true, dynamicPartPixelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DynamicPartPixel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := DynamicPartPixels().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testDynamicPartPixelsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dynamicPartPixelOne := &DynamicPartPixel{}
	dynamicPartPixelTwo := &DynamicPartPixel{}
	if err = randomize.Struct(seed, dynamicPartPixelOne, dynamicPartPixelDBTypes, false, dynamicPartPixelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DynamicPartPixel struct: %s", err)
	}
	if err = randomize.Struct(seed, dynamicPartPixelTwo, dynamicPartPixelDBTypes, false, dynamicPartPixelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DynamicPartPixel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = dynamicPartPixelOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = dynamicPartPixelTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := DynamicPartPixels().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testDynamicPartPixelsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	dynamicPartPixelOne := &DynamicPartPixel{}
	dynamicPartPixelTwo := &DynamicPartPixel{}
	if err = randomize.Struct(seed, dynamicPartPixelOne, dynamicPartPixelDBTypes, false, dynamicPartPixelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DynamicPartPixel struct: %s", err)
	}
	if err = randomize.Struct(seed, dynamicPartPixelTwo, dynamicPartPixelDBTypes, false, dynamicPartPixelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DynamicPartPixel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = dynamicPartPixelOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = dynamicPartPixelTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := DynamicPartPixels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func dynamicPartPixelBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *DynamicPartPixel) error {
	*o = DynamicPartPixel{}
	return nil
}

func dynamicPartPixelAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *DynamicPartPixel) error {
	*o = DynamicPartPixel{}
	return nil
}

func dynamicPartPixelAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *DynamicPartPixel) error {
	*o = DynamicPartPixel{}
	return nil
}

func dynamicPartPixelBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *DynamicPartPixel) error {
	*o = DynamicPartPixel{}
	return nil
}

func dynamicPartPixelAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *DynamicPartPixel) error {
	*o = DynamicPartPixel{}
	return nil
}

func dynamicPartPixelBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *DynamicPartPixel) error {
	*o = DynamicPartPixel{}
	return nil
}

func dynamicPartPixelAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *DynamicPartPixel) error {
	*o = DynamicPartPixel{}
	return nil
}

func dynamicPartPixelBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *DynamicPartPixel) error {
	*o = DynamicPartPixel{}
	return nil
}

func dynamicPartPixelAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *DynamicPartPixel) error {
	*o = DynamicPartPixel{}
	return nil
}

func testDynamicPartPixelsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &DynamicPartPixel{}
	o := &DynamicPartPixel{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, dynamicPartPixelDBTypes, false); err != nil {
		t.Errorf("Unable to randomize DynamicPartPixel object: %s", err)
	}

	AddDynamicPartPixelHook(boil.BeforeInsertHook, dynamicPartPixelBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	dynamicPartPixelBeforeInsertHooks = []DynamicPartPixelHook{}

	AddDynamicPartPixelHook(boil.AfterInsertHook, dynamicPartPixelAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	dynamicPartPixelAfterInsertHooks = []DynamicPartPixelHook{}

	AddDynamicPartPixelHook(boil.AfterSelectHook, dynamicPartPixelAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	dynamicPartPixelAfterSelectHooks = []DynamicPartPixelHook{}

	AddDynamicPartPixelHook(boil.BeforeUpdateHook, dynamicPartPixelBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	dynamicPartPixelBeforeUpdateHooks = []DynamicPartPixelHook{}

	AddDynamicPartPixelHook(boil.AfterUpdateHook, dynamicPartPixelAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	dynamicPartPixelAfterUpdateHooks = []DynamicPartPixelHook{}

	AddDynamicPartPixelHook(boil.BeforeDeleteHook, dynamicPartPixelBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	dynamicPartPixelBeforeDeleteHooks = []DynamicPartPixelHook{}

	AddDynamicPartPixelHook(boil.AfterDeleteHook, dynamicPartPixelAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	dynamicPartPixelAfterDeleteHooks = []DynamicPartPixelHook{}

	AddDynamicPartPixelHook(boil.BeforeUpsertHook, dynamicPartPixelBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	dynamicPartPixelBeforeUpsertHooks = []DynamicPartPixelHook{}

	AddDynamicPartPixelHook(boil.AfterUpsertHook, dynamicPartPixelAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	dynamicPartPixelAfterUpsertHooks = []DynamicPartPixelHook{}
}

func testDynamicPartPixelsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DynamicPartPixel{}
	if err = randomize.Struct(seed, o, dynamicPartPixelDBTypes, true, dynamicPartPixelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DynamicPartPixel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := DynamicPartPixels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testDynamicPartPixelsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DynamicPartPixel{}
	if err = randomize.Struct(seed, o, dynamicPartPixelDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DynamicPartPixel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(dynamicPartPixelColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := DynamicPartPixels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testDynamicPartPixelToOneColorStringUsingColorString(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local DynamicPartPixel
	var foreign ColorString

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, dynamicPartPixelDBTypes, false, dynamicPartPixelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DynamicPartPixel struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, colorStringDBTypes, false, colorStringColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ColorString struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.ColorStringID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.ColorString().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := DynamicPartPixelSlice{&local}
	if err = local.L.LoadColorString(ctx, tx, false, (*[]*DynamicPartPixel)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.ColorString == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.ColorString = nil
	if err = local.L.LoadColorString(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.ColorString == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testDynamicPartPixelToOneDynamicPartUsingDynamicPart(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local DynamicPartPixel
	var foreign DynamicPart

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, dynamicPartPixelDBTypes, false, dynamicPartPixelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DynamicPartPixel struct: %s", err)
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

	slice := DynamicPartPixelSlice{&local}
	if err = local.L.LoadDynamicPart(ctx, tx, false, (*[]*DynamicPartPixel)(&slice), nil); err != nil {
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

func testDynamicPartPixelToOneSetOpColorStringUsingColorString(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a DynamicPartPixel
	var b, c ColorString

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, dynamicPartPixelDBTypes, false, strmangle.SetComplement(dynamicPartPixelPrimaryKeyColumns, dynamicPartPixelColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, colorStringDBTypes, false, strmangle.SetComplement(colorStringPrimaryKeyColumns, colorStringColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, colorStringDBTypes, false, strmangle.SetComplement(colorStringPrimaryKeyColumns, colorStringColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*ColorString{&b, &c} {
		err = a.SetColorString(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.ColorString != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.DynamicPartPixels[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.ColorStringID != x.ID {
			t.Error("foreign key was wrong value", a.ColorStringID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.ColorStringID))
		reflect.Indirect(reflect.ValueOf(&a.ColorStringID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.ColorStringID != x.ID {
			t.Error("foreign key was wrong value", a.ColorStringID, x.ID)
		}
	}
}
func testDynamicPartPixelToOneSetOpDynamicPartUsingDynamicPart(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a DynamicPartPixel
	var b, c DynamicPart

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, dynamicPartPixelDBTypes, false, strmangle.SetComplement(dynamicPartPixelPrimaryKeyColumns, dynamicPartPixelColumnsWithoutDefault)...); err != nil {
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

		if x.R.DynamicPartPixels[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.DynamicPartID != x.ID {
			t.Error("foreign key was wrong value", a.DynamicPartID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.DynamicPartID))
		reflect.Indirect(reflect.ValueOf(&a.DynamicPartID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.DynamicPartID != x.ID {
			t.Error("foreign key was wrong value", a.DynamicPartID, x.ID)
		}
	}
}

func testDynamicPartPixelsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DynamicPartPixel{}
	if err = randomize.Struct(seed, o, dynamicPartPixelDBTypes, true, dynamicPartPixelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DynamicPartPixel struct: %s", err)
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

func testDynamicPartPixelsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DynamicPartPixel{}
	if err = randomize.Struct(seed, o, dynamicPartPixelDBTypes, true, dynamicPartPixelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DynamicPartPixel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := DynamicPartPixelSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testDynamicPartPixelsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DynamicPartPixel{}
	if err = randomize.Struct(seed, o, dynamicPartPixelDBTypes, true, dynamicPartPixelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DynamicPartPixel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := DynamicPartPixels().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	dynamicPartPixelDBTypes = map[string]string{`ID`: `integer`, `ColorStringID`: `integer`, `DynamicPartID`: `integer`, `X`: `smallint`, `Y`: `smallint`}
	_                       = bytes.MinRead
)

func testDynamicPartPixelsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(dynamicPartPixelPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(dynamicPartPixelAllColumns) == len(dynamicPartPixelPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &DynamicPartPixel{}
	if err = randomize.Struct(seed, o, dynamicPartPixelDBTypes, true, dynamicPartPixelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DynamicPartPixel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := DynamicPartPixels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, dynamicPartPixelDBTypes, true, dynamicPartPixelPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DynamicPartPixel struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testDynamicPartPixelsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(dynamicPartPixelAllColumns) == len(dynamicPartPixelPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &DynamicPartPixel{}
	if err = randomize.Struct(seed, o, dynamicPartPixelDBTypes, true, dynamicPartPixelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DynamicPartPixel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := DynamicPartPixels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, dynamicPartPixelDBTypes, true, dynamicPartPixelPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DynamicPartPixel struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(dynamicPartPixelAllColumns, dynamicPartPixelPrimaryKeyColumns) {
		fields = dynamicPartPixelAllColumns
	} else {
		fields = strmangle.SetComplement(
			dynamicPartPixelAllColumns,
			dynamicPartPixelPrimaryKeyColumns,
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

	slice := DynamicPartPixelSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testDynamicPartPixelsUpsert(t *testing.T) {
	t.Parallel()

	if len(dynamicPartPixelAllColumns) == len(dynamicPartPixelPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := DynamicPartPixel{}
	if err = randomize.Struct(seed, &o, dynamicPartPixelDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DynamicPartPixel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert DynamicPartPixel: %s", err)
	}

	count, err := DynamicPartPixels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, dynamicPartPixelDBTypes, false, dynamicPartPixelPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DynamicPartPixel struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert DynamicPartPixel: %s", err)
	}

	count, err = DynamicPartPixels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
