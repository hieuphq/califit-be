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

// Center is an object representing the database table.
type Center struct {
	ID        string      `boil:"id" json:"id" toml:"id" yaml:"id"`
	Name      null.String `boil:"name" json:"name,omitempty" toml:"name" yaml:"name,omitempty"`
	AddressID null.String `boil:"address_id" json:"address_id,omitempty" toml:"address_id" yaml:"address_id,omitempty"`
	CityID    null.String `boil:"city_id" json:"city_id,omitempty" toml:"city_id" yaml:"city_id,omitempty"`
	OpenAt    null.String `boil:"open_at" json:"open_at,omitempty" toml:"open_at" yaml:"open_at,omitempty"`
	CloseAt   null.String `boil:"close_at" json:"close_at,omitempty" toml:"close_at" yaml:"close_at,omitempty"`
	CreatedAt time.Time   `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt time.Time   `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`

	R *centerR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L centerL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var CenterColumns = struct {
	ID        string
	Name      string
	AddressID string
	CityID    string
	OpenAt    string
	CloseAt   string
	CreatedAt string
	UpdatedAt string
}{
	ID:        "id",
	Name:      "name",
	AddressID: "address_id",
	CityID:    "city_id",
	OpenAt:    "open_at",
	CloseAt:   "close_at",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// CenterRels is where relationship names are stored.
var CenterRels = struct {
	Address   string
	City      string
	Schedules string
}{
	Address:   "Address",
	City:      "City",
	Schedules: "Schedules",
}

// centerR is where relationships are stored.
type centerR struct {
	Address   *Address
	City      *City
	Schedules ScheduleSlice
}

// NewStruct creates a new relationship struct
func (*centerR) NewStruct() *centerR {
	return &centerR{}
}

// centerL is where Load methods for each relationship are stored.
type centerL struct{}

var (
	centerColumns               = []string{"id", "name", "address_id", "city_id", "open_at", "close_at", "created_at", "updated_at"}
	centerColumnsWithoutDefault = []string{"id", "name", "address_id", "city_id", "open_at", "close_at"}
	centerColumnsWithDefault    = []string{"created_at", "updated_at"}
	centerPrimaryKeyColumns     = []string{"id"}
)

type (
	// CenterSlice is an alias for a slice of pointers to Center.
	// This should generally be used opposed to []Center.
	CenterSlice []*Center
	// CenterHook is the signature for custom Center hook methods
	CenterHook func(context.Context, boil.ContextExecutor, *Center) error

	centerQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	centerType                 = reflect.TypeOf(&Center{})
	centerMapping              = queries.MakeStructMapping(centerType)
	centerPrimaryKeyMapping, _ = queries.BindMapping(centerType, centerMapping, centerPrimaryKeyColumns)
	centerInsertCacheMut       sync.RWMutex
	centerInsertCache          = make(map[string]insertCache)
	centerUpdateCacheMut       sync.RWMutex
	centerUpdateCache          = make(map[string]updateCache)
	centerUpsertCacheMut       sync.RWMutex
	centerUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
)

var centerBeforeInsertHooks []CenterHook
var centerBeforeUpdateHooks []CenterHook
var centerBeforeDeleteHooks []CenterHook
var centerBeforeUpsertHooks []CenterHook

var centerAfterInsertHooks []CenterHook
var centerAfterSelectHooks []CenterHook
var centerAfterUpdateHooks []CenterHook
var centerAfterDeleteHooks []CenterHook
var centerAfterUpsertHooks []CenterHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Center) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range centerBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Center) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range centerBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Center) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range centerBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Center) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range centerBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Center) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range centerAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Center) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range centerAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Center) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range centerAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Center) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range centerAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Center) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range centerAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddCenterHook registers your hook function for all future operations.
func AddCenterHook(hookPoint boil.HookPoint, centerHook CenterHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		centerBeforeInsertHooks = append(centerBeforeInsertHooks, centerHook)
	case boil.BeforeUpdateHook:
		centerBeforeUpdateHooks = append(centerBeforeUpdateHooks, centerHook)
	case boil.BeforeDeleteHook:
		centerBeforeDeleteHooks = append(centerBeforeDeleteHooks, centerHook)
	case boil.BeforeUpsertHook:
		centerBeforeUpsertHooks = append(centerBeforeUpsertHooks, centerHook)
	case boil.AfterInsertHook:
		centerAfterInsertHooks = append(centerAfterInsertHooks, centerHook)
	case boil.AfterSelectHook:
		centerAfterSelectHooks = append(centerAfterSelectHooks, centerHook)
	case boil.AfterUpdateHook:
		centerAfterUpdateHooks = append(centerAfterUpdateHooks, centerHook)
	case boil.AfterDeleteHook:
		centerAfterDeleteHooks = append(centerAfterDeleteHooks, centerHook)
	case boil.AfterUpsertHook:
		centerAfterUpsertHooks = append(centerAfterUpsertHooks, centerHook)
	}
}

// One returns a single center record from the query.
func (q centerQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Center, error) {
	o := &Center{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for center")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Center records from the query.
func (q centerQuery) All(ctx context.Context, exec boil.ContextExecutor) (CenterSlice, error) {
	var o []*Center

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Center slice")
	}

	if len(centerAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Center records in the query.
func (q centerQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count center rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q centerQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if center exists")
	}

	return count > 0, nil
}

// Address pointed to by the foreign key.
func (o *Center) Address(mods ...qm.QueryMod) addressQuery {
	queryMods := []qm.QueryMod{
		qm.Where("id=?", o.AddressID),
	}

	queryMods = append(queryMods, mods...)

	query := Addresses(queryMods...)
	queries.SetFrom(query.Query, "\"address\"")

	return query
}

// City pointed to by the foreign key.
func (o *Center) City(mods ...qm.QueryMod) cityQuery {
	queryMods := []qm.QueryMod{
		qm.Where("id=?", o.CityID),
	}

	queryMods = append(queryMods, mods...)

	query := Cities(queryMods...)
	queries.SetFrom(query.Query, "\"city\"")

	return query
}

// Schedules retrieves all the schedule's Schedules with an executor.
func (o *Center) Schedules(mods ...qm.QueryMod) scheduleQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"schedule\".\"center_id\"=?", o.ID),
	)

	query := Schedules(queryMods...)
	queries.SetFrom(query.Query, "\"schedule\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"schedule\".*"})
	}

	return query
}

// LoadAddress allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (centerL) LoadAddress(ctx context.Context, e boil.ContextExecutor, singular bool, maybeCenter interface{}, mods queries.Applicator) error {
	var slice []*Center
	var object *Center

	if singular {
		object = maybeCenter.(*Center)
	} else {
		slice = *maybeCenter.(*[]*Center)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &centerR{}
		}
		if !queries.IsNil(object.AddressID) {
			args = append(args, object.AddressID)
		}

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &centerR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.AddressID) {
					continue Outer
				}
			}

			if !queries.IsNil(obj.AddressID) {
				args = append(args, obj.AddressID)
			}

		}
	}

	query := NewQuery(qm.From(`address`), qm.WhereIn(`id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Address")
	}

	var resultSlice []*Address
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Address")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for address")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for address")
	}

	if len(centerAfterSelectHooks) != 0 {
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
		object.R.Address = foreign
		if foreign.R == nil {
			foreign.R = &addressR{}
		}
		foreign.R.Centers = append(foreign.R.Centers, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if queries.Equal(local.AddressID, foreign.ID) {
				local.R.Address = foreign
				if foreign.R == nil {
					foreign.R = &addressR{}
				}
				foreign.R.Centers = append(foreign.R.Centers, local)
				break
			}
		}
	}

	return nil
}

