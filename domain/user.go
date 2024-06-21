package domain

import (
	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

type User struct {
	ID    string `valid:"uuid" json:"id"`
	Name  string `valid:"notnull" json:"name"`
	Email string `valid:"email" json:"email"`
	Roles []Role `valid:"notnull" json:"roles"`
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

// Validate validates the User struct based on field tags
func (u *User) Validate() error {
	_, err := govalidator.ValidateStruct(u)

	if err != nil {
		return err
	}

	return nil
}

// SetRoles sets the user's roles with logic to ensure a valid role combination
func (us *User) SetRoles(roles []Role) []Role {
	hasAdmin := false
	hasModifier := false

	for _, role := range roles {
		if role == Admin {
			hasAdmin = true
			// break to avoid unnecessary check
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
		// Default role is Watcher if no valid roles provided
		return []Role{Watcher}
	}
}

func NewUser(name, email string, roles []Role) (User, error) {
	user := User{
		ID:    uuid.NewV4().String(),
		Name:  name,
		Email: email,
	}

	user.Roles = user.SetRoles(roles)

	err := user.Validate()

	if err != nil {
		return User{}, err
	}

	return user, nil
}
