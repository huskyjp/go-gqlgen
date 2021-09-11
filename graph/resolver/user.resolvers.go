package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	domain "go-gqlgen/domain/entity"
	"go-gqlgen/graph/model"
	"log"
)

func (r *mutationResolver) Register(ctx context.Context, input model.RegisterInput) (*model.AuthRegisterResponse, error) {
	res, err := r.AuthRepository.Register(ctx, domain.AuthRegisterInput{
		Email:        input.Email,
		Username:     input.Username,
		Password:     input.Password,
		AuthPassword: input.AuthPassword,
	})
	if err != nil {
		log.Printf("error happend when registration: %v", err)
		return nil, formatErrorResponse(ctx, err)
	}

	return mapAuthRegisterResponse(res), nil
}

func (r *mutationResolver) Login(ctx context.Context, input model.LoginInput) (*model.AuthRegisterResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Me(ctx context.Context) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func mapAuthRegisterResponse(r domain.AuthRegisterResponse) *model.AuthRegisterResponse {
	return &model.AuthRegisterResponse{
		AccessToken: r.AccessToken,
		User: &model.User{
			ID:        r.User.ID,
			Email:     r.User.Email,
			Username:  r.User.Username,
			CreatedAt: r.User.CreatedAt,
		},
	}
}
