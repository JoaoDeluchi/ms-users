package repositories

import (
	"fmt"

	"github.com/joaodeluchi/ms-users/domain"
)

type UserRepository interface {
	Insert(user *domain.User) (*domain.User, error)
	GetAll(id string) ([]*domain.User, error)
	FindById(id string) (*domain.User, error)
	FindByEmail(email string) *domain.User
	UpdateRoles(userId string, roles []domain.Role) (*domain.User, error)
	DeleteUser(id string) ([]*domain.User, error)
}

type UserRepositoryDb struct {
	dbInMemory []*domain.User
}

func (ur UserRepositoryDb) Insert(user *domain.User) (*domain.User, error) {
	for _, existingUser := range ur.dbInMemory {
		if existingUser.ID == user.ID {
			return nil, fmt.Errorf("user with ID %s already exists", user.ID)
		}
	}

	ur.dbInMemory = append(ur.dbInMemory, user)

	return user, nil
}

func (ur UserRepositoryDb) GetAll(id string) ([]*domain.User, error) {
	if id == "" {
		return ur.dbInMemory, nil
	}

	user, err := ur.FindById(id)
	if err != nil {
		return nil, err
	}

	return []*domain.User{user}, nil
}

func (ur UserRepositoryDb) FindById(id string) (*domain.User, error) {
	for _, user := range ur.dbInMemory {
		if user.ID == id {
			return user, nil
		}
	}
	return nil, fmt.Errorf("user with ID %s not found", id)
}

func (ur UserRepositoryDb) FindByEmail(email string) *domain.User {
	for _, user := range ur.dbInMemory {
		if user.Email == email {
			return user
		}
	}

	return nil
}

func (ur UserRepositoryDb) UpdateRoles(userId string, roles []domain.Role) (*domain.User, error) {
	foundIndex := -1
	for i, existingUser := range ur.dbInMemory {
		if existingUser.ID == userId {
			foundIndex = i
			break
		}
	}

	if foundIndex == -1 {
		return nil, fmt.Errorf("user with ID %s not found", userId)
	}

	ur.dbInMemory[foundIndex].SetRoles(roles)

	return ur.dbInMemory[foundIndex], nil
}

func (ur UserRepositoryDb) DeleteUser(id string) ([]*domain.User, error) {
	targetIndex := -1
	for i, user := range ur.dbInMemory {
		if user.ID == id {
			targetIndex = i
			break
		}
	}

	if targetIndex == -1 {
		return ur.dbInMemory, fmt.Errorf("user with ID %s not found", id)
	}

	ur.dbInMemory = append(ur.dbInMemory[:targetIndex], ur.dbInMemory[targetIndex+1:]...)

	return ur.dbInMemory, nil
}

func NewUserRepository() UserRepository {
	db := []*domain.User{}
	return UserRepositoryDb{
		dbInMemory: db,
	}
}
