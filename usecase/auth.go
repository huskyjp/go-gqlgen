package usecase

import (
	"context"
	"errors"
	"fmt"
	"go-gqlgen/domain/apperror"
	domain "go-gqlgen/domain/entity"
	repository "go-gqlgen/domain/interface"

	"golang.org/x/crypto/bcrypt"
)

// want data from database
// type UsecaseAuthImpl struct {
// 	authNameRepository repository.AuthRepository
// }

// func NewUsecaseAuthImpl(
// 	authNameRepository repository.AuthRepository,
// ) usecase.UsecaseAuth {
// 	return &UsecaseAuthImpl{
// 		authNameRepository: authNameRepository,
// 	}
// }

// domainとの繋ぎこみ
type AuthService struct {
	UserRepo repository.UserRepository
}

func NewAuthService(u repository.UserRepository) *AuthService {
	return &AuthService{
		UserRepo: u,
	}
}

// サーバーに登録するのでUsecaseに入れる
func (au *AuthService) Register(ctx context.Context, input domain.AuthRegisterInput) (domain.AuthRegisterResponse, error) {
	// check
	input.Initialize()

	// check if validate
	if err := input.Validation(); err != nil {
		return domain.AuthRegisterResponse{}, err
	}

	// check if username is still available
	if _, err := au.UserRepo.GetByUserName(ctx, input.Username); !errors.Is(err, apperror.ErrNotFound) {
		return domain.AuthRegisterResponse{}, apperror.ErrUserNameIsTaken
	}

	// check if email is still available
	if _, err := au.UserRepo.GetByEmail(ctx, input.Email); !errors.Is(err, apperror.ErrNotFound) {
		return domain.AuthRegisterResponse{}, apperror.ErrUserNameIsTaken
	}

	// assign passed input
	user := domain.User{
		Email:    input.Email,
		Username: input.Username,
	}

	// hash password
	cryptedPass, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return domain.AuthRegisterResponse{}, fmt.Errorf("error happened when encrypting password: %v ", err)
	}

	user.Password = string(cryptedPass)

	user, err = au.UserRepo.GenerateUser(ctx, user)
	if err != nil {
		return domain.AuthRegisterResponse{}, fmt.Errorf("error happened when generating user: %v ", err)
	}
	return domain.AuthRegisterResponse{
		AccessToken: "Access Token",
		User:        user,
	}, nil
}
