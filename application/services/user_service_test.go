package services

import (
	"fmt"
	"testing"

	"github.com/joaodeluchi/ms-users/application/repositories"
	"github.com/joaodeluchi/ms-users/domain"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestAlreadyHaveTheUser(t *testing.T) {
	type args struct {
		email string
	}

	type testCase struct {
		name        string
		fields      repositories.UserRepository
		args        args
		want        bool
		wantError   bool
		wantedError string
	}

	tests := []testCase{
		{
			name: "When repository returns an user must return true and error",
			args: args{
				email: "email@valid.com",
			},
			fields:      UserRepositoryMock{},
			want:        true,
			wantError:   true,
			wantedError: "user with email test@test.com already exist",
		},
		{
			name: "When repository dont return an user must return false",
			args: args{
				email: "email@valid.com",
			},
			fields: UserRepositoryMock{
				wantErr: false,
			},
			want:      false,
			wantError: false,
		},
	}

	for _, tt := range tests {
		service := NewUserService(tt.fields)

		got, err := service.alreadyHaveTheUser(tt.args.email)
		fmt.Println(got)
		if tt.wantError {
			require.EqualError(t, err, tt.wantedError)
			return
		}
		require.Equal(t, got, tt.want)
	}
}

// func TestCreateUser(t *testing.T) {
// 	type args struct {
// 		user domain.User
// 	}

// 	type testCase struct {
// 		testCase string
// 		fields   repositories.UserRepository
// 		args     args
// 		want     error
// 	}

// 	tests := []testCase{
// 		{
// 			testCase: "When dont have any error must return nil",
// 			args: args{
// 				domain.User{
// 					Email: "testing-@email.com",
// 				},
// 			},
// 			fields: UserRepositoryMock{
// 				wantErr: false,
// 			},
// 			want: nil,
// 		},
// 	}

// 	for _, tt := range tests {
// 		service := NewUserService(tt.fields)

// 		got := service.CreateUser(tt.args.user)
// 		require.Equal(t, tt.want, got.Error())
// 	}
// }

type UserRepositoryMock struct {
	wantErr bool
}

// DeleteUser implements repositories.UserRepository.
func (usm UserRepositoryMock) DeleteUser(id string) ([]*domain.User, error) {
	if usm.wantErr {
		return nil, fmt.Errorf("repository error")
	}

	return []*domain.User{
		{
			ID:    uuid.NewV4().String(),
			Name:  "test",
			Email: "test@test.com",
			Roles: []domain.Role{
				"Admin",
			},
		},
	}, nil
}

// FindById implements repositories.UserRepository.
func (usm UserRepositoryMock) FindById(id string) (*domain.User, error) {
	if usm.wantErr {
		return nil, fmt.Errorf("repository error")
	}

	return &domain.User{
		ID:    uuid.NewV4().String(),
		Name:  "test",
		Email: "test2@test.com",
		Roles: []domain.Role{
			"Admin",
		},
	}, nil
}

// GetAll implements repositories.UserRepository.
func (usm UserRepositoryMock) GetAll() []*domain.User {
	if usm.wantErr {
		return nil
	}

	return []*domain.User{
		{
			ID:    uuid.NewV4().String(),
			Name:  "test",
			Email: "test3@test.com",
			Roles: []domain.Role{
				"Admin",
			},
		},
	}
}

// Insert implements repositories.UserRepository.
func (usm UserRepositoryMock) Insert(user *domain.User) (*domain.User, error) {
	if usm.wantErr {
		return nil, fmt.Errorf("repository error")
	}

	return &domain.User{
		ID:    uuid.NewV4().String(),
		Name:  "test",
		Email: "test@test.com",
		Roles: []domain.Role{
			"Admin",
		},
	}, nil
}

// UpdateRoles implements repositories.UserRepository.
func (usm UserRepositoryMock) UpdateRoles(userId string, roles []domain.Role) (*domain.User, error) {
	if usm.wantErr {
		return nil, fmt.Errorf("repository error")
	}

	return &domain.User{
		ID:    uuid.NewV4().String(),
		Name:  "test",
		Email: "test@test.com",
		Roles: []domain.Role{
			"Admin",
		},
	}, nil
}

func (usm UserRepositoryMock) FindByEmail(email string) (*domain.User, error) {
	if usm.wantErr {
		return nil, fmt.Errorf("repository error")
	}

	return &domain.User{
		ID:    uuid.NewV4().String(),
		Name:  "test",
		Email: "test1@test.com",
		Roles: []domain.Role{
			"Admin",
		},
	}, nil
}
