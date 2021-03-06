// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/sphinx-plugin/pkg/db/ent/coininfo"
	"github.com/NpoolPlatform/sphinx-plugin/pkg/db/ent/review"
	"github.com/NpoolPlatform/sphinx-plugin/pkg/db/ent/transaction"
)

// TransactionCreate is the builder for creating a Transaction entity.
type TransactionCreate struct {
	config
	mutation *TransactionMutation
	hooks    []Hook
}

// SetAmountInt sets the "amount_int" field.
func (tc *TransactionCreate) SetAmountInt(i int) *TransactionCreate {
	tc.mutation.SetAmountInt(i)
	return tc
}

// SetAmountDigits sets the "amount_digits" field.
func (tc *TransactionCreate) SetAmountDigits(i int) *TransactionCreate {
	tc.mutation.SetAmountDigits(i)
	return tc
}

// SetNillableAmountDigits sets the "amount_digits" field if the given value is not nil.
func (tc *TransactionCreate) SetNillableAmountDigits(i *int) *TransactionCreate {
	if i != nil {
		tc.SetAmountDigits(*i)
	}
	return tc
}

// SetAddressFrom sets the "address_from" field.
func (tc *TransactionCreate) SetAddressFrom(s string) *TransactionCreate {
	tc.mutation.SetAddressFrom(s)
	return tc
}

// SetAddressTo sets the "address_to" field.
func (tc *TransactionCreate) SetAddressTo(s string) *TransactionCreate {
	tc.mutation.SetAddressTo(s)
	return tc
}

// SetNeedManualReview sets the "need_manual_review" field.
func (tc *TransactionCreate) SetNeedManualReview(b bool) *TransactionCreate {
	tc.mutation.SetNeedManualReview(b)
	return tc
}

// SetNillableNeedManualReview sets the "need_manual_review" field if the given value is not nil.
func (tc *TransactionCreate) SetNillableNeedManualReview(b *bool) *TransactionCreate {
	if b != nil {
		tc.SetNeedManualReview(*b)
	}
	return tc
}

// SetType sets the "type" field.
func (tc *TransactionCreate) SetType(t transaction.Type) *TransactionCreate {
	tc.mutation.SetType(t)
	return tc
}

// SetTransactionIDInsite sets the "transaction_id_insite" field.
func (tc *TransactionCreate) SetTransactionIDInsite(s string) *TransactionCreate {
	tc.mutation.SetTransactionIDInsite(s)
	return tc
}

// SetTransactionIDChain sets the "transaction_id_chain" field.
func (tc *TransactionCreate) SetTransactionIDChain(s string) *TransactionCreate {
	tc.mutation.SetTransactionIDChain(s)
	return tc
}

// SetStatus sets the "status" field.
func (tc *TransactionCreate) SetStatus(t transaction.Status) *TransactionCreate {
	tc.mutation.SetStatus(t)
	return tc
}

// SetMutex sets the "mutex" field.
func (tc *TransactionCreate) SetMutex(b bool) *TransactionCreate {
	tc.mutation.SetMutex(b)
	return tc
}

// SetNillableMutex sets the "mutex" field if the given value is not nil.
func (tc *TransactionCreate) SetNillableMutex(b *bool) *TransactionCreate {
	if b != nil {
		tc.SetMutex(*b)
	}
	return tc
}

// SetCreatetimeUtc sets the "createtime_utc" field.
func (tc *TransactionCreate) SetCreatetimeUtc(i int) *TransactionCreate {
	tc.mutation.SetCreatetimeUtc(i)
	return tc
}

// SetUpdatetimeUtc sets the "updatetime_utc" field.
func (tc *TransactionCreate) SetUpdatetimeUtc(i int) *TransactionCreate {
	tc.mutation.SetUpdatetimeUtc(i)
	return tc
}

// SetID sets the "id" field.
func (tc *TransactionCreate) SetID(i int32) *TransactionCreate {
	tc.mutation.SetID(i)
	return tc
}

// SetCoinID sets the "coin" edge to the CoinInfo entity by ID.
func (tc *TransactionCreate) SetCoinID(id int32) *TransactionCreate {
	tc.mutation.SetCoinID(id)
	return tc
}

// SetCoin sets the "coin" edge to the CoinInfo entity.
func (tc *TransactionCreate) SetCoin(c *CoinInfo) *TransactionCreate {
	return tc.SetCoinID(c.ID)
}

// AddReviewIDs adds the "review" edge to the Review entity by IDs.
func (tc *TransactionCreate) AddReviewIDs(ids ...int32) *TransactionCreate {
	tc.mutation.AddReviewIDs(ids...)
	return tc
}

