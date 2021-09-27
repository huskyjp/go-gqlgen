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

type UsecaseAuthImpl struct {
	userRepository repository.UserRepository
}

func NewUsecaseAuthImpl(userRepository repository.UserRepository) *UsecaseAuthImpl {
	return &UsecaseAuthImpl{
		userRepository: userRepository,
	}
}

// サーバーに登録するのでUsecaseに入れる
func (ur *UsecaseAuthImpl) Register(ctx context.Context, input domain.AuthRegisterInput) (domain.AuthRegisterResponse, error) {
	// check
	input.Initialize()

	// check if validate
	if err := input.Validation(); err != nil {
		return domain.AuthRegisterResponse{}, err
	}

	// check if username is still available
	_, err := ur.userRepository.GetByUserName(ctx, input.Username)
	if err != nil {
		return domain.AuthRegisterResponse{}, apperror.ErrUserNameIsTaken
	}
	// check if email is still available
	if _, err := ur.userRepository.GetByEmail(ctx, input.Email); !errors.Is(err, apperror.ErrNotFound) {
		return domain.AuthRegisterResponse{}, apperror.ErrEmailIsTaken
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

	user, err = ur.userRepository.GenerateUser(ctx, user)
	if err != nil {
		return domain.AuthRegisterResponse{}, fmt.Errorf("error happened when generating user: %v ", err)
	}
	return domain.AuthRegisterResponse{
		AccessToken: "Access Token",
		User:        user,
	}, nil
}
