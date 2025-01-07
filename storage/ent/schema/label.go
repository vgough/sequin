package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Label holds the schema definition for the Label entity.
type Label struct {
	ent.Schema
}

// Fields of the Label.
func (Label) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("value"),
	}
}

func (Label) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("name", "value").
			Unique(),
	}
}

// Edges of the Label.
func (Label) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("operation", Operation.Type).
			Ref("labels"),
	}
}
