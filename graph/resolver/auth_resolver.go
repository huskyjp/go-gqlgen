package graph

import (
	"context"
	"fmt"
	domain "go-gqlgen/domain/entity"
	"go-gqlgen/graph/model"
	"log"
)

func (r *mutationResolver) Register(ctx context.Context, input model.RegisterInput) (*domain.AuthRegisterResponse, error) {
	res, err := r.AuthRepository.Register(ctx, domain.AuthRegisterInput{
		Email:        input.Email,
		Username:     input.Username,
		Password:     input.Password,
		AuthPassword: input.AuthPassword,
	})
	if err != nil {
		log.Fatalln("error happend when registration: %v", err)

	}

	return mapAuthResponse(res), nil
}

func mapAuthResponse(r domain.AuthRegisterResponse) *domain.AuthRegisterResponse {
	return &domain.AuthRegisterResponse{
		AccessToken: r.AccessToken,
		User:        *mapUser(r.User),
	}
}

func mapUser(u domain.User) *domain.User {
	return &domain.User{
		ID:        u.ID,
		Email:     u.Email,
		Username:  u.Username,
		CreatedAt: u.CreatedAt,
	}
}

func (r *mutationResolver) Login(ctx context.Context, input model.LoginInput) (*model.AuthResponse, error) {
	panic(fmt.Errorf("not implemented"))
}
