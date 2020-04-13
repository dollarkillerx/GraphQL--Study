package resolver

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

import (
	"context"
	"graphql_demo/demo/gqlgen_demo/graph"
	"graphql_demo/demo/gqlgen_demo/model"
)

type Resolver struct{}

func (r *mutationResolver) CreateTodo(ctx context.Context, input *model.NewTodo) (*model.Todo, error) {
	panic("not implemented")
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	panic("not implemented")
}

// Mutation returns graph.MutationResolver implementation.
func (r *Resolver) Mutation() graph.MutationResolver { return &mutationResolver{r} }

// Query returns graph.QueryResolver implementation.
func (r *Resolver) Query() graph.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
