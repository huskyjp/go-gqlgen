package domain

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAuthRegisterInput_Validation(t *testing.T) {

	test := []struct {
		name  string
		input AuthRegisterInput
		err   error
	}{
		{
			name: "validation",
			input: AuthRegisterInput{
				Username:     "flutter_gopher",
				Email:        "husky@husky.com",
				Password:     "password_password",
				AuthPassword: "password_password",
			},
			err: nil,
		}, {
			name: "error",
			input: AuthRegisterInput{
				Username:     "error_graph",
				Email:        "err@err.com",
				Password:     "err",
				AuthPassword: "err",
			},
			err: nil,
		},
	}

	for _, test := range test {
		t.Run(test.name, func(t *testing.T) {
			err := test.input.Validation()

			if test.err != nil {
				require.ErrorIs(t, err, test.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
