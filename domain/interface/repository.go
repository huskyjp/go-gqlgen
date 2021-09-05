package domain

import (
	"context"
	domain "go-gqlgen/domain/entity"
)

// ここで宣言するFunctionはUsecaseに実態がある
// Modelに入れ込まれたデータを使ってくれる、interfaceで橋渡し

type AuthRepository interface {
	Register(ctx context.Context, input domain.AuthRegisterInput) (domain.AuthRegisterResponse, error)
	Login(ctx context.Context, input domain.LoginInput) (domain.AuthRegisterResponse, error)
}

type UserRepository interface {
	GenerateUser(ctx context.Context, user domain.User) (domain.User, error)
	GetByUserName(ctx context.Context, username string) (domain.User, error)
	GetByEmail(ctx context.Context, email string) (domain.User, error)
}
