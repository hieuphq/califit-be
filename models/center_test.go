// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/randomize"
	"github.com/volatiletech/sqlboiler/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testCenters(t *testing.T) {
	t.Parallel()

	query := Centers()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testCentersDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Center{}
	if err = randomize.Struct(seed, o, centerDBTypes, true, centerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Center struct: %s", err)
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

	count, err := Centers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testCentersQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Center{}
	if err = randomize.Struct(seed, o, centerDBTypes, true, centerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Center struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Centers().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Centers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testCentersSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Center{}
	if err = randomize.Struct(seed, o, centerDBTypes, true, centerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Center struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := CenterSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Centers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testCentersExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Center{}
	if err = randomize.Struct(seed, o, centerDBTypes, true, centerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Center struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := CenterExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Center exists: %s", err)
	}
	if !e {
		t.Errorf("Expected CenterExists to return true, but got false.")
	}
}

func testCentersFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Center{}
	if err = randomize.Struct(seed, o, centerDBTypes, true, centerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Center struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	centerFound, err := FindCenter(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if centerFound == nil {
		t.Error("want a record, got nil")
	}
}

func testCentersBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Center{}
	if err = randomize.Struct(seed, o, centerDBTypes, true, centerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Center struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Centers().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testCentersOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Center{}
	if err = randomize.Struct(seed, o, centerDBTypes, true, centerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Center struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Centers().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testCentersAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	centerOne := &Center{}
	centerTwo := &Center{}
	if err = randomize.Struct(seed, centerOne, centerDBTypes, false, centerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Center struct: %s", err)
	}
	if err = randomize.Struct(seed, centerTwo, centerDBTypes, false, centerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Center struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = centerOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = centerTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Centers().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testCentersCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	centerOne := &Center{}
	centerTwo := &Center{}
	if err = randomize.Struct(seed, centerOne, centerDBTypes, false, centerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Center struct: %s", err)
	}
	if err = randomize.Struct(seed, centerTwo, centerDBTypes, false, centerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Center struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = centerOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = centerTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Centers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func centerBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Center) error {
	*o = Center{}
	return nil
}

func centerAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Center) error {
	*o = Center{}
	return nil
}

func centerAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Center) error {
	*o = Center{}
	return nil
}

func centerBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Center) error {
	*o = Center{}
	return nil
}

func centerAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Center) error {
	*o = Center{}
	return nil
}

func centerBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Center) error {
	*o = Center{}
	return nil
}

func centerAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Center) error {
	*o = Center{}
	return nil
}

func centerBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Center) error {
	*o = Center{}
	return nil
}

func centerAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Center) error {
	*o = Center{}
	return nil
}

func testCentersHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Center{}
	o := &Center{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, centerDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Center object: %s", err)
	}

	AddCenterHook(boil.BeforeInsertHook, centerBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	centerBeforeInsertHooks = []CenterHook{}

	AddCenterHook(boil.AfterInsertHook, centerAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	centerAfterInsertHooks = []CenterHook{}

	AddCenterHook(boil.AfterSelectHook, centerAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	centerAfterSelectHooks = []CenterHook{}

	AddCenterHook(boil.BeforeUpdateHook, centerBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	centerBeforeUpdateHooks = []CenterHook{}

	AddCenterHook(boil.AfterUpdateHook, centerAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	centerAfterUpdateHooks = []CenterHook{}

	AddCenterHook(boil.BeforeDeleteHook, centerBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	centerBeforeDeleteHooks = []CenterHook{}

	AddCenterHook(boil.AfterDeleteHook, centerAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	centerAfterDeleteHooks = []CenterHook{}

	AddCenterHook(boil.BeforeUpsertHook, centerBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	centerBeforeUpsertHooks = []CenterHook{}

	AddCenterHook(boil.AfterUpsertHook, centerAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	centerAfterUpsertHooks = []CenterHook{}
}

func testCentersInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Center{}
	if err = randomize.Struct(seed, o, centerDBTypes, true, centerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Center struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Centers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testCentersInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Center{}
	if err = randomize.Struct(seed, o, centerDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Center struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(centerColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Centers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testCenterToManySchedules(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Center
	var b, c Schedule

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, centerDBTypes, true, centerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Center struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, scheduleDBTypes, false, scheduleColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, scheduleDBTypes, false, scheduleColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&b.CenterID, a.ID)
	queries.Assign(&c.CenterID, a.ID)
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	schedule, err := a.Schedules().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range schedule {
		if queries.Equal(v.CenterID, b.CenterID) {
			bFound = true
		}
		if queries.Equal(v.CenterID, c.CenterID) {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := CenterSlice{&a}
	if err = a.L.LoadSchedules(ctx, tx, false, (*[]*Center)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Schedules); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.Schedules = nil
	if err = a.L.LoadSchedules(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Schedules); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", schedule)
	}
}

func testCenterToManyAddOpSchedules(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Center
	var b, c, d, e Schedule

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, centerDBTypes, false, strmangle.SetComplement(centerPrimaryKeyColumns, centerColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Schedule{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, scheduleDBTypes, false, strmangle.SetComplement(schedulePrimaryKeyColumns, scheduleColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*Schedule{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddSchedules(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if !queries.Equal(a.ID, first.CenterID) {
			t.Error("foreign key was wrong value", a.ID, first.CenterID)
		}
		if !queries.Equal(a.ID, second.CenterID) {
			t.Error("foreign key was wrong value", a.ID, second.CenterID)
		}

		if first.R.Center != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Center != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.Schedules[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.Schedules[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.Schedules().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testCenterToManySetOpSchedules(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Center
	var b, c, d, e Schedule

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, centerDBTypes, false, strmangle.SetComplement(centerPrimaryKeyColumns, centerColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Schedule{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, scheduleDBTypes, false, strmangle.SetComplement(schedulePrimaryKeyColumns, scheduleColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err = a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.SetSchedules(ctx, tx, false, &b, &c)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.Schedules().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	err = a.SetSchedules(ctx, tx, true, &d, &e)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.Schedules().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if !queries.IsValuerNil(b.CenterID) {
		t.Error("want b's foreign key value to be nil")
	}
	if !queries.IsValuerNil(c.CenterID) {
		t.Error("want c's foreign key value to be nil")
	}
	if !queries.Equal(a.ID, d.CenterID) {
		t.Error("foreign key was wrong value", a.ID, d.CenterID)
	}
	if !queries.Equal(a.ID, e.CenterID) {
		t.Error("foreign key was wrong value", a.ID, e.CenterID)
	}

	if b.R.Center != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.Center != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.Center != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}
	if e.R.Center != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}

	if a.R.Schedules[0] != &d {
		t.Error("relationship struct slice not set to correct value")
	}
	if a.R.Schedules[1] != &e {
		t.Error("relationship struct slice not set to correct value")
	}
}

func testCenterToManyRemoveOpSchedules(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Center
	var b, c, d, e Schedule

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, centerDBTypes, false, strmangle.SetComplement(centerPrimaryKeyColumns, centerColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Schedule{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, scheduleDBTypes, false, strmangle.SetComplement(schedulePrimaryKeyColumns, scheduleColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.AddSchedules(ctx, tx, true, foreigners...)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.Schedules().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 4 {
		t.Error("count was wrong:", count)
	}

	err = a.RemoveSchedules(ctx, tx, foreigners[:2]...)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.Schedules().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if !queries.IsValuerNil(b.CenterID) {
		t.Error("want b's foreign key value to be nil")
	}
	if !queries.IsValuerNil(c.CenterID) {
		t.Error("want c's foreign key value to be nil")
	}

	if b.R.Center != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.Center != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.Center != &a {
		t.Error("relationship to a should have been preserved")
	}
	if e.R.Center != &a {
		t.Error("relationship to a should have been preserved")
	}

	if len(a.R.Schedules) != 2 {
		t.Error("should have preserved two relationships")
	}

	// Removal doesn't do a stable deletion for performance so we have to flip the order
	if a.R.Schedules[1] != &d {
		t.Error("relationship to d should have been preserved")
	}
	if a.R.Schedules[0] != &e {
		t.Error("relationship to e should have been preserved")
	}
}

func testCenterToOneAddressUsingAddress(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local Center
	var foreign Address

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, centerDBTypes, true, centerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Center struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, addressDBTypes, false, addressColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Address struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&local.AddressID, foreign.ID)
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Address().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if !queries.Equal(check.ID, foreign.ID) {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := CenterSlice{&local}
	if err = local.L.LoadAddress(ctx, tx, false, (*[]*Center)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Address == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Address = nil
	if err = local.L.LoadAddress(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Address == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testCenterToOneCityUsingCity(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local Center
	var foreign City

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, centerDBTypes, true, centerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Center struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, cityDBTypes, false, cityColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize City struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&local.CityID, foreign.ID)
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.City().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if !queries.Equal(check.ID, foreign.ID) {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := CenterSlice{&local}
	if err = local.L.LoadCity(ctx, tx, false, (*[]*Center)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.City == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.City = nil
	if err = local.L.LoadCity(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.City == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testCenterToOneSetOpAddressUsingAddress(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Center
	var b, c Address

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, centerDBTypes, false, strmangle.SetComplement(centerPrimaryKeyColumns, centerColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, addressDBTypes, false, strmangle.SetComplement(addressPrimaryKeyColumns, addressColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, addressDBTypes, false, strmangle.SetComplement(addressPrimaryKeyColumns, addressColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Address{&b, &c} {
		err = a.SetAddress(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Address != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.Centers[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if !queries.Equal(a.AddressID, x.ID) {
			t.Error("foreign key was wrong value", a.AddressID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.AddressID))
		reflect.Indirect(reflect.ValueOf(&a.AddressID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if !queries.Equal(a.AddressID, x.ID) {
			t.Error("foreign key was wrong value", a.AddressID, x.ID)
		}
	}
}

func testCenterToOneRemoveOpAddressUsingAddress(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Center
	var b Address

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, centerDBTypes, false, strmangle.SetComplement(centerPrimaryKeyColumns, centerColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, addressDBTypes, false, strmangle.SetComplement(addressPrimaryKeyColumns, addressColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err = a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = a.SetAddress(ctx, tx, true, &b); err != nil {
		t.Fatal(err)
	}

	if err = a.RemoveAddress(ctx, tx, &b); err != nil {
		t.Error("failed to remove relationship")
	}

	count, err := a.Address().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 0 {
		t.Error("want no relationships remaining")
	}

	if a.R.Address != nil {
		t.Error("R struct entry should be nil")
	}

	if !queries.IsValuerNil(a.AddressID) {
		t.Error("foreign key value should be nil")
	}

	if len(b.R.Centers) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testCenterToOneSetOpCityUsingCity(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Center
	var b, c City

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, centerDBTypes, false, strmangle.SetComplement(centerPrimaryKeyColumns, centerColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, cityDBTypes, false, strmangle.SetComplement(cityPrimaryKeyColumns, cityColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, cityDBTypes, false, strmangle.SetComplement(cityPrimaryKeyColumns, cityColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*City{&b, &c} {
		err = a.SetCity(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.City != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.Centers[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if !queries.Equal(a.CityID, x.ID) {
			t.Error("foreign key was wrong value", a.CityID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.CityID))
		reflect.Indirect(reflect.ValueOf(&a.CityID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if !queries.Equal(a.CityID, x.ID) {
			t.Error("foreign key was wrong value", a.CityID, x.ID)
		}
	}
}

func testCenterToOneRemoveOpCityUsingCity(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Center
	var b City

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, centerDBTypes, false, strmangle.SetComplement(centerPrimaryKeyColumns, centerColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, cityDBTypes, false, strmangle.SetComplement(cityPrimaryKeyColumns, cityColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err = a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = a.SetCity(ctx, tx, true, &b); err != nil {
		t.Fatal(err)
	}

	if err = a.RemoveCity(ctx, tx, &b); err != nil {
		t.Error("failed to remove relationship")
	}

	count, err := a.City().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 0 {
		t.Error("want no relationships remaining")
	}

	if a.R.City != nil {
		t.Error("R struct entry should be nil")
	}

	if !queries.IsValuerNil(a.CityID) {
		t.Error("foreign key value should be nil")
	}

	if len(b.R.Centers) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testCentersReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Center{}
	if err = randomize.Struct(seed, o, centerDBTypes, true, centerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Center struct: %s", err)
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

func testCentersReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Center{}
	if err = randomize.Struct(seed, o, centerDBTypes, true, centerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Center struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := CenterSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testCentersSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Center{}
	if err = randomize.Struct(seed, o, centerDBTypes, true, centerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Center struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Centers().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	centerDBTypes = map[string]string{`AddressID`: `uuid`, `CityID`: `uuid`, `CloseAt`: `character varying`, `CreatedAt`: `timestamp with time zone`, `ID`: `uuid`, `Name`: `character varying`, `OpenAt`: `character varying`, `UpdatedAt`: `timestamp with time zone`}
	_             = bytes.MinRead
)

func testCentersUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(centerPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(centerColumns) == len(centerPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Center{}
	if err = randomize.Struct(seed, o, centerDBTypes, true, centerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Center struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Centers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, centerDBTypes, true, centerPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Center struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testCentersSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(centerColumns) == len(centerPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Center{}
	if err = randomize.Struct(seed, o, centerDBTypes, true, centerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Center struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Centers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, centerDBTypes, true, centerPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Center struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(centerColumns, centerPrimaryKeyColumns) {
		fields = centerColumns
	} else {
		fields = strmangle.SetComplement(
			centerColumns,
			centerPrimaryKeyColumns,
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

	slice := CenterSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testCentersUpsert(t *testing.T) {
	t.Parallel()

	if len(centerColumns) == len(centerPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Center{}
	if err = randomize.Struct(seed, &o, centerDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Center struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Center: %s", err)
	}

	count, err := Centers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, centerDBTypes, false, centerPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Center struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Center: %s", err)
	}

	count, err = Centers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
