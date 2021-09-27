package usecase

import (
	"context"
	domain "go-gqlgen/domain/entity"
)

// resolver
type UsecaseAuthUser interface {
	Create(ctx context.Context, user domain.User) (domain.User, error)
	GetByUserName(ctx context.Context, username string) (domain.User, error)
	GetByEmail(ctx context.Context, email string) (domain.User, error)
}
