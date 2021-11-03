// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/sphinx-service/pkg/db/ent/coininfo"
	"github.com/NpoolPlatform/sphinx-service/pkg/db/ent/predicate"
	"github.com/NpoolPlatform/sphinx-service/pkg/db/ent/review"
	"github.com/NpoolPlatform/sphinx-service/pkg/db/ent/transaction"
)

// ReviewUpdate is the builder for updating Review entities.
type ReviewUpdate struct {
	config
	hooks    []Hook
	mutation *ReviewMutation
}

// Where appends a list predicates to the ReviewUpdate builder.
func (ru *ReviewUpdate) Where(ps ...predicate.Review) *ReviewUpdate {
	ru.mutation.Where(ps...)
	return ru
}

// SetIsApproved sets the "is_approved" field.
func (ru *ReviewUpdate) SetIsApproved(b bool) *ReviewUpdate {
	ru.mutation.SetIsApproved(b)
	return ru
}

// SetNillableIsApproved sets the "is_approved" field if the given value is not nil.
func (ru *ReviewUpdate) SetNillableIsApproved(b *bool) *ReviewUpdate {
	if b != nil {
		ru.SetIsApproved(*b)
	}
	return ru
}

// SetOperatorNote sets the "operator_note" field.
func (ru *ReviewUpdate) SetOperatorNote(s string) *ReviewUpdate {
	ru.mutation.SetOperatorNote(s)
	return ru
}

// SetCreatetimeUtc sets the "createtime_utc" field.
func (ru *ReviewUpdate) SetCreatetimeUtc(i int) *ReviewUpdate {
	ru.mutation.ResetCreatetimeUtc()
	ru.mutation.SetCreatetimeUtc(i)
	return ru
}

// AddCreatetimeUtc adds i to the "createtime_utc" field.
func (ru *ReviewUpdate) AddCreatetimeUtc(i int) *ReviewUpdate {
	ru.mutation.AddCreatetimeUtc(i)
	return ru
}

// SetUpdatetimeUtc sets the "updatetime_utc" field.
func (ru *ReviewUpdate) SetUpdatetimeUtc(i int) *ReviewUpdate {
	ru.mutation.ResetUpdatetimeUtc()
	ru.mutation.SetUpdatetimeUtc(i)
	return ru
}

// AddUpdatetimeUtc adds i to the "updatetime_utc" field.
func (ru *ReviewUpdate) AddUpdatetimeUtc(i int) *ReviewUpdate {
	ru.mutation.AddUpdatetimeUtc(i)
	return ru
}

// SetTransactionID sets the "transaction" edge to the Transaction entity by ID.
func (ru *ReviewUpdate) SetTransactionID(id int) *ReviewUpdate {
	ru.mutation.SetTransactionID(id)
	return ru
}

// SetTransaction sets the "transaction" edge to the Transaction entity.
func (ru *ReviewUpdate) SetTransaction(t *Transaction) *ReviewUpdate {
	return ru.SetTransactionID(t.ID)
}

// SetCoinID sets the "coin" edge to the CoinInfo entity by ID.
func (ru *ReviewUpdate) SetCoinID(id int) *ReviewUpdate {
	ru.mutation.SetCoinID(id)
	return ru
}

// SetCoin sets the "coin" edge to the CoinInfo entity.
func (ru *ReviewUpdate) SetCoin(c *CoinInfo) *ReviewUpdate {
	return ru.SetCoinID(c.ID)
}

// Mutation returns the ReviewMutation object of the builder.
func (ru *ReviewUpdate) Mutation() *ReviewMutation {
	return ru.mutation
}

// ClearTransaction clears the "transaction" edge to the Transaction entity.
func (ru *ReviewUpdate) ClearTransaction() *ReviewUpdate {
	ru.mutation.ClearTransaction()
	return ru
}

