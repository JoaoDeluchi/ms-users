package services

import (
	"fmt"

	"github.com/joaodeluchi/ms-users/application/repositories"
	"github.com/joaodeluchi/ms-users/domain"
)

type IUserService interface {
	CreateUser(user domain.User) error
	GetUser(userID string) (domain.User, error)
	UpdateUserRoles(userID string, roles []string) error
	DeleteUser(userID string) error
}

type UserService struct {
	UserRepository repositories.UserRepository
}

func (us UserService) CreateUser(user domain.User) error {
	u := us.UserRepository.FindByEmail(user.Email)

	if u.ID != "" {
		return fmt.Errorf("user already exist")
	}

	if err := user.Validate(); err != nil {
		return err
	}

	us.UserRepository.Insert(&user)

	return nil
}

func (us UserService) GetUser(userID string) ([]*domain.User, error) {
	if userID != "" {
		user, err := us.UserRepository.FindById(userID)

		if err != nil {
			return nil, err
		}

		return []*domain.User{user}, nil
	}

	user := us.UserRepository.GetAll()

	return user, nil
}

func (us UserService) UpdateUserRoles(userID string, roles []domain.Role) error {
	_, err := us.UserRepository.UpdateRoles(userID, roles)

	if err != nil {
		return err
	}

	return nil
}

func (us UserService) DeleteUser(userID string) error {
	_, err := us.UserRepository.DeleteUser(userID)

	if err != nil {
		return err
	}

	return nil
}

func NewUserService(userRepo repositories.UserRepository) UserService {
	return UserService{
		UserRepository: userRepo,
	}
}
