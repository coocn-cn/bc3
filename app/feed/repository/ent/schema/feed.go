package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
)

// Feed holds the schema definition for the Feed entity.
type Feed struct {
	ent.Schema
}

// Fields returns schema fields.
func (Feed) Fields() []ent.Field {
	return []ent.Field{
		field.String("title"),
		field.String("feed_url"),
		field.String("site_url"),
		field.Time("checked_at").Default(time.Now),
		field.String("etag_header"),
		field.String("last_modified_header").Default(""),
		field.String("parsing_error_msg").Default(""),
		field.Int("parsing_error_count").Default(0),
	}
}

// Indexes of the schema.
func (Feed) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("feed_url").Edges("user").Unique(),
	}
}

// Edges returns schema edges.
func (Feed) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("feed").Required().Unique(),
		edge.From("categor", Categor.Type).Ref("feed").Required().Unique(),

		edge.To("entry", Entry.Type).StorageKey(edge.Column("feed_id")),
		edge.To("icon", Icon.Type).StorageKey(edge.Columns("feed_id", "icon_id")),
	}
}

// Mixin of the schema.
func (Feed) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}
