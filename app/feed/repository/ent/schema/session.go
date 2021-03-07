package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
)

// Session holds the schema definition for the Session entity.
type Session struct {
	ent.Schema
}

// Fields of the schema.
func (Session) Fields() []ent.Field {
	return []ent.Field{
		field.String("token").Unique(),
		field.String("user_agent"),
		field.String("ip"),
	}
}

// Indexes of the schema.
func (Session) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("token").Edges("user").Unique(),
	}
}

// Edges of the schema.
func (Session) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("session").Required().Unique(),
	}
}

// Mixin of the schema.
func (Session) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}
