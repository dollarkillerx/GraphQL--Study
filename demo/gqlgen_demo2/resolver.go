package gqlgen_demo2

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

import (
	"context"
	"graphql_demo/demo/gqlgen_demo2/graph"
	"graphql_demo/demo/gqlgen_demo2/model"
)

type Resolver struct{}

func (r *mutationResolver) CreateTodo(ctx context.Context, input *model.NewTodo) (*model.Todo, error) {
	panic("not implemented")
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	panic("not implemented")
}

func (r *todoResolver) User(ctx context.Context, obj *model.Todo) (*model.User, error) {
	panic("not implemented")
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
