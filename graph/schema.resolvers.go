package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"graphql_demo/graph/generated"
	"graphql_demo/graph/model"
)

func (r *queryResolver) People(ctx context.Context, id string) (*model.People, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Peoples(ctx context.Context, first *int, after *string) (*model.PeopleConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