// ClearCoin clears the "coin" edge to the CoinInfo entity.
func (ru *ReviewUpdate) ClearCoin() *ReviewUpdate {
	ru.mutation.ClearCoin()
	return ru
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ru *ReviewUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ru.hooks) == 0 {
		if err = ru.check(); err != nil {
			return 0, err
		}
		affected, err = ru.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ReviewMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ru.check(); err != nil {
				return 0, err
			}
			ru.mutation = mutation
			affected, err = ru.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ru.hooks) - 1; i >= 0; i-- {
			if ru.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ru.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ru.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ru *ReviewUpdate) SaveX(ctx context.Context) int {
	affected, err := ru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ru *ReviewUpdate) Exec(ctx context.Context) error {
	_, err := ru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ru *ReviewUpdate) ExecX(ctx context.Context) {
	if err := ru.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ru *ReviewUpdate) check() error {
	if v, ok := ru.mutation.OperatorNote(); ok {
		if err := review.OperatorNoteValidator(v); err != nil {
			return &ValidationError{Name: "operator_note", err: fmt.Errorf("ent: validator failed for field \"operator_note\": %w", err)}
		}
	}
	if _, ok := ru.mutation.TransactionID(); ru.mutation.TransactionCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"transaction\"")
	}
	if _, ok := ru.mutation.CoinID(); ru.mutation.CoinCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"coin\"")
	}
	return nil
}

