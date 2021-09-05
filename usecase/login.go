package usecase

import (
	"context"
	"errors"
	apperror "go-gqlgen/domain/apperror"
	domain "go-gqlgen/domain/entity"

	"golang.org/x/crypto/bcrypt"
)

func (au *AuthService) Login(ctx context.Context, input domain.LoginInput) (domain.AuthRegisterResponse, error) {

	input.Initialize()

	if err := input.Validation(); err != nil {
		return domain.AuthRegisterResponse{}, err
	}

	user, err := au.UserRepo.GetByEmail(ctx, input.Email)
	if err != nil {
		switch {
		case errors.Is(err, apperror.ErrNotFound):
			return domain.AuthRegisterResponse{}, apperror.ErrNotFoundUserNamePassword
		default:
			return domain.AuthRegisterResponse{}, err
		}
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		return domain.AuthRegisterResponse{}, apperror.ErrNotFoundUserNamePassword
	}

	return domain.AuthRegisterResponse{
		AccessToken: "Login Token",
		User:        user,
	}, nil
}
