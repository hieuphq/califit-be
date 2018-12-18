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

func testRoles(t *testing.T) {
	t.Parallel()

	query := Roles()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testRolesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Role{}
	if err = randomize.Struct(seed, o, roleDBTypes, true, roleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Role struct: %s", err)
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

	count, err := Roles().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testRolesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Role{}
	if err = randomize.Struct(seed, o, roleDBTypes, true, roleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Role struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Roles().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Roles().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testRolesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Role{}
	if err = randomize.Struct(seed, o, roleDBTypes, true, roleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Role struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := RoleSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Roles().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testRolesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Role{}
	if err = randomize.Struct(seed, o, roleDBTypes, true, roleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Role struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := RoleExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Role exists: %s", err)
	}
	if !e {
		t.Errorf("Expected RoleExists to return true, but got false.")
	}
}

func testRolesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Role{}
	if err = randomize.Struct(seed, o, roleDBTypes, true, roleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Role struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	roleFound, err := FindRole(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if roleFound == nil {
		t.Error("want a record, got nil")
	}
}

func testRolesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Role{}
	if err = randomize.Struct(seed, o, roleDBTypes, true, roleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Role struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Roles().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testRolesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Role{}
	if err = randomize.Struct(seed, o, roleDBTypes, true, roleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Role struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Roles().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testRolesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	roleOne := &Role{}
	roleTwo := &Role{}
	if err = randomize.Struct(seed, roleOne, roleDBTypes, false, roleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Role struct: %s", err)
	}
	if err = randomize.Struct(seed, roleTwo, roleDBTypes, false, roleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Role struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = roleOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = roleTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Roles().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testRolesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	roleOne := &Role{}
	roleTwo := &Role{}
	if err = randomize.Struct(seed, roleOne, roleDBTypes, false, roleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Role struct: %s", err)
	}
	if err = randomize.Struct(seed, roleTwo, roleDBTypes, false, roleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Role struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = roleOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = roleTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Roles().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func roleBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Role) error {
	*o = Role{}
	return nil
}

func roleAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Role) error {
	*o = Role{}
	return nil
}

func roleAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Role) error {
	*o = Role{}
	return nil
}

func roleBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Role) error {
	*o = Role{}
	return nil
}

func roleAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Role) error {
	*o = Role{}
	return nil
}

func roleBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Role) error {
	*o = Role{}
	return nil
}

func roleAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Role) error {
	*o = Role{}
	return nil
}

func roleBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Role) error {
	*o = Role{}
	return nil
}

func roleAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Role) error {
	*o = Role{}
	return nil
}

func testRolesHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Role{}
	o := &Role{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, roleDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Role object: %s", err)
	}

	AddRoleHook(boil.BeforeInsertHook, roleBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	roleBeforeInsertHooks = []RoleHook{}

	AddRoleHook(boil.AfterInsertHook, roleAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	roleAfterInsertHooks = []RoleHook{}

	AddRoleHook(boil.AfterSelectHook, roleAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	roleAfterSelectHooks = []RoleHook{}

	AddRoleHook(boil.BeforeUpdateHook, roleBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	roleBeforeUpdateHooks = []RoleHook{}

	AddRoleHook(boil.AfterUpdateHook, roleAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	roleAfterUpdateHooks = []RoleHook{}

	AddRoleHook(boil.BeforeDeleteHook, roleBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	roleBeforeDeleteHooks = []RoleHook{}

	AddRoleHook(boil.AfterDeleteHook, roleAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	roleAfterDeleteHooks = []RoleHook{}

	AddRoleHook(boil.BeforeUpsertHook, roleBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	roleBeforeUpsertHooks = []RoleHook{}

	AddRoleHook(boil.AfterUpsertHook, roleAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	roleAfterUpsertHooks = []RoleHook{}
}

func testRolesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Role{}
	if err = randomize.Struct(seed, o, roleDBTypes, true, roleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Role struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Roles().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testRolesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Role{}
	if err = randomize.Struct(seed, o, roleDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Role struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(roleColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Roles().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testRoleToManyRolePermissions(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Role
	var b, c RolePermission

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, roleDBTypes, true, roleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Role struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, rolePermissionDBTypes, false, rolePermissionColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, rolePermissionDBTypes, false, rolePermissionColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&b.RoleID, a.ID)
	queries.Assign(&c.RoleID, a.ID)
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	rolePermission, err := a.RolePermissions().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range rolePermission {
		if queries.Equal(v.RoleID, b.RoleID) {
			bFound = true
		}
		if queries.Equal(v.RoleID, c.RoleID) {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := RoleSlice{&a}
	if err = a.L.LoadRolePermissions(ctx, tx, false, (*[]*Role)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.RolePermissions); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.RolePermissions = nil
	if err = a.L.LoadRolePermissions(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.RolePermissions); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", rolePermission)
	}
}

func testRoleToManyUserRoles(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Role
	var b, c UserRole

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, roleDBTypes, true, roleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Role struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, userRoleDBTypes, false, userRoleColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, userRoleDBTypes, false, userRoleColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&b.RoleID, a.ID)
	queries.Assign(&c.RoleID, a.ID)
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	userRole, err := a.UserRoles().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range userRole {
		if queries.Equal(v.RoleID, b.RoleID) {
			bFound = true
		}
		if queries.Equal(v.RoleID, c.RoleID) {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := RoleSlice{&a}
	if err = a.L.LoadUserRoles(ctx, tx, false, (*[]*Role)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.UserRoles); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.UserRoles = nil
	if err = a.L.LoadUserRoles(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.UserRoles); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", userRole)
	}
}

func testRoleToManyAddOpRolePermissions(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Role
	var b, c, d, e RolePermission

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, roleDBTypes, false, strmangle.SetComplement(rolePrimaryKeyColumns, roleColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*RolePermission{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, rolePermissionDBTypes, false, strmangle.SetComplement(rolePermissionPrimaryKeyColumns, rolePermissionColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*RolePermission{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddRolePermissions(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if !queries.Equal(a.ID, first.RoleID) {
			t.Error("foreign key was wrong value", a.ID, first.RoleID)
		}
		if !queries.Equal(a.ID, second.RoleID) {
			t.Error("foreign key was wrong value", a.ID, second.RoleID)
		}

		if first.R.Role != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Role != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.RolePermissions[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.RolePermissions[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.RolePermissions().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testRoleToManySetOpRolePermissions(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Role
	var b, c, d, e RolePermission

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, roleDBTypes, false, strmangle.SetComplement(rolePrimaryKeyColumns, roleColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*RolePermission{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, rolePermissionDBTypes, false, strmangle.SetComplement(rolePermissionPrimaryKeyColumns, rolePermissionColumnsWithoutDefault)...); err != nil {
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

	err = a.SetRolePermissions(ctx, tx, false, &b, &c)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.RolePermissions().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	err = a.SetRolePermissions(ctx, tx, true, &d, &e)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.RolePermissions().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if !queries.IsValuerNil(b.RoleID) {
		t.Error("want b's foreign key value to be nil")
	}
	if !queries.IsValuerNil(c.RoleID) {
		t.Error("want c's foreign key value to be nil")
	}
	if !queries.Equal(a.ID, d.RoleID) {
		t.Error("foreign key was wrong value", a.ID, d.RoleID)
	}
	if !queries.Equal(a.ID, e.RoleID) {
		t.Error("foreign key was wrong value", a.ID, e.RoleID)
	}

	if b.R.Role != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.Role != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.Role != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}
	if e.R.Role != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}

	if a.R.RolePermissions[0] != &d {
		t.Error("relationship struct slice not set to correct value")
	}
	if a.R.RolePermissions[1] != &e {
		t.Error("relationship struct slice not set to correct value")
	}
}

func testRoleToManyRemoveOpRolePermissions(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Role
	var b, c, d, e RolePermission

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, roleDBTypes, false, strmangle.SetComplement(rolePrimaryKeyColumns, roleColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*RolePermission{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, rolePermissionDBTypes, false, strmangle.SetComplement(rolePermissionPrimaryKeyColumns, rolePermissionColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.AddRolePermissions(ctx, tx, true, foreigners...)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.RolePermissions().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 4 {
		t.Error("count was wrong:", count)
	}

	err = a.RemoveRolePermissions(ctx, tx, foreigners[:2]...)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.RolePermissions().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if !queries.IsValuerNil(b.RoleID) {
		t.Error("want b's foreign key value to be nil")
	}
	if !queries.IsValuerNil(c.RoleID) {
		t.Error("want c's foreign key value to be nil")
	}

	if b.R.Role != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.Role != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.Role != &a {
		t.Error("relationship to a should have been preserved")
	}
	if e.R.Role != &a {
		t.Error("relationship to a should have been preserved")
	}

	if len(a.R.RolePermissions) != 2 {
		t.Error("should have preserved two relationships")
	}

	// Removal doesn't do a stable deletion for performance so we have to flip the order
	if a.R.RolePermissions[1] != &d {
		t.Error("relationship to d should have been preserved")
	}
	if a.R.RolePermissions[0] != &e {
		t.Error("relationship to e should have been preserved")
	}
}

func testRoleToManyAddOpUserRoles(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Role
	var b, c, d, e UserRole

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, roleDBTypes, false, strmangle.SetComplement(rolePrimaryKeyColumns, roleColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*UserRole{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, userRoleDBTypes, false, strmangle.SetComplement(userRolePrimaryKeyColumns, userRoleColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*UserRole{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddUserRoles(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if !queries.Equal(a.ID, first.RoleID) {
			t.Error("foreign key was wrong value", a.ID, first.RoleID)
		}
		if !queries.Equal(a.ID, second.RoleID) {
			t.Error("foreign key was wrong value", a.ID, second.RoleID)
		}

		if first.R.Role != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Role != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.UserRoles[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.UserRoles[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.UserRoles().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testRoleToManySetOpUserRoles(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Role
	var b, c, d, e UserRole

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, roleDBTypes, false, strmangle.SetComplement(rolePrimaryKeyColumns, roleColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*UserRole{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, userRoleDBTypes, false, strmangle.SetComplement(userRolePrimaryKeyColumns, userRoleColumnsWithoutDefault)...); err != nil {
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

	err = a.SetUserRoles(ctx, tx, false, &b, &c)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.UserRoles().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	err = a.SetUserRoles(ctx, tx, true, &d, &e)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.UserRoles().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if !queries.IsValuerNil(b.RoleID) {
		t.Error("want b's foreign key value to be nil")
	}
	if !queries.IsValuerNil(c.RoleID) {
		t.Error("want c's foreign key value to be nil")
	}
	if !queries.Equal(a.ID, d.RoleID) {
		t.Error("foreign key was wrong value", a.ID, d.RoleID)
	}
	if !queries.Equal(a.ID, e.RoleID) {
		t.Error("foreign key was wrong value", a.ID, e.RoleID)
	}

	if b.R.Role != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.Role != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.Role != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}
	if e.R.Role != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}

	if a.R.UserRoles[0] != &d {
		t.Error("relationship struct slice not set to correct value")
	}
	if a.R.UserRoles[1] != &e {
		t.Error("relationship struct slice not set to correct value")
	}
}

func testRoleToManyRemoveOpUserRoles(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Role
	var b, c, d, e UserRole

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, roleDBTypes, false, strmangle.SetComplement(rolePrimaryKeyColumns, roleColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*UserRole{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, userRoleDBTypes, false, strmangle.SetComplement(userRolePrimaryKeyColumns, userRoleColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.AddUserRoles(ctx, tx, true, foreigners...)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.UserRoles().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 4 {
		t.Error("count was wrong:", count)
	}

	err = a.RemoveUserRoles(ctx, tx, foreigners[:2]...)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.UserRoles().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if !queries.IsValuerNil(b.RoleID) {
		t.Error("want b's foreign key value to be nil")
	}
	if !queries.IsValuerNil(c.RoleID) {
		t.Error("want c's foreign key value to be nil")
	}

	if b.R.Role != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.Role != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.Role != &a {
		t.Error("relationship to a should have been preserved")
	}
	if e.R.Role != &a {
		t.Error("relationship to a should have been preserved")
	}

	if len(a.R.UserRoles) != 2 {
		t.Error("should have preserved two relationships")
	}

	// Removal doesn't do a stable deletion for performance so we have to flip the order
	if a.R.UserRoles[1] != &d {
		t.Error("relationship to d should have been preserved")
	}
	if a.R.UserRoles[0] != &e {
		t.Error("relationship to e should have been preserved")
	}
}

func testRolesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Role{}
	if err = randomize.Struct(seed, o, roleDBTypes, true, roleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Role struct: %s", err)
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

func testRolesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Role{}
	if err = randomize.Struct(seed, o, roleDBTypes, true, roleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Role struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := RoleSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testRolesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Role{}
	if err = randomize.Struct(seed, o, roleDBTypes, true, roleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Role struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Roles().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	roleDBTypes = map[string]string{`CreatedAt`: `timestamp with time zone`, `ID`: `uuid`, `Name`: `character varying`, `UpdatedAt`: `timestamp with time zone`}
	_           = bytes.MinRead
)

func testRolesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(rolePrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(roleColumns) == len(rolePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Role{}
	if err = randomize.Struct(seed, o, roleDBTypes, true, roleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Role struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Roles().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, roleDBTypes, true, rolePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Role struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testRolesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(roleColumns) == len(rolePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Role{}
	if err = randomize.Struct(seed, o, roleDBTypes, true, roleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Role struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Roles().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, roleDBTypes, true, rolePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Role struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(roleColumns, rolePrimaryKeyColumns) {
		fields = roleColumns
	} else {
		fields = strmangle.SetComplement(
			roleColumns,
			rolePrimaryKeyColumns,
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

	slice := RoleSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testRolesUpsert(t *testing.T) {
	t.Parallel()

	if len(roleColumns) == len(rolePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Role{}
	if err = randomize.Struct(seed, &o, roleDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Role struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Role: %s", err)
	}

	count, err := Roles().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, roleDBTypes, false, rolePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Role struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Role: %s", err)
	}

	count, err = Roles().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
