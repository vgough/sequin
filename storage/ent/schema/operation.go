package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
)

// Operation holds the schema definition for the Operation entity.
type Operation struct {
	ent.Schema
}

// Fields of the Operation.
func (Operation) Fields() []ent.Field {
	return []ent.Field{
		field.String("request_id"),
		field.Int64("shard").Default(0),
		field.Bytes("detail"),
		field.Time("next_check_at").Optional(),
		field.Bytes("state").Optional(),
		field.Bytes("result").Optional(),
		field.String("submitter"),
		field.Time("started_at").Optional(),
		field.Time("finished_at").Optional(),
	}
}

func (Operation) Index() []ent.Index {
	return []ent.Index{
		index.Fields("shard", "request_id"),
		index.Fields("shard", "next_check_at").Annotations(
			entsql.IndexWhere("next_check_at IS NOT NULL"),
		),
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
