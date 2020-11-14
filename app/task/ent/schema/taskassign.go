package schema

import (
	"github.com/coocn-cn/bc3/pkg/db/mixin"
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
)

// TaskAssign holds the schema definition for the TaskAssign entity.
type TaskAssign struct {
	ent.Schema
}

// Fields of the TaskAssign.
func (TaskAssign) Fields() []ent.Field {
	return nil
}

// Edges of the TaskAssign.
func (TaskAssign) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("task", Task.Type).Ref("assign").Required().Unique(),
		edge.From("user", User.Type).Ref("task_assign").Required().Unique(),
		edge.From("operator", User.Type).Ref("task_assign_operator").Required().Unique(),
	}
}

// Mixin of the Task.
func (TaskAssign) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.CreateTime{},
	}
}