// LoadCity allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (centerL) LoadCity(ctx context.Context, e boil.ContextExecutor, singular bool, maybeCenter interface{}, mods queries.Applicator) error {
	var slice []*Center
	var object *Center

	if singular {
		object = maybeCenter.(*Center)
	} else {
		slice = *maybeCenter.(*[]*Center)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &centerR{}
		}
		if !queries.IsNil(object.CityID) {
			args = append(args, object.CityID)
		}

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &centerR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.CityID) {
					continue Outer
				}
			}

			if !queries.IsNil(obj.CityID) {
				args = append(args, obj.CityID)
			}

		}
	}

	query := NewQuery(qm.From(`city`), qm.WhereIn(`id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load City")
	}

	var resultSlice []*City
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice City")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for city")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for city")
	}

	if len(centerAfterSelectHooks) != 0 {
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
		object.R.City = foreign
		if foreign.R == nil {
			foreign.R = &cityR{}
		}
		foreign.R.Centers = append(foreign.R.Centers, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if queries.Equal(local.CityID, foreign.ID) {
				local.R.City = foreign
				if foreign.R == nil {
					foreign.R = &cityR{}
				}
				foreign.R.Centers = append(foreign.R.Centers, local)
				break
			}
		}
	}

	return nil
}

// LoadSchedules allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (centerL) LoadSchedules(ctx context.Context, e boil.ContextExecutor, singular bool, maybeCenter interface{}, mods queries.Applicator) error {
	var slice []*Center
	var object *Center

	if singular {
		object = maybeCenter.(*Center)
	} else {
		slice = *maybeCenter.(*[]*Center)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &centerR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &centerR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.ID) {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	query := NewQuery(qm.From(`schedule`), qm.WhereIn(`center_id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load schedule")
	}

	var resultSlice []*Schedule
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice schedule")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on schedule")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for schedule")
	}

	if len(scheduleAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.Schedules = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &scheduleR{}
			}
			foreign.R.Center = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if queries.Equal(local.ID, foreign.CenterID) {
				local.R.Schedules = append(local.R.Schedules, foreign)
				if foreign.R == nil {
					foreign.R = &scheduleR{}
				}
				foreign.R.Center = local
				break
			}
		}
	}

	return nil
}

