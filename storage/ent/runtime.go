// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/vgough/sequin/storage/ent/operation"
	"github.com/vgough/sequin/storage/ent/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	operationMixin := schema.Operation{}.Mixin()
	operationMixinFields0 := operationMixin[0].Fields()
	_ = operationMixinFields0
	operationFields := schema.Operation{}.Fields()
	_ = operationFields
	// operationDescCreateTime is the schema descriptor for create_time field.
	operationDescCreateTime := operationMixinFields0[0].Descriptor()
	// operation.DefaultCreateTime holds the default value on creation for the create_time field.
	operation.DefaultCreateTime = operationDescCreateTime.Default.(func() time.Time)
	// operationDescUpdateTime is the schema descriptor for update_time field.
	operationDescUpdateTime := operationMixinFields0[1].Descriptor()
	// operation.DefaultUpdateTime holds the default value on creation for the update_time field.
	operation.DefaultUpdateTime = operationDescUpdateTime.Default.(func() time.Time)
	// operation.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	operation.UpdateDefaultUpdateTime = operationDescUpdateTime.UpdateDefault.(func() time.Time)
	// operationDescShard is the schema descriptor for shard field.
	operationDescShard := operationFields[1].Descriptor()
	// operation.DefaultShard holds the default value on creation for the shard field.
	operation.DefaultShard = operationDescShard.Default.(int64)
}
