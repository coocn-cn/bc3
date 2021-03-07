package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
)

// Entry holds the schema definition for the Feed entity.
type Entry struct {
	ent.Schema
}

// Fields returns schema fields.
func (Entry) Fields() []ent.Field {
	return []ent.Field{
		field.String("hash"),
		field.Time("published_at"),
		field.String("title"),
		field.String("url"),
		field.String("author"),
		field.String("content"),
		field.Enum("status").Values("unread", "read", "removed").Default("unread"),
	}
}

// Indexes of the schema.
func (Entry) Indexes() []ent.Index {
	return []ent.Index{
		index.Edges("feed"),
		index.Fields("hash").Edges("feed").Unique(),
	}
}

// Edges returns schema edges.
func (Entry) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("entry").Required().Unique(),
		edge.From("feed", Feed.Type).Ref("entry").Required().Unique(),

		edge.To("enclosur", Enclosur.Type).StorageKey(edge.Column("entry_id")),
	}
}

// Mixin of the schema.
func (Entry) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}