// SetAddress of the center to the related item.
// Sets o.R.Address to related.
// Adds o to related.R.Centers.
func (o *Center) SetAddress(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Address) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"center\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"address_id"}),
		strmangle.WhereClause("\"", "\"", 2, centerPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	queries.Assign(&o.AddressID, related.ID)
	if o.R == nil {
		o.R = &centerR{
			Address: related,
		}
	} else {
		o.R.Address = related
	}

	if related.R == nil {
		related.R = &addressR{
			Centers: CenterSlice{o},
		}
	} else {
		related.R.Centers = append(related.R.Centers, o)
	}

	return nil
}

// RemoveAddress relationship.
// Sets o.R.Address to nil.
// Removes o from all passed in related items' relationships struct (Optional).
func (o *Center) RemoveAddress(ctx context.Context, exec boil.ContextExecutor, related *Address) error {
	var err error

	queries.SetScanner(&o.AddressID, nil)
	if _, err = o.Update(ctx, exec, boil.Whitelist("address_id")); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.R.Address = nil
	if related == nil || related.R == nil {
		return nil
	}

	for i, ri := range related.R.Centers {
		if queries.Equal(o.AddressID, ri.AddressID) {
			continue
		}

		ln := len(related.R.Centers)
		if ln > 1 && i < ln-1 {
			related.R.Centers[i] = related.R.Centers[ln-1]
		}
		related.R.Centers = related.R.Centers[:ln-1]
		break
	}
	return nil
}

// SetCity of the center to the related item.
// Sets o.R.City to related.
// Adds o to related.R.Centers.
func (o *Center) SetCity(ctx context.Context, exec boil.ContextExecutor, insert bool, related *City) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"center\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"city_id"}),
		strmangle.WhereClause("\"", "\"", 2, centerPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	queries.Assign(&o.CityID, related.ID)
	if o.R == nil {
		o.R = &centerR{
			City: related,
		}
	} else {
		o.R.City = related
	}

	if related.R == nil {
		related.R = &cityR{
			Centers: CenterSlice{o},
		}
	} else {
		related.R.Centers = append(related.R.Centers, o)
	}

	return nil
}

// RemoveCity relationship.
// Sets o.R.City to nil.
// Removes o from all passed in related items' relationships struct (Optional).
func (o *Center) RemoveCity(ctx context.Context, exec boil.ContextExecutor, related *City) error {
	var err error

	queries.SetScanner(&o.CityID, nil)
	if _, err = o.Update(ctx, exec, boil.Whitelist("city_id")); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.R.City = nil
	if related == nil || related.R == nil {
		return nil
	}

	for i, ri := range related.R.Centers {
		if queries.Equal(o.CityID, ri.CityID) {
			continue
		}

		ln := len(related.R.Centers)
		if ln > 1 && i < ln-1 {
			related.R.Centers[i] = related.R.Centers[ln-1]
		}
		related.R.Centers = related.R.Centers[:ln-1]
		break
	}
	return nil
}

// AddSchedules adds the given related objects to the existing relationships
// of the center, optionally inserting them as new records.
// Appends related to o.R.Schedules.
// Sets related.R.Center appropriately.
func (o *Center) AddSchedules(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Schedule) error {
	var err error
	for _, rel := range related {
		if insert {
			queries.Assign(&rel.CenterID, o.ID)
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"schedule\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"center_id"}),
				strmangle.WhereClause("\"", "\"", 2, schedulePrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.DebugMode {
				fmt.Fprintln(boil.DebugWriter, updateQuery)
				fmt.Fprintln(boil.DebugWriter, values)
			}

			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			queries.Assign(&rel.CenterID, o.ID)
		}
	}

	if o.R == nil {
		o.R = &centerR{
			Schedules: related,
		}
	} else {
		o.R.Schedules = append(o.R.Schedules, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &scheduleR{
				Center: o,
			}
		} else {
			rel.R.Center = o
		}
	}
	return nil
}

// SetSchedules removes all previously related items of the
// center replacing them completely with the passed
// in related items, optionally inserting them as new records.
// Sets o.R.Center's Schedules accordingly.
// Replaces o.R.Schedules with related.
// Sets related.R.Center's Schedules accordingly.
func (o *Center) SetSchedules(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Schedule) error {
	query := "update \"schedule\" set \"center_id\" = null where \"center_id\" = $1"
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
		for _, rel := range o.R.Schedules {
			queries.SetScanner(&rel.CenterID, nil)
			if rel.R == nil {
				continue
			}

			rel.R.Center = nil
		}

		o.R.Schedules = nil
	}
	return o.AddSchedules(ctx, exec, insert, related...)
}

