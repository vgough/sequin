// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/vgough/sequin/storage/ent/label"
	"github.com/vgough/sequin/storage/ent/operation"
	"github.com/vgough/sequin/storage/ent/predicate"
)

// LabelUpdate is the builder for updating Label entities.
type LabelUpdate struct {
	config
	hooks    []Hook
	mutation *LabelMutation
}

// Where appends a list predicates to the LabelUpdate builder.
func (lu *LabelUpdate) Where(ps ...predicate.Label) *LabelUpdate {
	lu.mutation.Where(ps...)
	return lu
}

// SetName sets the "name" field.
func (lu *LabelUpdate) SetName(s string) *LabelUpdate {
	lu.mutation.SetName(s)
	return lu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (lu *LabelUpdate) SetNillableName(s *string) *LabelUpdate {
	if s != nil {
		lu.SetName(*s)
	}
	return lu
}

// SetValue sets the "value" field.
func (lu *LabelUpdate) SetValue(s string) *LabelUpdate {
	lu.mutation.SetValue(s)
	return lu
}

// SetNillableValue sets the "value" field if the given value is not nil.
func (lu *LabelUpdate) SetNillableValue(s *string) *LabelUpdate {
	if s != nil {
		lu.SetValue(*s)
	}
	return lu
}

// AddOperationIDs adds the "operation" edge to the Operation entity by IDs.
func (lu *LabelUpdate) AddOperationIDs(ids ...int) *LabelUpdate {
	lu.mutation.AddOperationIDs(ids...)
	return lu
}

// AddOperation adds the "operation" edges to the Operation entity.
func (lu *LabelUpdate) AddOperation(o ...*Operation) *LabelUpdate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return lu.AddOperationIDs(ids...)
}

// Mutation returns the LabelMutation object of the builder.
func (lu *LabelUpdate) Mutation() *LabelMutation {
	return lu.mutation
}

// ClearOperation clears all "operation" edges to the Operation entity.
func (lu *LabelUpdate) ClearOperation() *LabelUpdate {
	lu.mutation.ClearOperation()
	return lu
}

// RemoveOperationIDs removes the "operation" edge to Operation entities by IDs.
func (lu *LabelUpdate) RemoveOperationIDs(ids ...int) *LabelUpdate {
	lu.mutation.RemoveOperationIDs(ids...)
	return lu
}

