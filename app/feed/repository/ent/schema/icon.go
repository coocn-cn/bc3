package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// Icon holds the schema definition for the Feed entity.
type Icon struct {
	ent.Schema
}

// Fields returns schema fields.
func (Icon) Fields() []ent.Field {
	return []ent.Field{
		field.String("hash").Unique(),
		field.String("mime_type"),
		field.Bytes("content"),
	}
}

// Indexes of the schema.
func (Icon) Indexes() []ent.Index {
	return []ent.Index{}
}

// Edges returns schema edges.
func (Icon) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("feed", Feed.Type).Ref("icon").Required(),
	}
}

// Mixin of the schema.
func (Icon) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}
