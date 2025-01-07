// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/vgough/sequin/storage/ent/label"
	"github.com/vgough/sequin/storage/ent/operation"
)

// LabelCreate is the builder for creating a Label entity.
type LabelCreate struct {
	config
	mutation *LabelMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (lc *LabelCreate) SetName(s string) *LabelCreate {
	lc.mutation.SetName(s)
	return lc
}

// SetValue sets the "value" field.
func (lc *LabelCreate) SetValue(s string) *LabelCreate {
	lc.mutation.SetValue(s)
	return lc
}

// AddOperationIDs adds the "operation" edge to the Operation entity by IDs.
func (lc *LabelCreate) AddOperationIDs(ids ...string) *LabelCreate {
	lc.mutation.AddOperationIDs(ids...)
	return lc
}

// AddOperation adds the "operation" edges to the Operation entity.
func (lc *LabelCreate) AddOperation(o ...*Operation) *LabelCreate {
	ids := make([]string, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return lc.AddOperationIDs(ids...)
}

// Mutation returns the LabelMutation object of the builder.
func (lc *LabelCreate) Mutation() *LabelMutation {
	return lc.mutation
}

// Save creates the Label in the database.
func (lc *LabelCreate) Save(ctx context.Context) (*Label, error) {
	return withHooks(ctx, lc.sqlSave, lc.mutation, lc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (lc *LabelCreate) SaveX(ctx context.Context) *Label {
	v, err := lc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (lc *LabelCreate) Exec(ctx context.Context) error {
	_, err := lc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lc *LabelCreate) ExecX(ctx context.Context) {
	if err := lc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (lc *LabelCreate) check() error {
	if _, ok := lc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Label.name"`)}
	}
	if _, ok := lc.mutation.Value(); !ok {
		return &ValidationError{Name: "value", err: errors.New(`ent: missing required field "Label.value"`)}
	}
	return nil
}

func (lc *LabelCreate) sqlSave(ctx context.Context) (*Label, error) {
	if err := lc.check(); err != nil {
		return nil, err
	}
	_node, _spec := lc.createSpec()
	if err := sqlgraph.CreateNode(ctx, lc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	lc.mutation.id = &_node.ID
	lc.mutation.done = true
	return _node, nil
}

func (lc *LabelCreate) createSpec() (*Label, *sqlgraph.CreateSpec) {
	var (
		_node = &Label{config: lc.config}
		_spec = sqlgraph.NewCreateSpec(label.Table, sqlgraph.NewFieldSpec(label.FieldID, field.TypeInt))
	)
	if value, ok := lc.mutation.Name(); ok {
		_spec.SetField(label.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := lc.mutation.Value(); ok {
		_spec.SetField(label.FieldValue, field.TypeString, value)
		_node.Value = value
	}
	if nodes := lc.mutation.OperationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   label.OperationTable,
			Columns: label.OperationPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(operation.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// LabelCreateBulk is the builder for creating many Label entities in bulk.
type LabelCreateBulk struct {
	config
	err      error
	builders []*LabelCreate
}

// Save creates the Label entities in the database.
func (lcb *LabelCreateBulk) Save(ctx context.Context) ([]*Label, error) {
	if lcb.err != nil {
		return nil, lcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(lcb.builders))
	nodes := make([]*Label, len(lcb.builders))
	mutators := make([]Mutator, len(lcb.builders))
	for i := range lcb.builders {
		func(i int, root context.Context) {
			builder := lcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*LabelMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, lcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, lcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, lcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (lcb *LabelCreateBulk) SaveX(ctx context.Context) []*Label {
	v, err := lcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (lcb *LabelCreateBulk) Exec(ctx context.Context) error {
	_, err := lcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lcb *LabelCreateBulk) ExecX(ctx context.Context) {
	if err := lcb.Exec(ctx); err != nil {
		panic(err)
	}
}
