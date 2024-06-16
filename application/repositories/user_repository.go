package repositories

import (
	"fmt"

	"github.com/joaodeluchi/ms-users/domain"
	"gorm.io/gorm"
)

type UserRepository interface {
	Insert(user *domain.User) (*domain.User, error)
	GetAll(id string) (*domain.User, error)
	FindById(id string) (*domain.User, error)
	FindByEmail(email string) (*domain.User, error)
	Update(user *domain.User) (*domain.User, error)
}

type UserRepositoryDb struct {
	DB *gorm.DB
}

func (ur UserRepositoryDb) Insert(user *domain.User) (*domain.User, error) {
	err := ur.DB.Create(user).Error

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (ur UserRepositoryDb) GetAll(id string) (*domain.User, error) {
	var user domain.User

	if id != "" {
		return ur.FindById(id)
	}

	ur.DB.Find(&user)

	if user.ID == "" {
		return nil, fmt.Errorf("user not found")
	}

	return &user, nil
}

func (ur UserRepositoryDb) FindById(id string) (*domain.User, error) {
	var user domain.User

	ur.DB.First(&user, "id = ?", id)

	// if user.id is empty, means that the user was not found and user id is invalid
	if user.ID == "" {
		return nil, fmt.Errorf("user not found")
	}

	return &user, nil
}

func (ur UserRepositoryDb) FindByEmail(email string) (*domain.User, error) {
	var user domain.User

	ur.DB.First(&user, "email = ?", email)

	// if user.id is empty, means that the user was not found and user id is invalid
	if user.ID == "" {
		return nil, fmt.Errorf("user not found")
	}

	return &user, nil
}

func (ur UserRepositoryDb) Update(user *domain.User) (*domain.User, error) {
	err := ur.DB.Save(&user).Error

	if err != nil {
		return nil, err
	}

	return user, nil
}

func NewUserRepository(Db *gorm.DB) UserRepository {
	return UserRepositoryDb{DB: Db}
}
