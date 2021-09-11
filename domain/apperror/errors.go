package apperror

import "errors"

var (
	ErrNotFound                 = errors.New("not found")
	ErrValidation               = errors.New("validation error")
	ErrUserNameIsTaken          = errors.New("username is already taken")
	ErrEmailIsTaken             = errors.New("email is already taen")
	ErrNotFoundUserNamePassword = errors.New("username and password does not match")
)
