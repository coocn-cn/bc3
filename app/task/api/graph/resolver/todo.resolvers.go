package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/coocn-cn/bc3/app/task/api/graph"
	"github.com/coocn-cn/bc3/app/task/ent"
)

func (r *todoResolver) Parent(ctx context.Context, obj *ent.Todo) (*ent.Todo, error) {
	parent, err := obj.Edges.ParentOrErr()
	return parent, ent.MaskNotFound(err)
}

func (r *todoResolver) Children(ctx context.Context, obj *ent.Todo) ([]*ent.Todo, error) {
	return obj.Edges.ChildrenOrErr()
}

// Todo returns graph.TodoResolver implementation.
func (r *Resolver) Todo() graph.TodoResolver { return &todoResolver{r} }

type todoResolver struct{ *Resolver }
