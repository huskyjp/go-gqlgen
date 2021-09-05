package domain

import (
	"fmt"
	"go-gqlgen/domain/apperror"
	"strings"
)

type LoginInput struct {
	Email    string
	Password string
}

func (in *LoginInput) Initialize() {
	in.Email = strings.TrimSpace(in.Email)
	in.Email = strings.ToLower(in.Email)
}

func (in LoginInput) Validation() error {
	if !emailRegexp.MatchString(in.Email) {
		return fmt.Errorf("%w: that is not email - please recheck your email", apperror.ErrValidation)
	}

	// we don't check if password is same as stored in database that's usecase job!
	if len(in.Password) < 1 {
		return fmt.Errorf("%w: password is not provided - please recheck your email", apperror.ErrValidation)
	}

	return nil
}
