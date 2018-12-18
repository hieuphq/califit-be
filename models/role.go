// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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

	"github.com/pkg/errors"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/strmangle"
)

// Role is an object representing the database table.
type Role struct {
	ID        string      `boil:"id" json:"id" toml:"id" yaml:"id"`
	Name      null.String `boil:"name" json:"name,omitempty" toml:"name" yaml:"name,omitempty"`
	CreatedAt time.Time   `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt time.Time   `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`

	R *roleR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L roleL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var RoleColumns = struct {
	ID        string
	Name      string
	CreatedAt string
	UpdatedAt string
}{
	ID:        "id",
	Name:      "name",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// RoleRels is where relationship names are stored.
var RoleRels = struct {
	RolePermissions string
	UserRoles       string
}{
	RolePermissions: "RolePermissions",
	UserRoles:       "UserRoles",
}

// roleR is where relationships are stored.
type roleR struct {
	RolePermissions RolePermissionSlice
	UserRoles       UserRoleSlice
}

// NewStruct creates a new relationship struct
func (*roleR) NewStruct() *roleR {
	return &roleR{}
}

// roleL is where Load methods for each relationship are stored.
type roleL struct{}

var (
	roleColumns               = []string{"id", "name", "created_at", "updated_at"}
	roleColumnsWithoutDefault = []string{"id", "name"}
	roleColumnsWithDefault    = []string{"created_at", "updated_at"}
	rolePrimaryKeyColumns     = []string{"id"}
)

type (
	// RoleSlice is an alias for a slice of pointers to Role.
	// This should generally be used opposed to []Role.
	RoleSlice []*Role
	// RoleHook is the signature for custom Role hook methods
	RoleHook func(context.Context, boil.ContextExecutor, *Role) error

	roleQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	roleType                 = reflect.TypeOf(&Role{})
	roleMapping              = queries.MakeStructMapping(roleType)
	rolePrimaryKeyMapping, _ = queries.BindMapping(roleType, roleMapping, rolePrimaryKeyColumns)
	roleInsertCacheMut       sync.RWMutex
	roleInsertCache          = make(map[string]insertCache)
	roleUpdateCacheMut       sync.RWMutex
	roleUpdateCache          = make(map[string]updateCache)
	roleUpsertCacheMut       sync.RWMutex
	roleUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
)

var roleBeforeInsertHooks []RoleHook
var roleBeforeUpdateHooks []RoleHook
var roleBeforeDeleteHooks []RoleHook
var roleBeforeUpsertHooks []RoleHook

var roleAfterInsertHooks []RoleHook
var roleAfterSelectHooks []RoleHook
var roleAfterUpdateHooks []RoleHook
var roleAfterDeleteHooks []RoleHook
var roleAfterUpsertHooks []RoleHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Role) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range roleBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Role) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range roleBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Role) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range roleBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Role) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range roleBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Role) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range roleAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Role) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range roleAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Role) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range roleAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Role) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range roleAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Role) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range roleAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddRoleHook registers your hook function for all future operations.
func AddRoleHook(hookPoint boil.HookPoint, roleHook RoleHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		roleBeforeInsertHooks = append(roleBeforeInsertHooks, roleHook)
	case boil.BeforeUpdateHook:
		roleBeforeUpdateHooks = append(roleBeforeUpdateHooks, roleHook)
	case boil.BeforeDeleteHook:
		roleBeforeDeleteHooks = append(roleBeforeDeleteHooks, roleHook)
	case boil.BeforeUpsertHook:
		roleBeforeUpsertHooks = append(roleBeforeUpsertHooks, roleHook)
	case boil.AfterInsertHook:
		roleAfterInsertHooks = append(roleAfterInsertHooks, roleHook)
	case boil.AfterSelectHook:
		roleAfterSelectHooks = append(roleAfterSelectHooks, roleHook)
	case boil.AfterUpdateHook:
		roleAfterUpdateHooks = append(roleAfterUpdateHooks, roleHook)
	case boil.AfterDeleteHook:
		roleAfterDeleteHooks = append(roleAfterDeleteHooks, roleHook)
	case boil.AfterUpsertHook:
		roleAfterUpsertHooks = append(roleAfterUpsertHooks, roleHook)
	}
}

// One returns a single role record from the query.
func (q roleQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Role, error) {
	o := &Role{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for role")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Role records from the query.
func (q roleQuery) All(ctx context.Context, exec boil.ContextExecutor) (RoleSlice, error) {
	var o []*Role

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Role slice")
	}

	if len(roleAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Role records in the query.
func (q roleQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count role rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q roleQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if role exists")
	}

	return count > 0, nil
}

// RolePermissions retrieves all the role_permission's RolePermissions with an executor.
func (o *Role) RolePermissions(mods ...qm.QueryMod) rolePermissionQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"role_permission\".\"role_id\"=?", o.ID),
	)

	query := RolePermissions(queryMods...)
	queries.SetFrom(query.Query, "\"role_permission\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"role_permission\".*"})
	}

	return query
}

// UserRoles retrieves all the user_role's UserRoles with an executor.
func (o *Role) UserRoles(mods ...qm.QueryMod) userRoleQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"user_role\".\"role_id\"=?", o.ID),
	)

	query := UserRoles(queryMods...)
	queries.SetFrom(query.Query, "\"user_role\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"user_role\".*"})
	}

	return query
}

// LoadRolePermissions allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (roleL) LoadRolePermissions(ctx context.Context, e boil.ContextExecutor, singular bool, maybeRole interface{}, mods queries.Applicator) error {
	var slice []*Role
	var object *Role

	if singular {
		object = maybeRole.(*Role)
	} else {
		slice = *maybeRole.(*[]*Role)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &roleR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &roleR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.ID) {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	query := NewQuery(qm.From(`role_permission`), qm.WhereIn(`role_id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load role_permission")
	}

	var resultSlice []*RolePermission
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice role_permission")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on role_permission")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for role_permission")
	}

	if len(rolePermissionAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.RolePermissions = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &rolePermissionR{}
			}
			foreign.R.Role = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if queries.Equal(local.ID, foreign.RoleID) {
				local.R.RolePermissions = append(local.R.RolePermissions, foreign)
				if foreign.R == nil {
					foreign.R = &rolePermissionR{}
				}
				foreign.R.Role = local
				break
			}
		}
	}

	return nil
}

