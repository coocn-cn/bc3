package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("username").NotEmpty().Unique(),
		field.String("password"),
		field.String("group").Default("user"),
		field.String("theme").Default("default"),
		field.String("timezone").Default("UTC"),
		field.String("language").Default("en_US"),
		field.Time("last_login_at").Default(time.Now).UpdateDefault(time.Now).Immutable(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("feed", Feed.Type).StorageKey(edge.Column("user_id")),
		edge.To("entry", Entry.Type).StorageKey(edge.Column("user_id")),
		edge.To("session", Session.Type).StorageKey(edge.Column("user_id")),
		edge.To("categor", Categor.Type).StorageKey(edge.Column("user_id")),
		edge.To("enclosur", Enclosur.Type).StorageKey(edge.Column("user_id")),
	}
}

// Mixin of the Task.
func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}
