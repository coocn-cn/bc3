package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/99designs/gqlgen/graphql"
	"github.com/coocn-cn/bc3/app/task/api/graph"
	"github.com/coocn-cn/bc3/app/task/ent"
	"github.com/coocn-cn/bc3/app/task/ent/task"
	"github.com/coocn-cn/bc3/app/task/ent/taskassign"
	"github.com/coocn-cn/bc3/app/task/ent/taskstatus"
	"github.com/coocn-cn/bc3/app/task/ent/todo"
	"github.com/coocn-cn/bc3/app/task/ent/user"
	"github.com/hashicorp/go-multierror"
	"github.com/vektah/gqlparser/v2/ast"
)

func (r *queryResolver) Node(ctx context.Context, id int) (ent.Noder, error) {
	field := graphql.GetFieldContext(ctx)
	tables := map[string]string{
		ent.TypeUser:       user.Table,
		ent.TypeTodo:       todo.Table,
		ent.TypeTask:       task.Table,
		ent.TypeTaskAssign: taskassign.Table,
		ent.TypeTaskStatus: taskstatus.Table,
	}

	var merr error
	for _, sel := range field.Field.SelectionSet {
		switch sel := sel.(type) {
		case *ast.Field:
			return r.client.Noder(ctx, id, ent.WithNodeType(tables[sel.Name]))
		case *ast.InlineFragment:
			noder, err := r.client.Noder(ctx, id, ent.WithNodeType(tables[sel.TypeCondition]))
			if ent.IsNotFound(err) {
				merr = multierror.Append(merr, err)
				continue
			}

			return noder, err
		case *ast.FragmentSpread:
			return nil, nil
		default:
			panic(fmt.Errorf("unsupported %T", sel))
		}
	}

	return nil, merr
}

func (r *queryResolver) Todos(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.TodoOrder) (*ent.TodoConnection, error) {
	return r.client.Todo.Query().
		Paginate(ctx, after, first, before, last,
			ent.WithTodoOrder(orderBy),
		)
}

// Query returns graph.QueryResolver implementation.
func (r *Resolver) Query() graph.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
