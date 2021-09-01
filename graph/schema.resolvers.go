package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	generated1 "go-gqlgen/graph/generated"
	"go-gqlgen/graph/model"
)

func (r *mutationResolver) CreateCountries(ctx context.Context, input model.NewCountries) (*model.Countries, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Countries(ctx context.Context) ([]*model.Countries, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated1.MutationResolver implementation.
func (r *Resolver) Mutation() generated1.MutationResolver { return &mutationResolver{r} }

// Query returns generated1.QueryResolver implementation.
func (r *Resolver) Query() generated1.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