// RemoveSchedules relationships from objects passed in.
// Removes related items from R.Schedules (uses pointer comparison, removal does not keep order)
// Sets related.R.Center.
func (o *Center) RemoveSchedules(ctx context.Context, exec boil.ContextExecutor, related ...*Schedule) error {
	var err error
	for _, rel := range related {
		queries.SetScanner(&rel.CenterID, nil)
		if rel.R != nil {
			rel.R.Center = nil
		}
		if _, err = rel.Update(ctx, exec, boil.Whitelist("center_id")); err != nil {
			return err
		}
	}
	if o.R == nil {
		return nil
	}

	for _, rel := range related {
		for i, ri := range o.R.Schedules {
			if rel != ri {
				continue
			}

			ln := len(o.R.Schedules)
			if ln > 1 && i < ln-1 {
				o.R.Schedules[i] = o.R.Schedules[ln-1]
			}
			o.R.Schedules = o.R.Schedules[:ln-1]
			break
		}
	}

	return nil
}

// Centers retrieves all the records using an executor.
func Centers(mods ...qm.QueryMod) centerQuery {
	mods = append(mods, qm.From("\"center\""))
	return centerQuery{NewQuery(mods...)}
}

// FindCenter retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindCenter(ctx context.Context, exec boil.ContextExecutor, iD string, selectCols ...string) (*Center, error) {
	centerObj := &Center{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"center\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, centerObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from center")
	}

	return centerObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Center) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no center provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(centerColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	centerInsertCacheMut.RLock()
	cache, cached := centerInsertCache[key]
	centerInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			centerColumns,
			centerColumnsWithDefault,
			centerColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(centerType, centerMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(centerType, centerMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"center\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"center\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into center")
	}

	if !cached {
		centerInsertCacheMut.Lock()
		centerInsertCache[key] = cache
		centerInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Center.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Center) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	currTime := time.Now().In(boil.GetLocation())

	o.UpdatedAt = currTime

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	centerUpdateCacheMut.RLock()
	cache, cached := centerUpdateCache[key]
	centerUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			centerColumns,
			centerPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update center, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"center\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, centerPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(centerType, centerMapping, append(wl, centerPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update center row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for center")
	}

	if !cached {
		centerUpdateCacheMut.Lock()
		centerUpdateCache[key] = cache
		centerUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q centerQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for center")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for center")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o CenterSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), centerPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"center\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, centerPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in center slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all center")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Center) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no center provided for upsert")
	}
	currTime := time.Now().In(boil.GetLocation())

	if o.CreatedAt.IsZero() {
		o.CreatedAt = currTime
	}
	o.UpdatedAt = currTime

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(centerColumnsWithDefault, o)

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

	centerUpsertCacheMut.RLock()
	cache, cached := centerUpsertCache[key]
	centerUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			centerColumns,
			centerColumnsWithDefault,
			centerColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			centerColumns,
			centerPrimaryKeyColumns,
		)

		if len(update) == 0 {
			return errors.New("models: unable to upsert center, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(centerPrimaryKeyColumns))
			copy(conflict, centerPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"center\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(centerType, centerMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(centerType, centerMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert center")
	}

	if !cached {
		centerUpsertCacheMut.Lock()
		centerUpsertCache[key] = cache
		centerUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Center record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Center) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Center provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), centerPrimaryKeyMapping)
	sql := "DELETE FROM \"center\" WHERE \"id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from center")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for center")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q centerQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no centerQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from center")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for center")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o CenterSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Center slice provided for delete all")
	}

	if len(o) == 0 {
		return 0, nil
	}

	if len(centerBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), centerPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"center\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, centerPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from center slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for center")
	}

	if len(centerAfterDeleteHooks) != 0 {
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
func (o *Center) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindCenter(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *CenterSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := CenterSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), centerPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"center\".* FROM \"center\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, centerPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in CenterSlice")
	}

	*o = slice

	return nil
}

// CenterExists checks if the Center row exists.
func CenterExists(ctx context.Context, exec boil.ContextExecutor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"center\" where \"id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}

	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if center exists")
	}

	return exists, nil
}
