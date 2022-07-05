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

func testColorStrings(t *testing.T) {
	t.Parallel()

	query := ColorStrings()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testColorStringsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ColorString{}
	if err = randomize.Struct(seed, o, colorStringDBTypes, true, colorStringColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ColorString struct: %s", err)
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

	count, err := ColorStrings().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testColorStringsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ColorString{}
	if err = randomize.Struct(seed, o, colorStringDBTypes, true, colorStringColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ColorString struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := ColorStrings().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := ColorStrings().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testColorStringsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ColorString{}
	if err = randomize.Struct(seed, o, colorStringDBTypes, true, colorStringColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ColorString struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ColorStringSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := ColorStrings().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testColorStringsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ColorString{}
	if err = randomize.Struct(seed, o, colorStringDBTypes, true, colorStringColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ColorString struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := ColorStringExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if ColorString exists: %s", err)
	}
	if !e {
		t.Errorf("Expected ColorStringExists to return true, but got false.")
	}
}

func testColorStringsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ColorString{}
	if err = randomize.Struct(seed, o, colorStringDBTypes, true, colorStringColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ColorString struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	colorStringFound, err := FindColorString(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if colorStringFound == nil {
		t.Error("want a record, got nil")
	}
}

func testColorStringsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ColorString{}
	if err = randomize.Struct(seed, o, colorStringDBTypes, true, colorStringColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ColorString struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = ColorStrings().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testColorStringsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ColorString{}
	if err = randomize.Struct(seed, o, colorStringDBTypes, true, colorStringColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ColorString struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := ColorStrings().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testColorStringsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	colorStringOne := &ColorString{}
	colorStringTwo := &ColorString{}
	if err = randomize.Struct(seed, colorStringOne, colorStringDBTypes, false, colorStringColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ColorString struct: %s", err)
	}
	if err = randomize.Struct(seed, colorStringTwo, colorStringDBTypes, false, colorStringColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ColorString struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = colorStringOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = colorStringTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := ColorStrings().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testColorStringsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	colorStringOne := &ColorString{}
	colorStringTwo := &ColorString{}
	if err = randomize.Struct(seed, colorStringOne, colorStringDBTypes, false, colorStringColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ColorString struct: %s", err)
	}
	if err = randomize.Struct(seed, colorStringTwo, colorStringDBTypes, false, colorStringColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ColorString struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = colorStringOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = colorStringTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ColorStrings().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func colorStringBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *ColorString) error {
	*o = ColorString{}
	return nil
}

func colorStringAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *ColorString) error {
	*o = ColorString{}
	return nil
}

func colorStringAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *ColorString) error {
	*o = ColorString{}
	return nil
}

func colorStringBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *ColorString) error {
	*o = ColorString{}
	return nil
}

func colorStringAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *ColorString) error {
	*o = ColorString{}
	return nil
}

func colorStringBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *ColorString) error {
	*o = ColorString{}
	return nil
}

func colorStringAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *ColorString) error {
	*o = ColorString{}
	return nil
}

func colorStringBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *ColorString) error {
	*o = ColorString{}
	return nil
}

func colorStringAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *ColorString) error {
	*o = ColorString{}
	return nil
}

func testColorStringsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &ColorString{}
	o := &ColorString{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, colorStringDBTypes, false); err != nil {
		t.Errorf("Unable to randomize ColorString object: %s", err)
	}

	AddColorStringHook(boil.BeforeInsertHook, colorStringBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	colorStringBeforeInsertHooks = []ColorStringHook{}

	AddColorStringHook(boil.AfterInsertHook, colorStringAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	colorStringAfterInsertHooks = []ColorStringHook{}

	AddColorStringHook(boil.AfterSelectHook, colorStringAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	colorStringAfterSelectHooks = []ColorStringHook{}

	AddColorStringHook(boil.BeforeUpdateHook, colorStringBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	colorStringBeforeUpdateHooks = []ColorStringHook{}

	AddColorStringHook(boil.AfterUpdateHook, colorStringAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	colorStringAfterUpdateHooks = []ColorStringHook{}

	AddColorStringHook(boil.BeforeDeleteHook, colorStringBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	colorStringBeforeDeleteHooks = []ColorStringHook{}

	AddColorStringHook(boil.AfterDeleteHook, colorStringAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	colorStringAfterDeleteHooks = []ColorStringHook{}

	AddColorStringHook(boil.BeforeUpsertHook, colorStringBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	colorStringBeforeUpsertHooks = []ColorStringHook{}

	AddColorStringHook(boil.AfterUpsertHook, colorStringAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	colorStringAfterUpsertHooks = []ColorStringHook{}
}

func testColorStringsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ColorString{}
	if err = randomize.Struct(seed, o, colorStringDBTypes, true, colorStringColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ColorString struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ColorStrings().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testColorStringsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ColorString{}
	if err = randomize.Struct(seed, o, colorStringDBTypes, true); err != nil {
		t.Errorf("Unable to randomize ColorString struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(colorStringColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := ColorStrings().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testColorStringToManyAnimationFramePixels(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a ColorString
	var b, c AnimationFramePixel

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, colorStringDBTypes, true, colorStringColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ColorString struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, animationFramePixelDBTypes, false, animationFramePixelColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, animationFramePixelDBTypes, false, animationFramePixelColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.ColorStringID = a.ID
	c.ColorStringID = a.ID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.AnimationFramePixels().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.ColorStringID == b.ColorStringID {
			bFound = true
		}
		if v.ColorStringID == c.ColorStringID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := ColorStringSlice{&a}
	if err = a.L.LoadAnimationFramePixels(ctx, tx, false, (*[]*ColorString)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.AnimationFramePixels); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.AnimationFramePixels = nil
	if err = a.L.LoadAnimationFramePixels(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.AnimationFramePixels); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testColorStringToManyDynamicPartMappingPixels(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a ColorString
	var b, c DynamicPartMappingPixel

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, colorStringDBTypes, true, colorStringColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ColorString struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, dynamicPartMappingPixelDBTypes, false, dynamicPartMappingPixelColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, dynamicPartMappingPixelDBTypes, false, dynamicPartMappingPixelColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.ColorStringID = a.ID
	c.ColorStringID = a.ID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.DynamicPartMappingPixels().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.ColorStringID == b.ColorStringID {
			bFound = true
		}
		if v.ColorStringID == c.ColorStringID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := ColorStringSlice{&a}
	if err = a.L.LoadDynamicPartMappingPixels(ctx, tx, false, (*[]*ColorString)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.DynamicPartMappingPixels); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.DynamicPartMappingPixels = nil
	if err = a.L.LoadDynamicPartMappingPixels(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.DynamicPartMappingPixels); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testColorStringToManyDynamicPartPixels(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a ColorString
	var b, c DynamicPartPixel

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, colorStringDBTypes, true, colorStringColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ColorString struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, dynamicPartPixelDBTypes, false, dynamicPartPixelColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, dynamicPartPixelDBTypes, false, dynamicPartPixelColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.ColorStringID = a.ID
	c.ColorStringID = a.ID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.DynamicPartPixels().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.ColorStringID == b.ColorStringID {
			bFound = true
		}
		if v.ColorStringID == c.ColorStringID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := ColorStringSlice{&a}
	if err = a.L.LoadDynamicPartPixels(ctx, tx, false, (*[]*ColorString)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.DynamicPartPixels); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.DynamicPartPixels = nil
	if err = a.L.LoadDynamicPartPixels(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.DynamicPartPixels); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testColorStringToManyAddOpAnimationFramePixels(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a ColorString
	var b, c, d, e AnimationFramePixel

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, colorStringDBTypes, false, strmangle.SetComplement(colorStringPrimaryKeyColumns, colorStringColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*AnimationFramePixel{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, animationFramePixelDBTypes, false, strmangle.SetComplement(animationFramePixelPrimaryKeyColumns, animationFramePixelColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*AnimationFramePixel{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddAnimationFramePixels(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.ColorStringID {
			t.Error("foreign key was wrong value", a.ID, first.ColorStringID)
		}
		if a.ID != second.ColorStringID {
			t.Error("foreign key was wrong value", a.ID, second.ColorStringID)
		}

		if first.R.ColorString != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.ColorString != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.AnimationFramePixels[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.AnimationFramePixels[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.AnimationFramePixels().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testColorStringToManyAddOpDynamicPartMappingPixels(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a ColorString
	var b, c, d, e DynamicPartMappingPixel

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, colorStringDBTypes, false, strmangle.SetComplement(colorStringPrimaryKeyColumns, colorStringColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*DynamicPartMappingPixel{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, dynamicPartMappingPixelDBTypes, false, strmangle.SetComplement(dynamicPartMappingPixelPrimaryKeyColumns, dynamicPartMappingPixelColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*DynamicPartMappingPixel{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddDynamicPartMappingPixels(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.ColorStringID {
			t.Error("foreign key was wrong value", a.ID, first.ColorStringID)
		}
		if a.ID != second.ColorStringID {
			t.Error("foreign key was wrong value", a.ID, second.ColorStringID)
		}

		if first.R.ColorString != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.ColorString != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.DynamicPartMappingPixels[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.DynamicPartMappingPixels[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.DynamicPartMappingPixels().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testColorStringToManyAddOpDynamicPartPixels(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a ColorString
	var b, c, d, e DynamicPartPixel

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, colorStringDBTypes, false, strmangle.SetComplement(colorStringPrimaryKeyColumns, colorStringColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*DynamicPartPixel{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, dynamicPartPixelDBTypes, false, strmangle.SetComplement(dynamicPartPixelPrimaryKeyColumns, dynamicPartPixelColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*DynamicPartPixel{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddDynamicPartPixels(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.ColorStringID {
			t.Error("foreign key was wrong value", a.ID, first.ColorStringID)
		}
		if a.ID != second.ColorStringID {
			t.Error("foreign key was wrong value", a.ID, second.ColorStringID)
		}

		if first.R.ColorString != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.ColorString != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.DynamicPartPixels[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.DynamicPartPixels[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.DynamicPartPixels().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testColorStringsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ColorString{}
	if err = randomize.Struct(seed, o, colorStringDBTypes, true, colorStringColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ColorString struct: %s", err)
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

func testColorStringsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ColorString{}
	if err = randomize.Struct(seed, o, colorStringDBTypes, true, colorStringColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ColorString struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ColorStringSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testColorStringsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ColorString{}
	if err = randomize.Struct(seed, o, colorStringDBTypes, true, colorStringColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ColorString struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := ColorStrings().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	colorStringDBTypes = map[string]string{`ID`: `integer`, `Hexstring`: `character varying`}
	_                  = bytes.MinRead
)

func testColorStringsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(colorStringPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(colorStringAllColumns) == len(colorStringPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &ColorString{}
	if err = randomize.Struct(seed, o, colorStringDBTypes, true, colorStringColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ColorString struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ColorStrings().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, colorStringDBTypes, true, colorStringPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ColorString struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testColorStringsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(colorStringAllColumns) == len(colorStringPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &ColorString{}
	if err = randomize.Struct(seed, o, colorStringDBTypes, true, colorStringColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ColorString struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ColorStrings().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, colorStringDBTypes, true, colorStringPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ColorString struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(colorStringAllColumns, colorStringPrimaryKeyColumns) {
		fields = colorStringAllColumns
	} else {
		fields = strmangle.SetComplement(
			colorStringAllColumns,
			colorStringPrimaryKeyColumns,
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

	slice := ColorStringSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testColorStringsUpsert(t *testing.T) {
	t.Parallel()

	if len(colorStringAllColumns) == len(colorStringPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := ColorString{}
	if err = randomize.Struct(seed, &o, colorStringDBTypes, true); err != nil {
		t.Errorf("Unable to randomize ColorString struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert ColorString: %s", err)
	}

	count, err := ColorStrings().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, colorStringDBTypes, false, colorStringPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ColorString struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert ColorString: %s", err)
	}

	count, err = ColorStrings().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}