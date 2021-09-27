package domain

import (
	"context"
	domain "go-gqlgen/domain/entity"
)

// ここで宣言するFunctionはUsecaseに実態がある
// Modelに入れ込まれたデータを使ってくれる、interfaceで橋渡し

// AuthService -> Registration & Login
type AuthService interface {
	Register(ctx context.Context, input domain.RegisterInput) (domain.RegisterResponse, error)
	Login(ctx context.Context, input domain.AuthLoginInput) (domain.AuthLoginResponse, error)
}

// UserRepository -> Find something
type UserRepository interface {
	Create(ctx context.Context, user domain.User) (domain.User, error)
	GetByUserName(ctx context.Context, username string) (domain.User, error)
	GetByEmail(ctx context.Context, email string) (domain.User, error)
}