func (ru *ReviewUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   review.Table,
			Columns: review.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: review.FieldID,
			},
		},
	}
	if ps := ru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ru.mutation.IsApproved(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: review.FieldIsApproved,
		})
	}
	if value, ok := ru.mutation.OperatorNote(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: review.FieldOperatorNote,
		})
	}
	if value, ok := ru.mutation.CreatetimeUtc(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: review.FieldCreatetimeUtc,
		})
	}
	if value, ok := ru.mutation.AddedCreatetimeUtc(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: review.FieldCreatetimeUtc,
		})
	}
	if value, ok := ru.mutation.UpdatetimeUtc(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: review.FieldUpdatetimeUtc,
		})
	}
	if value, ok := ru.mutation.AddedUpdatetimeUtc(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: review.FieldUpdatetimeUtc,
		})
	}
	if ru.mutation.TransactionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   review.TransactionTable,
			Columns: []string{review.TransactionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: transaction.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.TransactionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   review.TransactionTable,
			Columns: []string{review.TransactionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: transaction.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ru.mutation.CoinCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   review.CoinTable,
			Columns: []string{review.CoinColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: coininfo.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.CoinIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   review.CoinTable,
			Columns: []string{review.CoinColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: coininfo.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{review.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// ReviewUpdateOne is the builder for updating a single Review entity.
type ReviewUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ReviewMutation
}

// SetIsApproved sets the "is_approved" field.
func (ruo *ReviewUpdateOne) SetIsApproved(b bool) *ReviewUpdateOne {
	ruo.mutation.SetIsApproved(b)
	return ruo
}

// SetNillableIsApproved sets the "is_approved" field if the given value is not nil.
func (ruo *ReviewUpdateOne) SetNillableIsApproved(b *bool) *ReviewUpdateOne {
	if b != nil {
		ruo.SetIsApproved(*b)
	}
	return ruo
}

// SetOperatorNote sets the "operator_note" field.
func (ruo *ReviewUpdateOne) SetOperatorNote(s string) *ReviewUpdateOne {
	ruo.mutation.SetOperatorNote(s)
	return ruo
}

// SetCreatetimeUtc sets the "createtime_utc" field.
func (ruo *ReviewUpdateOne) SetCreatetimeUtc(i int) *ReviewUpdateOne {
	ruo.mutation.ResetCreatetimeUtc()
	ruo.mutation.SetCreatetimeUtc(i)
	return ruo
}

// AddCreatetimeUtc adds i to the "createtime_utc" field.
func (ruo *ReviewUpdateOne) AddCreatetimeUtc(i int) *ReviewUpdateOne {
	ruo.mutation.AddCreatetimeUtc(i)
	return ruo
}

// SetUpdatetimeUtc sets the "updatetime_utc" field.
func (ruo *ReviewUpdateOne) SetUpdatetimeUtc(i int) *ReviewUpdateOne {
	ruo.mutation.ResetUpdatetimeUtc()
	ruo.mutation.SetUpdatetimeUtc(i)
	return ruo
}

// AddUpdatetimeUtc adds i to the "updatetime_utc" field.
func (ruo *ReviewUpdateOne) AddUpdatetimeUtc(i int) *ReviewUpdateOne {
	ruo.mutation.AddUpdatetimeUtc(i)
	return ruo
}

// SetTransactionID sets the "transaction" edge to the Transaction entity by ID.
func (ruo *ReviewUpdateOne) SetTransactionID(id int) *ReviewUpdateOne {
	ruo.mutation.SetTransactionID(id)
	return ruo
}

// SetTransaction sets the "transaction" edge to the Transaction entity.
func (ruo *ReviewUpdateOne) SetTransaction(t *Transaction) *ReviewUpdateOne {
	return ruo.SetTransactionID(t.ID)
}

// SetCoinID sets the "coin" edge to the CoinInfo entity by ID.
func (ruo *ReviewUpdateOne) SetCoinID(id int) *ReviewUpdateOne {
	ruo.mutation.SetCoinID(id)
	return ruo
}

// SetCoin sets the "coin" edge to the CoinInfo entity.
func (ruo *ReviewUpdateOne) SetCoin(c *CoinInfo) *ReviewUpdateOne {
	return ruo.SetCoinID(c.ID)
}

// Mutation returns the ReviewMutation object of the builder.
func (ruo *ReviewUpdateOne) Mutation() *ReviewMutation {
	return ruo.mutation
}

// ClearTransaction clears the "transaction" edge to the Transaction entity.
func (ruo *ReviewUpdateOne) ClearTransaction() *ReviewUpdateOne {
	ruo.mutation.ClearTransaction()
	return ruo
}

// ClearCoin clears the "coin" edge to the CoinInfo entity.
func (ruo *ReviewUpdateOne) ClearCoin() *ReviewUpdateOne {
	ruo.mutation.ClearCoin()
	return ruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ruo *ReviewUpdateOne) Select(field string, fields ...string) *ReviewUpdateOne {
	ruo.fields = append([]string{field}, fields...)
	return ruo
}

// Save executes the query and returns the updated Review entity.
func (ruo *ReviewUpdateOne) Save(ctx context.Context) (*Review, error) {
	var (
		err  error
		node *Review
	)
	if len(ruo.hooks) == 0 {
		if err = ruo.check(); err != nil {
			return nil, err
		}
		node, err = ruo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ReviewMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ruo.check(); err != nil {
				return nil, err
			}
			ruo.mutation = mutation
			node, err = ruo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ruo.hooks) - 1; i >= 0; i-- {
			if ruo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ruo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ruo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ruo *ReviewUpdateOne) SaveX(ctx context.Context) *Review {
	node, err := ruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ruo *ReviewUpdateOne) Exec(ctx context.Context) error {
	_, err := ruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ruo *ReviewUpdateOne) ExecX(ctx context.Context) {
	if err := ruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ruo *ReviewUpdateOne) check() error {
	if v, ok := ruo.mutation.OperatorNote(); ok {
		if err := review.OperatorNoteValidator(v); err != nil {
			return &ValidationError{Name: "operator_note", err: fmt.Errorf("ent: validator failed for field \"operator_note\": %w", err)}
		}
	}
	if _, ok := ruo.mutation.TransactionID(); ruo.mutation.TransactionCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"transaction\"")
	}
	if _, ok := ruo.mutation.CoinID(); ruo.mutation.CoinCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"coin\"")
	}
	return nil
}

func (ruo *ReviewUpdateOne) sqlSave(ctx context.Context) (_node *Review, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   review.Table,
			Columns: review.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: review.FieldID,
			},
		},
	}
	id, ok := ruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Review.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := ruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, review.FieldID)
		for _, f := range fields {
			if !review.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != review.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ruo.mutation.IsApproved(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: review.FieldIsApproved,
		})
	}
	if value, ok := ruo.mutation.OperatorNote(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: review.FieldOperatorNote,
		})
	}
	if value, ok := ruo.mutation.CreatetimeUtc(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: review.FieldCreatetimeUtc,
		})
	}
	if value, ok := ruo.mutation.AddedCreatetimeUtc(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: review.FieldCreatetimeUtc,
		})
	}
	if value, ok := ruo.mutation.UpdatetimeUtc(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: review.FieldUpdatetimeUtc,
		})
	}
	if value, ok := ruo.mutation.AddedUpdatetimeUtc(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: review.FieldUpdatetimeUtc,
		})
	}
	if ruo.mutation.TransactionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   review.TransactionTable,
			Columns: []string{review.TransactionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: transaction.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.TransactionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   review.TransactionTable,
			Columns: []string{review.TransactionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: transaction.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ruo.mutation.CoinCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   review.CoinTable,
			Columns: []string{review.CoinColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: coininfo.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.CoinIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   review.CoinTable,
			Columns: []string{review.CoinColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: coininfo.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Review{config: ruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{review.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
