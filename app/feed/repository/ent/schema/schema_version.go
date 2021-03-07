package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// SchemaVersion holds the schema definition for the TaskStatus entity.
type SchemaVersion struct {
	ent.Schema
}

// Fields of the SchemaVersion.
func (SchemaVersion) Fields() []ent.Field {
	return []ent.Field{
		field.String("version"),
	}
}

// Edges of the SchemaVersion.
func (SchemaVersion) Edges() []ent.Edge {
	return []ent.Edge{}
}

// Mixin of the Task.
func (SchemaVersion) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}
