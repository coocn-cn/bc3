package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/coocn-cn/bc3/app/task/api/graph"
	"github.com/coocn-cn/bc3/app/task/api/graph/model"
	"github.com/coocn-cn/bc3/app/task/ent"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, todo model.TodoInput) (*ent.Todo, error) {
	client := ent.FromContext(ctx)
	return client.Todo.
		Create().
		SetStatus(todo.Status).
		SetNillablePriority(todo.Priority).
		SetText(todo.Text).
		SetNillableParentID(todo.Parent).
		Save(ctx)
}

func (r *mutationResolver) ClearTodos(ctx context.Context) (int, error) {
	client := ent.FromContext(ctx)
	return client.Todo.
		Delete().
		Exec(ctx)
}

func (r *queryResolver) Node(ctx context.Context, id int) (ent.Noder, error) {
	return r.client.Noder(ctx, id)
}

func (r *queryResolver) Todos(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.TodoOrder) (*ent.TodoConnection, error) {
	return r.client.Todo.Query().
		Paginate(ctx, after, first, before, last,
			ent.WithTodoOrder(orderBy),
		)
}

func (r *todoResolver) Parent(ctx context.Context, obj *ent.Todo) (*ent.Todo, error) {
	parent, err := obj.Edges.ParentOrErr()
	return parent, ent.MaskNotFound(err)
}

func (r *todoResolver) Children(ctx context.Context, obj *ent.Todo) ([]*ent.Todo, error) {
	return obj.Edges.ChildrenOrErr()
}

// Mutation returns graph.MutationResolver implementation.
func (r *Resolver) Mutation() graph.MutationResolver { return &mutationResolver{r} }

// Query returns graph.QueryResolver implementation.
func (r *Resolver) Query() graph.QueryResolver { return &queryResolver{r} }

// Todo returns graph.TodoResolver implementation.
func (r *Resolver) Todo() graph.TodoResolver { return &todoResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type todoResolver struct{ *Resolver }
