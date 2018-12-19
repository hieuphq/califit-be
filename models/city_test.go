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

func testCities(t *testing.T) {
	t.Parallel()

	query := Cities()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testCitiesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &City{}
	if err = randomize.Struct(seed, o, cityDBTypes, true, cityColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize City struct: %s", err)
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

	count, err := Cities().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testCitiesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &City{}
	if err = randomize.Struct(seed, o, cityDBTypes, true, cityColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize City struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Cities().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Cities().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testCitiesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &City{}
	if err = randomize.Struct(seed, o, cityDBTypes, true, cityColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize City struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := CitySlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Cities().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testCitiesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &City{}
	if err = randomize.Struct(seed, o, cityDBTypes, true, cityColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize City struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := CityExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if City exists: %s", err)
	}
	if !e {
		t.Errorf("Expected CityExists to return true, but got false.")
	}
}

func testCitiesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &City{}
	if err = randomize.Struct(seed, o, cityDBTypes, true, cityColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize City struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	cityFound, err := FindCity(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if cityFound == nil {
		t.Error("want a record, got nil")
	}
}

func testCitiesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &City{}
	if err = randomize.Struct(seed, o, cityDBTypes, true, cityColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize City struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Cities().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testCitiesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &City{}
	if err = randomize.Struct(seed, o, cityDBTypes, true, cityColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize City struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Cities().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testCitiesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	cityOne := &City{}
	cityTwo := &City{}
	if err = randomize.Struct(seed, cityOne, cityDBTypes, false, cityColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize City struct: %s", err)
	}
	if err = randomize.Struct(seed, cityTwo, cityDBTypes, false, cityColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize City struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = cityOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = cityTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Cities().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testCitiesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	cityOne := &City{}
	cityTwo := &City{}
	if err = randomize.Struct(seed, cityOne, cityDBTypes, false, cityColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize City struct: %s", err)
	}
	if err = randomize.Struct(seed, cityTwo, cityDBTypes, false, cityColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize City struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = cityOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = cityTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Cities().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func cityBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *City) error {
	*o = City{}
	return nil
}

func cityAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *City) error {
	*o = City{}
	return nil
}

func cityAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *City) error {
	*o = City{}
	return nil
}

func cityBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *City) error {
	*o = City{}
	return nil
}

func cityAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *City) error {
	*o = City{}
	return nil
}

func cityBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *City) error {
	*o = City{}
	return nil
}

func cityAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *City) error {
	*o = City{}
	return nil
}

func cityBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *City) error {
	*o = City{}
	return nil
}

func cityAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *City) error {
	*o = City{}
	return nil
}

func testCitiesHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &City{}
	o := &City{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, cityDBTypes, false); err != nil {
		t.Errorf("Unable to randomize City object: %s", err)
	}

	AddCityHook(boil.BeforeInsertHook, cityBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	cityBeforeInsertHooks = []CityHook{}

	AddCityHook(boil.AfterInsertHook, cityAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	cityAfterInsertHooks = []CityHook{}

	AddCityHook(boil.AfterSelectHook, cityAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	cityAfterSelectHooks = []CityHook{}

	AddCityHook(boil.BeforeUpdateHook, cityBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	cityBeforeUpdateHooks = []CityHook{}

	AddCityHook(boil.AfterUpdateHook, cityAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	cityAfterUpdateHooks = []CityHook{}

	AddCityHook(boil.BeforeDeleteHook, cityBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	cityBeforeDeleteHooks = []CityHook{}

	AddCityHook(boil.AfterDeleteHook, cityAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	cityAfterDeleteHooks = []CityHook{}

	AddCityHook(boil.BeforeUpsertHook, cityBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	cityBeforeUpsertHooks = []CityHook{}

	AddCityHook(boil.AfterUpsertHook, cityAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	cityAfterUpsertHooks = []CityHook{}
}

func testCitiesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &City{}
	if err = randomize.Struct(seed, o, cityDBTypes, true, cityColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize City struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Cities().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testCitiesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &City{}
	if err = randomize.Struct(seed, o, cityDBTypes, true); err != nil {
		t.Errorf("Unable to randomize City struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(cityColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Cities().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testCityToManyCenters(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a City
	var b, c Center

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, cityDBTypes, true, cityColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize City struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, centerDBTypes, false, centerColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, centerDBTypes, false, centerColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&b.CityID, a.ID)
	queries.Assign(&c.CityID, a.ID)
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	center, err := a.Centers().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range center {
		if queries.Equal(v.CityID, b.CityID) {
			bFound = true
		}
		if queries.Equal(v.CityID, c.CityID) {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := CitySlice{&a}
	if err = a.L.LoadCenters(ctx, tx, false, (*[]*City)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Centers); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.Centers = nil
	if err = a.L.LoadCenters(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Centers); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", center)
	}
}

func testCityToManyAddOpCenters(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a City
	var b, c, d, e Center

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, cityDBTypes, false, strmangle.SetComplement(cityPrimaryKeyColumns, cityColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Center{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, centerDBTypes, false, strmangle.SetComplement(centerPrimaryKeyColumns, centerColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*Center{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddCenters(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if !queries.Equal(a.ID, first.CityID) {
			t.Error("foreign key was wrong value", a.ID, first.CityID)
		}
		if !queries.Equal(a.ID, second.CityID) {
			t.Error("foreign key was wrong value", a.ID, second.CityID)
		}

		if first.R.City != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.City != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.Centers[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.Centers[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.Centers().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testCityToManySetOpCenters(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a City
	var b, c, d, e Center

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, cityDBTypes, false, strmangle.SetComplement(cityPrimaryKeyColumns, cityColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Center{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, centerDBTypes, false, strmangle.SetComplement(centerPrimaryKeyColumns, centerColumnsWithoutDefault)...); err != nil {
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

	err = a.SetCenters(ctx, tx, false, &b, &c)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.Centers().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	err = a.SetCenters(ctx, tx, true, &d, &e)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.Centers().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if !queries.IsValuerNil(b.CityID) {
		t.Error("want b's foreign key value to be nil")
	}
	if !queries.IsValuerNil(c.CityID) {
		t.Error("want c's foreign key value to be nil")
	}
	if !queries.Equal(a.ID, d.CityID) {
		t.Error("foreign key was wrong value", a.ID, d.CityID)
	}
	if !queries.Equal(a.ID, e.CityID) {
		t.Error("foreign key was wrong value", a.ID, e.CityID)
	}

	if b.R.City != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.City != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.City != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}
	if e.R.City != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}

	if a.R.Centers[0] != &d {
		t.Error("relationship struct slice not set to correct value")
	}
	if a.R.Centers[1] != &e {
		t.Error("relationship struct slice not set to correct value")
	}
}

func testCityToManyRemoveOpCenters(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a City
	var b, c, d, e Center

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, cityDBTypes, false, strmangle.SetComplement(cityPrimaryKeyColumns, cityColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Center{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, centerDBTypes, false, strmangle.SetComplement(centerPrimaryKeyColumns, centerColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.AddCenters(ctx, tx, true, foreigners...)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.Centers().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 4 {
		t.Error("count was wrong:", count)
	}

	err = a.RemoveCenters(ctx, tx, foreigners[:2]...)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.Centers().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if !queries.IsValuerNil(b.CityID) {
		t.Error("want b's foreign key value to be nil")
	}
	if !queries.IsValuerNil(c.CityID) {
		t.Error("want c's foreign key value to be nil")
	}

	if b.R.City != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.City != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.City != &a {
		t.Error("relationship to a should have been preserved")
	}
	if e.R.City != &a {
		t.Error("relationship to a should have been preserved")
	}

	if len(a.R.Centers) != 2 {
		t.Error("should have preserved two relationships")
	}

	// Removal doesn't do a stable deletion for performance so we have to flip the order
	if a.R.Centers[1] != &d {
		t.Error("relationship to d should have been preserved")
	}
	if a.R.Centers[0] != &e {
		t.Error("relationship to e should have been preserved")
	}
}

func testCitiesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &City{}
	if err = randomize.Struct(seed, o, cityDBTypes, true, cityColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize City struct: %s", err)
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

func testCitiesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &City{}
	if err = randomize.Struct(seed, o, cityDBTypes, true, cityColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize City struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := CitySlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testCitiesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &City{}
	if err = randomize.Struct(seed, o, cityDBTypes, true, cityColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize City struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Cities().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	cityDBTypes = map[string]string{`CreatedAt`: `timestamp with time zone`, `ID`: `uuid`, `Name`: `character varying`, `UpdatedAt`: `timestamp with time zone`}
	_           = bytes.MinRead
)

func testCitiesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(cityPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(cityColumns) == len(cityPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &City{}
	if err = randomize.Struct(seed, o, cityDBTypes, true, cityColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize City struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Cities().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, cityDBTypes, true, cityPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize City struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testCitiesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(cityColumns) == len(cityPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &City{}
	if err = randomize.Struct(seed, o, cityDBTypes, true, cityColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize City struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Cities().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, cityDBTypes, true, cityPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize City struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(cityColumns, cityPrimaryKeyColumns) {
		fields = cityColumns
	} else {
		fields = strmangle.SetComplement(
			cityColumns,
			cityPrimaryKeyColumns,
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

	slice := CitySlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testCitiesUpsert(t *testing.T) {
	t.Parallel()

	if len(cityColumns) == len(cityPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := City{}
	if err = randomize.Struct(seed, &o, cityDBTypes, true); err != nil {
		t.Errorf("Unable to randomize City struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert City: %s", err)
	}

	count, err := Cities().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, cityDBTypes, false, cityPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize City struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert City: %s", err)
	}

	count, err = Cities().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
