package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Pet struct {
	ent.Schema
}

func (Pet) Fields() []ent.Field {
	return []ent.Field{
		field.Int("age").Positive(),
		field.String("name").Default("unknown"),
	}
}

func (Pet) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("pets", Pet.Type),
	}
}
