package schema

import (
	"github.com/coocn-cn/bc3/pkg/db/mixin"
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("email"),
		field.String("phone"),
		field.String("password"),
		field.String("nickname"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("task_assign", TaskAssign.Type).StorageKey(edge.Column("user_id")),

		edge.To("task_operator", Task.Type).StorageKey(edge.Column("operator_id")),
		edge.To("task_assign_operator", TaskAssign.Type).StorageKey(edge.Column("operator_id")),
		edge.To("task_status_operator", TaskStatus.Type).StorageKey(edge.Column("operator_id")),
	}
}

// Mixin of the Task.
func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}