// LoadUserRoles allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (roleL) LoadUserRoles(ctx context.Context, e boil.ContextExecutor, singular bool, maybeRole interface{}, mods queries.Applicator) error {
	var slice []*Role
	var object *Role

	if singular {
		object = maybeRole.(*Role)
	} else {
		slice = *maybeRole.(*[]*Role)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &roleR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &roleR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.ID) {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	query := NewQuery(qm.From(`user_role`), qm.WhereIn(`role_id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load user_role")
	}

	var resultSlice []*UserRole
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice user_role")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on user_role")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for user_role")
	}

	if len(userRoleAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.UserRoles = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &userRoleR{}
			}
			foreign.R.Role = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if queries.Equal(local.ID, foreign.RoleID) {
				local.R.UserRoles = append(local.R.UserRoles, foreign)
				if foreign.R == nil {
					foreign.R = &userRoleR{}
				}
				foreign.R.Role = local
				break
			}
		}
	}

	return nil
}

// AddRolePermissions adds the given related objects to the existing relationships
// of the role, optionally inserting them as new records.
// Appends related to o.R.RolePermissions.
// Sets related.R.Role appropriately.
func (o *Role) AddRolePermissions(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*RolePermission) error {
	var err error
	for _, rel := range related {
		if insert {
			queries.Assign(&rel.RoleID, o.ID)
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"role_permission\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"role_id"}),
				strmangle.WhereClause("\"", "\"", 2, rolePermissionPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.DebugMode {
				fmt.Fprintln(boil.DebugWriter, updateQuery)
				fmt.Fprintln(boil.DebugWriter, values)
			}

			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			queries.Assign(&rel.RoleID, o.ID)
		}
	}

	if o.R == nil {
		o.R = &roleR{
			RolePermissions: related,
		}
	} else {
		o.R.RolePermissions = append(o.R.RolePermissions, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &rolePermissionR{
				Role: o,
			}
		} else {
			rel.R.Role = o
		}
	}
	return nil
}

// SetRolePermissions removes all previously related items of the
// role replacing them completely with the passed
// in related items, optionally inserting them as new records.
// Sets o.R.Role's RolePermissions accordingly.
// Replaces o.R.RolePermissions with related.
// Sets related.R.Role's RolePermissions accordingly.
func (o *Role) SetRolePermissions(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*RolePermission) error {
	query := "update \"role_permission\" set \"role_id\" = null where \"role_id\" = $1"
	values := []interface{}{o.ID}
	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	_, err := exec.ExecContext(ctx, query, values...)
	if err != nil {
		return errors.Wrap(err, "failed to remove relationships before set")
	}

	if o.R != nil {
		for _, rel := range o.R.RolePermissions {
			queries.SetScanner(&rel.RoleID, nil)
			if rel.R == nil {
				continue
			}

			rel.R.Role = nil
		}

		o.R.RolePermissions = nil
	}
	return o.AddRolePermissions(ctx, exec, insert, related...)
}

// RemoveRolePermissions relationships from objects passed in.
// Removes related items from R.RolePermissions (uses pointer comparison, removal does not keep order)
// Sets related.R.Role.
func (o *Role) RemoveRolePermissions(ctx context.Context, exec boil.ContextExecutor, related ...*RolePermission) error {
	var err error
	for _, rel := range related {
		queries.SetScanner(&rel.RoleID, nil)
		if rel.R != nil {
			rel.R.Role = nil
		}
		if _, err = rel.Update(ctx, exec, boil.Whitelist("role_id")); err != nil {
			return err
		}
	}
	if o.R == nil {
		return nil
	}

	for _, rel := range related {
		for i, ri := range o.R.RolePermissions {
			if rel != ri {
				continue
			}

			ln := len(o.R.RolePermissions)
			if ln > 1 && i < ln-1 {
				o.R.RolePermissions[i] = o.R.RolePermissions[ln-1]
			}
			o.R.RolePermissions = o.R.RolePermissions[:ln-1]
			break
		}
	}

	return nil
}

// AddUserRoles adds the given related objects to the existing relationships
// of the role, optionally inserting them as new records.
// Appends related to o.R.UserRoles.
// Sets related.R.Role appropriately.
func (o *Role) AddUserRoles(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*UserRole) error {
	var err error
	for _, rel := range related {
		if insert {
			queries.Assign(&rel.RoleID, o.ID)
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"user_role\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"role_id"}),
				strmangle.WhereClause("\"", "\"", 2, userRolePrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.DebugMode {
				fmt.Fprintln(boil.DebugWriter, updateQuery)
				fmt.Fprintln(boil.DebugWriter, values)
			}

			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			queries.Assign(&rel.RoleID, o.ID)
		}
	}

	if o.R == nil {
		o.R = &roleR{
			UserRoles: related,
		}
	} else {
		o.R.UserRoles = append(o.R.UserRoles, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &userRoleR{
				Role: o,
			}
		} else {
			rel.R.Role = o
		}
	}
	return nil
}

// SetUserRoles removes all previously related items of the
// role replacing them completely with the passed
// in related items, optionally inserting them as new records.
// Sets o.R.Role's UserRoles accordingly.
// Replaces o.R.UserRoles with related.
// Sets related.R.Role's UserRoles accordingly.
func (o *Role) SetUserRoles(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*UserRole) error {
	query := "update \"user_role\" set \"role_id\" = null where \"role_id\" = $1"
	values := []interface{}{o.ID}
	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	_, err := exec.ExecContext(ctx, query, values...)
	if err != nil {
		return errors.Wrap(err, "failed to remove relationships before set")
	}

	if o.R != nil {
		for _, rel := range o.R.UserRoles {
			queries.SetScanner(&rel.RoleID, nil)
			if rel.R == nil {
				continue
			}

			rel.R.Role = nil
		}

		o.R.UserRoles = nil
	}
	return o.AddUserRoles(ctx, exec, insert, related...)
}

// RemoveUserRoles relationships from objects passed in.
// Removes related items from R.UserRoles (uses pointer comparison, removal does not keep order)
// Sets related.R.Role.
func (o *Role) RemoveUserRoles(ctx context.Context, exec boil.ContextExecutor, related ...*UserRole) error {
	var err error
	for _, rel := range related {
		queries.SetScanner(&rel.RoleID, nil)
		if rel.R != nil {
			rel.R.Role = nil
		}
		if _, err = rel.Update(ctx, exec, boil.Whitelist("role_id")); err != nil {
			return err
		}
	}
	if o.R == nil {
		return nil
	}

	for _, rel := range related {
		for i, ri := range o.R.UserRoles {
			if rel != ri {
				continue
			}

			ln := len(o.R.UserRoles)
			if ln > 1 && i < ln-1 {
				o.R.UserRoles[i] = o.R.UserRoles[ln-1]
			}
			o.R.UserRoles = o.R.UserRoles[:ln-1]
			break
		}
	}

	return nil
}

// Roles retrieves all the records using an executor.
func Roles(mods ...qm.QueryMod) roleQuery {
	mods = append(mods, qm.From("\"role\""))
	return roleQuery{NewQuery(mods...)}
}

// FindRole retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindRole(ctx context.Context, exec boil.ContextExecutor, iD string, selectCols ...string) (*Role, error) {
	roleObj := &Role{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"role\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, roleObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from role")
	}

	return roleObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Role) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no role provided for insertion")
	}

	var err error
	currTime := time.Now().In(boil.GetLocation())

	if o.CreatedAt.IsZero() {
		o.CreatedAt = currTime
	}
	if o.UpdatedAt.IsZero() {
		o.UpdatedAt = currTime
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(roleColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	roleInsertCacheMut.RLock()
	cache, cached := roleInsertCache[key]
	roleInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			roleColumns,
			roleColumnsWithDefault,
			roleColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(roleType, roleMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(roleType, roleMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"role\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"role\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into role")
	}

	if !cached {
		roleInsertCacheMut.Lock()
		roleInsertCache[key] = cache
		roleInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Role.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Role) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	currTime := time.Now().In(boil.GetLocation())

	o.UpdatedAt = currTime

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	roleUpdateCacheMut.RLock()
	cache, cached := roleUpdateCache[key]
	roleUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			roleColumns,
			rolePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update role, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"role\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, rolePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(roleType, roleMapping, append(wl, rolePrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update role row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for role")
	}

	if !cached {
		roleUpdateCacheMut.Lock()
		roleUpdateCache[key] = cache
		roleUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q roleQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for role")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for role")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o RoleSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), rolePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"role\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, rolePrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in role slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all role")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Role) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no role provided for upsert")
	}
	currTime := time.Now().In(boil.GetLocation())

	if o.CreatedAt.IsZero() {
		o.CreatedAt = currTime
	}
	o.UpdatedAt = currTime

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(roleColumnsWithDefault, o)

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

	roleUpsertCacheMut.RLock()
	cache, cached := roleUpsertCache[key]
	roleUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			roleColumns,
			roleColumnsWithDefault,
			roleColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			roleColumns,
			rolePrimaryKeyColumns,
		)

		if len(update) == 0 {
			return errors.New("models: unable to upsert role, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(rolePrimaryKeyColumns))
			copy(conflict, rolePrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"role\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(roleType, roleMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(roleType, roleMapping, ret)
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

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert role")
	}

	if !cached {
		roleUpsertCacheMut.Lock()
		roleUpsertCache[key] = cache
		roleUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Role record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Role) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Role provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), rolePrimaryKeyMapping)
	sql := "DELETE FROM \"role\" WHERE \"id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from role")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for role")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q roleQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no roleQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from role")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for role")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o RoleSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Role slice provided for delete all")
	}

	if len(o) == 0 {
		return 0, nil
	}

	if len(roleBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), rolePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"role\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, rolePrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from role slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for role")
	}

	if len(roleAfterDeleteHooks) != 0 {
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
func (o *Role) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindRole(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *RoleSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := RoleSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), rolePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"role\".* FROM \"role\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, rolePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in RoleSlice")
	}

	*o = slice

	return nil
}

// RoleExists checks if the Role row exists.
func RoleExists(ctx context.Context, exec boil.ContextExecutor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"role\" where \"id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}

	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if role exists")
	}

	return exists, nil
}