package usecase

import (
	"context"
	"errors"
	apperror "go-gqlgen/domain/apperror"
	domain "go-gqlgen/domain/entity"

	"golang.org/x/crypto/bcrypt"
)

func (au *UsecaseAuthImpl) Login(ctx context.Context, input domain.AuthLoginInput) (domain.AuthLoginResponse, error) {

	input.Initialize()

	if err := input.Validation(); err != nil {
		return domain.AuthLoginResponse{}, err
	}

	user, err := au.userRepository.GetByEmail(ctx, input.Email)
	if err != nil {
		switch {
		case errors.Is(err, apperror.ErrNotFound):
			return domain.AuthLoginResponse{}, apperror.ErrNotFoundUserNamePassword
		default:
			return domain.AuthLoginResponse{}, err
		}
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		return domain.AuthLoginResponse{}, apperror.ErrNotFoundUserNamePassword
	}

	return domain.AuthLoginResponse{
		AccessToken: "Login Token",
		User:        user,
	}, nil
}
