package schema

import (
	"github.com/coocn-cn/bc3/pkg/db/mixin"

	"github.com/facebook/ent"
	"github.com/facebook/ent/dialect"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// Task holds the schema definition for the Task entity.
type Task struct {
	ent.Schema
}

// Fields of the Task.
func (Task) Fields() []ent.Field {
	return []ent.Field{
		field.String("title"),
		field.String("content"),
		field.Time("end_time").SchemaType(map[string]string{dialect.Postgres: "date"}).Nillable().Optional(),
		field.Time("start_time").SchemaType(map[string]string{dialect.Postgres: "date"}).Nillable().Optional(),
		field.Int("man_hour").Default(0),
	}
}

// Edges of the Task.
func (Task) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("child", Task.Type).StorageKey(edge.Column("parent_id")).From("parent").Unique(),
		edge.To("assign", TaskAssign.Type).StorageKey(edge.Column("task_id")),

		edge.From("status", TaskStatus.Type).Ref("task").Required().Unique(),
		edge.From("operator", User.Type).Ref("task_operator").Required().Unique(),
	}
}

// Mixin of the Task.
func (Task) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.CreateTime{},
	}
}
