// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/sphinx-plugin/pkg/db/ent/coininfo"
	"github.com/NpoolPlatform/sphinx-plugin/pkg/db/ent/keystore"
	"github.com/NpoolPlatform/sphinx-plugin/pkg/db/ent/review"
	"github.com/NpoolPlatform/sphinx-plugin/pkg/db/ent/transaction"
	"github.com/NpoolPlatform/sphinx-plugin/pkg/db/ent/walletnode"
)

// CoinInfoCreate is the builder for creating a CoinInfo entity.
type CoinInfoCreate struct {
	config
	mutation *CoinInfoMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (cic *CoinInfoCreate) SetName(s string) *CoinInfoCreate {
	cic.mutation.SetName(s)
	return cic
}

// SetUnit sets the "unit" field.
func (cic *CoinInfoCreate) SetUnit(s string) *CoinInfoCreate {
	cic.mutation.SetUnit(s)
	return cic
}

// SetNeedSigninfo sets the "need_signinfo" field.
func (cic *CoinInfoCreate) SetNeedSigninfo(b bool) *CoinInfoCreate {
	cic.mutation.SetNeedSigninfo(b)
	return cic
}

// SetNillableNeedSigninfo sets the "need_signinfo" field if the given value is not nil.
func (cic *CoinInfoCreate) SetNillableNeedSigninfo(b *bool) *CoinInfoCreate {
	if b != nil {
		cic.SetNeedSigninfo(*b)
	}
	return cic
}

// SetID sets the "id" field.
func (cic *CoinInfoCreate) SetID(i int32) *CoinInfoCreate {
	cic.mutation.SetID(i)
	return cic
}

// AddKeyIDs adds the "keys" edge to the KeyStore entity by IDs.
func (cic *CoinInfoCreate) AddKeyIDs(ids ...int32) *CoinInfoCreate {
	cic.mutation.AddKeyIDs(ids...)
	return cic
}

// AddKeys adds the "keys" edges to the KeyStore entity.
func (cic *CoinInfoCreate) AddKeys(k ...*KeyStore) *CoinInfoCreate {
	ids := make([]int32, len(k))
	for i := range k {
		ids[i] = k[i].ID
	}
	return cic.AddKeyIDs(ids...)
}

// AddTransactionIDs adds the "transactions" edge to the Transaction entity by IDs.
func (cic *CoinInfoCreate) AddTransactionIDs(ids ...int32) *CoinInfoCreate {
	cic.mutation.AddTransactionIDs(ids...)
	return cic
}

// AddTransactions adds the "transactions" edges to the Transaction entity.
func (cic *CoinInfoCreate) AddTransactions(t ...*Transaction) *CoinInfoCreate {
	ids := make([]int32, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return cic.AddTransactionIDs(ids...)
}

// AddReviewIDs adds the "reviews" edge to the Review entity by IDs.
func (cic *CoinInfoCreate) AddReviewIDs(ids ...int32) *CoinInfoCreate {
	cic.mutation.AddReviewIDs(ids...)
	return cic
}

// AddReviews adds the "reviews" edges to the Review entity.
func (cic *CoinInfoCreate) AddReviews(r ...*Review) *CoinInfoCreate {
	ids := make([]int32, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return cic.AddReviewIDs(ids...)
}

// AddWalletNodeIDs adds the "wallet_nodes" edge to the WalletNode entity by IDs.
func (cic *CoinInfoCreate) AddWalletNodeIDs(ids ...int32) *CoinInfoCreate {
	cic.mutation.AddWalletNodeIDs(ids...)
	return cic
}

// AddWalletNodes adds the "wallet_nodes" edges to the WalletNode entity.
func (cic *CoinInfoCreate) AddWalletNodes(w ...*WalletNode) *CoinInfoCreate {
	ids := make([]int32, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return cic.AddWalletNodeIDs(ids...)
}

// Mutation returns the CoinInfoMutation object of the builder.
func (cic *CoinInfoCreate) Mutation() *CoinInfoMutation {
	return cic.mutation
}

// Save creates the CoinInfo in the database.
func (cic *CoinInfoCreate) Save(ctx context.Context) (*CoinInfo, error) {
	var (
		err  error
		node *CoinInfo
	)
	cic.defaults()
	if len(cic.hooks) == 0 {
		if err = cic.check(); err != nil {
			return nil, err
		}
		node, err = cic.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CoinInfoMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cic.check(); err != nil {
				return nil, err
			}
			cic.mutation = mutation
			if node, err = cic.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(cic.hooks) - 1; i >= 0; i-- {
			if cic.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cic.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cic.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (cic *CoinInfoCreate) SaveX(ctx context.Context) *CoinInfo {
	v, err := cic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cic *CoinInfoCreate) Exec(ctx context.Context) error {
	_, err := cic.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cic *CoinInfoCreate) ExecX(ctx context.Context) {
	if err := cic.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cic *CoinInfoCreate) defaults() {
	if _, ok := cic.mutation.NeedSigninfo(); !ok {
		v := coininfo.DefaultNeedSigninfo
		cic.mutation.SetNeedSigninfo(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cic *CoinInfoCreate) check() error {
	if _, ok := cic.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "name"`)}
	}
	if v, ok := cic.mutation.Name(); ok {
		if err := coininfo.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "name": %w`, err)}
		}
	}
	if _, ok := cic.mutation.Unit(); !ok {
		return &ValidationError{Name: "unit", err: errors.New(`ent: missing required field "unit"`)}
	}
	if v, ok := cic.mutation.Unit(); ok {
		if err := coininfo.UnitValidator(v); err != nil {
			return &ValidationError{Name: "unit", err: fmt.Errorf(`ent: validator failed for field "unit": %w`, err)}
		}
	}
	if _, ok := cic.mutation.NeedSigninfo(); !ok {
		return &ValidationError{Name: "need_signinfo", err: errors.New(`ent: missing required field "need_signinfo"`)}
	}
	return nil
}

func (cic *CoinInfoCreate) sqlSave(ctx context.Context) (*CoinInfo, error) {
	_node, _spec := cic.createSpec()
	if err := sqlgraph.CreateNode(ctx, cic.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int32(id)
	}
	return _node, nil
}

func (cic *CoinInfoCreate) createSpec() (*CoinInfo, *sqlgraph.CreateSpec) {
	var (
		_node = &CoinInfo{config: cic.config}
		_spec = &sqlgraph.CreateSpec{
			Table: coininfo.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt32,
				Column: coininfo.FieldID,
			},
		}
	)
	if id, ok := cic.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := cic.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: coininfo.FieldName,
		})
		_node.Name = value
	}
	if value, ok := cic.mutation.Unit(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: coininfo.FieldUnit,
		})
		_node.Unit = value
	}
	if value, ok := cic.mutation.NeedSigninfo(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: coininfo.FieldNeedSigninfo,
		})
		_node.NeedSigninfo = value
	}
	if nodes := cic.mutation.KeysIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   coininfo.KeysTable,
			Columns: []string{coininfo.KeysColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt32,
					Column: keystore.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cic.mutation.TransactionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   coininfo.TransactionsTable,
			Columns: []string{coininfo.TransactionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt32,
					Column: transaction.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cic.mutation.ReviewsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   coininfo.ReviewsTable,
			Columns: []string{coininfo.ReviewsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt32,
					Column: review.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cic.mutation.WalletNodesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   coininfo.WalletNodesTable,
			Columns: []string{coininfo.WalletNodesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt32,
					Column: walletnode.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// CoinInfoCreateBulk is the builder for creating many CoinInfo entities in bulk.
type CoinInfoCreateBulk struct {
	config
	builders []*CoinInfoCreate
}

// Save creates the CoinInfo entities in the database.
func (cicb *CoinInfoCreateBulk) Save(ctx context.Context) ([]*CoinInfo, error) {
	specs := make([]*sqlgraph.CreateSpec, len(cicb.builders))
	nodes := make([]*CoinInfo, len(cicb.builders))
	mutators := make([]Mutator, len(cicb.builders))
	for i := range cicb.builders {
		func(i int, root context.Context) {
			builder := cicb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CoinInfoMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, cicb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, cicb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int32(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, cicb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (cicb *CoinInfoCreateBulk) SaveX(ctx context.Context) []*CoinInfo {
	v, err := cicb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cicb *CoinInfoCreateBulk) Exec(ctx context.Context) error {
	_, err := cicb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cicb *CoinInfoCreateBulk) ExecX(ctx context.Context) {
	if err := cicb.Exec(ctx); err != nil {
		panic(err)
	}
}
