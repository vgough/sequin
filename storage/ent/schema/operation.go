package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// Operation holds the schema definition for the Operation entity.
type Operation struct {
	ent.Schema
}

// Fields of the Operation.
func (Operation) Fields() []ent.Field {
	return []ent.Field{
		field.String("id"),
		field.Bytes("detail"),
		field.Bytes("status"),
		field.Bool("is_done"),
	}
}

// Edges of the Operation.
func (Operation) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("labels", Label.Type),
	}
}

func (Operation) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}
