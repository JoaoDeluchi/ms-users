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

// func to set fields as required ASAP
func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

// validate struct using govalidator
func (u *User) Validate() error {
	_, err := govalidator.ValidateStruct(u)

	if err != nil {
		return err
	}

	return nil
}

func SetRoles(roles []Role) []Role {
	hasAdmin := false
	hasModifier := false

	// Check for Admin and Modifier roles in the input
	for _, role := range roles {
		if role == Admin {
			hasAdmin = true
			break // Early exit - dont need to iterate if has Admin
		}
		if role == Modifier {
			hasModifier = true
		}
	}

	// Return roles based on the presence of Admin and Modifier
	switch {
	case hasAdmin:
		return []Role{Admin, Modifier, Watcher}
	case hasModifier:
		return []Role{Modifier, Watcher}
	default:
		return []Role{Watcher}
	}
}

// This function is like a "constructor of the class"
func NewUser(name, email string, roles []Role) (*User, error) {

	user := &User{
		ID:       uuid.NewV4().String(),
		Name:     name,
		Email:    email,
		Roles:    SetRoles(roles),
		IsActive: true,
	}

	err := user.Validate()

	if err != nil {
		return nil, err
	}

	return user, nil
}
