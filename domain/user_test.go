package domain

import (
	"testing"

	uuid "github.com/satori/go.uuid"
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
		{
			testCase: "When name is invalid must return an error",
			args: args{
				name:  "",
				email: "test@email.com",
				roles: []Role{
					"Admin",
				},
			},
			want:                 nil,
			errorMessageExpected: "name: Missing required field",
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

func TestSetRoles(t *testing.T) {
	user := User{}

	tests := []struct {
		testCase string
		roles    []Role
		expected []Role
	}{
		{
			testCase: "When Admin is present must return all roles",
			roles:    []Role{"Admin", "Modifier", "Watcher"},
			expected: []Role{"Admin", "Modifier", "Watcher"},
		},
		{
			testCase: "When Modifier is present must return Modifier and Watcher and ignore invalid roles",
			roles:    []Role{"Manager", "Modifier", "Viewer"},
			expected: []Role{"Modifier", "Watcher"},
		},
		{
			testCase: "When no valid roles are provided must return Watcher",
			roles:    []Role{"Manager", "Viewer"},
			expected: []Role{"Watcher"},
		},
		{
			testCase: "Empty Slice",
			roles:    []Role{},
			expected: []Role{"Watcher"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.testCase, func(t *testing.T) {
			actual := user.SetRoles(tt.roles)
			if !compareSlices(actual, tt.expected) {
				t.Errorf("Test case: %s - Expected roles: %v, got: %v", tt.testCase, tt.expected, actual)
			}
		})
	}
}

func compareSlices(s1, s2 []Role) bool {
	if len(s1) != len(s2) {
		return false
	}

	for i, v := range s1 {
		if v != s2[i] {
			return false
		}
	}

	return true
}

func TestValidate(t *testing.T) {
	validId := uuid.NewV4().String()
	type testCase struct {
		name        string
		user        User
		wantErr     bool
		wantedError string
	}

	tests := []testCase{
		{
			name: "Valid User",
			user: User{
				ID:    validId,
				Name:  "Test User",
				Email: "test@example.com",
				Roles: []Role{Admin},
			},
			wantErr: false,
		},
		{
			name: "When id is not provided must return an error",
			user: User{
				Name:  "Test User",
				Email: "test@example.com",
				Roles: []Role{Admin},
			},
			wantErr:     true,
			wantedError: "id: Missing required field",
		},
		{
			name: "When email is not valid must return an error",
			user: User{
				ID:    validId,
				Name:  "Test User",
				Email: "invalid_email",
				Roles: []Role{Admin},
			},
			wantErr:     true,
			wantedError: "email: invalid_email does not validate as email",
		},
		{
			name: "When name is empty must return an error",
			user: User{
				ID:    validId,
				Name:  "",
				Email: "test@example.com",
				Roles: []Role{Admin},
			},
			wantErr:     true,
			wantedError: "name: Missing required field",
		},
		{
			name: "When are not roles provided return an error",
			user: User{
				ID:    validId,
				Name:  "Test User",
				Email: "test@example.com",
				Roles: []Role{},
			},
			wantErr:     true,
			wantedError: "roles: Missing required field",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.user.Validate()

			if tt.wantErr {
				require.EqualError(t, err, tt.wantedError)
				return
			}
			require.Equal(t, err, nil)
		})
	}
}
