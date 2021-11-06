// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/sphinx-service/pkg/db/ent/coininfo"
	"github.com/NpoolPlatform/sphinx-service/pkg/db/ent/predicate"
	"github.com/NpoolPlatform/sphinx-service/pkg/db/ent/walletnode"
)

// WalletNodeQuery is the builder for querying WalletNode entities.
type WalletNodeQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.WalletNode
	// eager-loading edges.
	withCoin *CoinInfoQuery
	withFKs  bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the WalletNodeQuery builder.
func (wnq *WalletNodeQuery) Where(ps ...predicate.WalletNode) *WalletNodeQuery {
	wnq.predicates = append(wnq.predicates, ps...)
	return wnq
}

// Limit adds a limit step to the query.
func (wnq *WalletNodeQuery) Limit(limit int) *WalletNodeQuery {
	wnq.limit = &limit
	return wnq
}

// Offset adds an offset step to the query.
func (wnq *WalletNodeQuery) Offset(offset int) *WalletNodeQuery {
	wnq.offset = &offset
	return wnq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (wnq *WalletNodeQuery) Unique(unique bool) *WalletNodeQuery {
	wnq.unique = &unique
	return wnq
}

// Order adds an order step to the query.
func (wnq *WalletNodeQuery) Order(o ...OrderFunc) *WalletNodeQuery {
	wnq.order = append(wnq.order, o...)
	return wnq
}

// QueryCoin chains the current query on the "coin" edge.
func (wnq *WalletNodeQuery) QueryCoin() *CoinInfoQuery {
	query := &CoinInfoQuery{config: wnq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := wnq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := wnq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(walletnode.Table, walletnode.FieldID, selector),
			sqlgraph.To(coininfo.Table, coininfo.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, walletnode.CoinTable, walletnode.CoinColumn),
		)
		fromU = sqlgraph.SetNeighbors(wnq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first WalletNode entity from the query.
// Returns a *NotFoundError when no WalletNode was found.
func (wnq *WalletNodeQuery) First(ctx context.Context) (*WalletNode, error) {
	nodes, err := wnq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{walletnode.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (wnq *WalletNodeQuery) FirstX(ctx context.Context) *WalletNode {
	node, err := wnq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first WalletNode ID from the query.
// Returns a *NotFoundError when no WalletNode ID was found.
func (wnq *WalletNodeQuery) FirstID(ctx context.Context) (id int32, err error) {
	var ids []int32
	if ids, err = wnq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{walletnode.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (wnq *WalletNodeQuery) FirstIDX(ctx context.Context) int32 {
	id, err := wnq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single WalletNode entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one WalletNode entity is not found.
// Returns a *NotFoundError when no WalletNode entities are found.
func (wnq *WalletNodeQuery) Only(ctx context.Context) (*WalletNode, error) {
	nodes, err := wnq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{walletnode.Label}
	default:
		return nil, &NotSingularError{walletnode.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (wnq *WalletNodeQuery) OnlyX(ctx context.Context) *WalletNode {
	node, err := wnq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only WalletNode ID in the query.
// Returns a *NotSingularError when exactly one WalletNode ID is not found.
// Returns a *NotFoundError when no entities are found.
func (wnq *WalletNodeQuery) OnlyID(ctx context.Context) (id int32, err error) {
	var ids []int32
	if ids, err = wnq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{walletnode.Label}
	default:
		err = &NotSingularError{walletnode.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (wnq *WalletNodeQuery) OnlyIDX(ctx context.Context) int32 {
	id, err := wnq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of WalletNodes.
func (wnq *WalletNodeQuery) All(ctx context.Context) ([]*WalletNode, error) {
	if err := wnq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return wnq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (wnq *WalletNodeQuery) AllX(ctx context.Context) []*WalletNode {
	nodes, err := wnq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of WalletNode IDs.
func (wnq *WalletNodeQuery) IDs(ctx context.Context) ([]int32, error) {
	var ids []int32
	if err := wnq.Select(walletnode.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (wnq *WalletNodeQuery) IDsX(ctx context.Context) []int32 {
	ids, err := wnq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (wnq *WalletNodeQuery) Count(ctx context.Context) (int, error) {
	if err := wnq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return wnq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (wnq *WalletNodeQuery) CountX(ctx context.Context) int {
	count, err := wnq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (wnq *WalletNodeQuery) Exist(ctx context.Context) (bool, error) {
	if err := wnq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return wnq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (wnq *WalletNodeQuery) ExistX(ctx context.Context) bool {
	exist, err := wnq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the WalletNodeQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (wnq *WalletNodeQuery) Clone() *WalletNodeQuery {
	if wnq == nil {
		return nil
	}
	return &WalletNodeQuery{
		config:     wnq.config,
		limit:      wnq.limit,
		offset:     wnq.offset,
		order:      append([]OrderFunc{}, wnq.order...),
		predicates: append([]predicate.WalletNode{}, wnq.predicates...),
		withCoin:   wnq.withCoin.Clone(),
		// clone intermediate query.
		sql:  wnq.sql.Clone(),
		path: wnq.path,
	}
}

// WithCoin tells the query-builder to eager-load the nodes that are connected to
// the "coin" edge. The optional arguments are used to configure the query builder of the edge.
func (wnq *WalletNodeQuery) WithCoin(opts ...func(*CoinInfoQuery)) *WalletNodeQuery {
	query := &CoinInfoQuery{config: wnq.config}
	for _, opt := range opts {
		opt(query)
	}
	wnq.withCoin = query
	return wnq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		UUID string `json:"uuid,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.WalletNode.Query().
//		GroupBy(walletnode.FieldUUID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (wnq *WalletNodeQuery) GroupBy(field string, fields ...string) *WalletNodeGroupBy {
	group := &WalletNodeGroupBy{config: wnq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := wnq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return wnq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		UUID string `json:"uuid,omitempty"`
//	}
//
//	client.WalletNode.Query().
//		Select(walletnode.FieldUUID).
//		Scan(ctx, &v)
//
func (wnq *WalletNodeQuery) Select(fields ...string) *WalletNodeSelect {
	wnq.fields = append(wnq.fields, fields...)
	return &WalletNodeSelect{WalletNodeQuery: wnq}
}

func (wnq *WalletNodeQuery) prepareQuery(ctx context.Context) error {
	for _, f := range wnq.fields {
		if !walletnode.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if wnq.path != nil {
		prev, err := wnq.path(ctx)
		if err != nil {
			return err
		}
		wnq.sql = prev
	}
	return nil
}

func (wnq *WalletNodeQuery) sqlAll(ctx context.Context) ([]*WalletNode, error) {
	var (
		nodes       = []*WalletNode{}
		withFKs     = wnq.withFKs
		_spec       = wnq.querySpec()
		loadedTypes = [1]bool{
			wnq.withCoin != nil,
		}
	)
	if wnq.withCoin != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, walletnode.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &WalletNode{config: wnq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, wnq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := wnq.withCoin; query != nil {
		ids := make([]int32, 0, len(nodes))
		nodeids := make(map[int32][]*WalletNode)
		for i := range nodes {
			if nodes[i].coin_info_wallet_nodes == nil {
				continue
			}
			fk := *nodes[i].coin_info_wallet_nodes
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(coininfo.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "coin_info_wallet_nodes" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Coin = n
			}
		}
	}

	return nodes, nil
}

func (wnq *WalletNodeQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := wnq.querySpec()
	return sqlgraph.CountNodes(ctx, wnq.driver, _spec)
}

func (wnq *WalletNodeQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := wnq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (wnq *WalletNodeQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   walletnode.Table,
			Columns: walletnode.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt32,
				Column: walletnode.FieldID,
			},
		},
		From:   wnq.sql,
		Unique: true,
	}
	if unique := wnq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := wnq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, walletnode.FieldID)
		for i := range fields {
			if fields[i] != walletnode.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := wnq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := wnq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := wnq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := wnq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (wnq *WalletNodeQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(wnq.driver.Dialect())
	t1 := builder.Table(walletnode.Table)
	columns := wnq.fields
	if len(columns) == 0 {
		columns = walletnode.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if wnq.sql != nil {
		selector = wnq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	for _, p := range wnq.predicates {
		p(selector)
	}
	for _, p := range wnq.order {
		p(selector)
	}
	if offset := wnq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := wnq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WalletNodeGroupBy is the group-by builder for WalletNode entities.
type WalletNodeGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (wngb *WalletNodeGroupBy) Aggregate(fns ...AggregateFunc) *WalletNodeGroupBy {
	wngb.fns = append(wngb.fns, fns...)
	return wngb
}

// Scan applies the group-by query and scans the result into the given value.
func (wngb *WalletNodeGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := wngb.path(ctx)
	if err != nil {
		return err
	}
	wngb.sql = query
	return wngb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (wngb *WalletNodeGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := wngb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (wngb *WalletNodeGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(wngb.fields) > 1 {
		return nil, errors.New("ent: WalletNodeGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := wngb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (wngb *WalletNodeGroupBy) StringsX(ctx context.Context) []string {
	v, err := wngb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (wngb *WalletNodeGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = wngb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{walletnode.Label}
	default:
		err = fmt.Errorf("ent: WalletNodeGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (wngb *WalletNodeGroupBy) StringX(ctx context.Context) string {
	v, err := wngb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (wngb *WalletNodeGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(wngb.fields) > 1 {
		return nil, errors.New("ent: WalletNodeGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := wngb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (wngb *WalletNodeGroupBy) IntsX(ctx context.Context) []int {
	v, err := wngb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (wngb *WalletNodeGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = wngb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{walletnode.Label}
	default:
		err = fmt.Errorf("ent: WalletNodeGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (wngb *WalletNodeGroupBy) IntX(ctx context.Context) int {
	v, err := wngb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (wngb *WalletNodeGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(wngb.fields) > 1 {
		return nil, errors.New("ent: WalletNodeGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := wngb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (wngb *WalletNodeGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := wngb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (wngb *WalletNodeGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = wngb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{walletnode.Label}
	default:
		err = fmt.Errorf("ent: WalletNodeGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (wngb *WalletNodeGroupBy) Float64X(ctx context.Context) float64 {
	v, err := wngb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (wngb *WalletNodeGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(wngb.fields) > 1 {
		return nil, errors.New("ent: WalletNodeGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := wngb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (wngb *WalletNodeGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := wngb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (wngb *WalletNodeGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = wngb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{walletnode.Label}
	default:
		err = fmt.Errorf("ent: WalletNodeGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (wngb *WalletNodeGroupBy) BoolX(ctx context.Context) bool {
	v, err := wngb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (wngb *WalletNodeGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range wngb.fields {
		if !walletnode.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := wngb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := wngb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (wngb *WalletNodeGroupBy) sqlQuery() *sql.Selector {
	selector := wngb.sql.Select()
	aggregation := make([]string, 0, len(wngb.fns))
	for _, fn := range wngb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(wngb.fields)+len(wngb.fns))
		for _, f := range wngb.fields {
			columns = append(columns, selector.C(f))
		}
		for _, c := range aggregation {
			columns = append(columns, c)
		}
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(wngb.fields...)...)
}

// WalletNodeSelect is the builder for selecting fields of WalletNode entities.
type WalletNodeSelect struct {
	*WalletNodeQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (wns *WalletNodeSelect) Scan(ctx context.Context, v interface{}) error {
	if err := wns.prepareQuery(ctx); err != nil {
		return err
	}
	wns.sql = wns.WalletNodeQuery.sqlQuery(ctx)
	return wns.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (wns *WalletNodeSelect) ScanX(ctx context.Context, v interface{}) {
	if err := wns.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (wns *WalletNodeSelect) Strings(ctx context.Context) ([]string, error) {
	if len(wns.fields) > 1 {
		return nil, errors.New("ent: WalletNodeSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := wns.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (wns *WalletNodeSelect) StringsX(ctx context.Context) []string {
	v, err := wns.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (wns *WalletNodeSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = wns.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{walletnode.Label}
	default:
		err = fmt.Errorf("ent: WalletNodeSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (wns *WalletNodeSelect) StringX(ctx context.Context) string {
	v, err := wns.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (wns *WalletNodeSelect) Ints(ctx context.Context) ([]int, error) {
	if len(wns.fields) > 1 {
		return nil, errors.New("ent: WalletNodeSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := wns.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (wns *WalletNodeSelect) IntsX(ctx context.Context) []int {
	v, err := wns.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (wns *WalletNodeSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = wns.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{walletnode.Label}
	default:
		err = fmt.Errorf("ent: WalletNodeSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (wns *WalletNodeSelect) IntX(ctx context.Context) int {
	v, err := wns.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (wns *WalletNodeSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(wns.fields) > 1 {
		return nil, errors.New("ent: WalletNodeSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := wns.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (wns *WalletNodeSelect) Float64sX(ctx context.Context) []float64 {
	v, err := wns.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (wns *WalletNodeSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = wns.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{walletnode.Label}
	default:
		err = fmt.Errorf("ent: WalletNodeSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (wns *WalletNodeSelect) Float64X(ctx context.Context) float64 {
	v, err := wns.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (wns *WalletNodeSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(wns.fields) > 1 {
		return nil, errors.New("ent: WalletNodeSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := wns.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (wns *WalletNodeSelect) BoolsX(ctx context.Context) []bool {
	v, err := wns.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (wns *WalletNodeSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = wns.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{walletnode.Label}
	default:
		err = fmt.Errorf("ent: WalletNodeSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (wns *WalletNodeSelect) BoolX(ctx context.Context) bool {
	v, err := wns.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (wns *WalletNodeSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := wns.sql.Query()
	if err := wns.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}