// RemoveOperation removes "operation" edges to Operation entities.
func (lu *LabelUpdate) RemoveOperation(o ...*Operation) *LabelUpdate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return lu.RemoveOperationIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (lu *LabelUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, lu.sqlSave, lu.mutation, lu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (lu *LabelUpdate) SaveX(ctx context.Context) int {
	affected, err := lu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (lu *LabelUpdate) Exec(ctx context.Context) error {
	_, err := lu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lu *LabelUpdate) ExecX(ctx context.Context) {
	if err := lu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (lu *LabelUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(label.Table, label.Columns, sqlgraph.NewFieldSpec(label.FieldID, field.TypeInt))
	if ps := lu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := lu.mutation.Name(); ok {
		_spec.SetField(label.FieldName, field.TypeString, value)
	}
	if value, ok := lu.mutation.Value(); ok {
		_spec.SetField(label.FieldValue, field.TypeString, value)
	}
	if lu.mutation.OperationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   label.OperationTable,
			Columns: label.OperationPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(operation.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lu.mutation.RemovedOperationIDs(); len(nodes) > 0 && !lu.mutation.OperationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   label.OperationTable,
			Columns: label.OperationPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(operation.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lu.mutation.OperationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   label.OperationTable,
			Columns: label.OperationPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(operation.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, lu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{label.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	lu.mutation.done = true
	return n, nil
}

// LabelUpdateOne is the builder for updating a single Label entity.
type LabelUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *LabelMutation
}

// SetName sets the "name" field.
func (luo *LabelUpdateOne) SetName(s string) *LabelUpdateOne {
	luo.mutation.SetName(s)
	return luo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (luo *LabelUpdateOne) SetNillableName(s *string) *LabelUpdateOne {
	if s != nil {
		luo.SetName(*s)
	}
	return luo
}

// SetValue sets the "value" field.
func (luo *LabelUpdateOne) SetValue(s string) *LabelUpdateOne {
	luo.mutation.SetValue(s)
	return luo
}

// SetNillableValue sets the "value" field if the given value is not nil.
func (luo *LabelUpdateOne) SetNillableValue(s *string) *LabelUpdateOne {
	if s != nil {
		luo.SetValue(*s)
	}
	return luo
}

// AddOperationIDs adds the "operation" edge to the Operation entity by IDs.
func (luo *LabelUpdateOne) AddOperationIDs(ids ...int) *LabelUpdateOne {
	luo.mutation.AddOperationIDs(ids...)
	return luo
}

// AddOperation adds the "operation" edges to the Operation entity.
func (luo *LabelUpdateOne) AddOperation(o ...*Operation) *LabelUpdateOne {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return luo.AddOperationIDs(ids...)
}

// Mutation returns the LabelMutation object of the builder.
func (luo *LabelUpdateOne) Mutation() *LabelMutation {
	return luo.mutation
}

// ClearOperation clears all "operation" edges to the Operation entity.
func (luo *LabelUpdateOne) ClearOperation() *LabelUpdateOne {
	luo.mutation.ClearOperation()
	return luo
}

// RemoveOperationIDs removes the "operation" edge to Operation entities by IDs.
func (luo *LabelUpdateOne) RemoveOperationIDs(ids ...int) *LabelUpdateOne {
	luo.mutation.RemoveOperationIDs(ids...)
	return luo
}

// RemoveOperation removes "operation" edges to Operation entities.
func (luo *LabelUpdateOne) RemoveOperation(o ...*Operation) *LabelUpdateOne {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return luo.RemoveOperationIDs(ids...)
}

// Where appends a list predicates to the LabelUpdate builder.
func (luo *LabelUpdateOne) Where(ps ...predicate.Label) *LabelUpdateOne {
	luo.mutation.Where(ps...)
	return luo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (luo *LabelUpdateOne) Select(field string, fields ...string) *LabelUpdateOne {
	luo.fields = append([]string{field}, fields...)
	return luo
}

// Save executes the query and returns the updated Label entity.
func (luo *LabelUpdateOne) Save(ctx context.Context) (*Label, error) {
	return withHooks(ctx, luo.sqlSave, luo.mutation, luo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (luo *LabelUpdateOne) SaveX(ctx context.Context) *Label {
	node, err := luo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (luo *LabelUpdateOne) Exec(ctx context.Context) error {
	_, err := luo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (luo *LabelUpdateOne) ExecX(ctx context.Context) {
	if err := luo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (luo *LabelUpdateOne) sqlSave(ctx context.Context) (_node *Label, err error) {
	_spec := sqlgraph.NewUpdateSpec(label.Table, label.Columns, sqlgraph.NewFieldSpec(label.FieldID, field.TypeInt))
	id, ok := luo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Label.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := luo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, label.FieldID)
		for _, f := range fields {
			if !label.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != label.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := luo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := luo.mutation.Name(); ok {
		_spec.SetField(label.FieldName, field.TypeString, value)
	}
	if value, ok := luo.mutation.Value(); ok {
		_spec.SetField(label.FieldValue, field.TypeString, value)
	}
	if luo.mutation.OperationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   label.OperationTable,
			Columns: label.OperationPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(operation.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := luo.mutation.RemovedOperationIDs(); len(nodes) > 0 && !luo.mutation.OperationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   label.OperationTable,
			Columns: label.OperationPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(operation.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := luo.mutation.OperationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   label.OperationTable,
			Columns: label.OperationPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(operation.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Label{config: luo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, luo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{label.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	luo.mutation.done = true
	return _node, nil
}
