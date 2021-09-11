package usecase

import (
	domain "go-gqlgen/domain/entity"
	"go-gqlgen/graph/model"
)

// resolver
type UsecaseAuthUser interface {
	GenerateUser(input model.RegisterInput) (*domain.User, error)
	// GetByUserName(ctx context.Context, username string) (domain.User, error)
	// GetByEmail(ctx context.Context, email string) (domain.User, error)
}
