package user

import (
	"github.com/google/uuid"
)

type UserFactorier interface {
	Create(UserName) (*User, error)
}

type UserFactory struct{}

func NewUserFactory() (*UserFactory, error) {
	return &UserFactory{}, nil
}

func (uf *UserFactory) Create(name UserName) (*User, error) {
	uuidV4, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	userId, err := NewUserId(uuidV4.String())
	if err != nil {
		return nil, err
	}

	user, err := NewUser(*userId, name)
	if err != nil {
		return nil, err
	}
	return user, nil
}
