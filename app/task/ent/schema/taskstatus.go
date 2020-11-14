package schema

import (
	"github.com/coocn-cn/bc3/pkg/db/mixin"

	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// TaskStatus holds the schema definition for the TaskStatus entity.
type TaskStatus struct {
	ent.Schema
}

// Fields of the TaskStatus.
func (TaskStatus) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.Enum("type").ValueMap(map[string]string{"0": "无效", "1": "通常", "2": "结束"}),
	}
}

// Edges of the TaskStatus.
func (TaskStatus) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("task", Task.Type).StorageKey(edge.Column("status_id")),

		edge.From("operator", User.Type).Ref("task_status_operator").Required().Unique(),
	}
}

// Mixin of the Task.
func (TaskStatus) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.CreateTime{},
	}
}
