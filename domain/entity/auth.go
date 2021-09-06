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
type AuthRegisterInput struct {
	Email        string
	Username     string
	Password     string
	AuthPassword string
}

// User Register Response
// the name is same as graphql
type AuthRegisterResponse struct {
	AccessToken string
	User        User
}

func (auth *AuthRegisterInput) Initialize() {
	// remove space
	auth.Email = strings.TrimSpace(auth.Email)
	auth.Email = strings.ToLower(auth.Email)

}

func (auth AuthRegisterInput) Validation() error {

	if len(auth.Username) < UserNameMinLength || !emailRegexp.MatchString(auth.Email) || len(auth.Password) < PasswordMinLength || auth.Password != auth.AuthPassword {
		return errors.New("validation errror when registration - you may want to fix email or password")
	}

	return nil

}