// AddReview adds the "review" edges to the Review entity.
func (tc *TransactionCreate) AddReview(r ...*Review) *TransactionCreate {
	ids := make([]int32, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return tc.AddReviewIDs(ids...)
}

// Mutation returns the TransactionMutation object of the builder.
func (tc *TransactionCreate) Mutation() *TransactionMutation {
	return tc.mutation
}

// Save creates the Transaction in the database.
func (tc *TransactionCreate) Save(ctx context.Context) (*Transaction, error) {
	var (
		err  error
		node *Transaction
	)
	tc.defaults()
	if len(tc.hooks) == 0 {
		if err = tc.check(); err != nil {
			return nil, err
		}
		node, err = tc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TransactionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tc.check(); err != nil {
				return nil, err
			}
			tc.mutation = mutation
			if node, err = tc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(tc.hooks) - 1; i >= 0; i-- {
			if tc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (tc *TransactionCreate) SaveX(ctx context.Context) *Transaction {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tc *TransactionCreate) Exec(ctx context.Context) error {
	_, err := tc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tc *TransactionCreate) ExecX(ctx context.Context) {
	if err := tc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tc *TransactionCreate) defaults() {
	if _, ok := tc.mutation.AmountDigits(); !ok {
		v := transaction.DefaultAmountDigits
		tc.mutation.SetAmountDigits(v)
	}
	if _, ok := tc.mutation.NeedManualReview(); !ok {
		v := transaction.DefaultNeedManualReview
		tc.mutation.SetNeedManualReview(v)
	}
	if _, ok := tc.mutation.Mutex(); !ok {
		v := transaction.DefaultMutex
		tc.mutation.SetMutex(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tc *TransactionCreate) check() error {
	if _, ok := tc.mutation.AmountInt(); !ok {
		return &ValidationError{Name: "amount_int", err: errors.New(`ent: missing required field "amount_int"`)}
	}
	if v, ok := tc.mutation.AmountInt(); ok {
		if err := transaction.AmountIntValidator(v); err != nil {
			return &ValidationError{Name: "amount_int", err: fmt.Errorf(`ent: validator failed for field "amount_int": %w`, err)}
		}
	}
	if _, ok := tc.mutation.AmountDigits(); !ok {
		return &ValidationError{Name: "amount_digits", err: errors.New(`ent: missing required field "amount_digits"`)}
	}
	if v, ok := tc.mutation.AmountDigits(); ok {
		if err := transaction.AmountDigitsValidator(v); err != nil {
			return &ValidationError{Name: "amount_digits", err: fmt.Errorf(`ent: validator failed for field "amount_digits": %w`, err)}
		}
	}
	if _, ok := tc.mutation.AddressFrom(); !ok {
		return &ValidationError{Name: "address_from", err: errors.New(`ent: missing required field "address_from"`)}
	}
	if v, ok := tc.mutation.AddressFrom(); ok {
		if err := transaction.AddressFromValidator(v); err != nil {
			return &ValidationError{Name: "address_from", err: fmt.Errorf(`ent: validator failed for field "address_from": %w`, err)}
		}
	}
	if _, ok := tc.mutation.AddressTo(); !ok {
		return &ValidationError{Name: "address_to", err: errors.New(`ent: missing required field "address_to"`)}
	}
	if v, ok := tc.mutation.AddressTo(); ok {
		if err := transaction.AddressToValidator(v); err != nil {
			return &ValidationError{Name: "address_to", err: fmt.Errorf(`ent: validator failed for field "address_to": %w`, err)}
		}
	}
	if _, ok := tc.mutation.NeedManualReview(); !ok {
		return &ValidationError{Name: "need_manual_review", err: errors.New(`ent: missing required field "need_manual_review"`)}
	}
	if _, ok := tc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "type"`)}
	}
	if v, ok := tc.mutation.GetType(); ok {
		if err := transaction.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "type": %w`, err)}
		}
	}
	if _, ok := tc.mutation.TransactionIDInsite(); !ok {
		return &ValidationError{Name: "transaction_id_insite", err: errors.New(`ent: missing required field "transaction_id_insite"`)}
	}
	if v, ok := tc.mutation.TransactionIDInsite(); ok {
		if err := transaction.TransactionIDInsiteValidator(v); err != nil {
			return &ValidationError{Name: "transaction_id_insite", err: fmt.Errorf(`ent: validator failed for field "transaction_id_insite": %w`, err)}
		}
	}
	if _, ok := tc.mutation.TransactionIDChain(); !ok {
		return &ValidationError{Name: "transaction_id_chain", err: errors.New(`ent: missing required field "transaction_id_chain"`)}
	}
	if v, ok := tc.mutation.TransactionIDChain(); ok {
		if err := transaction.TransactionIDChainValidator(v); err != nil {
			return &ValidationError{Name: "transaction_id_chain", err: fmt.Errorf(`ent: validator failed for field "transaction_id_chain": %w`, err)}
		}
	}
	if _, ok := tc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "status"`)}
	}
	if v, ok := tc.mutation.Status(); ok {
		if err := transaction.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "status": %w`, err)}
		}
	}
	if _, ok := tc.mutation.Mutex(); !ok {
		return &ValidationError{Name: "mutex", err: errors.New(`ent: missing required field "mutex"`)}
	}
	if _, ok := tc.mutation.CreatetimeUtc(); !ok {
		return &ValidationError{Name: "createtime_utc", err: errors.New(`ent: missing required field "createtime_utc"`)}
	}
	if _, ok := tc.mutation.UpdatetimeUtc(); !ok {
		return &ValidationError{Name: "updatetime_utc", err: errors.New(`ent: missing required field "updatetime_utc"`)}
	}
	if _, ok := tc.mutation.CoinID(); !ok {
		return &ValidationError{Name: "coin", err: errors.New("ent: missing required edge \"coin\"")}
	}
	return nil
}

