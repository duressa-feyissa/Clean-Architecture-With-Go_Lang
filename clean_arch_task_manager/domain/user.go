package domain

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)
const (
	CollectionUser = "Users"
)

type User struct {
	ID       string `validate:"-"`
	Username string `validate:"required"`
	Password string `validate:"required"`
	Role     string `validate:"required"`
}

type Login struct {
	ID string `validate:"required"`
	Password string `validate:"required"`
}

func (l *Login) ValidateLogin() error {
	validate := validator.New()
	err := validate.Struct(l)
	if err != nil {
		return errors.New(err.Error())
	}
	return nil
}

func (u *User) Validate() error {
	validate := validator.New()
	err := validate.Struct(u)
	if err != nil {
		return errors.New(err.Error())
	}
	return nil
}

// NewUser creates a new user with a unique ID.
func NewUser(username, password, role string) *User {
	userUUID := uuid.New().String()
	// Create a unique ID by combining the user's ID and the task's UUID.
	uniqueUserID := fmt.Sprintf("%s-%s", username, userUUID)
	return &User{
		ID:       uniqueUserID,
		Username: username,
		Password: password,
		Role:     role,
	}
}
type UserRepository interface {
	RegisterUser(user *User) error
	GetUserByID(id string) (*User, error)
}
type UserUsecase interface {
	RegisterUser(user *User) error
	LoginUser(user *User) (string, error)
	GetUserByID(id string) (*User, error)
}
