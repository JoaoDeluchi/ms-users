package domain

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewUser(t *testing.T) {

	type args struct {
		name  string
		email string
		roles []Role
	}

	tests := []struct {
		testCase string
		args
		want                 *User
		wantError            bool
		errorMessageExpected string
	}{
		{
			testCase: "When user is provided must return an user with the right values",
			args: args{
				name:  "test_name",
				email: "test_email@abc.com",
				roles: []Role{
					"Admin",
				},
			},
			want: &User{
				Name:  "test_name",
				Email: "test_email@abc.com",
				Roles: []Role{
					"Admin",
					"Modifier",
					"Watcher",
				},
			},
			wantError: false,
		},
		{
			testCase: "When email provided is invalid must return an error",
			args: args{
				name:  "test_name",
				email: "test_email",
				roles: []Role{
					"Admin",
				},
			},
			want:                 nil,
			errorMessageExpected: "email: test_email does not validate as email",
			wantError:            true,
		},
	}
	for _, tt := range tests {
		got, err := NewUser(tt.args.name, tt.args.email, tt.args.roles)
		if tt.wantError {
			require.EqualError(t, err, tt.errorMessageExpected)
			return
		}
		require.Equal(t, got.Email, tt.want.Email)
		require.Equal(t, got.Name, tt.want.Name)
		require.Equal(t, got.Roles, tt.want.Roles)
	}
}
