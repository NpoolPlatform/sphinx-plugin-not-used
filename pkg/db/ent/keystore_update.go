// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/sphinx-plugin/pkg/db/ent/coininfo"
	"github.com/NpoolPlatform/sphinx-plugin/pkg/db/ent/keystore"
	"github.com/NpoolPlatform/sphinx-plugin/pkg/db/ent/predicate"
)

// KeyStoreUpdate is the builder for updating KeyStore entities.
type KeyStoreUpdate struct {
	config
	hooks    []Hook
	mutation *KeyStoreMutation
}

// Where appends a list predicates to the KeyStoreUpdate builder.
func (ksu *KeyStoreUpdate) Where(ps ...predicate.KeyStore) *KeyStoreUpdate {
	ksu.mutation.Where(ps...)
	return ksu
}

// SetAddress sets the "address" field.
func (ksu *KeyStoreUpdate) SetAddress(s string) *KeyStoreUpdate {
	ksu.mutation.SetAddress(s)
	return ksu
}

// SetPrivateKey sets the "private_key" field.
func (ksu *KeyStoreUpdate) SetPrivateKey(s string) *KeyStoreUpdate {
	ksu.mutation.SetPrivateKey(s)
	return ksu
}

// SetCoinID sets the "coin" edge to the CoinInfo entity by ID.
func (ksu *KeyStoreUpdate) SetCoinID(id int32) *KeyStoreUpdate {
	ksu.mutation.SetCoinID(id)
	return ksu
}

// SetCoin sets the "coin" edge to the CoinInfo entity.
func (ksu *KeyStoreUpdate) SetCoin(c *CoinInfo) *KeyStoreUpdate {
	return ksu.SetCoinID(c.ID)
}

// Mutation returns the KeyStoreMutation object of the builder.
func (ksu *KeyStoreUpdate) Mutation() *KeyStoreMutation {
	return ksu.mutation
}

// ClearCoin clears the "coin" edge to the CoinInfo entity.
func (ksu *KeyStoreUpdate) ClearCoin() *KeyStoreUpdate {
	ksu.mutation.ClearCoin()
	return ksu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ksu *KeyStoreUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ksu.hooks) == 0 {
		if err = ksu.check(); err != nil {
			return 0, err
		}
		affected, err = ksu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*KeyStoreMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ksu.check(); err != nil {
				return 0, err
			}
			ksu.mutation = mutation
			affected, err = ksu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ksu.hooks) - 1; i >= 0; i-- {
			if ksu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ksu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ksu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ksu *KeyStoreUpdate) SaveX(ctx context.Context) int {
	affected, err := ksu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ksu *KeyStoreUpdate) Exec(ctx context.Context) error {
	_, err := ksu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ksu *KeyStoreUpdate) ExecX(ctx context.Context) {
	if err := ksu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ksu *KeyStoreUpdate) check() error {
	if v, ok := ksu.mutation.Address(); ok {
		if err := keystore.AddressValidator(v); err != nil {
			return &ValidationError{Name: "address", err: fmt.Errorf("ent: validator failed for field \"address\": %w", err)}
		}
	}
	if v, ok := ksu.mutation.PrivateKey(); ok {
		if err := keystore.PrivateKeyValidator(v); err != nil {
			return &ValidationError{Name: "private_key", err: fmt.Errorf("ent: validator failed for field \"private_key\": %w", err)}
		}
	}
	if _, ok := ksu.mutation.CoinID(); ksu.mutation.CoinCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"coin\"")
	}
	return nil
}

