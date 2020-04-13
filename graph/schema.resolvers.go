package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"graphql_demo/graph/generated"
	"graphql_demo/graph/model"
	"sync"
)
const (
	todo = "TODO"
)

var (
	dbMu sync.Mutex
	db map[string]interface{}
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	dbMu.Lock()
	defer dbMu.Unlock()
	ic:
	i,b := db[todo]
	if !b {
		db[todo] = make([]*model.Todo,0)
		goto ic
	}
	item := &model.Todo{
		ID: input.UserID,
		Text: input.Text,
		Done: true,
		User: nil,
	}
	ic := i.([]*model.Todo)
	ic = append(ic,item)
	db[todo] = ic
	return item,nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	dbMu.Lock()
	defer dbMu.Unlock()
	i,b := db[todo]
	if !b {
		return nil,errors.New("not data")
	}
	ic,c := i.([]*model.Todo)
	if !c {
		return nil,errors.New("not data")
	}
	return ic,nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
