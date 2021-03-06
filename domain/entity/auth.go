package domain

import (
	"errors"
	"regexp"
	"strings"
)

var (
	UserNameMinLength = 4
	PasswordMinLength = 8
)

// email validation
var emailRegexp = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

// User Input
type RegisterInput struct {
	Email        string
	Username     string
	Password     string
	AuthPassword string
}

// User Register Response
// the name is same as graphql
type RegisterResponse struct {
	AccessToken string
	User        User
}

func (input *RegisterInput) Initialize() {
	// remove space
	input.Email = strings.TrimSpace(input.Email)
	input.Email = strings.ToLower(input.Email)

}

func (auth RegisterInput) Validation() error {

	if len(auth.Username) < UserNameMinLength || !emailRegexp.MatchString(auth.Email) || len(auth.Password) < PasswordMinLength || auth.Password != auth.AuthPassword {
		return errors.New("validation errror when registration - you may want to fix email or password")
	}

	return nil

}