func (ksu *KeyStoreUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   keystore.Table,
			Columns: keystore.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt32,
				Column: keystore.FieldID,
			},
		},
	}
	if ps := ksu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ksu.mutation.Address(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: keystore.FieldAddress,
		})
	}
	if value, ok := ksu.mutation.PrivateKey(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: keystore.FieldPrivateKey,
		})
	}
	if ksu.mutation.CoinCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   keystore.CoinTable,
			Columns: []string{keystore.CoinColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt32,
					Column: coininfo.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ksu.mutation.CoinIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   keystore.CoinTable,
			Columns: []string{keystore.CoinColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ksu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{keystore.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// KeyStoreUpdateOne is the builder for updating a single KeyStore entity.
type KeyStoreUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *KeyStoreMutation
}

// SetAddress sets the "address" field.
func (ksuo *KeyStoreUpdateOne) SetAddress(s string) *KeyStoreUpdateOne {
	ksuo.mutation.SetAddress(s)
	return ksuo
}

// SetPrivateKey sets the "private_key" field.
func (ksuo *KeyStoreUpdateOne) SetPrivateKey(s string) *KeyStoreUpdateOne {
	ksuo.mutation.SetPrivateKey(s)
	return ksuo
}

// SetCoinID sets the "coin" edge to the CoinInfo entity by ID.
func (ksuo *KeyStoreUpdateOne) SetCoinID(id int32) *KeyStoreUpdateOne {
	ksuo.mutation.SetCoinID(id)
	return ksuo
}

// SetCoin sets the "coin" edge to the CoinInfo entity.
func (ksuo *KeyStoreUpdateOne) SetCoin(c *CoinInfo) *KeyStoreUpdateOne {
	return ksuo.SetCoinID(c.ID)
}

// Mutation returns the KeyStoreMutation object of the builder.
func (ksuo *KeyStoreUpdateOne) Mutation() *KeyStoreMutation {
	return ksuo.mutation
}

// ClearCoin clears the "coin" edge to the CoinInfo entity.
func (ksuo *KeyStoreUpdateOne) ClearCoin() *KeyStoreUpdateOne {
	ksuo.mutation.ClearCoin()
	return ksuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ksuo *KeyStoreUpdateOne) Select(field string, fields ...string) *KeyStoreUpdateOne {
	ksuo.fields = append([]string{field}, fields...)
	return ksuo
}

// Save executes the query and returns the updated KeyStore entity.
func (ksuo *KeyStoreUpdateOne) Save(ctx context.Context) (*KeyStore, error) {
	var (
		err  error
		node *KeyStore
	)
	if len(ksuo.hooks) == 0 {
		if err = ksuo.check(); err != nil {
			return nil, err
		}
		node, err = ksuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*KeyStoreMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ksuo.check(); err != nil {
				return nil, err
			}
			ksuo.mutation = mutation
			node, err = ksuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ksuo.hooks) - 1; i >= 0; i-- {
			if ksuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ksuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ksuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ksuo *KeyStoreUpdateOne) SaveX(ctx context.Context) *KeyStore {
	node, err := ksuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ksuo *KeyStoreUpdateOne) Exec(ctx context.Context) error {
	_, err := ksuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ksuo *KeyStoreUpdateOne) ExecX(ctx context.Context) {
	if err := ksuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ksuo *KeyStoreUpdateOne) check() error {
	if v, ok := ksuo.mutation.Address(); ok {
		if err := keystore.AddressValidator(v); err != nil {
			return &ValidationError{Name: "address", err: fmt.Errorf("ent: validator failed for field \"address\": %w", err)}
		}
	}
	if v, ok := ksuo.mutation.PrivateKey(); ok {
		if err := keystore.PrivateKeyValidator(v); err != nil {
			return &ValidationError{Name: "private_key", err: fmt.Errorf("ent: validator failed for field \"private_key\": %w", err)}
		}
	}
	if _, ok := ksuo.mutation.CoinID(); ksuo.mutation.CoinCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"coin\"")
	}
	return nil
}

func (ksuo *KeyStoreUpdateOne) sqlSave(ctx context.Context) (_node *KeyStore, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   keystore.Table,
			Columns: keystore.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt32,
				Column: keystore.FieldID,
			},
		},
	}
	id, ok := ksuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing KeyStore.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := ksuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, keystore.FieldID)
		for _, f := range fields {
			if !keystore.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != keystore.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ksuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ksuo.mutation.Address(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: keystore.FieldAddress,
		})
	}
	if value, ok := ksuo.mutation.PrivateKey(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: keystore.FieldPrivateKey,
		})
	}
	if ksuo.mutation.CoinCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   keystore.CoinTable,
			Columns: []string{keystore.CoinColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt32,
					Column: coininfo.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ksuo.mutation.CoinIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   keystore.CoinTable,
			Columns: []string{keystore.CoinColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &KeyStore{config: ksuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ksuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{keystore.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
