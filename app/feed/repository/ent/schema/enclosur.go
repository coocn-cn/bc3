package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// Enclosur holds the schema definition for the Feed entity.
type Enclosur struct {
	ent.Schema
}

// Fields returns schema fields.
func (Enclosur) Fields() []ent.Field {
	return []ent.Field{
		field.String("url"),
		field.Int("size").Default(0),
		field.String("mime_type").Default(""),
	}
}

// Indexes of the schema.
func (Enclosur) Indexes() []ent.Index {
	return []ent.Index{}
}

// Edges returns schema edges.
func (Enclosur) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("enclosur").Required().Unique(),
		edge.From("entry", Entry.Type).Ref("enclosur").Required().Unique(),
	}
}

// Mixin of the schema.
func (Enclosur) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}
