package repositories

import (
	"reflect"
	"testing"

	"github.com/joaodeluchi/ms-users/domain"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func getTestDatabase() []*domain.User {
	return []*domain.User{
		{
			ID:    "123123",
			Name:  "test-name",
			Email: "test@email.com",
			Roles: []domain.Role{
				domain.Watcher,
			},
		},
		{
			ID:    "321321",
			Name:  "test-name",
			Email: "test-1@email.com",
			Roles: []domain.Role{
				domain.Watcher,
			},
		},
	}
}

func TestNewUserRepository(t *testing.T) {
	db := []*domain.User{}
	tests := []struct {
		name string
		want UserRepository
		args []*domain.User
	}{
		{
			name: "Must return a repository instance",
			args: db,
			want: &UserRepositoryDb{
				dbInMemory: db,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserRepository(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserRepositoryDb_DeleteUser(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		args    args
		want    []*domain.User
		wantErr bool
	}{
		{
			name: "Must remove user from database when the id is provided",

			args: args{"123123"},
			want: []*domain.User{
				{
					ID:    "321321",
					Name:  "test-name",
					Email: "test@email.com",
					Roles: []domain.Role{
						domain.Watcher,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "Must return an error when the id provided is not found",
			args: args{"456456"},
			want: []*domain.User{
				{
					ID:    "123123",
					Name:  "test-name",
					Email: "test@email.com",
					Roles: []domain.Role{
						domain.Watcher,
					},
				},
				{
					ID:    "321321",
					Name:  "test-name",
					Email: "test@email.com",
					Roles: []domain.Role{
						domain.Watcher,
					},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		db := getTestDatabase()

		t.Run(tt.name, func(t *testing.T) {
			ur := &UserRepositoryDb{
				dbInMemory: db,
			}

			require.Equal(t, 2, len(ur.dbInMemory))
			got, err := ur.DeleteUser(tt.args.id)

			if !tt.wantErr {
				require.Equal(t, 1, len(ur.dbInMemory))
				require.Equal(t, tt.want, got)
				return
			}
			require.Equal(t, tt.want, got)
			require.EqualError(t, err, "user with ID 456456 not found")
		})
	}
}

func TestUserRepositoryDb_FindByEmail(t *testing.T) {
	type args struct {
		email string
	}
	tests := []struct {
		name    string
		args    args
		want    *domain.User
		wantErr bool
	}{
		{
			name: "Must return error when email is not found",
			args: args{
				email: "test-not-used@email.com",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Must return the user with the same email when email is found",
			args: args{
				email: "test@email.com",
			},
			want: &domain.User{
				ID:    "123123",
				Name:  "test-name",
				Email: "test@email.com",
				Roles: []domain.Role{
					domain.Watcher,
				},
			},
		},
	}
	for _, tt := range tests {
		db := getTestDatabase()
		t.Run(tt.name, func(t *testing.T) {
			ur := &UserRepositoryDb{
				dbInMemory: db,
			}
			got, err := ur.FindByEmail(tt.args.email)
			if !tt.wantErr {
				require.Equal(t, tt.want, got)
				return
			}
			require.EqualError(t, err, "user with email test-not-used@email.com not found")

		})
	}
}

func TestUserRepositoryDb_FindById(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		args    args
		want    *domain.User
		wantErr bool
	}{
		{
			name: "Must return error when id is not found",
			args: args{
				id: "456456",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Must return the user with the same id when id is found",
			args: args{
				id: "123123",
			},
			want: &domain.User{
				ID:    "123123",
				Name:  "test-name",
				Email: "test@email.com",
				Roles: []domain.Role{
					domain.Watcher,
				},
			},
		},
	}
	for _, tt := range tests {
		db := getTestDatabase()
		t.Run(tt.name, func(t *testing.T) {
			ur := &UserRepositoryDb{
				dbInMemory: db,
			}
			got, err := ur.FindById(tt.args.id)
			if !tt.wantErr {
				require.Equal(t, tt.want, got)
				return
			}
			require.EqualError(t, err, "user with ID 456456 not found")

		})
	}
}

func TestUserRepositoryDb_GetAll(t *testing.T) {

	tests := []struct {
		name string
		want []*domain.User
	}{
		{
			name: "Must return the database with the right values",
			want: []*domain.User{
				{
					ID:    "123123",
					Name:  "test-name",
					Email: "test@email.com",
					Roles: []domain.Role{
						domain.Watcher,
					},
				},
				{
					ID:    "321321",
					Name:  "test-name",
					Email: "test-1@email.com",
					Roles: []domain.Role{
						domain.Watcher,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		db := getTestDatabase()

		t.Run(tt.name, func(t *testing.T) {
			ur := &UserRepositoryDb{
				dbInMemory: db,
			}
			got := ur.GetAll()
			require.Equal(t, got, tt.want)
			require.Len(t, got, 2)
		})
	}
}

func TestUserRepositoryDb_Insert(t *testing.T) {
	testUuid := uuid.NewV4().String()
	type args struct {
		user *domain.User
	}
	tests := []struct {
		name    string
		args    args
		want    *domain.User
		wantErr bool
	}{
		{
			name: "Must return a new user when a valid user is provided and insert a new user in database",
			args: args{
				&domain.User{
					Name: "valid-test-user",
					Roles: []domain.Role{
						domain.Watcher,
					},
					Email: "test.new.user@email.com",
					ID:    testUuid,
				},
			},
			want: &domain.User{
				Name: "valid-test-user",
				Roles: []domain.Role{
					domain.Watcher,
				},
				Email: "test.new.user@email.com",
				ID:    testUuid,
			},
		},
		{
			name:    "Must return an error when the user provided is invalid",
			args:    args{nil},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		db := getTestDatabase()

		t.Run(tt.name, func(t *testing.T) {
			ur := &UserRepositoryDb{
				dbInMemory: db,
			}
			got, err := ur.Insert(tt.args.user)
			if !tt.wantErr {
				require.Len(t, ur.dbInMemory, 3)
				require.Equal(t, tt.want, got)
				return
			}
			require.EqualError(t, err, "user cannot be nil")
		})
	}
}

func TestUserRepositoryDb_UpdateRoles(t *testing.T) {
	type args struct {
		userId string
		roles  []domain.Role
	}
	tests := []struct {
		name    string
		args    args
		want    *domain.User
		wantErr bool
	}{
		{
			name: "Must return an error when user provided is not found",
			args: args{
				userId: "99999",
				roles: []domain.Role{
					domain.Modifier,
					domain.Watcher,
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Must return an user with the right roles when user id and roles are valid",
			args: args{
				userId: "123123",
				roles: []domain.Role{
					domain.Modifier,
				},
			},
			want: &domain.User{
				ID:    "123123",
				Name:  "test-name",
				Email: "test@email.com",
				Roles: []domain.Role{
					domain.Modifier,
					domain.Watcher,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		db := getTestDatabase()

		t.Run(tt.name, func(t *testing.T) {
			ur := &UserRepositoryDb{
				dbInMemory: db,
			}

			got, err := ur.UpdateRoles(tt.args.userId, tt.args.roles)
			if !tt.wantErr {
				require.Equal(t, tt.want, got)
				return
			}
			require.EqualError(t, err, "user with ID 99999 not found")
		})
	}
}