func (tc *TransactionCreate) sqlSave(ctx context.Context) (*Transaction, error) {
	_node, _spec := tc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
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

func (tc *TransactionCreate) createSpec() (*Transaction, *sqlgraph.CreateSpec) {
	var (
		_node = &Transaction{config: tc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: transaction.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt32,
				Column: transaction.FieldID,
			},
		}
	)
	if id, ok := tc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := tc.mutation.AmountInt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: transaction.FieldAmountInt,
		})
		_node.AmountInt = value
	}
	if value, ok := tc.mutation.AmountDigits(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: transaction.FieldAmountDigits,
		})
		_node.AmountDigits = value
	}
	if value, ok := tc.mutation.AddressFrom(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: transaction.FieldAddressFrom,
		})
		_node.AddressFrom = value
	}
	if value, ok := tc.mutation.AddressTo(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: transaction.FieldAddressTo,
		})
		_node.AddressTo = value
	}
	if value, ok := tc.mutation.NeedManualReview(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: transaction.FieldNeedManualReview,
		})
		_node.NeedManualReview = value
	}
	if value, ok := tc.mutation.GetType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: transaction.FieldType,
		})
		_node.Type = value
	}
	if value, ok := tc.mutation.TransactionIDInsite(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: transaction.FieldTransactionIDInsite,
		})
		_node.TransactionIDInsite = value
	}
	if value, ok := tc.mutation.TransactionIDChain(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: transaction.FieldTransactionIDChain,
		})
		_node.TransactionIDChain = value
	}
	if value, ok := tc.mutation.Status(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: transaction.FieldStatus,
		})
		_node.Status = value
	}
	if value, ok := tc.mutation.Mutex(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: transaction.FieldMutex,
		})
		_node.Mutex = value
	}
	if value, ok := tc.mutation.CreatetimeUtc(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: transaction.FieldCreatetimeUtc,
		})
		_node.CreatetimeUtc = value
	}
	if value, ok := tc.mutation.UpdatetimeUtc(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: transaction.FieldUpdatetimeUtc,
		})
		_node.UpdatetimeUtc = value
	}
	if nodes := tc.mutation.CoinIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   transaction.CoinTable,
			Columns: []string{transaction.CoinColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt32,
					Column: coininfo.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.coin_info_transactions = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tc.mutation.ReviewIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   transaction.ReviewTable,
			Columns: []string{transaction.ReviewColumn},
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
	return _node, _spec
}

// TransactionCreateBulk is the builder for creating many Transaction entities in bulk.
type TransactionCreateBulk struct {
	config
	builders []*TransactionCreate
}

// Save creates the Transaction entities in the database.
func (tcb *TransactionCreateBulk) Save(ctx context.Context) ([]*Transaction, error) {
	specs := make([]*sqlgraph.CreateSpec, len(tcb.builders))
	nodes := make([]*Transaction, len(tcb.builders))
	mutators := make([]Mutator, len(tcb.builders))
	for i := range tcb.builders {
		func(i int, root context.Context) {
			builder := tcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TransactionMutation)
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
					_, err = mutators[i+1].Mutate(root, tcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, tcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tcb *TransactionCreateBulk) SaveX(ctx context.Context) []*Transaction {
	v, err := tcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tcb *TransactionCreateBulk) Exec(ctx context.Context) error {
	_, err := tcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcb *TransactionCreateBulk) ExecX(ctx context.Context) {
	if err := tcb.Exec(ctx); err != nil {
		panic(err)
	}
}
