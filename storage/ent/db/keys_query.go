// Code generated by entc, DO NOT EDIT.

package db

import (
	"context"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/concourse/dex/storage/ent/db/keys"
	"github.com/concourse/dex/storage/ent/db/predicate"
)

// KeysQuery is the builder for querying Keys entities.
type KeysQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.Keys
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the KeysQuery builder.
func (kq *KeysQuery) Where(ps ...predicate.Keys) *KeysQuery {
	kq.predicates = append(kq.predicates, ps...)
	return kq
}

// Limit adds a limit step to the query.
func (kq *KeysQuery) Limit(limit int) *KeysQuery {
	kq.limit = &limit
	return kq
}

// Offset adds an offset step to the query.
func (kq *KeysQuery) Offset(offset int) *KeysQuery {
	kq.offset = &offset
	return kq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (kq *KeysQuery) Unique(unique bool) *KeysQuery {
	kq.unique = &unique
	return kq
}

// Order adds an order step to the query.
func (kq *KeysQuery) Order(o ...OrderFunc) *KeysQuery {
	kq.order = append(kq.order, o...)
	return kq
}

// First returns the first Keys entity from the query.
// Returns a *NotFoundError when no Keys was found.
func (kq *KeysQuery) First(ctx context.Context) (*Keys, error) {
	nodes, err := kq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{keys.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (kq *KeysQuery) FirstX(ctx context.Context) *Keys {
	node, err := kq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Keys ID from the query.
// Returns a *NotFoundError when no Keys ID was found.
func (kq *KeysQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = kq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{keys.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (kq *KeysQuery) FirstIDX(ctx context.Context) string {
	id, err := kq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Keys entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one Keys entity is not found.
// Returns a *NotFoundError when no Keys entities are found.
func (kq *KeysQuery) Only(ctx context.Context) (*Keys, error) {
	nodes, err := kq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{keys.Label}
	default:
		return nil, &NotSingularError{keys.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (kq *KeysQuery) OnlyX(ctx context.Context) *Keys {
	node, err := kq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Keys ID in the query.
// Returns a *NotSingularError when exactly one Keys ID is not found.
// Returns a *NotFoundError when no entities are found.
func (kq *KeysQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = kq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{keys.Label}
	default:
		err = &NotSingularError{keys.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (kq *KeysQuery) OnlyIDX(ctx context.Context) string {
	id, err := kq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of KeysSlice.
func (kq *KeysQuery) All(ctx context.Context) ([]*Keys, error) {
	if err := kq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return kq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (kq *KeysQuery) AllX(ctx context.Context) []*Keys {
	nodes, err := kq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Keys IDs.
func (kq *KeysQuery) IDs(ctx context.Context) ([]string, error) {
	var ids []string
	if err := kq.Select(keys.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (kq *KeysQuery) IDsX(ctx context.Context) []string {
	ids, err := kq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (kq *KeysQuery) Count(ctx context.Context) (int, error) {
	if err := kq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return kq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (kq *KeysQuery) CountX(ctx context.Context) int {
	count, err := kq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (kq *KeysQuery) Exist(ctx context.Context) (bool, error) {
	if err := kq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return kq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (kq *KeysQuery) ExistX(ctx context.Context) bool {
	exist, err := kq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the KeysQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (kq *KeysQuery) Clone() *KeysQuery {
	if kq == nil {
		return nil
	}
	return &KeysQuery{
		config:     kq.config,
		limit:      kq.limit,
		offset:     kq.offset,
		order:      append([]OrderFunc{}, kq.order...),
		predicates: append([]predicate.Keys{}, kq.predicates...),
		// clone intermediate query.
		sql:  kq.sql.Clone(),
		path: kq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		VerificationKeys []storage.VerificationKey `json:"verification_keys,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Keys.Query().
//		GroupBy(keys.FieldVerificationKeys).
//		Aggregate(db.Count()).
//		Scan(ctx, &v)
//
func (kq *KeysQuery) GroupBy(field string, fields ...string) *KeysGroupBy {
	group := &KeysGroupBy{config: kq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := kq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return kq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		VerificationKeys []storage.VerificationKey `json:"verification_keys,omitempty"`
//	}
//
//	client.Keys.Query().
//		Select(keys.FieldVerificationKeys).
//		Scan(ctx, &v)
//
func (kq *KeysQuery) Select(fields ...string) *KeysSelect {
	kq.fields = append(kq.fields, fields...)
	return &KeysSelect{KeysQuery: kq}
}

func (kq *KeysQuery) prepareQuery(ctx context.Context) error {
	for _, f := range kq.fields {
		if !keys.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("db: invalid field %q for query", f)}
		}
	}
	if kq.path != nil {
		prev, err := kq.path(ctx)
		if err != nil {
			return err
		}
		kq.sql = prev
	}
	return nil
}

func (kq *KeysQuery) sqlAll(ctx context.Context) ([]*Keys, error) {
	var (
		nodes = []*Keys{}
		_spec = kq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &Keys{config: kq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("db: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, kq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (kq *KeysQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := kq.querySpec()
	return sqlgraph.CountNodes(ctx, kq.driver, _spec)
}

func (kq *KeysQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := kq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("db: check existence: %w", err)
	}
	return n > 0, nil
}

func (kq *KeysQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   keys.Table,
			Columns: keys.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: keys.FieldID,
			},
		},
		From:   kq.sql,
		Unique: true,
	}
	if unique := kq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := kq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, keys.FieldID)
		for i := range fields {
			if fields[i] != keys.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := kq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := kq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := kq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := kq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (kq *KeysQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(kq.driver.Dialect())
	t1 := builder.Table(keys.Table)
	columns := kq.fields
	if len(columns) == 0 {
		columns = keys.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if kq.sql != nil {
		selector = kq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	for _, p := range kq.predicates {
		p(selector)
	}
	for _, p := range kq.order {
		p(selector)
	}
	if offset := kq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := kq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// KeysGroupBy is the group-by builder for Keys entities.
type KeysGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (kgb *KeysGroupBy) Aggregate(fns ...AggregateFunc) *KeysGroupBy {
	kgb.fns = append(kgb.fns, fns...)
	return kgb
}

// Scan applies the group-by query and scans the result into the given value.
func (kgb *KeysGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := kgb.path(ctx)
	if err != nil {
		return err
	}
	kgb.sql = query
	return kgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (kgb *KeysGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := kgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (kgb *KeysGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(kgb.fields) > 1 {
		return nil, errors.New("db: KeysGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := kgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (kgb *KeysGroupBy) StringsX(ctx context.Context) []string {
	v, err := kgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (kgb *KeysGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = kgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{keys.Label}
	default:
		err = fmt.Errorf("db: KeysGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (kgb *KeysGroupBy) StringX(ctx context.Context) string {
	v, err := kgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (kgb *KeysGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(kgb.fields) > 1 {
		return nil, errors.New("db: KeysGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := kgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (kgb *KeysGroupBy) IntsX(ctx context.Context) []int {
	v, err := kgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (kgb *KeysGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = kgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{keys.Label}
	default:
		err = fmt.Errorf("db: KeysGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (kgb *KeysGroupBy) IntX(ctx context.Context) int {
	v, err := kgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (kgb *KeysGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(kgb.fields) > 1 {
		return nil, errors.New("db: KeysGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := kgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (kgb *KeysGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := kgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (kgb *KeysGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = kgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{keys.Label}
	default:
		err = fmt.Errorf("db: KeysGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (kgb *KeysGroupBy) Float64X(ctx context.Context) float64 {
	v, err := kgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (kgb *KeysGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(kgb.fields) > 1 {
		return nil, errors.New("db: KeysGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := kgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (kgb *KeysGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := kgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (kgb *KeysGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = kgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{keys.Label}
	default:
		err = fmt.Errorf("db: KeysGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (kgb *KeysGroupBy) BoolX(ctx context.Context) bool {
	v, err := kgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (kgb *KeysGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range kgb.fields {
		if !keys.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := kgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := kgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (kgb *KeysGroupBy) sqlQuery() *sql.Selector {
	selector := kgb.sql.Select()
	aggregation := make([]string, 0, len(kgb.fns))
	for _, fn := range kgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(kgb.fields)+len(kgb.fns))
		for _, f := range kgb.fields {
			columns = append(columns, selector.C(f))
		}
		for _, c := range aggregation {
			columns = append(columns, c)
		}
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(kgb.fields...)...)
}

// KeysSelect is the builder for selecting fields of Keys entities.
type KeysSelect struct {
	*KeysQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ks *KeysSelect) Scan(ctx context.Context, v interface{}) error {
	if err := ks.prepareQuery(ctx); err != nil {
		return err
	}
	ks.sql = ks.KeysQuery.sqlQuery(ctx)
	return ks.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ks *KeysSelect) ScanX(ctx context.Context, v interface{}) {
	if err := ks.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (ks *KeysSelect) Strings(ctx context.Context) ([]string, error) {
	if len(ks.fields) > 1 {
		return nil, errors.New("db: KeysSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := ks.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ks *KeysSelect) StringsX(ctx context.Context) []string {
	v, err := ks.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (ks *KeysSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ks.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{keys.Label}
	default:
		err = fmt.Errorf("db: KeysSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ks *KeysSelect) StringX(ctx context.Context) string {
	v, err := ks.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (ks *KeysSelect) Ints(ctx context.Context) ([]int, error) {
	if len(ks.fields) > 1 {
		return nil, errors.New("db: KeysSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := ks.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ks *KeysSelect) IntsX(ctx context.Context) []int {
	v, err := ks.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (ks *KeysSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ks.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{keys.Label}
	default:
		err = fmt.Errorf("db: KeysSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ks *KeysSelect) IntX(ctx context.Context) int {
	v, err := ks.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (ks *KeysSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(ks.fields) > 1 {
		return nil, errors.New("db: KeysSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := ks.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ks *KeysSelect) Float64sX(ctx context.Context) []float64 {
	v, err := ks.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (ks *KeysSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ks.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{keys.Label}
	default:
		err = fmt.Errorf("db: KeysSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ks *KeysSelect) Float64X(ctx context.Context) float64 {
	v, err := ks.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (ks *KeysSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(ks.fields) > 1 {
		return nil, errors.New("db: KeysSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := ks.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ks *KeysSelect) BoolsX(ctx context.Context) []bool {
	v, err := ks.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (ks *KeysSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ks.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{keys.Label}
	default:
		err = fmt.Errorf("db: KeysSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ks *KeysSelect) BoolX(ctx context.Context) bool {
	v, err := ks.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ks *KeysSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ks.sql.Query()
	if err := ks.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
