package domain

import (
	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

type User struct {
	ID       string `valid:"uuid"`
	Name     string `valid:"notnull"`
	Email    string `valid:"email"`
	Roles    []Role `valid:"notnull"`
	IsActive bool   `valid:"required"`
}

type Role string

const (
	Admin    Role = "Admin"
	Modifier Role = "Modifier"
	Watcher  Role = "Watcher"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

func (u *User) Validate() error {
	_, err := govalidator.ValidateStruct(u)

	if err != nil {
		return err
	}

	return nil
}

func (us *User) SetRoles(roles []Role) []Role {
	hasAdmin := false
	hasModifier := false

	for _, role := range roles {
		if role == Admin {
			hasAdmin = true
			break
		}
		if role == Modifier {
			hasModifier = true
		}
	}

	switch {
	case hasAdmin:
		return []Role{Admin, Modifier, Watcher}
	case hasModifier:
		return []Role{Modifier, Watcher}
	default:
		return []Role{Watcher}
	}
}

func NewUser(name, email string, roles []Role) (*User, error) {

	user := &User{
		ID:       uuid.NewV4().String(),
		Name:     name,
		Email:    email,
		IsActive: true,
	}

	user.Roles = user.SetRoles(roles)

	err := user.Validate()

	if err != nil {
		return nil, err
	}

	return user, nil
}
