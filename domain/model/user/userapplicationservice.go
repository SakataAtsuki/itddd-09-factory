package user

import (
	"fmt"
	"log"
)

type UserApplicationService struct {
	userFactory    UserFactorier
	userRepository UserRepositorier
	userService    UserService
}

func NewUserApplicationService(userFactory UserFactorier, userRepository UserRepositorier, userService UserService) (*UserApplicationService, error) {
	return &UserApplicationService{userFactory: userFactory, userRepository: userRepository, userService: userService}, nil
}

func (uas *UserApplicationService) Register(name string) (err error) {
	defer func() {
		if err != nil {
			err = &RegisterError{Name: name, Message: fmt.Sprintf("userapplicationservice.Register err: %s", err), Err: err}
		}
	}()
	userName, err := NewUserName(name)
	if err != nil {
		return err
	}

	user, err := uas.userFactory.Create(*userName)
	if err != nil {
		return err
	}

	isUserExists, err := uas.userService.Exists(user)
	if err != nil {
		return err
	}
	if isUserExists {
		return fmt.Errorf("user name of %s is already exists", name)
	}

	if err := uas.userRepository.Save(user); err != nil {
		return err
	}

	log.Printf("user name of %s is successfully saved", name)
	return nil
}

type RegisterError struct {
	Name    string
	Message string
	Err     error
}

func (err *RegisterError) Error() string {
	return err.Message
}
