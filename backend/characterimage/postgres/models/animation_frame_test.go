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

func testAnimationFrames(t *testing.T) {
	t.Parallel()

	query := AnimationFrames()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testAnimationFramesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AnimationFrame{}
	if err = randomize.Struct(seed, o, animationFrameDBTypes, true, animationFrameColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AnimationFrame struct: %s", err)
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

	count, err := AnimationFrames().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testAnimationFramesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AnimationFrame{}
	if err = randomize.Struct(seed, o, animationFrameDBTypes, true, animationFrameColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AnimationFrame struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := AnimationFrames().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := AnimationFrames().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testAnimationFramesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AnimationFrame{}
	if err = randomize.Struct(seed, o, animationFrameDBTypes, true, animationFrameColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AnimationFrame struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := AnimationFrameSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := AnimationFrames().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testAnimationFramesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AnimationFrame{}
	if err = randomize.Struct(seed, o, animationFrameDBTypes, true, animationFrameColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AnimationFrame struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := AnimationFrameExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if AnimationFrame exists: %s", err)
	}
	if !e {
		t.Errorf("Expected AnimationFrameExists to return true, but got false.")
	}
}

func testAnimationFramesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AnimationFrame{}
	if err = randomize.Struct(seed, o, animationFrameDBTypes, true, animationFrameColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AnimationFrame struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	animationFrameFound, err := FindAnimationFrame(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if animationFrameFound == nil {
		t.Error("want a record, got nil")
	}
}

func testAnimationFramesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AnimationFrame{}
	if err = randomize.Struct(seed, o, animationFrameDBTypes, true, animationFrameColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AnimationFrame struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = AnimationFrames().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testAnimationFramesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AnimationFrame{}
	if err = randomize.Struct(seed, o, animationFrameDBTypes, true, animationFrameColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AnimationFrame struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := AnimationFrames().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testAnimationFramesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	animationFrameOne := &AnimationFrame{}
	animationFrameTwo := &AnimationFrame{}
	if err = randomize.Struct(seed, animationFrameOne, animationFrameDBTypes, false, animationFrameColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AnimationFrame struct: %s", err)
	}
	if err = randomize.Struct(seed, animationFrameTwo, animationFrameDBTypes, false, animationFrameColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AnimationFrame struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = animationFrameOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = animationFrameTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := AnimationFrames().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testAnimationFramesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	animationFrameOne := &AnimationFrame{}
	animationFrameTwo := &AnimationFrame{}
	if err = randomize.Struct(seed, animationFrameOne, animationFrameDBTypes, false, animationFrameColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AnimationFrame struct: %s", err)
	}
	if err = randomize.Struct(seed, animationFrameTwo, animationFrameDBTypes, false, animationFrameColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AnimationFrame struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = animationFrameOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = animationFrameTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := AnimationFrames().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func animationFrameBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *AnimationFrame) error {
	*o = AnimationFrame{}
	return nil
}

func animationFrameAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *AnimationFrame) error {
	*o = AnimationFrame{}
	return nil
}

func animationFrameAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *AnimationFrame) error {
	*o = AnimationFrame{}
	return nil
}

func animationFrameBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *AnimationFrame) error {
	*o = AnimationFrame{}
	return nil
}

func animationFrameAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *AnimationFrame) error {
	*o = AnimationFrame{}
	return nil
}

func animationFrameBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *AnimationFrame) error {
	*o = AnimationFrame{}
	return nil
}

func animationFrameAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *AnimationFrame) error {
	*o = AnimationFrame{}
	return nil
}

func animationFrameBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *AnimationFrame) error {
	*o = AnimationFrame{}
	return nil
}

func animationFrameAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *AnimationFrame) error {
	*o = AnimationFrame{}
	return nil
}

func testAnimationFramesHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &AnimationFrame{}
	o := &AnimationFrame{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, animationFrameDBTypes, false); err != nil {
		t.Errorf("Unable to randomize AnimationFrame object: %s", err)
	}

	AddAnimationFrameHook(boil.BeforeInsertHook, animationFrameBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	animationFrameBeforeInsertHooks = []AnimationFrameHook{}

	AddAnimationFrameHook(boil.AfterInsertHook, animationFrameAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	animationFrameAfterInsertHooks = []AnimationFrameHook{}

	AddAnimationFrameHook(boil.AfterSelectHook, animationFrameAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	animationFrameAfterSelectHooks = []AnimationFrameHook{}

	AddAnimationFrameHook(boil.BeforeUpdateHook, animationFrameBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	animationFrameBeforeUpdateHooks = []AnimationFrameHook{}

	AddAnimationFrameHook(boil.AfterUpdateHook, animationFrameAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	animationFrameAfterUpdateHooks = []AnimationFrameHook{}

	AddAnimationFrameHook(boil.BeforeDeleteHook, animationFrameBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	animationFrameBeforeDeleteHooks = []AnimationFrameHook{}

	AddAnimationFrameHook(boil.AfterDeleteHook, animationFrameAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	animationFrameAfterDeleteHooks = []AnimationFrameHook{}

	AddAnimationFrameHook(boil.BeforeUpsertHook, animationFrameBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	animationFrameBeforeUpsertHooks = []AnimationFrameHook{}

	AddAnimationFrameHook(boil.AfterUpsertHook, animationFrameAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	animationFrameAfterUpsertHooks = []AnimationFrameHook{}
}

func testAnimationFramesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AnimationFrame{}
	if err = randomize.Struct(seed, o, animationFrameDBTypes, true, animationFrameColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AnimationFrame struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := AnimationFrames().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testAnimationFramesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AnimationFrame{}
	if err = randomize.Struct(seed, o, animationFrameDBTypes, true); err != nil {
		t.Errorf("Unable to randomize AnimationFrame struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(animationFrameColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := AnimationFrames().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testAnimationFrameToManyAnimationFramePixels(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a AnimationFrame
	var b, c AnimationFramePixel

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, animationFrameDBTypes, true, animationFrameColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AnimationFrame struct: %s", err)
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

	b.AnimationFrameID = a.ID
	c.AnimationFrameID = a.ID

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
		if v.AnimationFrameID == b.AnimationFrameID {
			bFound = true
		}
		if v.AnimationFrameID == c.AnimationFrameID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := AnimationFrameSlice{&a}
	if err = a.L.LoadAnimationFramePixels(ctx, tx, false, (*[]*AnimationFrame)(&slice), nil); err != nil {
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

func testAnimationFrameToManyAnimationFramePropPositions(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a AnimationFrame
	var b, c AnimationFramePropPosition

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, animationFrameDBTypes, true, animationFrameColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AnimationFrame struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, animationFramePropPositionDBTypes, false, animationFramePropPositionColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, animationFramePropPositionDBTypes, false, animationFramePropPositionColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.AnimationFrameID = a.ID
	c.AnimationFrameID = a.ID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.AnimationFramePropPositions().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.AnimationFrameID == b.AnimationFrameID {
			bFound = true
		}
		if v.AnimationFrameID == c.AnimationFrameID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := AnimationFrameSlice{&a}
	if err = a.L.LoadAnimationFramePropPositions(ctx, tx, false, (*[]*AnimationFrame)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.AnimationFramePropPositions); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.AnimationFramePropPositions = nil
	if err = a.L.LoadAnimationFramePropPositions(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.AnimationFramePropPositions); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testAnimationFrameToManyAnimationFrameStaticPositions(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a AnimationFrame
	var b, c AnimationFrameStaticPosition

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, animationFrameDBTypes, true, animationFrameColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AnimationFrame struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, animationFrameStaticPositionDBTypes, false, animationFrameStaticPositionColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, animationFrameStaticPositionDBTypes, false, animationFrameStaticPositionColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.AnimationFrameID = a.ID
	c.AnimationFrameID = a.ID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.AnimationFrameStaticPositions().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.AnimationFrameID == b.AnimationFrameID {
			bFound = true
		}
		if v.AnimationFrameID == c.AnimationFrameID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := AnimationFrameSlice{&a}
	if err = a.L.LoadAnimationFrameStaticPositions(ctx, tx, false, (*[]*AnimationFrame)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.AnimationFrameStaticPositions); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.AnimationFrameStaticPositions = nil
	if err = a.L.LoadAnimationFrameStaticPositions(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.AnimationFrameStaticPositions); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testAnimationFrameToManyAddOpAnimationFramePixels(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a AnimationFrame
	var b, c, d, e AnimationFramePixel

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, animationFrameDBTypes, false, strmangle.SetComplement(animationFramePrimaryKeyColumns, animationFrameColumnsWithoutDefault)...); err != nil {
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

		if a.ID != first.AnimationFrameID {
			t.Error("foreign key was wrong value", a.ID, first.AnimationFrameID)
		}
		if a.ID != second.AnimationFrameID {
			t.Error("foreign key was wrong value", a.ID, second.AnimationFrameID)
		}

		if first.R.AnimationFrame != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.AnimationFrame != &a {
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
func testAnimationFrameToManyAddOpAnimationFramePropPositions(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a AnimationFrame
	var b, c, d, e AnimationFramePropPosition

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, animationFrameDBTypes, false, strmangle.SetComplement(animationFramePrimaryKeyColumns, animationFrameColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*AnimationFramePropPosition{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, animationFramePropPositionDBTypes, false, strmangle.SetComplement(animationFramePropPositionPrimaryKeyColumns, animationFramePropPositionColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*AnimationFramePropPosition{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddAnimationFramePropPositions(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.AnimationFrameID {
			t.Error("foreign key was wrong value", a.ID, first.AnimationFrameID)
		}
		if a.ID != second.AnimationFrameID {
			t.Error("foreign key was wrong value", a.ID, second.AnimationFrameID)
		}

		if first.R.AnimationFrame != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.AnimationFrame != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.AnimationFramePropPositions[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.AnimationFramePropPositions[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.AnimationFramePropPositions().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testAnimationFrameToManyAddOpAnimationFrameStaticPositions(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a AnimationFrame
	var b, c, d, e AnimationFrameStaticPosition

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, animationFrameDBTypes, false, strmangle.SetComplement(animationFramePrimaryKeyColumns, animationFrameColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*AnimationFrameStaticPosition{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, animationFrameStaticPositionDBTypes, false, strmangle.SetComplement(animationFrameStaticPositionPrimaryKeyColumns, animationFrameStaticPositionColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*AnimationFrameStaticPosition{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddAnimationFrameStaticPositions(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.AnimationFrameID {
			t.Error("foreign key was wrong value", a.ID, first.AnimationFrameID)
		}
		if a.ID != second.AnimationFrameID {
			t.Error("foreign key was wrong value", a.ID, second.AnimationFrameID)
		}

		if first.R.AnimationFrame != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.AnimationFrame != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.AnimationFrameStaticPositions[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.AnimationFrameStaticPositions[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.AnimationFrameStaticPositions().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testAnimationFrameToOneAnimationUsingAnimation(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local AnimationFrame
	var foreign Animation

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, animationFrameDBTypes, false, animationFrameColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AnimationFrame struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, animationDBTypes, false, animationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Animation struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.AnimationID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Animation().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := AnimationFrameSlice{&local}
	if err = local.L.LoadAnimation(ctx, tx, false, (*[]*AnimationFrame)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Animation == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Animation = nil
	if err = local.L.LoadAnimation(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Animation == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testAnimationFrameToOneSetOpAnimationUsingAnimation(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a AnimationFrame
	var b, c Animation

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, animationFrameDBTypes, false, strmangle.SetComplement(animationFramePrimaryKeyColumns, animationFrameColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, animationDBTypes, false, strmangle.SetComplement(animationPrimaryKeyColumns, animationColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, animationDBTypes, false, strmangle.SetComplement(animationPrimaryKeyColumns, animationColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Animation{&b, &c} {
		err = a.SetAnimation(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Animation != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.AnimationFrames[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.AnimationID != x.ID {
			t.Error("foreign key was wrong value", a.AnimationID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.AnimationID))
		reflect.Indirect(reflect.ValueOf(&a.AnimationID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.AnimationID != x.ID {
			t.Error("foreign key was wrong value", a.AnimationID, x.ID)
		}
	}
}

func testAnimationFramesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AnimationFrame{}
	if err = randomize.Struct(seed, o, animationFrameDBTypes, true, animationFrameColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AnimationFrame struct: %s", err)
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

func testAnimationFramesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AnimationFrame{}
	if err = randomize.Struct(seed, o, animationFrameDBTypes, true, animationFrameColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AnimationFrame struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := AnimationFrameSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testAnimationFramesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AnimationFrame{}
	if err = randomize.Struct(seed, o, animationFrameDBTypes, true, animationFrameColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AnimationFrame struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := AnimationFrames().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	animationFrameDBTypes = map[string]string{`ID`: `integer`, `AnimationID`: `integer`, `FrameIndex`: `integer`, `Expression`: `enum.expression_type('NEUTRAL')`}
	_                     = bytes.MinRead
)

func testAnimationFramesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(animationFramePrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(animationFrameAllColumns) == len(animationFramePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &AnimationFrame{}
	if err = randomize.Struct(seed, o, animationFrameDBTypes, true, animationFrameColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AnimationFrame struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := AnimationFrames().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, animationFrameDBTypes, true, animationFramePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize AnimationFrame struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testAnimationFramesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(animationFrameAllColumns) == len(animationFramePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &AnimationFrame{}
	if err = randomize.Struct(seed, o, animationFrameDBTypes, true, animationFrameColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AnimationFrame struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := AnimationFrames().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, animationFrameDBTypes, true, animationFramePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize AnimationFrame struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(animationFrameAllColumns, animationFramePrimaryKeyColumns) {
		fields = animationFrameAllColumns
	} else {
		fields = strmangle.SetComplement(
			animationFrameAllColumns,
			animationFramePrimaryKeyColumns,
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

	slice := AnimationFrameSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testAnimationFramesUpsert(t *testing.T) {
	t.Parallel()

	if len(animationFrameAllColumns) == len(animationFramePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := AnimationFrame{}
	if err = randomize.Struct(seed, &o, animationFrameDBTypes, true); err != nil {
		t.Errorf("Unable to randomize AnimationFrame struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert AnimationFrame: %s", err)
	}

	count, err := AnimationFrames().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, animationFrameDBTypes, false, animationFramePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize AnimationFrame struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert AnimationFrame: %s", err)
	}

	count, err = AnimationFrames().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
