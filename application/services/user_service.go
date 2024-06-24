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

func (us UserService) alreadyHaveTheUser(email string) (bool, error) {
	user, _ := us.UserRepository.FindByEmail(email)

	if user != nil {
		return true, fmt.Errorf("user with email %s already exist", user.Email)
	}

	return false, nil
}

func (us UserService) CreateUser(user domain.User) error {
	// Check if User Already exist
	haveUser, err := us.alreadyHaveTheUser(user.Email)
	if haveUser || err != nil {
		return err
	}
	// Create and validate new User
	newUser, err := domain.NewUser(user.Name, user.Email, user.Roles)

	if err != nil {
		return err
	}

	// call the repository inserting the user and the inMemory DataBase
	_, err = us.UserRepository.Insert(&newUser)

	if err != nil {
		return err
	}

	return nil
}

func (us UserService) GetUser(userID string) ([]*domain.User, error) {
	if userID != "" {
		// if id is provided, call the repository to get one user by id
		user, err := us.UserRepository.FindById(userID)

		if err != nil {
			return nil, err
		}

		return []*domain.User{user}, nil
	}
	// if not id is provided, get all users
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
